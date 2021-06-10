package main

import (
	"math/rand"
	"os"
	"strings"
	"time"
)

const deckSize = 52

type Deck struct {
	Cards []*Card
}

func NewDeck() *Deck {
	d := Deck{
		Cards: make([]*Card, 0, deckSize),
	}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, s := range cardSuits {
		for _, v := range cardValues {
			d.Cards = append(d.Cards, NewCard(v, s))
		}
	}

	return &d
}

func NewDeckFromFile(n string) (*Deck, error) {
	bs, err := os.ReadFile(n)
	if err != nil {
		return nil, err
	}

	d := &Deck{
		Cards: make([]*Card, 0),
	}

	cs := strings.Split(string(bs), "\n")
	for _, s := range cs {
		c, err := NewCardFromString(s)
		if err != nil {
			return nil, err
		}

		d.Cards = append(d.Cards, c)
	}

	return d, nil
}

func (d *Deck) Shuffle() {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	for i := range d.Cards {
		np := r.Intn(len(d.Cards) - 1)

		d.Cards[i], d.Cards[np] = d.Cards[np], d.Cards[i]
	}
}

func (d *Deck) MoveCardsFromDeckToHand(h *Hand, cc uint8) {
	var i uint8
	for i = 0; i < cc; i++ {
		h.AddCardToHand(d.Cards[i])
	}

	d.Cards = d.Cards[cc:]
}

func (d Deck) toString() string {
	cs := make([]string, 0, len(d.Cards))

	for _, c := range d.Cards {
		cs = append(cs, c.toString())
	}

	return strings.Join(cs, "\n")
}

func (d Deck) toFile(n string) error {
	return os.WriteFile(n, []byte(d.toString()), 0666)
}
