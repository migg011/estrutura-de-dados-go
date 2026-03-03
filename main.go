package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.OpenFile("app.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		fmt.Println("Erro ao abrir log:", err)
		return
	}
	defer f.Close()

	_, err = f.WriteString("nova linha adicionada")
	if err != nil {
		fmt.Println("Erro ao escrever log:", err)
		return
	}
}
