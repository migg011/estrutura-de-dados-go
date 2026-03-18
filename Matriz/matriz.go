package main

import "fmt"

func main() {
	vetor := []int{6, 1, 6, 8, 4}

	for i := 0; i < len(vetor); i++ {
		if i%2 != 0 {
			fmt.Println(vetor[i])
		}
	}
}
