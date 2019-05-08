package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	// Crear un channel unbuffered

	// Usar wg para manejar la concurrencia

	// Poner dos tenistas a jugar

	// Empezar el partido

	// Esperar que termine el partido

}

// player() simula el un jugador de tenis
func player(name string, court chan int) {
	for {

		// Esperar que nos llegue la pelota

		// Elegir un numero aleatorio y ver si le pegamos a la pelota

		// Imprimir que el jugador le pego a la pelota y aumentar la cuenta

		// Devolverle la pelota al oponente
	}
}
