package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type deck []string

func (d deck) Print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

func (d deck) Deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func NewDeck() deck {
	cards := deck{}
	types := [4]string{"Diamonds", "Spades", "Hearts", "Clubs"}
	values := [13]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, cardType := range types {
		for _, value := range values {
			cards = append(cards, value+" of "+cardType)
		}
	}

	return cards
}

func (d deck) toString() string {
	return strings.Join([]string(d), "\n")
}

func (d deck) SaveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func ShuffleDeck(d deck) (deck, error) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d), func(i, j int) { d[i], d[j] = d[j], d[i] })
	return d, nil
}
