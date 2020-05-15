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

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)

	exp := Card{Rank: Ace, Suit: Spade}

	if cards[0] != exp {
		t.Error("Expected Ave of Spades as first card. recieved:", cards[0])
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))

	exp := Card{Rank: Ace, Suit: Spade}

	if cards[0] != exp {
		t.Error("Expected Ave of Spades as first card. recieved:", cards[0])
	}
}

func TestJokers(t *testing.T) {
	cards := New(Jokers(3))
	count := 0
	for _, card := range cards {
		if card.Suit == Joker {
			count++
		}
	}

	if count != 3 {
		t.Error("Expected 3 Jokers, recieved:", count)
	}
}

func TestFilter(t *testing.T) {

	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}

	cards := New(Filter(filter))

	for _, c := range cards {
		if c.Rank == Two || c.Rank == Three {
			t.Error("Expeceted all twos and threes to be filtered out.")
		}
	}
}
