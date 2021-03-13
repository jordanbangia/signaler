package signaler_test

import (
	"sync"
	"testing"
	"time"

	"github.com/jordanbangia/signaler"
)

func TestSignaler(t *testing.T) {
	s := signaler.New()

	// start some clients
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-s.Subscribe()
		}()
	}

	// wait a second for all the goroutines to subscribe
	time.Sleep(100 * time.Millisecond)

	s.Trigger()

	wg.Wait()
}

func TestTriggerWithNoSubscribers(t *testing.T) {
	// given a new signaler
	s := signaler.New()

	// when trigger is called
	s.Trigger()

	// then no action happens, new subscriptions shouldn't wait
	select {
	case <-s.Subscribe():
		t.Fatal("subscirbe should not complete")
	case <-time.After(10 * time.Millisecond):
		// subscribe should not complete
	}
}
