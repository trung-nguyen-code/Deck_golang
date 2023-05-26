package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuit := []string{"Spades", "Diamonds", "Heart", "Club"}
	cardValue := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven"}

	for _, suit := range cardSuit {
		for _, val := range cardValue {
			cards = append(cards, val+"of"+suit)
		}
	}

	return cards

}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
