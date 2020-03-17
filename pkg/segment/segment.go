package segment

import "math"

// Segment - interface for segment tree operation
type Segment interface {
	Data() []Item
}

type Item struct {
	StartIndex int
	EndIndex   int
	Val        int
}

type seg struct {
	data       []Item
	comparable func(int, int) int
}

// Data - return interval tree
func (s *seg) Data() []Item {
	return s.data
}

//        0
//    1        2
//  3   4    5   6
func parent(i int) int {
	return (i - 1) / 2
}

func leftChild(i int) int {
	return i*2 + 1
}

func rightChild(i int) int {
	return i*2 + 2
}

func (s *seg) construct(data []int, position, start, end int) (int, int) {
	if start == end {
		s.data[position] = Item{
			StartIndex: start,
			EndIndex:   end,
			Val:        data[start],
		}
		return start, end
	}

	mid := start + (end-start)/2
	leftChild := leftChild(position)
	rightChild := rightChild(position)

	l, _ := s.construct(data, leftChild, start, mid)
	_, r := s.construct(data, rightChild, mid+1, end)

	s.data[position] = Item{
		StartIndex: l,
		EndIndex:   r,
		Val:        s.comparable(s.data[leftChild].Val, s.data[rightChild].Val),
	}
	return l, r
}

func New(data []int, f func(int, int) int) Segment {
	length := len(data)
	if length == 0 {
		return &seg{
			data:       make([]Item, 0),
			comparable: f,
		}
	}

	// l: total leaf count, considering tree height, total needed space
	// is 2*l-1 (e.g. l = 4, second level with 2 nodes, root with 1 node)

	s := &seg{
		data:       make([]Item, 2*length-1),
		comparable: f,
	}

	for i := range s.data {
		s.data[i].Val = math.MaxInt32
	}

	s.construct(data, 0, 0, length-1)

	return s
}
