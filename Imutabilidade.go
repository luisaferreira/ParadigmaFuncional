package main

import "fmt"

// Definindo uma estrutura imutável
type Point struct {
    X, Y int
}

// Método para criar um novo Point com valores incrementados
func (p Point) Add(dx, dy int) Point {
    return Point{p.X + dx, p.Y + dy}
}

func main() {
    // Criando um Point
    originalPoint := Point{X: 1, Y: 2}
    fmt.Println("Point original:", originalPoint)

    // Criando um novo Point com valores incrementados
    newPoint := originalPoint.Add(3, 4)
    fmt.Println("Novo Point:", newPoint)

    // O originalPoint permanece inalterado
    fmt.Println("Point original após adição:", originalPoint)
}
