package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"Lucas", "Batista", 27, 175}
	fmt.Println(p1)

	e1 := estudante{p1, "Analise e Desenvolvimento de Sistemas", "Estacio"}
	fmt.Println(e1)
}
