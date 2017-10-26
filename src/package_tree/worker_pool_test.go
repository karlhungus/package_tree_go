package main

import "testing"

func TestWorkerPool(t *testing.T) {
	pkg := NewPackager()
	msg := NewMessage("INDEX|foo|")

	w := Work{
		input: msg,
		workable: func(m InputData) Result {
			return pkg.Process(m.(Message))
		},
  }
	pool := NewWorkerPool(1)
	pool.input <- &w
	<- pool.output
	if "OK\n" != w.result {
		t.Error("Expected OK got", w.result)
	}
}

