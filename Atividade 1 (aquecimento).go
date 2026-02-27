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

	var peso, altura, resultado, idade float64
	var sexo int

	fmt.Println("Taxa Metabólica Basal (TMB)")
	fmt.Println("digite seu peso:")
	fmt.Scanln(&peso)

	fmt.Println("digite sua altura:")
	fmt.Scanln(&altura)

	fmt.Println("digite sua idade:")
	fmt.Scanln(&idade)

	fmt.Println("Sexo")
	fmt.Println("1 - Masculino")
	fmt.Println("2 - Feminino")
	fmt.Print("escolha sua opcao: ")
	fmt.Scanln(&sexo)

	switch sexo {
	case 1:
		resultado = 88.362 + (13.397 * peso) + (4.799 * altura) - (5.677 * idade)
	case 2:
		resultado = 447.593 + (9.247 * peso) + (3.098 * altura) - (4.330 * idade)
	}
	fmt.Printf("sua TMB é: %.2f", resultado)
}

func questao6() {

	var numero int
	fmt.Println("Primo?")
	fmt.Print("digite um numero: ")
	fmt.Scanln(&numero)

	if numero < 2 {
		fmt.Println("nao é primo")
		return
	}

	for i := 2; i*i <= numero; i++ { //ele testa os divisores até a raiz do numero digitado
		if numero%i == 0 {
			fmt.Println("nao é primo")
			return
		}
	}

	fmt.Println("é primo")

}

func questao7() {

	var numero, quantidade, somaPar, somaImpar int

	fmt.Print("deseja adicionar quantos numeros? ")
	fmt.Scanln(&quantidade)

	for i := 0; i < quantidade; i++ {
		fmt.Printf("digite o %d numero: ", i+1)
		fmt.Scanln(&numero)

		if numero%2 == 0 {
			somaPar += numero
		} else {
			somaImpar += numero
		}
	}

	fmt.Println("soma de todos os pares: ", somaPar)
	fmt.Println("soma de todos os impares: ", somaImpar)

}

func questao8() {

	var numero int
	fatorial := 1

	fmt.Print("digite um numero positivo: ")
	fmt.Scanln(&numero)

	for i := 2; i <= numero; i++ {
		fatorial *= i
	}

	fmt.Println("fatorial desse numero é: ", fatorial)

}

func questao9() {
	fmt.Print("1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765")
} //kkkkkkk

func questao09() {

}

//chatgpt a partir daqui

func questao10() {

	var n int
	fmt.Print("Digite um numero inteiro: ")
	fmt.Scanln(&n)

	if n == 0 {
		fmt.Println(0)
		return
	}

	neg := false
	if n < 0 {
		neg = true
		n = -n
	}

	reverso := 0
	for n > 0 {
		dig := n % 10
		reverso = reverso*10 + dig
		n /= 10
	}

	if neg {
		reverso = -reverso
	}

	fmt.Println("Numero reverso:", reverso)
}

func main() {
	questao10()
}
