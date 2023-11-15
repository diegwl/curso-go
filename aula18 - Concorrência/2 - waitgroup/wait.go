package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2) // 2

	go func() {
		escrever("Olá Mundo!")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever("Programando em Go!")
		waitGroup.Done() // -1
	}()

	waitGroup.Wait() // 0
}

func escrever(texto string) {
	for i := 0; i < 5; i++{
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}