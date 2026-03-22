package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardsSuits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	cardsValues := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	for _, suit := range cardsSuits {
		for _, value := range cardsValues {
			cards = append(cards, value+" of "+suit)
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

func (d deck) toString() string {
	sliceString := []string(d)
	myString := strings.Join(sliceString, ",")
	return myString
}

func (d deck) saveToFile(s string) error {
	myString := d.toString()
	myByteArray := []byte(myString)
	return os.WriteFile(s, myByteArray, 0666)
}

func newDeckFromFile(fileName string) deck {
	byteArray, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	stringItem := string(byteArray)
	stringSlice := strings.Split(stringItem, ",")
	deck := deck{}
	for _, item := range stringSlice {
		deck = append(deck, item)
	}
	return deck
}

func (d deck) shuffle() {
	for i := range d {
		newPosition := rand.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
