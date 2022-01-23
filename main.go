package main

import (
	"FilterWorkerService/internal/IoC"
	"FilterWorkerService/internal/IoC/golobby"
	IController "FilterWorkerService/internal/controller"
	contorller "FilterWorkerService/internal/controller/kafka"
	"github.com/joho/godotenv"
	"log"
	"runtime"
	"sync"
)

func main() {
	runtime.MemProfileRate = 0

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	wg := sync.WaitGroup{}

	IoC.InjectContainers(golobby.InjectionConstructor())
	IController.StartInsertListener(&wg, contorller.KafkaControllerConstructor())
	wg.Wait()
}