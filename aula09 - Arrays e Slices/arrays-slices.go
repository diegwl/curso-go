package main

import (
	"fmt"
	"reflect"
)

func main() {
	// arrays
	var array1 [5]int
	array1[0] = 1
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)
	
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	// slices
	slice1 := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice1)

	fmt.Println(reflect.TypeOf(array1))
	fmt.Println(reflect.TypeOf(slice1))

	slice1 = append(slice1, 18)
	fmt.Println(slice1)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posição alterada"
	fmt.Println(slice2)

	// arrays internos -> cria um array e retorna um slice que referencia ele
	// make([]tipo, tamanho, opcional(capacidade))
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice3 = append(slice3, 5)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice3 = append(slice3, 6)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}