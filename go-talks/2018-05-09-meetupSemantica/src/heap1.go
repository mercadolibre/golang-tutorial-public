package main

//go:noinline
func shareUpTheStack() *int{
	i := 0
	return &i
}

func main() {
	shareUpTheStack()
}