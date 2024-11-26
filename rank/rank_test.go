package rank

import "fmt"

func ExampleRank_String() {
	for r := Two; r <= Ace; r++ {
		fmt.Println(r)
	}

	// Output:
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
	// 9
	// 10
	// J
	// Q
	// K
	// A
}
