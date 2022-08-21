package workerpool

type PoolOption func(*Pool)

func WithMaxWorker(maxWorkers int) PoolOption {
	return func(p *Pool) {
		p.maxWorkers = maxWorkers
	}
}
