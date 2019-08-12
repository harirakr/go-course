package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Data struct {
	name     string
	input    []int
	expected []int
}

func TestSortFunctions(t *testing.T) {

	testSet := []Data{
		{"Empty set",
			[]int{},
			[]int{},
		},
		{"Single element set",
			[]int{10},
			[]int{10},
		},
		{"Two element set",
			[]int{20, 10},
			[]int{10, 20},
		},
		{"Worst case",
			[]int{200, 111, 108, 93, 89, 76, 68, 51, 42, 35},
			[]int{35, 42, 51, 68, 76, 89, 93, 108, 111, 200},
		},
		{"Average case 1",
			[]int{4, 8, 3, 11, 1, 0, 5},
			[]int{0, 1, 3, 4, 5, 8, 11},
		},
		{"Average case 2",
			[]int{3, 2, 1, 5},
			[]int{1, 2, 3, 5},
		},
		{"Best case",
			[]int{20, 31, 48, 53, 69, 76, 88, 91, 102, 115},
			[]int{20, 31, 48, 53, 69, 76, 88, 91, 102, 115},
		},
	}

	for _, v := range testSet {
		assert.Equal(t, v.expected, bubble(v.input))
		assert.Equal(t, v.expected, insertion(v.input))
	}
}


