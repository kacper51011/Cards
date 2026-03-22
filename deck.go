package main

type deck []string

func newDeck() deck {
	cards := deck{}
	cardsSuits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	cardsValues := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	for _, suit := range cardsSuits {
		for _, value := range cardsValues {
			cards = append(cards, suit+" of "+value)
		}
	}
	return cards
}
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
func (d deck) print() {
	for i, c := range d {
		println("on number ", i, "there is ", c)
	}
}
