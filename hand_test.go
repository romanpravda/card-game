package main

import (
	"reflect"
	"testing"
)

func TestAddCardToHand(t *testing.T) {
	c1 := &Card{
		Suit:      "Ace",
		TextValue: "Spades",
		Value:     1,
	}
	c2 := &Card{
		Suit:      "Ten",
		TextValue: "Diamonds",
		Value:     10,
	}

	h := NewHand(2)

	h.AddCardToHand(c1)
	h.AddCardToHand(c2)

	hasC1 := false
	hasC2 := false

	for _, c := range h.Cards {
		if reflect.DeepEqual(c, c1) {
			hasC1 = true
		} else if reflect.DeepEqual(c, c2) {
			hasC2 = true
		}
	}

	if !hasC1 {
		t.Errorf("there is no first card in the hand")
	}

	if !hasC2 {
		t.Errorf("there is no second card in the hand")
	}
}

func TestHasBlackjack(t *testing.T) {
	c1 := &Card{
		Suit:      "Spades",
		TextValue: "Ace",
		Value:     1,
	}
	c2 := &Card{
		Suit:      "Diamonds",
		TextValue: "Ten",
		Value:     10,
	}
	c3 := &Card{
		Suit:      "Hearts",
		TextValue: "Nine",
		Value:     9,
	}

	hbj := NewHand(2)
	hnbj := NewHand(2)

	hbj.AddCardToHand(c1)
	hbj.AddCardToHand(c2)

	hnbj.AddCardToHand(c1)
	hnbj.AddCardToHand(c3)

	if !hbj.HasBlackjack() {
		t.Errorf("there is no blackjack with blackjack cards")
	}

	if hnbj.HasBlackjack() {
		t.Errorf("there is blackjack with not blackjack cards")
	}
}

func TestHasBlackjackCards(t *testing.T) {
	c1 := &Card{
		Suit:      "Spades",
		TextValue: "Ace",
		Value:     1,
	}
	c2 := &Card{
		Suit:      "Clubs",
		TextValue: "Two",
		Value:     2,
	}
	c3 := &Card{
		Suit:      "Hearts",
		TextValue: "Nine",
		Value:     9,
	}

	hbj := NewHand(2)
	hnbj := NewHand(2)

	hbj.AddCardToHand(c1)
	hbj.AddCardToHand(c2)

	hnbj.AddCardToHand(c2)
	hnbj.AddCardToHand(c3)

	if !hbj.HasBlackjackCards() {
		t.Errorf("there is no blackjack card when excepts it")
	}

	if hnbj.HasBlackjackCards() {
		t.Errorf("there is blackjack cards when not excepts it")
	}

}

func TestValueOfHandTwoAces(t *testing.T) {
	exSum := uint8(2)

	c1 := &Card{
		Suit:      "Spades",
		TextValue: "Ace",
		Value:     1,
	}
	c2 := &Card{
		Suit:      "Clubs",
		TextValue: "Ace",
		Value:     1,
	}

	h := NewHand(2)

	h.AddCardToHand(c1)
	h.AddCardToHand(c2)

	if h.Sum() != exSum {
		t.Errorf("wrong sum\nexpected - %d\ngot - %d", exSum, h.Sum())
	}
}

func TestValueOfHandWithAceAndTen(t *testing.T) {
	exSum := uint8(21)

	c1 := &Card{
		Suit:      "Spades",
		TextValue: "Ace",
		Value:     1,
	}
	c2 := &Card{
		Suit:      "Clubs",
		TextValue: "Ten",
		Value:     10,
	}

	h := NewHand(2)

	h.AddCardToHand(c1)
	h.AddCardToHand(c2)

	if h.Sum() != exSum {
		t.Errorf("wrong sum\nexpected - %d\ngot - %d", exSum, h.Sum())
	}
}

func TestValueOfAceTwoAndThree(t *testing.T) {
	exSum := uint8(16)

	c1 := &Card{
		Suit:      "Spades",
		TextValue: "Ace",
		Value:     1,
	}
	c2 := &Card{
		Suit:      "Clubs",
		TextValue: "Two",
		Value:     2,
	}
	c3 := &Card{
		Suit:      "Diamonds",
		TextValue: "Three",
		Value:     3,
	}

	h := NewHand(3)

	h.AddCardToHand(c1)
	h.AddCardToHand(c2)
	h.AddCardToHand(c3)

	if h.Sum() != exSum {
		t.Errorf("wrong sum\nexpected - %d\ngot - %d", exSum, h.Sum())
	}
}

func TestValueOfAceSevenAndEight(t *testing.T) {
	exSum := uint8(16)

	c1 := &Card{
		Suit:      "Spades",
		TextValue: "Ace",
		Value:     1,
	}
	c2 := &Card{
		Suit:      "Clubs",
		TextValue: "Seven",
		Value:     7,
	}
	c3 := &Card{
		Suit:      "Diamonds",
		TextValue: "Eight",
		Value:     8,
	}

	h := NewHand(3)

	h.AddCardToHand(c1)
	h.AddCardToHand(c2)
	h.AddCardToHand(c3)

	if h.Sum() != exSum {
		t.Errorf("wrong sum\nexpected - %d\ngot - %d", exSum, h.Sum())
	}
}

func TestPrintHand(t *testing.T) {
	c1 := &Card{
		Suit:      "Spades",
		TextValue: "Ace",
		Value:     1,
	}
	c2 := &Card{
		Suit:      "Clubs",
		TextValue: "Ace",
		Value:     1,
	}

	h := NewHand(2)

	h.AddCardToHand(c1)
	h.AddCardToHand(c2)

	cns := h.GetCardsNames()

	if cns[0] != c1.toString() {
		t.Errorf("Expected card %s, got %s.", c1.toString(), cns[0])
	}

	if cns[1] != c2.toString() {
		t.Errorf("Expected card %s, got %s.", c2.toString(), cns[1])
	}
}
