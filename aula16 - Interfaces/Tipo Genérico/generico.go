package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func escrever(interf ...interface{}) {
	fmt.Println(interf...)
}

func main() {
	generica("String")
	generica(1)
	generica(true)

	escrever("Olá", "Meu", "Nome", "é", "Diego,", "Tenho", 19, "anos")

	mapa := map[interface{}]interface{} {
		1: "String",
		float32(100): true,
		"String": 23,
		true: float32(12),
	}
	escrever(mapa)
}