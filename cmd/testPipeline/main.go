package main

import (
	"fmt"
	"github.com/leosykes117/tasker/pkg/pipeline"
	"runtime"
	"time"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	go func() {
		for {
			// Se resta 2 porque una gorutine corresponde a la función main y la segunda es esta
			// función de monitoreo
			fmt.Printf("[main] Total current goroutine: %d\n", runtime.NumGoroutine()-2)
			time.Sleep(1 * time.Second)
		}
	}()
	start := time.Now()
	//sequential()
	concurrent()
	fmt.Printf("Took: %s\n", time.Since(start))
}

func concurrent() {
	outC := pipeline.New(func(inC chan interface{}) {
		defer close(inC)
		for i := 0; i < 10; i++ {
			inC <- i
		}
	}).
		Pipe(func(in interface{}) (interface{}, error) {
			return multiplyTwo(in.(int)), nil
		}).
		Pipe(func(in interface{}) (interface{}, error) {
			return square(in.(int)), nil
		}).
		Pipe(func(in interface{}) (interface{}, error) {
			return addQuote(in.(int)), nil
		}).
		Pipe(func(in interface{}) (interface{}, error) {
			return addFoo(in.(string)), nil
		}).
		Merge()
	for result := range outC {
		fmt.Println(result)
	}
}

func sequential() {
	for i := 0; i < 10; i++ {
		println(addFoo(addQuote(square(multiplyTwo(i)))))
	}
}

func multiplyTwo(v int) int {
	time.Sleep(1 * time.Second)
	return v * 2
}

func square(v int) int {
	time.Sleep(2 * time.Second)
	return v * v
}
func addQuote(v int) string {
	time.Sleep(1 * time.Second)
	return fmt.Sprintf("'%d'", v)
}
func addFoo(v string) string {
	time.Sleep(2 * time.Second)
	return fmt.Sprintf("%s - Foo", v)
}
