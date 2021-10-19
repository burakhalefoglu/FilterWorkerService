package main

import (
	"FilterWorkerService/internal/dataAccess/mongodb"
	"runtime"
)
 
func main() {
	_ = make([]byte, 10<<30) 
	runtime.MemProfileRate = 0 
	mongodb.ConnectMongoDb()
}

