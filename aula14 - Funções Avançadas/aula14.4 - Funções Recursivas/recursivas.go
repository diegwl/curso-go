package main

import (
	"fmt"
	"time"
)

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao - 2) + fibonacci(posicao -1)
}

func main() {
	var posicao uint = 15

	for i := uint(1); i <= posicao; i++ {
		fmt.Println(fibonacci(i))
		time.Sleep(time.Second)
	}
}