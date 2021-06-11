package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	d := NewDeck()

	playBlackJack(d)
}

func printHand(h *Hand) {
	for _, c := range h.Cards {
		println(c.toString())
	}
}

func playBlackJack(d *Deck) {
	d.Shuffle()

	ph := NewHand(2)
	dh := NewHand(2)

	d.MoveCardsFromDeckToHand(ph, 2)
	d.MoveCardsFromDeckToHand(dh, 1)

	fmt.Println("Your hand:")
	printHand(ph)

	if ph.HasBlackjack() && !dh.HasBlackjackCards() {
		fmt.Println("You won!")
		return
	}

	fmt.Printf("Your sum: %d", ph.Sum())

	if ph.Sum() < 21 {
		fmt.Printf(". Would you take another card? (1 - Yes, 2 - No): ")
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			a := scanner.Text()

			if a == "2" {
				break
			} else if a != "1" {
				fmt.Println("Wrong enter. Would you take another card? (1 - Yes, 2 - No): ")
				continue
			}

			d.MoveCardsFromDeckToHand(ph, 1)

			fmt.Println("Your hand:")
			printHand(ph)

			if ph.Sum() < 21 {
				fmt.Printf("Your sum: %d. Would you take another card? (1 - Yes, 2 - No): ", ph.Sum())
			} else {
				fmt.Printf("Your sum: %d\n", ph.Sum())
				break
			}
		}
	} else {
		fmt.Println()
	}

	if ph.Sum() > 21 {
		fmt.Println("You lost!")
		return
	}

	d.MoveCardsFromDeckToHand(dh, 1)
	if dh.HasBlackjack() {
		fmt.Println("You lost!")
		return
	}

	for dh.Sum() < 17 {
		d.MoveCardsFromDeckToHand(dh, 1)
	}

	fmt.Println("Dealer hand:")
	printHand(dh)

	fmt.Printf("Dealer sum: %d\n", dh.Sum())

	if dh.Sum() > 21 {
		fmt.Println("You won!")
		return
	}

	if ph.Sum() > dh.Sum() {
		fmt.Println("You won!")
	} else if ph.Sum() == dh.Sum() {
		fmt.Println("You even!")
	} else {
		fmt.Println("You lost!")
	}
}
