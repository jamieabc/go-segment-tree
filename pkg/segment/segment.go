package segment

// Segment - interface for segment tree operation
type Segment interface {
	Data() []Item
	Query(int, int, int) int
}

type Item struct {
	StartIndex int
	EndIndex   int
	Val        int
}

type seg struct {
	data       []Item
	comparator func(int, int) int
	defaultVal int
}

// Data - return interval tree
func (s *seg) Data() []Item {
	return s.data
}

// Query - retrieve value from specific range
func (s *seg) Query(start, end, position int) int {
	item := s.data[position]

	// no match, return default value
	if item.StartIndex > end || item.EndIndex < start || end < start {
		return s.defaultVal
	}

	// exact match,, return value
	if item.StartIndex == start && item.EndIndex == end {
		return item.Val
	}

	left := leftChild(position)
	right := rightChild(position)

	// query range in left sub-domain
	if s.data[left].EndIndex >= end {
		return s.Query(start, end, left)
	}

	// query range in right sub-domain
	if s.data[right].StartIndex <= start {
		return s.Query(start, end, right)
	}

	// query range in between separation line
	return s.comparator(
		s.Query(start, s.data[left].EndIndex, left),
		s.Query(s.data[right].StartIndex, end, right),
	)
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

// position - the array index that stores compared value
// start / end - range that stored value represents
func (s *seg) construct(data []int, position, start, end int) {
	if start == end {
		s.data[position] = Item{
			StartIndex: start,
			EndIndex:   end,
			Val:        data[start],
		}
		return
	}

	mid := start + (end-start)/2
	leftChild := leftChild(position)
	rightChild := rightChild(position)

	s.construct(data, leftChild, start, mid)
	s.construct(data, rightChild, mid+1, end)

	s.data[position] = Item{
		StartIndex: start,
		EndIndex:   end,
		Val:        s.comparator(s.data[leftChild].Val, s.data[rightChild].Val),
	}
}

func New(data []int, f func(int, int) int, defaultValue int) Segment {
	length := len(data)
	if length == 0 {
		return &seg{
			data:       make([]Item, 0),
			comparator: f,
			defaultVal: defaultValue,
		}
	}

	// l: total leaf count, considering tree height, total needed space
	// is 2*l-1 (e.g. l = 4, second level with 2 nodes, root with 1 node)

	s := &seg{
		data:       make([]Item, 2*length-1),
		comparator: f,
		defaultVal: defaultValue,
	}

	for i := range s.data {
		s.data[i].Val = defaultValue
	}

	s.construct(data, 0, 0, length-1)

	return s
}
