package main

import "fmt"

func main() {
	months := [...]string{1: "January", 2:"February",3:"March", 4: "April" ,5:"May", 6:"June",
		7:"July",8:"August",9:"September",10:"October",11:"November", 12: "December"}

	summer := months[6:9]

	Q3yQ4 := months[7:]
	Q3yQ4 = append(Q3yQ4, "New_month")
	Q3yQ4[1] = "new_August"
	fmt.Println(Q3yQ4)
	fmt.Println(summer)
}
