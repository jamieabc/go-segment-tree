package segment_test

import (
	"github.com/jamieabc/go-segment-tree/pkg/segment"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewWhenLengthPowerOf2(t *testing.T) {
	s := segment.New([]int{1, 3, 2, 9, 7, 6, 12, -1})
	values := make([]int, len(s.Data))
	for i, d := range s.Data {
		values[i] = d.Val
	}

	assert.Equal(t, []int{-1, 1, -1, 1, 2, 6, -1, 1, 3, 2, 9, 7, 6, 12, -1}, values, "wrong construction")
}

func TestNewWhenLengthNotPowerOf2(t *testing.T) {
	s := segment.New([]int{1, 3, 2, 9, 7, 6, -1})
	values := make([]int, len(s.Data))
	for i, d := range s.Data {
		values[i] = d.Val
	}

	assert.Equal(t, []int{-1, 1, -1, 1, 2, 6, -1, 1, 3, 2, 9, 7, 6}, values, "wrong construction")
}
