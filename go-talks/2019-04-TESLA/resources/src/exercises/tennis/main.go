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
	court := make(chan int)

	// Usar wg para manejar la concurrencia
	var wg sync.WaitGroup
	wg.Add(2)

	// Poner dos tenistas a jugar
	go func() {
		player("Del Potro ", court)
		wg.Done()
	}()

	go func() {
		player("Djokovic", court)
		wg.Done()
	}()

	// Empezar el partido
	court <- 1

	// Esperar que termine el partido
	wg.Wait()
}

func player(name string, court chan int) {
	for {

		// Esperar que nos llegue la pelota
		ball, wd := <-court
		if !wd {

			fmt.Printf("Player %s Won\n", name)
			return
		}

		// Elegir un numero aleatorio y ver si le pegamos a la pelota
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)

			close(court)
			return
		}

		// Imprimir que el jugador le pego a la pelota y aumentar la cuenta
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		// Devolverle la pelota al oponente
		court <- ball
	}
}
