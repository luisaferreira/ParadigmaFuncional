package main

import "fmt"

// Função recursiva para calcular o fatorial
func fatorial(n int) int {
    if n <= 0 {
        return 1
    }
    return n * fatorial(n-1)
}

func main() {
    num := 5
    resultado := fatorial(num)
    fmt.Printf("O fatorial de %d é %d\n", num, resultado)
}