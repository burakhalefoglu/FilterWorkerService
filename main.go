package main

import (
	"FilterWorkerService/internal/IoC"
	"FilterWorkerService/internal/IoC/golobby"
	IController "FilterWorkerService/internal/controller"
	contorller "FilterWorkerService/internal/controller/kafka"
	"FilterWorkerService/pkg/helper"
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	defer helper.DeleteHealthFile()
	runtime.MemProfileRate = 0
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	for {
		startConsumer()
		time.Sleep(time.Second * 5)
	}

}

func startConsumer() {
	wg := sync.WaitGroup{}
	IoC.InjectContainers(golobby.InjectionConstructor())
	IController.StartInsertListener(&wg, contorller.KafkaControllerConstructor())
	wg.Wait()
}
