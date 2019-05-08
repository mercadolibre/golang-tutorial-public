package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "Omitir tailing")

var sep = flag.String("s", " ", "Separador")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
