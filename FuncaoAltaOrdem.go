package main

import (
    "fmt"
	"math"
)

//Função retornando uma função
func CalcularVolumeEsfera() func(radius float64) float64 {
	resultado := func(radius float64) float64 {
		volume := 4 /3 * math.Pi * radius * radius * radius
		return volume
	} 

	return resultado
}

func main() {
	sVol := CalcularVolumeEsfera()
	fmt.Println("Volume da esfera é:", sVol(5))
}