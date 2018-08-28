package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email  string
	mobile int
}

func main() {
	cis := contactInfo{
		"bhupi@gmail.com", 9899778404,
	}
	pst := person{
		firstName:   "Bhupinder",
		lastName:    "Bisht",
		contactInfo: cis,
	}
	pst.update("JIM")
	fmt.Println(pst)

}

func (p *person) update(firstName string) {
	fmt.Printf("%p \n", p)
	(*p).contactInfo.email = firstName
}
