package model

import "fmt"

type Student struct {
	Name   string
	name   string
	age    int
	Age    int
	height float64

	do  Dog
	Dog Dog
}

type Dog struct {
	Name string
	name string
}

func init() {
	fmt.Println("init================")
}
