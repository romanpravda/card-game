package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	d := newDeck()

	playBlackjack(d)
}

func playBlackjack(d deck) {
	ps := 0
	ds := 0

	d.shuffle()

	ph, ld := deal(d, 2)
	dh, ld := deal(ld, 1)

	fmt.Println("Your hand:")
	for _, card := range ph {
		fmt.Println(card)
	}

	pbj := (strings.Contains(ph[0], "Ace") && (strings.Contains(ph[1], "Ten") || strings.Contains(ph[1], "Jack") || strings.Contains(ph[1], "Queen") || strings.Contains(ph[1], "King"))) || ((strings.Contains(ph[0], "Ten") || strings.Contains(ph[0], "Jack") || strings.Contains(ph[0], "Queen") || strings.Contains(ph[0], "King")) && strings.Contains(ph[1], "Ace"))
	pdbj := strings.Contains(dh[0], "Ace") || strings.Contains(dh[0], "Ten") || strings.Contains(dh[0], "Jack") || strings.Contains(dh[0], "Queen") || strings.Contains(dh[0], "King")

	if pbj && !pdbj {
		fmt.Println("You won!")
		return
	}

	ps = calcSum(ph)

	if ps < 21 {
		fmt.Printf("Your sum: %d. Would you take another card? (1 - Yes, 2 - No): ", ps)
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			a := scanner.Text()

			if a == "2" {
				break
			}

			cth, nld := deal(ld, 1)
			ld = nld

			ph = append(ph, cth[0])
			ps = calcSum(ph)

			fmt.Println("Your hand:")
			for _, card := range ph {
				fmt.Println(card)
			}

			if ps < 21 {
				fmt.Printf("Your sum: %d. Would you take another card? (1 - Yes, 2 - No): ", ps)
			} else {
				fmt.Printf("Your sum: %d\n", ps)
				break
			}
		}
	} else {
		fmt.Printf("Your sum: %d\n", ps)
	}

	if ps > 21 {
		fmt.Println("You lost!")
		return
	}

	ctd, ld := deal(ld, 1)

	if (strings.Contains(dh[0], "Ace") && (strings.Contains(ctd[0], "Ten") || strings.Contains(ctd[0], "Jack") || strings.Contains(ctd[0], "Queen") || strings.Contains(ctd[0], "King"))) || ((strings.Contains(dh[0], "Ten") || strings.Contains(dh[0], "Jack") || strings.Contains(dh[0], "Queen") || strings.Contains(dh[0], "King")) && strings.Contains(ctd[0], "Ace")) {
		fmt.Println("You lost!")
		return
	}

	dh = append(dh, ctd[0])
	ds = calcSum(dh)

	for ds < 17 {
		ctd, ld = deal(ld, 1)
		dh = append(dh, ctd[0])
		ds = calcSum(dh)
	}

	fmt.Println("Dealer hand:")
	for _, card := range dh {
		fmt.Println(card)
	}

	fmt.Printf("Dealer sum: %d\n", ds)

	if ds > 21 {
		fmt.Println("You won!")
		return
	}

	if ps > ds {
		fmt.Println("You won!")
	} else if ps == ds {
		fmt.Println("You even!")
	} else {
		fmt.Println("You lost!")
	}
}

func calcSum(d deck) int {
	s := 0
	a := 0

	for _, card := range d {
		if strings.Contains(card, "Ace") {
			a++
		} else if strings.Contains(card, "Two") {
			s += 2
		} else if strings.Contains(card, "Three") {
			s += 3
		} else if strings.Contains(card, "Four") {
			s += 4
		} else if strings.Contains(card, "Five") {
			s += 5
		} else if strings.Contains(card, "Six") {
			s += 6
		} else if strings.Contains(card, "Seven") {
			s += 7
		} else if strings.Contains(card, "Eight") {
			s += 8
		} else if strings.Contains(card, "Nine") {
			s += 9
		} else {
			s += 10
		}
	}

	m := true
	i := 0

	for s < 21 && a > 0 && m {
		if s+11 > 21 {
			m = false
		} else {
			s += 11
			a--
			i++
		}
	}

	if a > 0 {
		s -= i * 11
		s += a + i
	}

	return s
}
