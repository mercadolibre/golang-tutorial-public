package main

 func main() {

    // Declare variable of type int with a value of 10.
    count := 10

    // Display the "value of" and "address of" count.
    println("count:\t Valor: ", count, "\tDireccion: ", &count)

    // Pass the "address of" count.
    // Code 5 START OMIT
    increment(&count)
    // Code 5 END OMIT

    println("count:\t Valor: ", count, "\tDireccion: ", &count)
 }

 //go:noinline
 func increment(inc *int) {
    // Increment the "value of" count that the "pointer points to". (dereferencing)
    // Code 6 START OMIT
    *inc++
    println("inc:\t Apunta: ", *inc, "\tValor:\t ", inc, "\tDireccion: ", &inc)
    // Code 6 END OMIT
}