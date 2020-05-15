//go:generate stringer -type=Suit,Rank

package shuffle

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

//Suit type
type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker // special case
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

//Rank type
type Rank uint8

const (
	_ Rank = iota
	Ace
	Two
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
)

const (
	minRank = Ace
	maxRank = King
)

//Card struct
type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

//New func
func New(opts ...func([]Card) []Card) []Card {
	var cards []Card

	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}

	for _, opt := range opts {
		cards = opt(cards)
	}

	return cards
}

//DefaultSort func
func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

//Less func
func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

//Sort func : Custom
func Sort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

func absRank(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}

//Shuffle func
func Shuffle(cards []Card) []Card {
	newCards := make([]Card, len(cards))

	r := rand.New(rand.NewSource(time.Now().Unix()))

	perm := r.Perm(len(cards))
	// perm = [0,1,4,2,3]
	for i, j := range perm {
		newCards[i] = cards[j]
	}

	return newCards
}

// Jokers func
func Jokers(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			cards = append(cards, Card{
				Rank: Rank(i),
				Suit: Joker,
			})
		}
		return cards
	}
}

//Filter func
func Filter(f func(card Card) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		var filtered []Card
		for _, c := range cards {
			if !f(c) {
				filtered = append(filtered, c)
			}
		}
		return filtered
	}
}

// Deck func
func Deck(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		var retCards []Card
		for i := 0; i < n; i++ {
			retCards = append(retCards, cards...)
		}
		return retCards
	}
}
