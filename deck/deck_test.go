package deck

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestNewDeck(t *testing.T) {
	deck := New()
	deck.AddJockers(3)

	spew.Dump(deck)

	tcs := []Card{
		{
			Suit:  Heart,
			Value: King,
		},
		{
			Suit:  Diamond,
			Value: Jack,
		},
	}

	for _, tc := range tcs {
		spew.Dump(tc.Rank())
	}

	deck.Shuffle()

	spew.Dump(deck)

	deck2 := deck.Filter(func(c Card) bool { return c.Value == Two })
	spew.Dump(deck2)
}
