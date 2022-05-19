package main

import (
	"fmt"
)

type MonthlyCosts struct {
	IUC            int
	inspecao       int
	disel          int
	seguro         int
	CarroElectrico int
	casa           int
	passe          int
	salario        int
	netSaves       int
}

func main() {
	calculateSvings()
}

func BasicCosts(carroDisel int, CarroElectrico int, casa int, passe int) MonthlyCosts {
	bc := MonthlyCosts{}
	if carroDisel == 1 { // https://www.e-konomista.pt/quanto-custa-ter-um-carro/
		bc.IUC = 18
		bc.inspecao = 3
		bc.disel = 200 // assumindo 6L /100km
		bc.seguro = 20
	}

	if CarroElectrico == 1 {
		bc.CarroElectrico = 16 // 1.5eur eletecidade /200km
		bc.inspecao = 3
		bc.seguro = 20
	}
	if casa == 1 {
		bc.casa = 500
	}
	if passe == 1 {
		bc.passe = 40
	}
	bc.salario = 1600
	bc.netSaves = bc.salario - (bc.IUC + bc.inspecao + bc.disel + bc.seguro + bc.CarroElectrico + bc.casa + bc.passe)
	return bc
}

func calculateSvings() {
	bc := BasicCosts(1, 0, 1, 0)
	fmt.Println("Svings como carro disel e casa = ", bc.netSaves)

	bc2 := BasicCosts(0, 1, 1, 0)
	fmt.Println("Svings como carro electrico e casa = ", bc2.netSaves)

	bc3 := BasicCosts(0, 0, 1, 1)
	fmt.Println("Svings como passe e casa = ", bc3.netSaves)

	bc4 := BasicCosts(0, 1, 1, 0)
	fmt.Println("Svings como carro electrico e casa = ", bc4.netSaves)

	// bc5 := BasicCosts(0, 1, 0, 0)
	// fmt.Println("Svings como carro electrico e sem casa = ", bc5.netSaves)
}
