package main

import (
	controller "FilterWorkerService/internal/controller"
	"FilterWorkerService/internal/dataAccess/mongodb"
	"runtime"
)

func main() {
	_ = make([]byte, 10<<30)
	runtime.MemProfileRate = 0
	mongodb.Init()
	controller.ListenLocations()
}
