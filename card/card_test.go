package card

import (
	"fmt"
	"maumau/rank"
	"maumau/suit"
)

func ExampleCard_String() {
	fmt.Println(Card{Suit: suit.Clubs, Rank: rank.Ace})
	fmt.Println(Card{Suit: suit.Hearts, Rank: rank.Two})

	// Output:
	// ♣A
	// ♥2
}

func ExampleCard_Image() {
	c1 := Card{Suit: suit.Clubs, Rank: rank.Ace}
	for _, line := range c1.Image() {
		fmt.Println(line)
	}

	c2 := Card{Suit: suit.Hearts, Rank: rank.Ten}

	for _, line := range c2.Image() {
		fmt.Println(line)
	}

	// Output:
	// ┌───────┐
	// │A      │
	// │       │
	// │   ♣   │
	// │       │
	// │      A│
	// └───────┘
	// ┌───────┐
	// │10     │
	// │       │
	// │   ♥   │
	// │       │
	// │     10│
	// └───────┘
}

func ExampleCard_Matches() {
	c1 := Card{Suit: suit.Clubs, Rank: rank.Ace}
	c2 := Card{Suit: suit.Clubs, Rank: rank.Ten}
	c3 := Card{Suit: suit.Hearts, Rank: rank.Ten}

	fmt.Println(c1.Matches(c2))
	fmt.Println(c1.Matches(c3))
	fmt.Println(c2.Matches(c3))
	fmt.Println(c2.Matches(c1))
	fmt.Println(c3.Matches(c1))
	fmt.Println(c3.Matches(c2))

	// Output:
	// true
	// false
	// true
	// true
	// false
	// true
}
