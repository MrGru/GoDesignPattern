package worker_pool

import (
	"fmt"
	"regexp"
	"sync"
	"testing"
)

func Test_Dispatcher(t *testing.T) {
	bufferSize := 100
	var dispatcher = NewDispatcher(bufferSize)
	workers := 3
	for i := 0; i < workers; i++ {
		var w = &PreffixSuffixWorker{
			prefixS: fmt.Sprintf("WorkerID: %d -> ", i),
			suffixS: " World",
			id:      i,
		}
		dispatcher.LaunchWorker(i, w)
	}
	requests := 10
	var wg sync.WaitGroup
	wg.Add(requests)
	for i := 0; i < requests; i++ {
		// req := NewStringRequest("(Msg_id: %d) -> Hello", i, &wg)
		// dispatcher.MakeRequest(req)

		req := Request{
			Data: fmt.Sprintf("(Msg_id: %d) -> Hello", i),
			Handler: func(i interface{}) {
				s, ok := i.(string)
				defer wg.Done()
				if !ok {
					t.Fail()
				}
				ok, err := regexp.Match(`WorkerID\: \d* -\> \(MSG_ID: \d*\) -> [A-Z]*\sWorld`, []byte(s))
				if !ok || err != nil {
					t.Fail()
				}
			},
		}
		dispatcher.MakeRequest(req)
	}
	dispatcher.Stop()
	wg.Wait()
}
