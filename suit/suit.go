package suit

// Suit ist ein Datentyp für Karten-Farben.
// I.W. int, aber durch die Konstanten unten werden die Werte eingeschränkt.
// Man kann nicht einfach jedes int verwenden, sondern nur die definierten Konstanten.
type Suit int

// Konstanten für die Karten-Farben.
const (
	Clubs Suit = iota
	Spades
	Hearts
	Diamonds
)

// Eine Map, die die Farben auf die Symbole abbildet.
var suitMap = map[Suit](string){
	Clubs:    "♣",
	Spades:   "♠",
	Hearts:   "♥",
	Diamonds: "♦",
}

// String gibt die Farbe als String zurück.
func (s Suit) String() string {
	return suitMap[s]
}
