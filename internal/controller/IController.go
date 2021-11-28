package IController

import "sync"

type IController interface {
	StartListen(waitGroup *sync.WaitGroup)
}

func StartInsertListener(waitGroup *sync.WaitGroup, listener IController){
	listener.StartListen(waitGroup)
}

