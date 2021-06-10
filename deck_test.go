package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := NewDeck()

	if len(d.Cards) != 52 {
		t.Errorf("Excpected deck length of 52, but got %d", len(d.Cards))
	}

	fc := d.Cards[0].toString()
	if fc != "Ace of Spades" {
		t.Errorf("Excpected first card of Ace of Spades, but got %s", fc)
	}

	lc := d.Cards[len(d.Cards)-1].toString()
	if lc != "King of Clubs" {
		t.Errorf("Excpected last card of King of Clubs, but got %s", lc)
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"
	err := os.Remove(filename)
	if err != nil && !os.IsNotExist(err) {
		t.Error(err)

		err = os.Remove(filename)
		if err != nil && !os.IsNotExist(err) {
			t.Error(err)
			return
		}

		return
	}

	d := NewDeck()
	err = d.toFile(filename)
	if err != nil {
		t.Error(err)

		err = os.Remove(filename)
		if err != nil && !os.IsNotExist(err) {
			t.Error(err)
			return
		}

		return
	}

	ld, err := NewDeckFromFile(filename)
	if err != nil {
		t.Error(err)

		err = os.Remove(filename)
		if err != nil && !os.IsNotExist(err) {
			t.Error(err)
			return
		}

		return
	}

	if len(ld.Cards) != 52 {
		t.Errorf("Excpected deck length of 52, but got %d", len(ld.Cards))
	}

	if d.toString() != ld.toString() {
		t.Errorf("Excpected original deck and deck from file to be same, but they aren't")
	}

	err = os.Remove(filename)
	if err != nil && !os.IsNotExist(err) {
		t.Error(err)
		return
	}
}

func TestDeckShuffle(t *testing.T) {
	od := NewDeck()
	d := NewDeck()

	d.Shuffle()
	ds := d.toString()

	if ds == od.toString() {
		t.Errorf("Excpected original deck and shuffled deck not to be same, but they are")
	}

	d.Shuffle()
	if ds == d.toString() {
		t.Errorf("Excpected shuffled deck and double shuffled deck not to be same, but they are")
	}
}

func TestDeckMoveCardsFromDeckToHand(t *testing.T) {
	d := NewDeck()
	d.Shuffle()

	fc := d.Cards[0].toString()

	h := NewHand(1)
	d.MoveCardsFromDeckToHand(h, 1)

	if len(h.Cards) != 1 {
		t.Errorf("Wrong hand size, expected 1 card, got - %d", len(h.Cards))
	}

	if len(d.Cards) != 51 {
		t.Errorf("Wrong deck size, expected 51 card, got - %d", len(d.Cards))
	}

	if h.Cards[0].toString() != fc {
		t.Errorf("Wrong card in hand, expected %s, got - %s", fc, h.Cards[0].toString())
	}
}
