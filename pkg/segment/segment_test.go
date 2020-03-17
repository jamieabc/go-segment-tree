package segment_test

import (
	"github.com/jamieabc/go-segment-tree/pkg/segment"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewWhenLengthPowerOf2(t *testing.T) {
	data := []int{1, 3, 2, 9, 7, 6, 12, -1}
	s := segment.New(data)
	items := s.Data()
	values := make([]int, len(items))
	for i, d := range s.Data() {
		values[i] = d.Val
	}

	assert.Equal(t, []int{-1, 1, -1, 1, 2, 6, -1, 1, 3, 2, 9, 7, 6, 12, -1}, values, "wrong construction")
	assert.Equal(t, 0, items[0].StartIndex, "wrong root start")
	assert.Equal(t, len(data)-1, items[0].EndIndex, "wrong root end")
	assert.Equal(t, 0, items[1].StartIndex, "wrong second node start")
	assert.Equal(t, 3, items[1].EndIndex, "wrong second node end")
	assert.Equal(t, 4, items[2].StartIndex, "wrong third node start")
	assert.Equal(t, 7, items[2].EndIndex, "wrong third node end")
}

func TestNewWhenLengthNotPowerOf2(t *testing.T) {
	s := segment.New([]int{1, 3, 2, 9, 7, 6, -1})
	values := make([]int, len(s.Data()))
	for i, d := range s.Data() {
		values[i] = d.Val
	}

	assert.Equal(t, []int{-1, 1, -1, 1, 2, 6, -1, 1, 3, 2, 9, 7, 6}, values, "wrong construction")
}
