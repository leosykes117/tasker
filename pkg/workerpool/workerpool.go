package workerpool

import (
	"runtime"
	"sync"
)

type WorkerPool interface {
	Run()
	AddTask(Worker)
	Shutdown()
}

type Pool struct {
	maxWorkers   int
	queuedTaskCh chan Worker
	wg           sync.WaitGroup
}

func NewWorkerPool(options ...PoolOption) WorkerPool {
	p := &Pool{
		maxWorkers:   runtime.NumCPU(),
		queuedTaskCh: make(chan Worker),
	}
	for _, option := range options {
		option(p)
	}
	return p
}

func (p *Pool) Run() {
	p.wg.Add(p.maxWorkers)
	for i := 0; i < p.maxWorkers; i++ {
		go func(workerID int) {
			defer p.wg.Done()
			for worker := range p.queuedTaskCh {
				worker.Task()
			}
		}(i + 1)
	}
}

func (p *Pool) AddTask(w Worker) {
	p.queuedTaskCh <- w
}

func (p *Pool) Shutdown() {
	close(p.queuedTaskCh)
	p.wg.Wait()
}
