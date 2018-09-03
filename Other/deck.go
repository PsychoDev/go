package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	var d deck
	ranks := []string{"Ace", "King", "Queen", "Joker"}
	groups := []string{"Spades", "Clubs", "Dimonds", "Hearts"}
	for _, group := range groups {
		for _, rank := range ranks {
			d = append(d, rank+" of "+group)
		}
	}
	return d
}

func (d deck) deal(size int) (deck, deck) {
	return d[:size], d[size:]
}

func (d deck) print() {
	for i, deck := range d {
		fmt.Printf("%d %s\n", i, deck)
	}
}

func (d deck) tostring() string {
	ss := []string(d)
	return strings.Join(ss, ",")
}

func (d deck) writeToFile(filename string) {
	s := d.tostring()
	bs := []byte(s)
	ioutil.WriteFile(filename, bs, 777)
}

func readFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	s := string(bs)
	ss := strings.Split(s, ",")
	return ss
}

func (d deck) shuffel() {
	for i := range d {
		s2 := rand.NewSource(time.Now().UnixNano())
		r2 := rand.New(s2)
		r := r2.Intn(len(d))
		d[i], d[r] = d[r], d[i]
	}
}
