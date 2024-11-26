package card

import (
	"fmt"
	"maumau/rank"
	"maumau/suit"
)

// Card repräsentiert eine Spielkarte mit Farbe und Wert.
type Card struct {
	Suit suit.Suit
	Rank rank.Rank
}

// String gibt die Karte in der Form <Suit><Rank> zurück.
func (c Card) String() string {
	return fmt.Sprintf("%s%s", c.Suit, c.Rank)
}

// Image liefert eine Liste von Strings, die ein Bild der Karte als AsciiArt darstellen.
func (c Card) Image() []string {
	return []string{
		"┌───────┐",
		fmt.Sprintf("│%-2s     │", c.Rank),
		"│       │",
		fmt.Sprintf("│   %s   │", c.Suit),
		"│       │",
		fmt.Sprintf("│     %2s│", c.Rank),
		"└───────┘",
	}
}

// Matches gibt zurück, ob die Karte zu einer anderen passt.
func (c Card) Matches(other Card) bool {
	return c.Suit == other.Suit || c.Rank == other.Rank
}
