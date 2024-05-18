package main

import (
	"fmt"
)

// cria uma constante pra limitar a capacidade do canal
const capacity = 101

// canal para verificar se algo foi feito
var done = make(chan bool)

// canal que recebe as "caixas" do deposito, com uma capacidade total de 100 itens
var deposit = make(chan int, capacity)

// funcao principal
func main() {
	go producer()

	go consumer()

	<-done
	
	//Limpa o canal
	clean()
	
	//Preenche o canal
	fill()
	
	for len(deposit) > 0 {
		// Aguarda até que o canal seja totalmente consumido
	}

	// Fecha o canal
	close(deposit)
}

// funcao produtor
func producer() {
	for box := 0; box < 101; box++ {
		select {
		case deposit <- box:
			fmt.Println("Caixa enviada:", box)
		default:
			fmt.Println("Caixa não enviada, o canal cheio")
		}
	}
	done <- true
}

// funcao consumidor
func consumer() {

	for {
		select {
		case pack, ok := <-deposit:
			if !ok {
				fmt.Println("Canal está vazio, não há itens para retirar.")
				return
			}

			fmt.Println("Recebeu caixa:", pack) // caixa "Consumida"

		}
	}
}

//  Funcao para limpar o canal
func clean() {
    for len(deposit) > 0 {
        <-deposit
    }
    fmt.Println("Canal Esvaziado.")
}

// Funcao para preencher o canal e fecha-lo no final
func fill(){
    for i := 0; i < capacity; i++{
        deposit <- i 
    }
    fmt.Println("Canal preenchido.")
}