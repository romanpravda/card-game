package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Excpected deck length of 52, but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Excpected first card of Ace of Spades, but got %s", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Excpected last card of King of Clubs, but got %s", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"
	os.Remove(filename)

	d := newDeck()
	d.saveToFile(filename)

	ld := newDeckFromFile(filename)

	if len(ld) != 52 {
		t.Errorf("Excpected deck length of 52, but got %d", len(ld))
	}

	if d.toString() != ld.toString() {
		t.Errorf("Excpected original deck and deck from file to be same, but they aren't")
	}

	os.Remove(filename)
}

func TestDeckShuffle(t *testing.T) {
	od := newDeck()
	d := newDeck()

	d.shuffle()
	ds := d.toString()

	if ds == od.toString() {
		t.Errorf("Excpected original deck and shuffled deck not to be same, but they are")
	}

	d.shuffle()
	if ds == d.toString() {
		t.Errorf("Excpected shuffled deck and double shuffled deck not to be same, but they are")
	}
}

func TestDeal(t *testing.T) {
	d := newDeck()

	fd, sd := deal(d, 2)

	if len(fd) != 2 {
		t.Errorf("Excpected first deck length of 2, but got %d", len(fd))
	}

	if len(sd) != 50 {
		t.Errorf("Excpected first deck length of 50, but got %d", len(sd))
	}
}
