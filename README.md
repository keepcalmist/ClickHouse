# Fibonacci

### **Installation**:
    * You should have working Redis on local machine
    1. You can run redis in docker-compose file *Run this command* docker-compose up
    To start server run 
    1. go mod download 
    2. go run ./cmd/server/main.go 
    To start client 
    1. go run ./cmd/client/main.go 
    
### **Features**:
 * You can count Fibonacci numbers between x and y 
 * If the requested number occurs for the first time server will save it to redis storage
 * The server support grpc and http protocol with addresses localhost:5300 (grpc) and localhost:5400 (http)

### **Soon**:
 * Add worker pool for background calculating (if a request takes a long time server will return 
 timeout error and will count this numbers in background and later server will return numbers without timeout)
