package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCipher(t *testing.T) {
	testCases := []struct {
		s string
		k int
		e string
	}{
		{
			s: "abcdefghijklmnopqrstuvwxyz",
			k: 3,
			e: "defghijklmnopqrstuvwxyzabc",
		},
		{
			s: "middle-Outz",
			k: 2,
			e: "okffng-Qwvb",
		},
	}

	for _, tc := range testCases {
		r, err := cipher(tc.s, tc.k)
		assert.NoError(t, err)
		assert.Equal(t, tc.e, r)
	}
}
