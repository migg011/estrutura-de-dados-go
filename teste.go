package main

import (
	"fmt"
	"math"
)

func questao1() {
	var salario float64
	var bonus float64

	fmt.Print("digite seu salario: ")
	fmt.Scanln(&salario)
	if salario < 1000 {
		bonus = salario * 0.20
	} else {
		bonus = salario * 0.10
	}
	fmt.Println("salario final com bonus:", salario+bonus)
}

func questao2() {

	var opcao int

	fmt.Println("cardapio")
	fmt.Println("1. Pizza")
	fmt.Println("2. salada")
	fmt.Println("3. sushi")
	fmt.Print("Escolha uma opção: ")
	fmt.Scanln(&opcao)

	switch opcao {
	case 1:
		fmt.Println("voce escolheu pizza")
	case 2:
		fmt.Println("voce escolheu salada")
	case 3:
		fmt.Println("voce escolheu sushi")
	}
}

func questao3() {

	var num1 float64
	var num2 float64

	fmt.Println("digite o primeiro numero: ")
	fmt.Scanln(&num1)
	fmt.Println("digite o segundo numero: ")
	fmt.Scanln(&num2)

	var opcao int

	fmt.Println("Menu de Operações")
	fmt.Println("1. Soma")
	fmt.Println("2. Substração")
	fmt.Println("3. Multiplicação")
	fmt.Println("4. Divisão")
	fmt.Println("5. Resto Inteiro")

	fmt.Print("Escolha uma opção: ")
	fmt.Scanln(&opcao)

	var resultado float64

	switch opcao {
	case 1:
		resultado = num1 + num2
		fmt.Println("resultado da soma: ", resultado)
	case 2:
		resultado = num1 - num2
		fmt.Println("resultado da subtração: ", resultado)
	case 3:
		resultado = num1 * num2
		fmt.Println("resultado da multiplicação: ", resultado)
	case 4:
		resultado = num1 / num2
		fmt.Println("resultado da divisao: ", resultado)
	case 5:
		resultado = math.Mod(num1, num2)
		fmt.Println("resto da divisao: ", resultado)
	}
}

func questao4() {

	var sexo string
	var peso, altura double
	var idade int

	fmt.Println("digite seu peso:")
	fmt.Scanln(&peso)

	fmt.Println("digite sua altura:")
	fmt.Scanln(&altura)

	fmt.Println("digite sua idade:")
	fmt.Scanln(&idade)

	fmt.Println("Sexo")
	fmt.Println("1 - Masculino")
	fmt.Println("2 - Feminino")
	fmt.Println("escolha sua opcao: ")
	fmt.Scanln(&sexo)

	switch sexo {

	}

}

func main() {
	questao4()
}
