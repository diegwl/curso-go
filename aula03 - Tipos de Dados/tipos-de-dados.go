package main

import (
	"errors"
	"fmt"
)

func main() {
	// int, int8, int16, int32, int64
	var numero int16 = 100
	var numero2 int = 10000000
	numero3 := -100000
	fmt.Println(numero, numero2, numero3)

	// uint -> int sem sinal (uint8, uint16, uint32, uint64)
	var numero4 uint = 100000
	fmt.Println(numero4)

	// alias int32 = RUNE
	var numero5 rune = 123456
	fmt.Println(numero5)

	// alias byte = UINT8
	var numero6 byte = 123
	fmt.Println(numero6)

	// float32, float64
	var numeroReal1 float32 = 123.45
	var numeroReal2 float64 = 12345.6789
	numeroReal3 := 1234.65
	fmt.Println(numeroReal1, numeroReal2, numeroReal3)

	// string
	var str string = "Texto"
	str2 := "Texto 2"
	fmt.Println(str, str2)

	// char -> numero da tabela ASCII - tipo int (não tipo existe char)
	char := 'B'
	fmt.Println(char)

	// valor inicial -> string = "" - int/float = 0
	var texto string
	var numero7 int
	fmt.Println(texto, numero7)

	// bool - valor inicial é false
	var boolean1 bool = false
	boolean2 := true
	fmt.Println(boolean1, boolean2)

	// error -> <nil> (null)
	var erro1 error
	var erro2 error = errors.New("Erro Interno")
	fmt.Println(erro1, erro2)
}