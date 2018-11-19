package worker_pool

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"
)

type Request struct {
	Data    interface{}
	Handler RequestHandler
}

type RequestHandler func(interface{})

func NewStringRequest(s string, id int, wg *sync.WaitGroup) Request {
	myRequest := Request{
		Data: fmt.Sprintf(s, id), Handler: func(i interface{}) {
			defer wg.Done()
			s, ok := i.(string)
			if !ok {
				log.Fatal("Invalid casting to string")
			}
			fmt.Println(s)
		},
	}
	return myRequest
}

type WorkerLauncher interface {
	LaunchWorker(in chan Request)
}

type Dispatcher interface {
	LaunchWorker(w WorkerLauncher)
	MakeRequest(Request)
	Stop()
}

type dispatcher struct {
	inCh chan Request
}

func (d *dispatcher) LaunchWorker(id int, w WorkerLauncher) {
	w.LaunchWorker(d.inCh)
}

func (d *dispatcher) Stop() {
	close(d.inCh)
}

func (d *dispatcher) MakeRequest(r Request) {
	select {
	case d.inCh <- r:
	case <-time.After(time.Second * 5):
		return
	}

}

func NewDispatcher(b int) *dispatcher {
	return &dispatcher{
		inCh: make(chan Request, b),
	}
}

type PreffixSuffixWorker struct {
	id      int
	prefixS string
	suffixS string
}

func (w *PreffixSuffixWorker) uppercase(in <-chan Request) <-chan Request {
	out := make(chan Request)
	go func() {
		for msg := range in {
			s, ok := msg.Data.(string)
			if !ok {
				msg.Handler(nil)
				continue
			}
			msg.Data = strings.ToUpper(s)
			out <- msg
		}
		close(out)
	}()
	return out
}

func (w *PreffixSuffixWorker) append(in <-chan Request) <-chan Request {
	out := make(chan Request)
	go func() {
		for msg := range in {
			uppercaseString, ok := msg.Data.(string)
			if !ok {
				msg.Handler(nil)
				continue
			}
			msg.Data = fmt.Sprintf("%s%s", uppercaseString, w.suffixS)
			out <- msg
		}
		close(out)
	}()
	return out
}

func (w *PreffixSuffixWorker) prefix(in <-chan Request) {
	go func() {
		for msg := range in {
			uppercaseStringWithSuffix, ok := msg.Data.(string)
			if !ok {
				msg.Handler(nil)
				continue
			}
			msg.Handler(fmt.Sprintf("%s%s", w.prefixS, uppercaseStringWithSuffix))
		}
	}()
}

func (w *PreffixSuffixWorker) LaunchWorker(in chan Request) {
	w.prefix(w.append(w.uppercase(in)))
}
