package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/RGood/v8-rpc/internal/generated/protos/clock"
	"github.com/RGood/v8-rpc/pkg/plugins"
	"github.com/RGood/v8-rpc/pkg/worker"
)

func main() {
	ref := worker.ServiceRef{
		// Description helps us build a JS reference to the service.
		// Alternatively we could just use reflection and the client by itself.
		Description: clock.Clock_ServiceDesc,

		// In an ideal world, this would be a client to an RCP interface instead of the service impl directly,
		// but that's difficult for a few reasons:
		// 1. Would need to implement some RPC based comms
		//   a. Either a local impl
		//   b. Or set up a grpc service to the clock on a separate process and hit it over some legitimate RPC
		// 2. The Description has a Handler that only knows how to talk to a service implementation,
		//    so we'd have to wrap the client in some proxy service (blah)
		Client: plugins.NewClock(),
	}

	// Instantiate a new worker and pass it the service reference(s)
	w := worker.NewWorker(ref)

	wg := sync.WaitGroup{}
	start := time.Now()

	reqCount := 1000
	totalTime := atomic.Uint64{}

	// Call the worker N times
	for i := 0; i < reqCount; i++ {
		wg.Add(1)
		go func() {
			rs := time.Now()
			// `clock.now` is polyfilled completely from the service description
			resChan, errChan := w.Run("clock.now({})", "main.js")

			// Wait for a success / error response (panic on error)
			select {
			case res := <-resChan:
				resStr, _ := w.Stringify(res)
				fmt.Printf("%v\n", resStr)
			case err := <-errChan:
				panic(err)
			}

			totalTime.Add(uint64(time.Since(rs)))
			wg.Done()
		}()
	}

	allStarted := time.Now()

	// Wait for all requests to complete
	wg.Wait()

	allDone := time.Now()

	totalDuration := allDone.Sub(start)

	fmt.Printf("All started in: %s\n", allStarted.Sub(start))
	fmt.Printf("Finished in %s\n", totalDuration)
	fmt.Printf("Average Req Duration: %s\n", time.Duration(totalTime.Load()/uint64(reqCount)))
	fmt.Printf("Average Clock Time: %s\n", totalDuration/time.Duration(reqCount))

}
