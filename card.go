package main

import (
	"errors"
	"strings"
)

type Card struct {
	Suit      string
	TextValue string
	Value     uint8
}

func NewCard(v, s string) *Card {
	var val uint8

	switch v {
	case "Ace":
	case "Two":
		val = 2
	case "Three":
		val = 3
	case "Four":
		val = 4
	case "Five":
		val = 5
	case "Six":
		val = 6
	case "Seven":
		val = 7
	case "Eight":
		val = 8
	case "Nine":
		val = 9
	default:
		val = 10
	}

	return &Card{
		Suit:      s,
		TextValue: v,
		Value:     val,
	}
}

func NewCardFromString(s string) (*Card, error) {
	if !strings.Contains(s, " of ") {
		return nil, errors.New(`Wrong string with card format. It should be "<suit> of <value>", got string ` + s + `.`)
	}

	ps := strings.Split(s, " of ")

	return NewCard(ps[0], ps[1]), nil
}

func (c Card) toString() string {
	return c.TextValue + " of " + c.Suit
}
