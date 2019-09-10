package main

import "testing"
import "github.com/stretchr/testify/assert"


func TestSum(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input1 int
		input2 int
		expected int
	}{
		{1, 1, 2},
		{1, 2, 3},
		{2, 2, 4},
		{5, 2, 7},
	}

	for _, test := range tests {
		assert.Equal(Sum(test.input1, test.input2), test.expected)
	}
	
}

func Testmain(t *testing.T) {
	main()

}