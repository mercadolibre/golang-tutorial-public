package main

import (
	"os"
	"fmt"
	"strconv"
)

type node struct{
	value int
	left *node
	right *node
}

func main(){
	//creating slice with cap = args count
	input:=make([]int,0,len(os.Args)-1)

	//append each int arg to the slice
	for i := 1; i < len(os.Args); i++ {
		v, err := strconv.Atoi(os.Args[i])
		if err == nil {
			input=append(input,v)
		}
	}
	//insert each value in tree
	var root *node
	for _,v:=range input{
		if root==nil{
			root=&node{v,nil,nil}
		}else{
			root.insert(v)
		}
	}

	fmt.Println("Original Slice:",input)
	fmt.Println("Sorted Slice:",root.asList())
}

func (n *node) insert(v int){
	if(v<=n.value){
		if(n.left==nil){
			n.left=&node{v,nil,nil}
		}else{
			n.left.insert(v)
		}
	}else{
		if(n.right==nil){
			n.right=&node{v,nil,nil}
		}else{
			n.right.insert(v)
		}
	}
}

func (n *node) asList() []int{
	var result []int
	if(n.left!=nil){
		result=append(result,n.left.asList()...)
	}
	result=append(result,n.value)
	if(n.right!=nil){
		result=append(result,n.right.asList()...)
	}
	return result
}