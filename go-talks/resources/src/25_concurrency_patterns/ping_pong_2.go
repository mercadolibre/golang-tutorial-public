package main

import (
"time"
"fmt")

func main() {
    var ball int
    table := make(chan int)
    for i:= 1; i<=6;i++ {
        go player(table, fmt.Sprintf("%s %v","Player ",i))
    }
    table <- ball
    time.Sleep(3 * time.Second)
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