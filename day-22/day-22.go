package main

import (
	"strings"
	. "utils"
)

func puzzle1(input string) int {
	deck1, deck2 := getDecksFromInput(input)

	for deck1.length() != 0 && deck2.length() != 0 {
		card1 := deck1.popCard()
		card2 := deck2.popCard()
		if card1 > card2 {
			deck1.appendCard(card1)
			deck1.appendCard(card2)
		} else {
			deck2.appendCard(card2)
			deck2.appendCard(card1)
		}
	}

	return getWinner(deck1, deck2).getScore()
}

func getDecksFromInput(input string) (Deck, Deck) {
	decks := []Deck{{cards: []int{}}, {cards: []int{}}}
	for i, s := range strings.Split(input, "\n\n") {
		for j, line := range strings.Split(s, "\n") {
			if j == 0 {
				continue
			}
			card := ConvertToInt(string(line))
			decks[i].appendCard(card)
		}
	}
	return decks[0], decks[1]
}

func getWinner(d1 Deck, d2 Deck) Deck {
	if d1.length() == 0 {
		return d2
	}
	return d1
}

type Deck struct {
	cards []int
}

func (d *Deck) appendCard(card int) {
	d.cards = append(d.cards, card)
}

func (d *Deck) popCard() (card int) {
	card = d.cards[0]
	d.cards = d.cards[1:]
	return
}

func (d Deck) length() int {
	return len(d.cards)
}

func (d Deck) getScore() (result int) {
	for i, card := range d.cards {
		result += card * (d.length() - i)
	}
	return
}
