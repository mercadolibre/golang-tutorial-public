package main

 func main() {

    // Declare variable of type int with a value of 10.
    // Code 1 START OMIT
    count := 10
    println("count:\t Valor: ", count, "\tDireccion: ", &count)
    // Code 1 END OMIT

    // Pass the "value of" the count.
    // Code 2 START OMIT
    increment(count)
    // Code 2 END OMIT
    
    // Code 4 START OMIT
    println("count:\t Valor: ", count, "\tDireccion: ", &count)
    // Code 4 END OMIT
 }

 //go:noinline
 func increment(inc int) {

    // Increment the "value of" inc.
    // Code 3 START OMIT
    inc++
    println("inc:\t Valor: ", inc, "\tDireccion: ", &inc)
    // Code 3 END OMIT
}