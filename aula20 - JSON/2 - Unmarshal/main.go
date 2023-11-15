package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome string `json:"nome"`
	Raca string `json:"raca"`
	Idade uint `json:"idade"`
}

func main () {
	cachorroJSON := `{"nome":"Rex","raca":"Dálmata","idade":3}`

	var c cachorro

	fmt.Println(c)

	if erro := json.Unmarshal([]byte(cachorroJSON), &c); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c)

	cachorroJSON2 := `{"nome":"Toby","raca":"Poodle"}`

	c2 := make(map[string]string)
	if erro := json.Unmarshal([]byte(cachorroJSON2), c2); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c2)
}