package shuffle

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Spade})
	fmt.Println(Card{Rank: Seven, Suit: Diamond})
	fmt.Println(Card{Rank: Three, Suit: Club})
	fmt.Println(Card{Suit: Joker})

	//Output:
	// Ace of Hearts
	// Two of Spades
	// Seven of Diamonds
	// Three of Clubs
	// Joker

}

func TestNew(t *testing.T) {
	cards := New()

	if len(cards) != 52 {
		t.Error("Wrong number of cards in a new Deck")
	}
}
