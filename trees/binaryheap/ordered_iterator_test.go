package binaryheap

import (
	"testing"

	testCommon "github.com/JonasMuehlmann/datastructures.go/tests"

	"github.com/JonasMuehlmann/datastructures.go/ds"

	"github.com/JonasMuehlmann/datastructures.go/utils"
	"github.com/stretchr/testify/assert"
)

const (
	NoMoveMagicPosition = 7869543205234798
)

func TestHeapOrderedIteratorIsValid(t *testing.T) {
	tests := []struct {
		name         string
		map_         *Heap[int]
		position     int
		isValid      bool
		iteratorInit func(*Heap[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
	}{
		{
			name:         "Empty",
			map_:         New[int](utils.BasicComparator[int]),
			isValid:      false,
			iteratorInit: (*Heap[int]).OrderedFirst,
		},

		{
			name:         "One element, first",
			map_:         NewFromSlice[int](utils.BasicComparator[int], []int{1}),
			position:     NoMoveMagicPosition,
			isValid:      true,
			iteratorInit: (*Heap[int]).OrderedFirst,
		},
		{
			name:         "3 elements, middle",
			map_:         NewFromSlice[int](utils.BasicComparator[int], []int{1, 2, 3}),
			position:     2,
			isValid:      true,
			iteratorInit: (*Heap[int]).OrderedFirst,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.map_)

			if test.position != NoMoveMagicPosition {
				it.MoveTo(test.position)
			}

			isValid := it.IsValid()

			assert.Equalf(t, test.isValid, isValid, test.name)
		})
	}
}

func TestHeapOrderedIteratorGet(t *testing.T) {
	tests := []struct {
		name     string
		map_     *Heap[int]
		position int
		value    int
		found    bool
	}{
		{
			name:     "Empty",
			map_:     New[int](utils.BasicComparator[int]),
			position: NoMoveMagicPosition,
			found:    false,
		},
		{
			name:     "One element, first",
			map_:     NewFromSlice[int](utils.BasicComparator[int], []int{1}),
			position: NoMoveMagicPosition,
			value:    1,
			found:    true,
		},
		{
			name:     "3 elements, first",
			map_:     NewFromSlice[int](utils.BasicComparator[int], []int{1, 2, 3}),
			position: 0,
			value:    1,
			found:    true,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it := test.map_.OrderedFirst()

			if test.position != NoMoveMagicPosition {
				it.MoveTo(test.position)
			}

			value, found := it.Get()

			assert.Equalf(t, test.found, found, test.name)
			assert.Equalf(t, test.value, value, test.name)
		})
	}
}

func TestHeapOrderedIteratorSet(t *testing.T) {
	tests := []struct {
		name        string
		map_        *Heap[int]
		position    int
		value       int
		successfull bool
	}{
		{
			name:        "Empty",
			map_:        New[int](utils.BasicComparator[int]),
			position:    NoMoveMagicPosition,
			value:       1,
			successfull: false,
		},
		{
			name:        "One element, first",
			map_:        NewFromSlice[int](utils.BasicComparator[int], []int{1}),
			position:    NoMoveMagicPosition,
			value:       1,
			successfull: true,
		},
		{
			name:        "3 elements, first",
			map_:        NewFromSlice[int](utils.BasicComparator[int], []int{1, 2, 3}),
			position:    0,
			value:       1,
			successfull: true,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it := test.map_.OrderedFirst()

			if test.position != NoMoveMagicPosition {
				it.MoveTo(test.position)
			}

			successfull := it.Set(test.value)

			assert.Equalf(t, test.successfull, successfull, test.name)
		})
	}
}

func TestHeapOrderedIteratorGetAt(t *testing.T) {
	tests := []struct {
		name     string
		map_     *Heap[int]
		position int
		value    int
		found    bool
	}{
		{
			name:     "Empty",
			map_:     New[int](utils.BasicComparator[int]),
			position: 0,
			found:    false,
		},

		{
			name:     "One element, first",
			map_:     NewFromSlice[int](utils.BasicComparator[int], []int{1}),
			position: 0,
			value:    1,
			found:    true,
		},
		{
			name:     "3 elements, first",
			map_:     NewFromSlice[int](utils.BasicComparator[int], []int{1, 2, 3}),
			position: 0,
			value:    1,
			found:    true,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it := test.map_.OrderedFirst()

			value, found := it.GetAt(test.position)

			assert.Equalf(t, test.found, found, test.name)
			assert.Equalf(t, test.value, value, test.name)
		})
	}
}

func TestHeapOrderedIteratorSetAt(t *testing.T) {
	tests := []struct {
		name        string
		map_        *Heap[int]
		position    int
		value       int
		successfull bool
	}{
		{
			name:        "Empty",
			map_:        New[int](utils.BasicComparator[int]),
			position:    0,
			value:       1,
			successfull: false,
		},

		{
			name:        "One element, first",
			map_:        NewFromSlice[int](utils.BasicComparator[int], []int{1}),
			position:    0,
			value:       1,
			successfull: true,
		},
		{
			name:        "3 elements, first",
			map_:        NewFromSlice[int](utils.BasicComparator[int], []int{1, 2, 3}),
			position:    0,
			value:       1,
			successfull: true,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it := test.map_.OrderedFirst()

			successfull := it.SetAt(test.position, test.value)

			assert.Equalf(t, test.successfull, successfull, test.name)
		})
	}
}

func TestHeapOrderedIteratorDistanceTo(t *testing.T) {
	tests := []struct {
		name     string
		map_     *Heap[int]
		pos1     int
		pos2     int
		distance int
	}{
		{
			name:     "Equal",
			map_:     NewFromSlice[int](utils.BasicComparator[int], []int{1, 2, 3}),
			pos1:     1,
			pos2:     1,
			distance: 0,
		},
		{
			name:     "First lower",
			map_:     NewFromSlice[int](utils.BasicComparator[int], []int{1, 2, 3}),
			pos1:     1,
			pos2:     2,
			distance: -1,
		},
		{
			name:     "Second lower",
			map_:     NewFromSlice[int](utils.BasicComparator[int], []int{1, 2, 3}),
			pos1:     2,
			pos2:     1,
			distance: 1,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it1 := test.map_.OrderedBegin()
			it2 := test.map_.OrderedBegin()

			it1.MoveTo(test.pos1)
			it2.MoveTo(test.pos2)

			distance := it1.DistanceTo(it2)

			assert.Equalf(t, test.distance, distance, test.name)
		})
	}
}

func TestHeapOrderedIteratorIsAfter(t *testing.T) {
	tests := []struct {
		name    string
		map_    *Heap[int]
		pos1    int
		pos2    int
		isAfter bool
	}{
		{
			name:    "Equal",
			map_:    NewFromSlice[int](utils.BasicComparator[int], []int{1, 2, 3}),
			pos1:    1,
			pos2:    1,
			isAfter: false,
		},
		{
			name:    "First lower",
			map_:    NewFromSlice[int](utils.BasicComparator[int], []int{1, 2, 3}),
			pos1:    1,
			pos2:    2,
			isAfter: false,
		},
		{
			name:    "Second lower",
			map_:    NewFromSlice[int](utils.BasicComparator[int], []int{1, 2, 3}),
			pos1:    2,
			pos2:    1,
			isAfter: true,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it1 := test.map_.OrderedBegin()
			it2 := test.map_.OrderedBegin()

			it1.MoveTo(test.pos1)
			it2.MoveTo(test.pos2)

			isAfter := it1.IsAfter(it2)

			assert.Equalf(t, test.isAfter, isAfter, test.name)
		})
	}
}

func TestHeapOrderedIteratorIsBefore(t *testing.T) {
	tests := []struct {
		name    string
		map_    *Heap[int]
		pos1    int
		pos2    int
		isAfter bool
	}{
		{
			name:    "Equal",
			map_:    NewFromSlice[int](utils.BasicComparator[int], []int{1, 2, 3}),
			pos1:    1,
			pos2:    1,
			isAfter: false,
		},
		{
			name:    "First lower",
			map_:    NewFromSlice[int](utils.BasicComparator[int], []int{1, 2, 3}),
			pos1:    1,
			pos2:    2,
			isAfter: true,
		},
		{
			name:    "Second lower",
			map_:    NewFromSlice[int](utils.BasicComparator[int], []int{1, 2, 3}),
			pos1:    2,
			pos2:    1,
			isAfter: false,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it1 := test.map_.OrderedBegin()
			it2 := test.map_.OrderedBegin()

			it1.MoveTo(test.pos1)
			it2.MoveTo(test.pos2)

			isAfter := it1.IsBefore(it2)

			assert.Equalf(t, test.isAfter, isAfter, test.name)
		})
	}
}

func TestHeapOrderedIteratorIsEqual(t *testing.T) {
	tests := []struct {
		name      string
		position1 int
		position2 int
		isAfter   bool
	}{
		{
			name:      "Equal",
			position1: 1,
			position2: 1,
			isAfter:   true,
		},
		{
			name:      "First lower",
			position1: 1,
			position2: 2,
			isAfter:   false,
		},
		{
			name:      "Second lower",
			position1: 2,
			position2: 1,
			isAfter:   false,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			m := NewFromSlice[int](utils.BasicComparator[int], []int{1, 2, 4, 5})

			it1 := m.OrderedFirst()
			it2 := m.OrderedFirst()

			it1.MoveTo(test.position1)
			it2.MoveTo(test.position2)

			isAfter := it1.IsEqual(it2)

			assert.Equalf(t, test.isAfter, isAfter, test.name+m.ToString())
		})
	}
}

func TestHeapOrderedIteratorIndex(t *testing.T) {
	tests := []struct {
		name         string
		map_         *Heap[int]
		index        int
		valid        bool
		iteratorInit func(*Heap[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
	}{
		{
			name:         "Empty",
			map_:         New[int](utils.BasicComparator[int]),
			valid:        false,
			iteratorInit: (*Heap[int]).OrderedBegin,
		},
		{
			name:         "One element, begin",
			map_:         NewFromSlice[int](utils.BasicComparator[int], []int{1}),
			valid:        false,
			iteratorInit: (*Heap[int]).OrderedBegin,
		},
		{
			name:         "One element, end",
			map_:         NewFromSlice[int](utils.BasicComparator[int], []int{1}),
			valid:        false,
			iteratorInit: (*Heap[int]).OrderedEnd,
		},
		{
			name:         "One element, first",
			map_:         NewFromSlice[int](utils.BasicComparator[int], []int{1}),
			index:        0,
			valid:        true,
			iteratorInit: (*Heap[int]).OrderedFirst,
		},
		{
			name:         "One element, last",
			map_:         NewFromSlice[int](utils.BasicComparator[int], []int{1}),
			index:        0,
			valid:        true,
			iteratorInit: (*Heap[int]).OrderedLast,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.map_)

			position, valid := it.Index()

			assert.Equalf(t, test.valid, valid, test.name)
			if test.valid {
				assert.Equalf(t, test.index, position, test.name)
			}
		})
	}
}

func TestHeapOrderedIteratorSize(t *testing.T) {
	tests := []struct {
		name string
		map_ *Heap[int]
		size int
	}{
		{
			name: "Empty",
			map_: New[int](utils.BasicComparator[int]),
			size: 0,
		},
		{
			name: "3 elements, middle",
			map_: NewFromSlice[int](utils.BasicComparator[int], []int{1, 2, 3}),
			size: 3,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it := test.map_.OrderedBegin()

			size := it.Size()

			assert.Equalf(t, test.size, size, test.name)
		})
	}
}

func TestHeapOrderedIteratorNext(t *testing.T) {
	tests := []struct {
		name          string
		map_          *Heap[int]
		position      int
		isValidBefore bool
		isValidAfter  bool
		iteratorInit  func(*Heap[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
	}{
		{
			name:          "Empty",
			map_:          New[int](utils.BasicComparator[int]),
			position:      NoMoveMagicPosition,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Heap[int]).OrderedBegin,
		},
		{
			name:          "One element, begin",
			map_:          NewFromSlice[int](utils.BasicComparator[int], []int{1}),
			position:      NoMoveMagicPosition,
			isValidBefore: false,
			isValidAfter:  true,
			iteratorInit:  (*Heap[int]).OrderedBegin,
		},
		{
			name:          "One element, end",
			map_:          NewFromSlice[int](utils.BasicComparator[int], []int{1}),
			position:      NoMoveMagicPosition,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Heap[int]).OrderedEnd,
		},
		{
			name:          "One element, first",
			map_:          NewFromSlice[int](utils.BasicComparator[int], []int{1}),
			position:      NoMoveMagicPosition,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Heap[int]).OrderedFirst,
		},
		{
			name:          "One element, last",
			map_:          NewFromSlice[int](utils.BasicComparator[int], []int{1}),
			position:      NoMoveMagicPosition,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Heap[int]).OrderedLast,
		},
		{
			name:          "3 elements, middle",
			map_:          NewFromSlice[int](utils.BasicComparator[int], []int{1, 2, 3}),
			position:      1,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Heap[int]).OrderedBegin,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.map_)

			if test.position != NoMoveMagicPosition {
				it.MoveTo(test.position)
			}

			isValidBefore := it.IsValid()
			assert.Equalf(t, test.isValidBefore, isValidBefore, test.name)

			it.Next()

			isValidAfter := it.IsValid()
			assert.Equalf(t, test.isValidAfter, isValidAfter, test.name)
		})
	}
}

func TestHeapOrderedIteratorNextN(t *testing.T) {
	tests := []struct {
		name          string
		map_          *Heap[int]
		position      int
		n             int
		isValidBefore bool
		isValidAfter  bool
		iteratorInit  func(*Heap[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
	}{
		{
			name:          "Empty",
			map_:          New[int](utils.BasicComparator[int]),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Heap[int]).OrderedBegin,
		},
		{
			name:          "One element, begin",
			map_:          NewFromSlice[int](utils.BasicComparator[int], []int{1}),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: false,
			isValidAfter:  true,
			iteratorInit:  (*Heap[int]).OrderedBegin,
		},
		{
			name:          "One element, end",
			map_:          NewFromSlice[int](utils.BasicComparator[int], []int{1}),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Heap[int]).OrderedEnd,
		},
		{
			name:          "One element, first",
			map_:          NewFromSlice[int](utils.BasicComparator[int], []int{1}),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Heap[int]).OrderedFirst,
		},
		{
			name:          "One element, last",
			map_:          NewFromSlice[int](utils.BasicComparator[int], []int{1}),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Heap[int]).OrderedLast,
		},
		{
			name:          "3 elements, middle",
			map_:          NewFromSlice[int](utils.BasicComparator[int], []int{1, 2, 3}),
			position:      1,
			n:             1,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Heap[int]).OrderedBegin,
		},
		{
			name:          "3 elements, middle, move out of bounds",
			map_:          NewFromSlice[int](utils.BasicComparator[int], []int{1, 2, 3}),
			position:      1,
			n:             5,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Heap[int]).OrderedBegin,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.map_)

			if test.position != NoMoveMagicPosition {
				it.MoveTo(test.position)
			}

			isValidBefore := it.IsValid()
			assert.Equalf(t, test.isValidBefore, isValidBefore, test.name)

			it.NextN(test.n)

			isValidAfter := it.IsValid()
			assert.Equalf(t, test.isValidAfter, isValidAfter, test.name)
		})
	}
}

func TestHeapOrderedIteratorPrevious(t *testing.T) {
	tests := []struct {
		name          string
		map_          *Heap[int]
		position      int
		isValidBefore bool
		isValidAfter  bool
		iteratorInit  func(*Heap[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
	}{
		{
			name:          "Empty",
			map_:          New[int](utils.BasicComparator[int]),
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Heap[int]).OrderedBegin,
		},
		{
			name:          "One element, begin",
			map_:          NewFromSlice[int](utils.BasicComparator[int], []int{1}),
			position:      NoMoveMagicPosition,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Heap[int]).OrderedBegin,
		},
		{
			name:          "One element, end",
			map_:          NewFromSlice[int](utils.BasicComparator[int], []int{1}),
			position:      NoMoveMagicPosition,
			isValidBefore: false,
			isValidAfter:  true,
			iteratorInit:  (*Heap[int]).OrderedEnd,
		},
		{
			name:          "One element, first",
			map_:          NewFromSlice[int](utils.BasicComparator[int], []int{1}),
			position:      NoMoveMagicPosition,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Heap[int]).OrderedFirst,
		},
		{
			name:          "One element, last",
			map_:          NewFromSlice[int](utils.BasicComparator[int], []int{1}),
			position:      NoMoveMagicPosition,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Heap[int]).OrderedLast,
		},
		{
			name:          "3 elements, middle",
			map_:          NewFromSlice[int](utils.BasicComparator[int], []int{1, 2, 3}),
			position:      1,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Heap[int]).OrderedBegin,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.map_)

			if test.position != NoMoveMagicPosition {
				it.MoveTo(test.position)
			}

			isValidBefore := it.IsValid()
			assert.Equalf(t, test.isValidBefore, isValidBefore, test.name)

			it.Previous()

			isValidAfter := it.IsValid()
			assert.Equalf(t, test.isValidAfter, isValidAfter, test.name)
		})
	}
}

func TestHeapOrderedIteratorPreviousN(t *testing.T) {
	tests := []struct {
		name          string
		map_          *Heap[int]
		position      int
		n             int
		isValidBefore bool
		isValidAfter  bool
		iteratorInit  func(*Heap[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
	}{
		{
			name:          "Empty",
			map_:          New[int](utils.BasicComparator[int]),
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Heap[int]).OrderedBegin,
		},
		{
			name:          "One element, begin",
			map_:          NewFromSlice[int](utils.BasicComparator[int], []int{1}),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Heap[int]).OrderedBegin,
		},
		{
			name:          "One element, end",
			map_:          NewFromSlice[int](utils.BasicComparator[int], []int{1}),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: false,
			isValidAfter:  true,
			iteratorInit:  (*Heap[int]).OrderedEnd,
		},
		{
			name:          "One element, first",
			map_:          NewFromSlice[int](utils.BasicComparator[int], []int{1}),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Heap[int]).OrderedFirst,
		},
		{
			name:          "One element, last",
			map_:          NewFromSlice[int](utils.BasicComparator[int], []int{1}),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Heap[int]).OrderedLast,
		},
		{
			name:          "3 elements, middle",
			map_:          NewFromSlice[int](utils.BasicComparator[int], []int{1, 2, 3}),
			position:      1,
			n:             1,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Heap[int]).OrderedBegin,
		},
		{
			name:          "3 elements, middle, move out of bounds",
			map_:          NewFromSlice[int](utils.BasicComparator[int], []int{1, 2, 3}),
			position:      1,
			n:             5,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Heap[int]).OrderedBegin,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.map_)

			if test.position != NoMoveMagicPosition {
				it.MoveTo(test.position)
			}

			isValidBefore := it.IsValid()
			assert.Equalf(t, test.isValidBefore, isValidBefore, test.name)

			it.PreviousN(test.n)

			isValidAfter := it.IsValid()
			assert.Equalf(t, test.isValidAfter, isValidAfter, test.name)
		})
	}
}

func TestHeapOrderedIteratorMoveBy(t *testing.T) {
	tests := []struct {
		name          string
		map_          *Heap[int]
		position      int
		n             int
		isValidBefore bool
		isValidAfter  bool
		iteratorInit  func(*Heap[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
	}{
		{
			name:          "Empty",
			map_:          New[int](utils.BasicComparator[int]),
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Heap[int]).OrderedBegin,
		},
		{
			name:          "One element, begin",
			map_:          NewFromSlice[int](utils.BasicComparator[int], []int{1}),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: false,
			isValidAfter:  true,
			iteratorInit:  (*Heap[int]).OrderedBegin,
		},
		{
			name:          "One element, end",
			map_:          NewFromSlice[int](utils.BasicComparator[int], []int{1}),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Heap[int]).OrderedEnd,
		},
		{
			name:          "One element, first",
			map_:          NewFromSlice[int](utils.BasicComparator[int], []int{1}),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Heap[int]).OrderedFirst,
		},
		{
			name:          "One element, last",
			map_:          NewFromSlice[int](utils.BasicComparator[int], []int{1}),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Heap[int]).OrderedLast,
		},
		{
			name:          "3 elements, middle",
			map_:          NewFromSlice[int](utils.BasicComparator[int], []int{1, 2, 3}),
			position:      1,
			n:             1,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Heap[int]).OrderedBegin,
		},
		{
			name:          "5 elements, middle, forward by 2",
			map_:          NewFromSlice[int](utils.BasicComparator[int], []int{1, 2, 3, 4, 5}),
			position:      2,
			n:             2,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Heap[int]).OrderedBegin,
		},
		{
			name:          "5 elements, middle, backward by 2",
			map_:          NewFromSlice[int](utils.BasicComparator[int], []int{1, 2, 3, 4, 5}),
			position:      2,
			n:             -2,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Heap[int]).OrderedBegin,
		},
		{
			name:          "3 elements, middle, move out of bounds",
			map_:          NewFromSlice[int](utils.BasicComparator[int], []int{1, 2, 3}),
			position:      1,
			n:             5,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Heap[int]).OrderedBegin,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)

			it := test.iteratorInit(test.map_)

			if test.position != NoMoveMagicPosition {
				it.MoveTo(test.position)
			}

			isValidBefore := it.IsValid()
			assert.Equalf(t, test.isValidBefore, isValidBefore, test.name)

			it.MoveBy(test.n)

			isValidAfter := it.IsValid()
			assert.Equalf(t, test.isValidAfter, isValidAfter, test.name)
		})
	}
}

func TestHeapOrderedIteratorMoveTo(t *testing.T) {
	tests := []struct {
		name         string
		map_         *Heap[int]
		position     int
		isValidAfter bool
		index        int
	}{
		{
			name:         "Empty",
			map_:         New[int](utils.BasicComparator[int]),
			isValidAfter: false,
		},
		{
			name:         "3 elements, first item",
			map_:         NewFromSlice[int](utils.BasicComparator[int], []int{1, 2, 3}),
			isValidAfter: true,
			index:        0,
		},
		{
			name:         "3 elements, middle item",
			map_:         NewFromSlice[int](utils.BasicComparator[int], []int{1, 2, 3}),
			isValidAfter: true,
			index:        1,
		},
		{
			name:         "3 elements, last item",
			map_:         NewFromSlice[int](utils.BasicComparator[int], []int{1, 2, 3}),
			isValidAfter: true,
			index:        2,
		},
	}

	for _, test := range tests {
		test := test

		testNameOrig := test.name
		for _, iteratorInit := range []struct {
			name string
			f    func(*Heap[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
		}{
			{", from first,", (*Heap[int]).OrderedFirst},
			{", from last,", (*Heap[int]).OrderedLast},
			{", from begin,", (*Heap[int]).OrderedBegin},
			{", from end,", (*Heap[int]).OrderedEnd},
		} {
			test.name = testNameOrig + iteratorInit.name

			t.Run(test.name, func(t *testing.T) {
				t.Parallel()
				defer testCommon.HandlePanic(t, test.name)

				repr := test.map_.ToString()
				assert.NotEmpty(t, repr)

				it := iteratorInit.f(test.map_)

				it.MoveTo(test.index)

				isValidAfter := it.IsValid()
				assert.Equalf(t, test.isValidAfter, isValidAfter, test.name+" valid after")

				if test.isValidAfter {
					index, _ := it.Index()
					assert.Equalf(t, test.index, index, test.name+" index")
				}
			})
		}
	}
}

func TestHeapOrderedIteratorIsBeginEndFirstLast(t *testing.T) {
	tests := []struct {
		name          string
		iteratorInit  func(*Heap[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
		iteratorCheck func(ds.ReadWriteOrdCompBidRandCollIterator[int, int]) bool
	}{
		{
			name:          "OrderedFirst",
			iteratorInit:  (*Heap[int]).OrderedFirst,
			iteratorCheck: (ds.ReadWriteOrdCompBidRandCollIterator[int, int]).IsFirst,
		}, {
			name:          "OrderedLast",
			iteratorInit:  (*Heap[int]).OrderedLast,
			iteratorCheck: (ds.ReadWriteOrdCompBidRandCollIterator[int, int]).IsLast,
		},
		{
			name:          "OrderedBegin",
			iteratorInit:  (*Heap[int]).OrderedBegin,
			iteratorCheck: (ds.ReadWriteOrdCompBidRandCollIterator[int, int]).IsBegin,
		},
		{
			name:          "OrderedEnd",
			iteratorInit:  (*Heap[int]).OrderedEnd,
			iteratorCheck: (ds.ReadWriteOrdCompBidRandCollIterator[int, int]).IsEnd,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(NewFromSlice[int](utils.BasicComparator[int], []int{1, 2, 4, 5}))
			assert.Truef(t, test.iteratorCheck(it), test.name)
		})
	}
}
