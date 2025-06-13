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

var memoryLeak [][]byte // package-level slice to retain data

func slowFunction2() {
    // Allocate 1MB memory 10 times (10 MB per call)
    for i := 0; i < 10; i++ {
        b := make([]byte, 1024*1024) // 1MB
        memoryLeak = append(memoryLeak, b) // retain to simulate memory use
    }
    fmt.Println("Allocated 10MB")
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
go func() {
		for {
			slowFunction2()
			time.Sleep(1 * time.Second)
		}
	}()

	// Prevent main from exiting
	select {}
}
