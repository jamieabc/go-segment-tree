package segment

import "math"

type Item struct {
	StartIndex int
	EndIndex   int
	Val        int
}

type Seg struct {
	Data []Item
}

//        0
//    1        2
//  3   4    5   6
func (s *Seg) parent(i int) int {
	return (i - 1) / 2
}

func (s *Seg) leftChild(i int) int {
	return i*2 + 1
}

func (s *Seg) rightChild(i int) int {
	return i*2 + 2
}

func (s *Seg) construct(data []int, position, start, end int) {
	if start == end {
		s.Data[position] = Item{
			StartIndex: start,
			EndIndex:   end,
			Val:        data[start],
		}
		return
	}

	mid := start + (end-start)/2
	leftChild := s.leftChild(position)
	rightChild := s.rightChild(position)

	s.construct(data, leftChild, start, mid)
	s.construct(data, rightChild, mid+1, end)

	s.Data[position] = Item{
		StartIndex: start,
		EndIndex:   end,
		Val:        min(s.Data[leftChild].Val, s.Data[rightChild].Val),
	}
}

func New(data []int) Seg {
	length := len(data)
	if length == 0 {
		return Seg{
			Data: make([]Item, 0),
		}
	}

	// l: total leaf count, considering tree height, total needed space
	// is 2*l-1 (e.g. l = 4, second level with 2 nodes, root with 1 node)

	s := Seg{
		Data: make([]Item, 2*length-1),
	}

	for i := range s.Data {
		s.Data[i].Val = math.MaxInt32
	}

	s.construct(data, 0, 0, length-1)

	return s
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}
