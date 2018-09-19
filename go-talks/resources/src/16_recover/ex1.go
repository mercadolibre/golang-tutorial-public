package main

func main() {
	println("After this, panic will start")
	panic("Panic occoured!")
	println("This line will not appear")
}
