package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func main() {
	tree := Tree{
		Left: &Tree{
			Left: &Tree{
				Value: 1,
			},
			Value: 1,
			Right: &Tree{
				Value: 3,
			},
		},
		Value: 3,
		Right: &Tree{
			Left: &Tree{
				Left: &Tree{
					Value: 512,
				},
				Value: 8,
			},
			Value: 5,
			Right: &Tree{
				Value: 13,
			},
		},
	}

	c := make(chan int)
	go Walk(&tree, c)
	for v := range c {
		fmt.Println(v)
	}

	bs := &Tree{}
	bs.Add(5)
	bs.Add(3)
	bs.Add(4)

	bs.Add(6)
	bs.Add(2)
	bs.Add(10)
	bs.Add(9)
	bs.Add(8)
	bs.Add(7)




	spew.Dump(bs)

	c = make(chan int)
	go Walk(bs, c)
	for v := range c {
		fmt.Println(v)
	}

	bs2 := &Tree{}

	for _, i := range []int{11, 6, 8, 19, 4, 10, 5, 17, 43, 49, 31} {
		bs2.Add(i)
	}
	spew.Dump(bs2)

	c = make(chan int)
	go Walk(bs2, c)
	for v := range c {
		fmt.Println(v)
	}

	tr := make(chan *Tree)
	go bs2.Find(3, tr)
	select {
	case leaf := <- tr:
		spew.Dump(leaf)
	}

}

func (t *Tree) Find(v int, c chan *Tree) {
	spew.Dump(t.Value)
	if v < t.Value {

		if t.Left == nil {
			c <- t
		} else {
		 t.Left.Find(v, c)
	 }
	}	else {
			if t.Right == nil {
				c <- t
			} else {
			t.Right.Find(v, c)
		}
	}
}


func (t *Tree) Add(v int) *Tree {
	if t == nil {
		t = &Tree{Value:v}
		return t
	}
	if t.Value == 0 {
		t.Value = v
		return t
	}
	if v < t.Value {
		if t.Left == nil {
			t.Left = &Tree{Value:v}
			return t
		}
		return t.Left.Add(v)
	} else {
		if t.Right == nil {
			t.Right = &Tree{Value:v}
			return t
		}
		return t.Right.Add(v)
	}
}


func Walk(t *Tree, c chan int) {
	rWalk(t, c)
	close(c)
}


func rWalk(t *Tree, c chan int) {
	if t != nil {
		rWalk(t.Left, c)
		c <- t.Value
		rWalk(t.Right, c)
	}
}
