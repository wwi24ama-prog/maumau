package hand

import (
	"fmt"
	"maumau/deck"
)

func ExampleNewFromDeck() {
	d := deck.New32()
	h := NewFromDeck(&d)

	for _, c := range h.Cards {
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
}

func ExampleHand_Show() {
	d := deck.New32()
	h := NewFromDeck(&d)

	h.Show()

	// Output:
	//┌───────┐┌───────┐┌───────┐┌───────┐┌───────┐┌───────┐┌───────┐
	//│7      ││8      ││9      ││10     ││J      ││Q      ││K      │
	//│       ││       ││       ││       ││       ││       ││       │
	//│   ♣   ││   ♣   ││   ♣   ││   ♣   ││   ♣   ││   ♣   ││   ♣   │
	//│       ││       ││       ││       ││       ││       ││       │
	//│      7││      8││      9││     10││      J││      Q││      K│
	//└───────┘└───────┘└───────┘└───────┘└───────┘└───────┘└───────┘
}
