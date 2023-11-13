package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":      "Diego",
		"sobrenome": "Castan",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario["sobrenome"] = "Castan Lopes"

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Diego",
			"ultimo": "Lopes",
		}, 
		"curso": {
			"nome": "ADS",
			"faculdade": "IFSP",
			"campus": "Hortolândia",
		},
	}

	fmt.Println(usuario2)

	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "libra",
	}
	fmt.Println(usuario2)
}