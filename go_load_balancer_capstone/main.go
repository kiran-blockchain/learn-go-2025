package main

import (
    "log"
    "net/http"

    "go-load-balancer/loadbalancer"
)

func main() {
    lb, err := loadbalancer.NewLoadBalancer("config/backends.json", "roundrobin")
    if err != nil {
        log.Fatalf("Failed to initialize load balancer: %v", err)
    }

    log.Println("Load Balancer started at :8000")
    http.ListenAndServe(":8000", lb)
}
