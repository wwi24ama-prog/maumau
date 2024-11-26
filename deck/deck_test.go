package deck

import (
	"fmt"
	"math/rand"
	"maumau/card"
	"maumau/rank"
	"maumau/suit"
)

func ExampleNew32() {
	for _, c := range New32().cards {
		fmt.Println(c)
	}

	// Output:
	// ♣7
	// ♣8
	// ♣9
	// ♣10
	// ♣J
	// ♣Q
	// ♣K
	// ♣A
	// ♠7
	// ♠8
	// ♠9
	// ♠10
	// ♠J
	// ♠Q
	// ♠K
	// ♠A
	// ♥7
	// ♥8
	// ♥9
	// ♥10
	// ♥J
	// ♥Q
	// ♥K
	// ♥A
	// ♦7
	// ♦8
	// ♦9
	// ♦10
	// ♦J
	// ♦Q
	// ♦K
	// ♦A
}

// testDeck liefert einen sehr einfachen Kartenstapel mit nur 4 Karten und vorhersehbarer Zufallsquelle.
func testDeck() Deck {
	// Erstelle einen leeren Kartenstapel mit 4 Plätzen.
	cards := []card.Card{
		{Suit: suit.Clubs, Rank: rank.Ace},
		{Suit: suit.Spades, Rank: rank.Ace},
		{Suit: suit.Hearts, Rank: rank.Ace},
		{Suit: suit.Diamonds, Rank: rank.Ace},
	}
	return Deck{cards: cards, randSrc: rand.NewSource(0)}
}

func ExampleDeck_Shuffle() {
	d := testDeck()
	for _, c := range d.cards {
		fmt.Println(c)
	}
	d.Shuffle()
	for _, c := range d.cards {
		fmt.Println(c)
	}

	// Output:
	// ♣A
	// ♠A
	// ♥A
	// ♦A
	// ♥A
	// ♠A
	// ♣A
	// ♦A
}

func ExampleDeck_Len() {
	d := testDeck()
	fmt.Println(d.Len())

	// Output:
	// 4
}

func ExampleDeck_Draw() {
	d := testDeck()
	c := d.Draw()

	fmt.Println(c)
	fmt.Println(d.Len())
	for i := 0; i < 3; i++ {
		fmt.Println(d.Draw())
	}

	// Output:
	// ♣A
	// 3
	// ♠A
	// ♥A
	// ♦A
}

func ExampleDeck_ShowTop() {
	d := testDeck()
	d.ShowTop()

	// Output:
	// ┌───────┐
	// │A      │
	// │       │
	// │   ♣   │
	// │       │
	// │      A│
	// └───────┘
}
