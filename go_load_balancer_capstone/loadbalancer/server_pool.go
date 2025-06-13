package loadbalancer

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http/httputil"
    "net/url"
    "os"
)

type Backend struct {
    URL          *url.URL
    Alive        bool
    ReverseProxy *httputil.ReverseProxy
}

type LoadBalancer struct {
    strategy LoadBalancerStrategy
}

func NewLoadBalancer(configPath string, strategyType string) (*LoadBalancer, error) {
    file, err := os.Open(configPath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    data, err := ioutil.ReadAll(file)
    if err != nil {
        return nil, err
    }

    var rawBackends []struct {
        URL string `json:"url"`
    }
    if err := json.Unmarshal(data, &rawBackends); err != nil {
        return nil, err
    }

    var backends []*Backend
    for _, b := range rawBackends {
        parsedURL, _ := url.Parse(b.URL)
        proxy := httputil.NewSingleHostReverseProxy(parsedURL)
        backends = append(backends, &Backend{
            URL:          parsedURL,
            Alive:        true,
            ReverseProxy: proxy,
        })
    }

    var strategy LoadBalancerStrategy
    switch strategyType {
    case "roundrobin":
        strategy = NewRoundRobinStrategy(backends)
    default:
        strategy = NewRoundRobinStrategy(backends)
    }

    return &LoadBalancer{strategy: strategy}, nil
}

func (lb *LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    backend := lb.strategy.Next()
    if backend == nil {
        http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
        return
    }
    backend.ReverseProxy.ServeHTTP(w, r)
}
