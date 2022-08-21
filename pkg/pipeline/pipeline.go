package pipeline

type Executor func(interface{}) (interface{}, error)

type Pipeline interface {
	Pipe(Executor) Pipeline
	Merge() <-chan interface{}
}

type pipeline struct {
	dataC     chan interface{}
	errC      chan error
	executors []Executor
}

func New(f func(chan interface{})) Pipeline {
	inC := make(chan interface{})
	go f(inC)
	return &pipeline{
		dataC:     inC,
		errC:      make(chan error),
		executors: []Executor{},
	}
}

// Pipe add Executor (stage) to our array of executors
func (p *pipeline) Pipe(executor Executor) Pipeline {
	p.executors = append(p.executors, executor)
	return p
}

// Merge run Executor(stage) one by one, and use the result as an input for the next Executor(stage)
func (p *pipeline) Merge() <-chan interface{} {
	for i := 0; i < len(p.executors); i++ {
		p.dataC, p.errC = run(p.dataC, p.executors[i])
	}
	return p.dataC
}

// run executes or start a Executor (stage)
func run(inC <-chan interface{}, f Executor) (chan interface{}, chan error) {
	outC := make(chan interface{})
	errC := make(chan error)
	// TODO: CONSIDERAR SI EN ESTA PARTE SE TOMA UN WORKER DE POOL
	go func() {
		defer close(outC)
		for v := range inC {
			res, err := f(v)
			if err != nil {
				errC <- err
				continue
			}
			outC <- res
		}
	}()
	return outC, errC
}
