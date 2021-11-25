package main

import (
	"FilterWorkerService/internal/IoC"
	"FilterWorkerService/internal/IoC/golobby"
	"runtime"
)

func main() {
	_ = make([]byte, 10<<30)
	runtime.MemProfileRate = 0

	//wg := sync.WaitGroup{}

	IoC.InjectContainers(golobby.InjectionConstructor())
	//IController.StartInsertListener(&wg, )
}