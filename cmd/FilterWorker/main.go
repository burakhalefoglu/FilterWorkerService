package main

import (
	"runtime"
)

func main() {
	_ = make([]byte, 10<<30)
	runtime.MemProfileRate = 0
}
