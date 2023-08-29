package main

import (
    "fmt"
	"math"
)

//Função como argumento
func CalculaVolumeEsfera(vol func(radius float64) float64) {
	fmt.Println("Volume da esfera é:", vol(3))
}


//Função retornando uma função
func CalcularVolumeEsfera() func(radius float64) float64 {
	resultado := func(radius float64) float64 {
		volume := 4 /3 * math.Pi * radius * radius * radius
		return volume
	} 

	return resultado
}

func main() {
	volumeEsfera := func(radius float64) float64 {
		resultado := 4 /3 * math.Pi * radius * radius * radius
		return resultado
	}

	CalculaVolumeEsfera(volumeEsfera)

	//sVol := CalcularVolumeEsfera()
	//fmt.Println("Volume da esfera é:", sVol(5))
}