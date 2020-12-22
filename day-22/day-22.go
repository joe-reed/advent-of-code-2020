package main

import (
	"strings"
	. "utils"
)

func puzzle1(input string) int {
	deck1, deck2 := getDecksFromInput(input)

	for deck1.length() != 0 && deck2.length() != 0 {
		deck1, deck2 = playRound(deck1, deck2)
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

func playRound(d1, d2 Deck) (r1, r2 Deck) {
	r1 = d1.clone()
	r2 = d2.clone()
	card1 := r1.popCard()
	card2 := r2.popCard()
	if card1 > card2 {
		r1.appendCard(card1)
		r1.appendCard(card2)
	} else {
		r2.appendCard(card2)
		r2.appendCard(card1)
	}
	return
}

func getWinner(d1, d2 Deck) Deck {
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

func (d Deck) clone() (result Deck) {
	cards := make([]int, len(d.cards))
	copy(cards, d.cards)
	result = d
	result.cards = cards
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
