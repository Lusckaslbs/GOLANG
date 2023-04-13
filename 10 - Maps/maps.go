package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Lucas",
		"sobrenome": "Batista",
	}
	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeito": "Lucas",
			"ultimo":   "Silva",
		},
		"curso": {
			"nome":   "Analise de Sistemas",
			"campus": "Estacio",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Gêmeos",
	}
	fmt.Println(usuario2)

}
