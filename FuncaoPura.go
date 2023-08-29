package main

import "fmt"

// Função pura que calcula a soma de dois inteiros
func adicionar(a, b int) int {
    return a + b
}

func main() {
    x, y := 5, 3
    soma := adicionar(x, y)
    fmt.Printf("A soma de %d e %d é %d\n", x, y, soma)
}