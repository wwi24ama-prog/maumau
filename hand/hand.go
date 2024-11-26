package hand

import (
	"fmt"
	"maumau/card"
	"maumau/deck"
)

// Hand repräsentiert eine Reihe von Handkarten.
// Wie bei Deck enthält es eine Liste von Karten.
// Hier gibt es aber keine Random-Source, weil die Karten nicht gemischt werden.
type Hand struct {
	Cards []card.Card
}

// NewFromDeck erwartet einen Kartenstapel und gibt eine Hand mit den obersten 7 Karten zurück.
func NewFromDeck(deck *deck.Deck) Hand {
	// Erstelle eine leere Hand mit Platz für 7 Karten.
	hand := Hand{Cards: make([]card.Card, 0, 7)}

	// Ziehe 7 Karten vom Stapel und füge sie zur Hand hinzu.
	for i := 0; i < 7; i++ {
		hand.Cards = append(hand.Cards, deck.Draw())
	}

	return hand
}

// Show zeigt die Hand auf der Konsole an.
func (h Hand) Show() {
	lines := []string{"", "", "", "", "", "", ""}
	for _, c := range h.Cards {
		for i, line := range c.Image() {
			lines[i] += line
		}
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}
