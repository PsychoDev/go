package main

type check string

func main() {
	//deck := newDeck()
	//deal, restDeck := deck.deal(3)
	//deal.print()
	//restDeck.print()
	//deck.writeToFile("deck_local_storage")
	d := readFromFile("deck_local_storage")
	// if deck == "sa" {
	// 	fmt.Println("WORKED")
	// }
	//fmt.Println(reflect.TypeOf(deck(d)))
	//d := deck{"aaaaa", "bbbbbb", "ccccccc", "dddd"}
	d.shuffel()
	d.print()
}
