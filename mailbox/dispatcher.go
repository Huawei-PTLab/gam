package mailbox

import "runtime"

type Dispatcher interface {
	Schedule(fn func())
	Throughput() int
	AfterStart()
	BeforeTerminate()
	BeforeBatchProcess()
	BeforeProcessingMessage(index *int)
}

type goroutineDispatcher int

func (goroutineDispatcher) Schedule(fn func()) {
	go fn()
}

func (d goroutineDispatcher) AfterStart() {
}
func (d goroutineDispatcher) BeforeTerminate() {
}
func (d goroutineDispatcher) Throughput() int {
	return int(d)
}

func (d goroutineDispatcher) BeforeProcessingMessage(index *int) {
	if *index > int(d) {
		*index = 0
		runtime.Gosched()
	}
}

func (d goroutineDispatcher) BeforeBatchProcess() {

}

func NewDefaultDispatcher(throughput int) Dispatcher {
	return goroutineDispatcher(throughput)
}
