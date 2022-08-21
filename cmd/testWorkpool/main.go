package main

import (
	"fmt"
	"github.com/leosykes117/tasker/internal/testworker"
	"github.com/leosykes117/tasker/pkg/workerpool"
	"runtime"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("CPUs:", runtime.NumCPU())
	go func() {
		for {
			// Se resta 2 porque una gorutine corresponde a la función main y la segunda es esta
			// función de monitoreo
			fmt.Printf("[main] Total current goroutine: %d\n", runtime.NumGoroutine()-2)
			time.Sleep(1 * time.Second)
		}
	}()
	concurrent()
	//sequential()
	fmt.Printf("Task time: %s", time.Since(start))
}

func concurrent() {
	wp := workerpool.NewWorkerPool()
	wp.Run()
	for i := 0; i < 10; i++ {
		task := testworker.NewWorkerTester(i + 1)
		wp.AddTask(task)
	}
	wp.Shutdown()
	fmt.Println("Tareas terminadas")
}

func sequential() {
	for i := 0; i < 10; i++ {
		task := testworker.NewWorkerTester(i + 1)
		task.Task()
	}
}
