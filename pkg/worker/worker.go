package worker

import (
	"context"
	"errors"
	"strings"

	"github.com/iancoleman/strcase"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
	"rogchap.com/v8go"
)

// Worker type encapsulates execution environment
type Worker struct {
	v8ctx *v8go.Context
}

// ServiceRef is used to instantiate access to external RPC references
type ServiceRef struct {
	Description grpc.ServiceDesc
	Client      interface{}
}

// NewWorker creates a client with a reference to the provided service references
func NewWorker(refs ...ServiceRef) *Worker {
	iso := v8go.NewIsolate()

	global := v8go.NewObjectTemplate(iso)

	worker := &Worker{}

	for _, ref := range refs {
		service := v8go.NewObjectTemplate(iso)
		serviceName := strcase.ToLowerCamel(strings.Split(ref.Description.ServiceName, ".")[1])

		for _, methodDesc := range ref.Description.Methods {
			service.Set(
				strcase.ToLowerCamel(methodDesc.MethodName),
				v8go.NewFunctionTemplate(iso, func(i *v8go.FunctionCallbackInfo) *v8go.Value {
					resolver, _ := v8go.NewPromiseResolver(worker.v8ctx)

					// Make async remote RPC call
					go func(info *v8go.FunctionCallbackInfo, r *v8go.PromiseResolver) {
						defer info.Release()
						defer r.Release()

						// Invoke the js RPC with that string expecting a string output
						out, err := methodDesc.Handler(ref.Client, context.Background(), func(in interface{}) error {
							i, iok := in.(protoreflect.ProtoMessage)
							if !iok {
								return errors.New("input type not proto message")
							}
							protojson.Unmarshal([]byte(info.Args()[0].String()), i)
							return nil
						}, nil)

						o, ook := out.(protoreflect.ProtoMessage)
						if !ook {
							if err != nil {
								v, _ := v8go.NewValue(iso, "output not a proto message")
								r.Reject(v)
								return
							}
						}

						res, err := protojson.Marshal(o)
						if err != nil {
							v, _ := v8go.NewValue(iso, err.Error())
							r.Reject(v)
							return
						}

						// Convert string output to JSON value
						v, err := v8go.JSONParse(worker.v8ctx, string(res))
						if err != nil {
							v, _ := v8go.NewValue(iso, err.Error())
							r.Reject(v)
							return
						}

						// Resolve promise with json value
						r.Resolve(v)
					}(i, resolver)

					return resolver.GetPromise().Value
				}),
			)
		}

		global.Set(serviceName, service)
	}

	worker.v8ctx = v8go.NewContext(iso, global)

	return worker
}

// Run executes the passed in script and returns the result via a channel
func (worker *Worker) Run(source, origin string) (chan *v8go.Value, chan error) {
	resChan := make(chan *v8go.Value)
	errChan := make(chan error)

	go func(rc chan *v8go.Value, ec chan error) {
		res, err := worker.v8ctx.RunScript(source, origin)
		if err != nil {
			ec <- err
			return
		}

		if p, err := res.AsPromise(); err != nil {
			// Handle non-promises
			rc <- res

		} else if p.State() == v8go.Pending {
			// Handle pending promises
			p.Then(
				func(info *v8go.FunctionCallbackInfo) *v8go.Value {
					rc <- info.Args()[0]
					return nil
				},
			).Catch(
				func(info *v8go.FunctionCallbackInfo) *v8go.Value {
					ec <- errors.New(info.Args()[0].String())
					return nil
				},
			)
		} else if p.State() == v8go.Fulfilled {
			// Handles resolved promises
			rc <- p.Result()
		} else {
			// Handle rejected promises
			ec <- errors.New(p.Result().String())
		}
	}(resChan, errChan)

	return resChan, errChan
}

// Resolve runs the microtask checkpoint until everything resolvable is resolved
func (worker *Worker) Resolve() {
	worker.v8ctx.PerformMicrotaskCheckpoint()
}

// Stringify uses the current v8 context to convert a v8 object to a string
func (worker *Worker) Stringify(v *v8go.Value) (string, error) {
	return v8go.JSONStringify(worker.v8ctx, v)
}
