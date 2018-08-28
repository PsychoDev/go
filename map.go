package main

import (
	"fmt"
)

type men map[string]string
type women map[string]string

type person interface {
	gender()
	getIdentity(string)
}

func main() {
	m := make(men)
	m["gender"] = "MALE"
	w := make(women)
	w["gender"] = "FEMALE"
	m.gender()
	w.gender()
}

func (m person) gender() {
	m.getIdentity(m["gender"])
}

func (w person) gender() {
	w.getIdentity(w["gender"])
}

func (p person) getIdentity(g string) {
	fmt.Println("I AM " + m["gender"])
}
