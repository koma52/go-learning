package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	pista := person{firstName: "Pista", lastName: "Kiss"}
	fmt.Println(pista)
}
