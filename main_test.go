package main

//type testCase struct {
//	cards deck
//	sum   int
//}
//
//func TestCalcSum(t *testing.T) {
//	tcs := []testCase{
//		{cards: deck{"Ace of Spades", "King of Clubs"}, sum: 21},
//		{cards: deck{"Two of Clubs", "Ace of Spades"}, sum: 13},
//		{cards: deck{"Ace of Spades", "Ace of Clubs"}, sum: 2},
//		{cards: deck{"Seven of Hearts", "Six of Clubs"}, sum: 13},
//		{cards: deck{"King of Clubs", "Jack of Diamonds"}, sum: 20},
//		{cards: deck{"Ten of Spades", "Five of Hearts"}, sum: 15},
//		{cards: deck{"Ace of Diamonds", "Ace of Spades", "Queen of Hearts"}, sum: 12},
//		{cards: deck{"Ace of Diamonds"}, sum: 11},
//	}
//
//	for caseNum, tc := range tcs {
//		sum := calcSum(tc.cards)
//		if sum != tc.sum {
//			t.Error("case:", caseNum,
//				"\n\tcards:", tc.cards,
//				"\n\tresult:  ", sum,
//				"\n\texpected:", tc.sum)
//		}
//	}
//}
