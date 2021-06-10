package main

import "fmt"

const (
	aceHighValue = 11
	aceLowValue  = 1
)

type Hand struct {
	Cards []*Card
	Value uint8
	aces  uint8
}

func NewHand(s uint8) *Hand {
	return &Hand{
		Cards: make([]*Card, 0, s),
		Value: 0,
	}
}

func (h *Hand) AddCardToHand(c *Card) {
	h.Cards = append(h.Cards, c)

	if c.TextValue == "Ace" {
		h.aces++
	} else {
		h.Value += c.Value
	}
}

func (h Hand) Print() {
	for _, c := range h.Cards {
		fmt.Println(c.toString())
	}
}

func (h Hand) HasBlackjack() bool {
	hasAces := false
	has10Val := false

	for _, c := range h.Cards {
		if c.TextValue == "Ace" {
			hasAces = true
		} else if c.Value == 10 {
			has10Val = true
		}

		if hasAces && has10Val {
			return true
		}
	}

	return false
}

func (h Hand) HasBlackjackCards() bool {
	for _, c := range h.Cards {
		if c.TextValue == "Ace" || c.Value == 10 {
			return true
		}
	}

	return false
}

func (h Hand) Sum() uint8 {
	aces := h.aces
	s := h.Value

	if s+(aceHighValue*aces) > 21 {
		return s + (aceLowValue * aces)
	} else {
		return s + (aceHighValue * aces)
	}
}
