package loadbalancer

import (
    "sync"
)

type LoadBalancerStrategy interface {
    Next() *Backend
}

type RoundRobinStrategy struct {
    backends []*Backend
    index    int
    mu       sync.Mutex
}

func NewRoundRobinStrategy(backends []*Backend) *RoundRobinStrategy {
    return &RoundRobinStrategy{backends: backends}
}

func (r *RoundRobinStrategy) Next() *Backend {
    r.mu.Lock()
    defer r.mu.Unlock()
    if len(r.backends) == 0 {
        return nil
    }
    backend := r.backends[r.index%len(r.backends)]
    r.index++
    return backend
}
