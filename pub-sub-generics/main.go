package main

import (
	"errors"
	"log"
	"math/rand"
	"sync"
	"time"
)

type fakePubSubServer struct {
	ticker             *time.Ticker
	subscribeCallbacks []func()
	mu                 sync.Mutex // Needed to ensure addtions to the list are thread safe.
}

func (server *fakePubSubServer) registerSubscriber(f func()) {
	server.mu.Lock()
	server.subscribeCallbacks = append(server.subscribeCallbacks, f)
	server.mu.Unlock()
}

func NewFakePubSubServer(intervalToCallSubscribersInSeconds int) *fakePubSubServer {
	ticker := time.NewTicker(time.Second * time.Duration(intervalToCallSubscribersInSeconds))

	server := &fakePubSubServer{
		ticker: ticker,
	}

	// running ther server in a different go routine so it doesn't block.

	go func() {
		for {
			// Waiting for each tick.
			<-server.ticker.C

			// Calling all subscribers.
			server.mu.Lock()
			for _, callback := range server.subscribeCallbacks {
				callback()
			}
			server.mu.Unlock()
		}
	}()

	return server
}

type AsyncResult[T any] struct {
	Result T
	Error  error
}

func main() {
	fakeServer := NewFakePubSubServer(2)

	resultCh := make(chan AsyncResult[int], 1)
	rand.Seed(time.Now().UnixNano())
	fakeServer.registerSubscriber(func() {
		// Deciding to return error or success.
		if rand.Int()%2 == 0 {
			resultCh <- AsyncResult[int]{
				Result: 1,
			}
		} else {
			resultCh <- AsyncResult[int]{
				Error: errors.New("error in the subscriber"),
			}
		}

		close(resultCh)
	})

	const timeToWaitForCallbackToBeCalledInSec = 3
	select {
	case result := <-resultCh:
		if result.Error != nil {
			log.Println("Callback was called with an error.", result.Error)
			return
		}

		log.Println("Callback was successfully called with a result value of: ", result.Result)
		return
	case <-time.After(time.Second * timeToWaitForCallbackToBeCalledInSec):
		log.Println("Timed out waiting for callback to be called.")
		return
	}
}
