package rank

// Rank ist ein Datentyp für Karten-Werte.
type Rank int

// Konstanten für die Karten-Werte.
const (
	Two Rank = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

// Eine Map, die die Werte auf die Symbole abbildet.
var rankMap = map[Rank]string{
	Two:   "2",
	Three: "3",
	Four:  "4",
	Five:  "5",
	Six:   "6",
	Seven: "7",
	Eight: "8",
	Nine:  "9",
	Ten:   "10",
	Jack:  "J",
	Queen: "Q",
	King:  "K",
	Ace:   "A",
}

// String gibt den Wert als String zurück.
func (r Rank) String() string {
	return rankMap[r]
}
