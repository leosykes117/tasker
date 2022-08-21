package testworker

import (
	"fmt"
	"time"
)

type WorkerTester struct {
	waitSec time.Duration
}

func NewWorkerTester(sec int) *WorkerTester {
	return &WorkerTester{
		waitSec: time.Duration(sec) * time.Second,
	}
}

func (wt *WorkerTester) Task() {
	fmt.Printf("Iniciando tarea de %s\n", wt.waitSec)
	time.Sleep(wt.waitSec)
	fmt.Printf("Terminando tarea de %s\n", wt.waitSec)
}
