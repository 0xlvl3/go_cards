package main

func main(){
	// Creating a slice.
	// Slice in go is an array that can append and pop
	cards := deck{"Hello", "hi", newCard()}
	cards = append(cards, "test")

	cards.print()
	// For loop.
	// Iterate.
	/*
	for index, variable := range slice{		# Slice here is a go array that can be append and pop.
		fmt.Println(index, variable)		# Will index and a card on each line as it iterates.
    }
	*/
/*
	for i, card := range cards{
		fmt.Println(i, card)
	}

*/
	
}

func newCard() string {
	return "Five of Diamonds"
}
