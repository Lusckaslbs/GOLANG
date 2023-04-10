package main

import (
	"errors"
	"fmt"
)

func main() {
	// A diferenca de cada int e float é a quantiade de bits que eles suportam

	// INT
	numero := 100000000000000
	fmt.Println(numero)

	// uint = usygned int
	var numero2 uint32 = 1000
	fmt.Println(numero2)

	// alias
	// INT32 = RUNE - INT8 = BYTE
	var numero3 rune = 1000
	fmt.Println(numero3)

	var numero4 byte = 123
	fmt.Println(numero4)

	// FIM INT

	// NÚMEROS REAIS
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123000000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)

	// FIM NÚMEROS REAIS

	// STRINGS
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)

	// FIM STRINGS

	var texto string
	fmt.Println(texto)

	// BOOLEAN
	var booleano1 bool = true
	fmt.Println(booleano1)

	var booleano2 bool
	fmt.Println(booleano2)

	// FIM BOOLEAN

	//ERRO
	//PARA UTILIZAR O ERRO É NECESSARIO IMPORTAR UM PACOTE INTERNO DO GO "ERRORS"
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

	// FIM ERRO

}
