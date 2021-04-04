# Fibonacci

###### **Installation**:
    
###### **Features**:
 [x] You can count Fibonacci numbers between x and y
 [x] If the requested number occurs for the first time server will save it to redis storage 
 [x] The server support grpc and http protocol with addresses localhost:5300 (grpc) and localhost:5400 (http)

###### **Soon**:
 [x] Add worker pool for background calculating (if a request takes a long time server will return 
 timeout error and will count this numbers in background and later server will return numbers without timeout)
