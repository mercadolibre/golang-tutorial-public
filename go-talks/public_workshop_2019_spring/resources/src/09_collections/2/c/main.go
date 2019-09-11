package main

import "fmt"

func main() {
	a1 := [...]int{1,2,3}
	s1 := a1[0:2] // same as s1:=a1[:2]
	s1 = append(s1, 5)
	fmt.Println(s1)

	var s2 []int  // len(s2) == 0 , s2== nil
	s2 = append(s2, 5)
	fmt.Println(s2)

	s3 := make([]int, 2) //len(s3) ==2 , cap==2. Same as make([]int,2,2)
	s3 = append(s3, 5)
	fmt.Println(s3)

	fmt.Println(s1 == nil)
	//fmt.Println(s1 == s2) //Error! - slices are not comparable
}
