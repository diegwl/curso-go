package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome string `json:"nome"`
	Raca string `json:"raca"`
	Idade uint `json:"idade"`
}

func main() {
	c := cachorro{"Rex", "Dálmata", 3}
	fmt.Println(c)

	cachorroJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroJSON)
	fmt.Println(bytes.NewBuffer(cachorroJSON))

	c2 := map[string]string {
		"nome": "Toby",
		"raca": "Poodle",
	}

	cachorroJSON2, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroJSON2)
	fmt.Println(bytes.NewBuffer(cachorroJSON2))
}