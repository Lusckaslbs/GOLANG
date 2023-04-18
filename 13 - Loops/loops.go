package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")

	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementando i")
		time.Sleep(time.Second)

	}

	fmt.Println(i)

	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando j", j)
		time.Sleep(time.Second)

	}

	nomes := [3]string{"João", "Davi", "Lucas"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for indece, letra := range "PALAVRA" {
		fmt.Println(indece, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Lucas",
		"sobrenome": "Batista",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}
