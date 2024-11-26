package deck

import (
	"fmt"
	"math/rand"
	"maumau/card"
	"maumau/rank"
	"maumau/suit"
	"time"
)

// Deck repräsentiert einen Kartenstapel.
// Es enthält i.W. eine Liste von Karten.
// Diese ist im Struct und nicht direkt definiert,
// weil Code, der das Deck manipuliert, dadurch einfacher/übersichtlicher wird.
// Zusätzlich enthält es eine Random-Source, die für das Mischen der Karten verwendet wird.
// Diese Random-Source wird im Konstruktor erstellt und dann für alle Operationen verwendet.
// Normalerweise liefert sie wirklich Zufallszahlen, für Tests kann sie aber durch
// etwas vorhersagbares ersetzt werden.
type Deck struct {
	cards   []card.Card
	randSrc rand.Source
}

// New32 gibt einen neuen Skat-Kartenstapel mit 32 Karten zurück.
func New32() Deck {
	// Erstelle einen leeren Kartenstapel mit 32 Plätzen.
	// Der zweite Parameter reserviert von Anfang an Platz für 32 Karten.
	// Dadurch kosten append-Operationen (fast) keine Performance.
	cards := make([]card.Card, 0, 32)

	// Füge alle Karten hinzu.
	// Die äußere Schleife läuft über alle Farben
	// und die innere Schleife über alle Werte.
	// So wird jede Kombination von Farbe und Wert einmal durchlaufen.
	for _, suit := range []suit.Suit{suit.Clubs, suit.Spades, suit.Hearts, suit.Diamonds} {
		for _, rank := range []rank.Rank{rank.Seven, rank.Eight, rank.Nine, rank.Ten, rank.Jack, rank.Queen, rank.King, rank.Ace} {
			cards = append(cards, card.Card{Suit: suit, Rank: rank})
		}
	}

	// Erstelle eine neue Random-Source.
	randSrc := rand.NewSource(time.Now().UnixNano())

	// Liefere ein neues Deck mit den erstellten Karten und der Random-Source.
	return Deck{cards: cards, randSrc: randSrc}
}

// Shuffle mischt den Kartenstapel mit Hilfe der internen Random-Source.
func (d *Deck) Shuffle() {
	// Erstelle einen neuen Zufallsgenerator mit der Random-Source.
	rng := rand.New(d.randSrc)

	// Mische die Karten mit Hilfe einer Swap-Funktion,
	// die zwei Karten im Deck vertauscht.
	rng.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

// Len gibt die Anzahl der Karten im Stapel zurück.
func (d Deck) Len() int {
	return len(d.cards)
}

// Draw zieht eine Karte vom Stapel und gibt sie zurück.
// Die Karte wird dabei aus dem Stapel entfernt.
// Wenn der Stapel leer ist, wird eine Panic ausgelöst.
// D.h. es muss vorher geprüft werden, ob noch Karten vorhanden sind.
func (d *Deck) Draw() card.Card {
	// Prüfe, ob der Stapel leer ist.
	if len(d.cards) == 0 {
		panic("no cards left")
	}

	// Ziehe die oberste Karte und entferne sie aus dem Stapel.
	c := d.cards[0]
	d.cards = d.cards[1:]

	// Liefere die gezogene Karte.
	return c
}

// ShowTop zeigt die oberste Karte des Stapels auf der Konsole an.
// Falls der Stapel leer ist, wird eine Panic ausgelöst.
func (d Deck) ShowTop() {
	// Prüfe, ob der Stapel leer ist.
	if len(d.cards) == 0 {
		panic("no cards left")
	}

	// Zeige die oberste Karte auf der Konsole an.
	c := d.cards[0]
	for _, line := range c.Image() {
		fmt.Println(line)
	}
}
