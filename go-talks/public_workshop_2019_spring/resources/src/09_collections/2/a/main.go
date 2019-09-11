package main

import "fmt"

func main() {
	months := [...]string{1: "January", 2:"February",3:"March", 4: "April" ,5:"May", 6:"June",
		           7:"July",8:"August",9:"September",10:"October",11:"November", 12: "December"}

	Q2 := months[4:7]
	fmt.Println(Q2)
	summer := months[6:9]
	fmt.Println(summer)

	Q2[2] = "New_month"
	fmt.Println(Q2)
	fmt.Println(summer)
	fmt.Println(months[6])

}