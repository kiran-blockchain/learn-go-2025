# Capstone Project: Load Balancer in Go

Core
Accept HTTP requests from clients (reverse proxy style)

Route requests to one of many backend servers

Read and forward full HTTP requests/responses

✅ Load Balancing Algorithms
Round Robin

Least Connections

Random

(Optional) IP Hash

✅ Health Checking
Periodic health checks (HTTP ping) to each backend

Mark servers as UP/DOWN

Route traffic only to healthy backends

✅ Logging
Log requests, target server, response time, and status code

Optional: log to a file or external log service

✅ Admin API (Optional)
Add/remove backend servers at runtime via REST API

Check backend status (health + connections)



## 📁 Structure
```
go-load-balancer/
├── main.go
├── loadbalancer/
│   ├── balancer.go
│   └── server_pool.go
├── config/
│   └── backends.json
```

##  How It Works
1. Load backends from `config/backends.json`
2. Initialize chosen strategy (default: round-robin)
3. Start HTTP server on `:8000`
4. Use `ReverseProxy` to forward requests

## 🧪 To Run
1. Start 3 sample servers on ports `8081`, `8082`, `8083`
2. Run the load balancer:
   ```
   go run main.go
   ```
3. Send requests to `localhost:8000`

## ✅ Future Improvements
- Health check goroutines
- Admin API to add/remove backends
- IP Hash or Least Connections strategy
- Logging & metrics (Prometheus)
