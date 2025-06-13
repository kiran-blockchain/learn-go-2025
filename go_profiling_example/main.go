package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func slowFunction() {
	sum := 0
	for i := 0; i < 1e7; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)
}

func main() {
	// Start the pprof server
	go func() {
		fmt.Println("pprof listening at http://localhost:6060/debug/pprof/")
		http.ListenAndServe("localhost:6060", nil)
	}()

	// Simulate workload in a goroutine
	go func() {
		for {
			slowFunction()
			time.Sleep(1 * time.Second)
		}
	}()

	// Prevent main from exiting
	select {}
}
