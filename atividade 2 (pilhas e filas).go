package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var lista []int

//Joao
func menuPilha() {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n======= MENU PILHA =======")
		fmt.Println("1) Push")
		fmt.Println("2) Pop")
		fmt.Println("3) Topo")
		fmt.Println("4) Base")
		fmt.Println("5) Mostrar pilha")
		fmt.Println("-1) Voltar")

		fmt.Print("Escolha: ")

		texto, _ := reader.ReadString('\n')
		texto = strings.TrimSpace(texto)
		opcao, _ := strconv.Atoi(texto)

		if opcao == -1 {
			logga("MENU | VOLTAR | origem=PILHA | OK")
			return
		}

		switch opcao {

		case 1:
			valor := rand.Intn(100)
			lista = append(lista, valor)

			fmt.Println("Pilha:", lista)
			logga(fmt.Sprintf("PILHA | PUSH | valor=%d | OK", valor))

		case 2:
			if len(lista) == 0 {
				fmt.Println("Pilha vazia")
				break
			}

			valor := lista[len(lista)-1]
			lista = lista[:len(lista)-1]

			fmt.Println("Pilha:", lista)
			logga(fmt.Sprintf("PILHA | POP | valor=%d | OK", valor))

		case 3:
			if len(lista) == 0 {
				fmt.Println("Pilha vazia")
				break
			}

			fmt.Println("Topo:", lista[len(lista)-1])
			logga("PILHA | TOPO | OK")

		case 4:
			if len(lista) == 0 {
				fmt.Println("Pilha vazia")
				break
			}

			fmt.Println("Base:", lista[0])
			logga("PILHA | BASE | OK")

		case 5:
			fmt.Println("Pilha completa:", lista)
			logga("PILHA | MOSTRAR | OK")

		default:
			fmt.Println("Opção inválida")
		}
	}
}

//Antônio
func subMenuLista() {
	for {
		var opcao int

		fmt.Println("1) Inserir elemento na lista (Número aleatório)\n2) Remover elemento da lista\n3) Mostrar a cabeça (primeiro elemento)\n4) Mostrar o rabo (último elemento)\n-1) Voltar ao menu principal")
		fmt.Scanln(&opcao)

		switch opcao {
		case 1:
			logga("usuario escolheu adicionar um numero aleatorio ")
			lista = append(lista, rand.Intn(100))
			fmt.Println("lista atualizada: ", lista, "\n")
		case 2:
			if len(lista) > 0 {
				logga("usuario escolheu remover um elemento da lista ")
				lista = lista[:len(lista)-1]
				fmt.Println("lista atualizada: ", lista, "\n")
			} else {
				fmt.Println("a lista esta vazia\n")
				continue
			}
		case 3:
			if len(lista) > 0 {
				logga("usuario escolheu mostrar o primeiro elemento da lista")
				primeiro := lista[0]
				fmt.Println("o primeiro elemento da lista é: ", primeiro, "\n")
			} else {
				fmt.Println("a lista esta vazia\n")
				continue
			}
		case 4:
			if len(lista) > 0 {
				logga("usuario escolheu mostrar o ultimo elemento da lista")
				ultimo := lista[len(lista)-1]
				fmt.Println("o ultimo elemento da lista é: ", ultimo, "\n")
			} else {
				fmt.Println("a lista esta vazia\n")
				continue
			}
		case -1:
			logga("usuario escolheu retornar ao menu principal")
			return
		default:
			logga("usuario escolheu uma opcao invalida em menu de lista")
			fmt.Println("opcao invalida")
		}
	}
}

func menuPrincipal() {

	var opcao int

	for opcao != -1 {

		fmt.Println("//1) Manipular Lista\n//2) Manipular Pilha\n//-1) Sair")
		fmt.Scanln(&opcao)

		switch opcao {

		case 1:
			logga("usuario escolheu manipular a lista")
			subMenuLista()
		case 2:
			logga("usuario escolheu manipular a pilhas")
			menuPilha()
		case -1:
			logga("usuario escolheu sair do sistema")
			fmt.Println("saindo do sistema...")
			break
		default:
			logga("usuario escolheu uma opcao invalida")
			fmt.Println("opcao invalida\n")
		}
	}
}

func logga(msg string) {

	agora := time.Now()
	timestamp := agora.Format("2006-01-02 15:04:05: ")

	file, _ := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	_, _ = file.WriteString(timestamp + msg + "\n")
}

func main() {
	menuPrincipal()
}
