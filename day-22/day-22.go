package main

import (
	"strings"
	. "utils"
)

func puzzle1(input string) int {
	deck1, deck2 := getDecksFromInput(input)
	var card1, card2 int
	for deck1.length() != 0 && deck2.length() != 0 {
		deck1, deck2, card1, card2 = drawCards(deck1, deck2)
		deck1, deck2 = playRound(deck1, deck2, card1, card2)
	}
	return getWinner(deck1, deck2).getScore()
}

func puzzle2(input string) int {
	deck1, deck2 := getDecksFromInput(input)
	return playRecursiveGame(deck1, deck2).getScore()
}

func getDecksFromInput(input string) (Deck, Deck) {
	decks := []Deck{{id: 1, cards: []int{}}, {id: 2, cards: []int{}}}
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

func playRecursiveGame(d1, d2 Deck) Deck {
	p1History := []Deck{}
	r1 := d1.clone()
	r2 := d2.clone()
	for r1.length() != 0 && r2.length() != 0 {
		r1, r2, p1History = playRecursiveRound(r1, r2, p1History)
		if len(p1History) == 0 {
			return r1
		}
	}
	return getWinner(r1, r2)
}

func playRecursiveRound(d1, d2 Deck, p1History []Deck) (r1, r2 Deck, updatedHistory []Deck) {
	if contains(p1History, d1) {
		return d1, d2, []Deck{}
	}
	updatedHistory = append(p1History, d1)

	var c1, c2 int
	r1, r2, c1, c2 = drawCards(d1, d2)
	if r1.length() >= c1 && r2.length() >= c2 {
		winner := playRecursiveGame(
			Deck{id: 1, cards: r1.cards[0:c1]},
			Deck{id: 2, cards: r2.cards[0:c2]},
		)
		r1, r2 = getOutcome(winner, r1, r2, c1, c2)
		return
	}
	r1, r2 = playRound(r1, r2, c1, c2)
	return
}

func drawCards(d1, d2 Deck) (r1, r2 Deck, c1, c2 int) {
	r1 = d1.clone()
	r2 = d2.clone()
	c1 = r1.popCard()
	c2 = r2.popCard()
	return
}

func playRound(d1, d2 Deck, c1, c2 int) (Deck, Deck) {
	winner := d1
	if c2 > c1 {
		winner = d2
	}
	return getOutcome(winner, d1, d2, c1, c2)
}

func getOutcome(winner, d1, d2 Deck, c1, c2 int) (r1, r2 Deck) {
	r1 = d1.clone()
	r2 = d2.clone()
	if winner.id == 1 {
		r1.appendCard(c1)
		r1.appendCard(c2)
	} else {
		r2.appendCard(c2)
		r2.appendCard(c1)
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
	id    int
	cards []int
}

func contains(decks []Deck, d Deck) bool {
	for _, deck := range decks {
		if deck.equals(d) {
			return true
		}
	}
	return false
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

func (d Deck) equals(deck Deck) bool {
	if d.length() != deck.length() {
		return false
	}
	for i, card := range d.cards {
		if deck.cards[i] != card {
			return false
		}
	}
	return true
}

func (d Deck) getScore() (result int) {
	for i, card := range d.cards {
		result += card * (d.length() - i)
	}
	return
}
