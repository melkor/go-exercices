package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountWords(t *testing.T) {
	testCases := []struct {
		CamlString string
		Count      int
	}{
		{
			CamlString: "saveChangesInTheEditor",
			Count:      5,
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.Count, countWords(tc.CamlString))
	}
}
