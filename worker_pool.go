package main

type InputData interface{}
type Result interface{}

type Work struct {
	workable func(InputData) Result
	input InputData
	result Result
}

type WorkerPool struct {
	input chan *Work
	output chan *Work
}

func NewWorkerPool(workers int) WorkerPool {
	pool := WorkerPool{
		input: make(chan *Work),
		output: make(chan *Work),
	}
	for i := 0; i< workers; i++ {
		go pool.work()
	}
	return pool
}

func (pool WorkerPool) work() {
	for {
		w := <- pool.input
		w.result = w.workable(w.input)
		pool.output <- w
	}
}

