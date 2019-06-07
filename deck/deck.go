package deck

import (
	"fmt"
	"math/rand"
	"time"
)

//Card represents a game card
type Card struct {
	Value Value
	Suit  Suit
	Score int
}

//Deck represents a lo of cards
type Deck struct {
	Cards []Card
}

//Value represents rank of a card
type Value int

const (
	Zero Value = iota
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

type Suit int

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker // this is a special case
)

var defaultSuitOrder = [...]Suit{Spade, Diamond, Club, Heart}

//New func return a Deck of cards
func New() Deck {
	d := Deck{}
	d.Cards = make([]Card, 0, len(defaultSuitOrder)*int(maxRank))
	for _, suit := range defaultSuitOrder {
		for score := minRank; score <= maxRank; score++ {
			d.Cards = append(d.Cards, Card{
				Value: score,
				Suit:  suit,
				Score: int(score),
			})
		}
	}
	return d
}

//Filter return a new Deck withrou filter cards
func (d *Deck) Filter(f func(c Card) bool) Deck {
	dd := Deck{}
	for _, card := range d.Cards {
		if !f(card) {
			dd.Cards = append(dd.Cards, card)
		}
	}
	return dd
}

func (c *Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Value.String(), c.Suit.String())
}

//Rank return rank in default order of a card
func (c *Card) Rank() int {
	return int(c.Suit)*int(maxRank) + int(c.Value)
}

//Less is less function for sort.Slice
func (d *Deck) Less() func(i, j int) bool {
	return func(i, j int) bool {
		return d.Cards[i].Rank() < d.Cards[j].Rank()
	}
}

//Shuffle cards into deck
func (d *Deck) Shuffle() {
	tmp := make([]Card, len(d.Cards))

	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(tmp))
	for i, j := range perm {
		tmp[i] = d.Cards[j]
	}
	d.Cards = tmp
}

func (d *Deck) AddJockers(n int) {
	for i := 0; i < n; i++ {
		d.Cards = append(d.Cards, Card{
			Value: Zero,
			Suit:  Joker,
			Score: 0,
		})
	}
}
