package main
import (
"time"
"fmt")

func main() {
    var ball int
    table := make(chan int)
    go player(table, "Player 1")
    go player(table, "Player 2")
    table <- ball
    time.Sleep(1 * time.Second)
    <-table
}
func player(table chan int, player string) {
    for {
        ball := <-table
        ball++
	
        time.Sleep(100 * time.Millisecond)
        table <- ball
	   fmt.Println(player, ball)
    }
}