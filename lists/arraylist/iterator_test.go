package arraylist

import (
	"testing"

	"github.com/JonasMuehlmann/datastructures.go/ds"
	"github.com/stretchr/testify/assert"
)

const (
	NoMoveMagicPosition = 7869543205234798
)

func TestArrayListIteratorIsValid(t *testing.T) {
	tests := []struct {
		name         string
		list         *List[int]
		position     int
		isValid      bool
		iteratorInit func(*List[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
	}{
		{
			name:         "Empty",
			list:         New[int](),
			position:     NoMoveMagicPosition,
			isValid:      false,
			iteratorInit: (*List[int]).Begin,
		},
		{
			name:         "One element, begin",
			list:         New[int](1),
			position:     NoMoveMagicPosition,
			isValid:      false,
			iteratorInit: (*List[int]).Begin,
		},
		{
			name:         "One element, end",
			list:         New[int](1),
			position:     NoMoveMagicPosition,
			isValid:      false,
			iteratorInit: (*List[int]).End,
		},
		{
			name:         "One element, first",
			list:         New[int](1),
			position:     NoMoveMagicPosition,
			isValid:      true,
			iteratorInit: (*List[int]).First,
		},
		{
			name:         "One element, last",
			list:         New[int](1),
			position:     NoMoveMagicPosition,
			isValid:      true,
			iteratorInit: (*List[int]).Last,
		},
		{
			name:         "3 elements, middle",
			list:         New[int](1, 2, 3),
			position:     1,
			isValid:      true,
			iteratorInit: (*List[int]).Begin,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			it := test.iteratorInit(test.list)

			if test.position != NoMoveMagicPosition {
				it.MoveTo(test.position)
			}

			isValid := it.IsValid()

			assert.Equalf(t, test.isValid, isValid, test.name)
		})
	}
}

func TestArrayListIteratorIndex(t *testing.T) {
	tests := []struct {
		name         string
		list         *List[int]
		position     int
		iteratorInit func(*List[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
	}{
		{
			name:         "Empty",
			list:         New[int](),
			position:     -1,
			iteratorInit: (*List[int]).Begin,
		},
		{
			name:         "One element, begin",
			list:         New[int](1),
			position:     -1,
			iteratorInit: (*List[int]).Begin,
		},
		{
			name:         "One element, end",
			list:         New[int](1),
			position:     1,
			iteratorInit: (*List[int]).End,
		},
		{
			name:         "One element, first",
			list:         New[int](1),
			position:     0,
			iteratorInit: (*List[int]).First,
		},
		{
			name:         "One element, last",
			list:         New[int](1),
			position:     0,
			iteratorInit: (*List[int]).Last,
		},
		{
			name:         "3 elements, middle",
			list:         New[int](1, 2, 3),
			position:     -1,
			iteratorInit: (*List[int]).Begin,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			it := test.iteratorInit(test.list)

			position, valid := it.Index()

			assert.Equalf(t, test.position, position, test.name)
			assert.Truef(t, valid, test.name)
		})
	}
}

func TestArrayListIteratorNext(t *testing.T) {
	tests := []struct {
		name          string
		list          *List[int]
		position      int
		isValidBefore bool
		isValidAfter  bool
		iteratorInit  func(*List[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
	}{
		{
			name:          "Empty",
			list:          New[int](),
			position:      NoMoveMagicPosition,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*List[int]).Begin,
		},
		{
			name:          "One element, begin",
			list:          New[int](1),
			position:      NoMoveMagicPosition,
			isValidBefore: false,
			isValidAfter:  true,
			iteratorInit:  (*List[int]).Begin,
		},
		{
			name:          "One element, end",
			list:          New[int](1),
			position:      NoMoveMagicPosition,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*List[int]).End,
		},
		{
			name:          "One element, first",
			list:          New[int](1),
			position:      NoMoveMagicPosition,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*List[int]).First,
		},
		{
			name:          "One element, last",
			list:          New[int](1),
			position:      NoMoveMagicPosition,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*List[int]).Last,
		},
		{
			name:          "3 elements, middle",
			list:          New[int](1, 2, 3),
			position:      1,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*List[int]).Begin,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			it := test.iteratorInit(test.list)

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

func TestArrayListIteratorNextN(t *testing.T) {
	tests := []struct {
		name          string
		list          *List[int]
		position      int
		n             int
		isValidBefore bool
		isValidAfter  bool
		iteratorInit  func(*List[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
	}{
		{
			name:          "Empty",
			list:          New[int](),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*List[int]).Begin,
		},
		{
			name:          "One element, begin",
			list:          New[int](1),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: false,
			isValidAfter:  true,
			iteratorInit:  (*List[int]).Begin,
		},
		{
			name:          "One element, end",
			list:          New[int](1),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*List[int]).End,
		},
		{
			name:          "One element, first",
			list:          New[int](1),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*List[int]).First,
		},
		{
			name:          "One element, last",
			list:          New[int](1),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*List[int]).Last,
		},
		{
			name:          "3 elements, middle",
			list:          New[int](1, 2, 3),
			position:      1,
			n:             1,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*List[int]).Begin,
		},
		{
			name:          "3 elements, middle, move out of bounds",
			list:          New[int](1, 2, 3),
			position:      1,
			n:             5,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*List[int]).Begin,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			it := test.iteratorInit(test.list)

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

func TestArrayListIteratorPrevious(t *testing.T) {
	tests := []struct {
		name          string
		list          *List[int]
		position      int
		isValidBefore bool
		isValidAfter  bool
		iteratorInit  func(*List[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
	}{
		{
			name:          "Empty",
			list:          New[int](),
			position:      NoMoveMagicPosition,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*List[int]).Begin,
		},
		{
			name:          "One element, begin",
			list:          New[int](1),
			position:      NoMoveMagicPosition,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*List[int]).Begin,
		},
		{
			name:          "One element, end",
			list:          New[int](1),
			position:      NoMoveMagicPosition,
			isValidBefore: false,
			isValidAfter:  true,
			iteratorInit:  (*List[int]).End,
		},
		{
			name:          "One element, first",
			list:          New[int](1),
			position:      NoMoveMagicPosition,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*List[int]).First,
		},
		{
			name:          "One element, last",
			list:          New[int](1),
			position:      NoMoveMagicPosition,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*List[int]).Last,
		},
		{
			name:          "3 elements, middle",
			list:          New[int](1, 2, 3),
			position:      1,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*List[int]).Begin,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			it := test.iteratorInit(test.list)

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

func TestArrayListIteratorPreviousN(t *testing.T) {
	tests := []struct {
		name          string
		list          *List[int]
		position      int
		n             int
		isValidBefore bool
		isValidAfter  bool
		iteratorInit  func(*List[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
	}{
		{
			name:          "Empty",
			list:          New[int](),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*List[int]).Begin,
		},
		{
			name:          "One element, begin",
			list:          New[int](1),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*List[int]).Begin,
		},
		{
			name:          "One element, end",
			list:          New[int](1),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: false,
			isValidAfter:  true,
			iteratorInit:  (*List[int]).End,
		},
		{
			name:          "One element, first",
			list:          New[int](1),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*List[int]).First,
		},
		{
			name:          "One element, last",
			list:          New[int](1),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*List[int]).Last,
		},
		{
			name:          "3 elements, middle",
			list:          New[int](1, 2, 3),
			position:      1,
			n:             1,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*List[int]).Begin,
		},
		{
			name:          "3 elements, middle, move out of bounds",
			list:          New[int](1, 2, 3),
			position:      1,
			n:             5,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*List[int]).Begin,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			it := test.iteratorInit(test.list)

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

func TestArrayListIteratorMoveBy(t *testing.T) {
	tests := []struct {
		name          string
		list          *List[int]
		position      int
		n             int
		isValidBefore bool
		isValidAfter  bool
		iteratorInit  func(*List[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
	}{
		{
			name:          "Empty",
			list:          New[int](),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*List[int]).Begin,
		},
		{
			name:          "One element, begin",
			list:          New[int](1),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: false,
			isValidAfter:  true,
			iteratorInit:  (*List[int]).Begin,
		},
		{
			name:          "One element, end",
			list:          New[int](1),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*List[int]).End,
		},
		{
			name:          "One element, first",
			list:          New[int](1),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*List[int]).First,
		},
		{
			name:          "One element, last",
			list:          New[int](1),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*List[int]).Last,
		},
		{
			name:          "3 elements, middle",
			list:          New[int](1, 2, 3),
			position:      1,
			n:             1,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*List[int]).Begin,
		},
		{
			name:          "5 elements, middle, forward by 2",
			list:          New[int](1, 2, 3, 4, 5),
			position:      2,
			n:             2,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*List[int]).Begin,
		}, {
			name:          "5 elements, middle, backward by 2",
			list:          New[int](1, 2, 3, 4, 5),
			position:      2,
			n:             -2,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*List[int]).Begin,
		},
		{
			name:          "3 elements, middle, move out of bounds",
			list:          New[int](1, 2, 3),
			position:      1,
			n:             5,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*List[int]).Begin,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			it := test.iteratorInit(test.list)

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

func TestArrayListIteratorGet(t *testing.T) {
	tests := []struct {
		name     string
		list     *List[int]
		position int
		value    int
		found    bool
	}{
		{
			name:     "Empty",
			list:     New[int](),
			position: NoMoveMagicPosition,
			found:    false,
		},
		{
			name:     "One element, begin",
			list:     New[int](1),
			position: NoMoveMagicPosition,
			found:    false,
		},
		{
			name:     "One element, first",
			list:     New[int](1),
			position: 0,
			value:    1,
			found:    true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			it := test.list.Begin()

			if test.position != NoMoveMagicPosition {
				it.MoveTo(test.position)
			}

			value, found := it.Get()

			assert.Equalf(t, test.found, found, test.name)
			assert.Equalf(t, test.value, value, test.name)
		})
	}
}

func TestArrayListIteratorSet(t *testing.T) {
	tests := []struct {
		name        string
		list        *List[int]
		position    int
		value       int
		successfull bool
	}{
		{
			name:        "Empty",
			list:        New[int](),
			position:    NoMoveMagicPosition,
			value:       1,
			successfull: false,
		},
		{
			name:        "One element, begin",
			list:        New[int](1),
			position:    NoMoveMagicPosition,
			value:       1,
			successfull: false,
		},
		{
			name:        "One element, first",
			list:        New[int](1),
			position:    0,
			value:       1,
			successfull: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			it := test.list.Begin()

			if test.position != NoMoveMagicPosition {
				it.MoveTo(test.position)
			}

			successfull := it.Set(test.value)

			assert.Equalf(t, test.successfull, successfull, test.name)
		})
	}
}

func TestArrayListIteratorGetAt(t *testing.T) {
	tests := []struct {
		name     string
		list     *List[int]
		position int
		value    int
		found    bool
	}{
		{
			name:     "Empty",
			list:     New[int](),
			position: 0,
			found:    false,
		},
		{
			name:     "One element, begin",
			list:     New[int](1),
			position: -1,
			found:    false,
		},
		{
			name:     "One element, first",
			list:     New[int](1),
			position: 0,
			value:    1,
			found:    true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			it := test.list.Begin()

			value, found := it.GetAt(test.position)

			assert.Equalf(t, test.found, found, test.name)
			assert.Equalf(t, test.value, value, test.name)
		})
	}
}

func TestArrayListIteratorSetAt(t *testing.T) {
	tests := []struct {
		name        string
		list        *List[int]
		position    int
		value       int
		successfull bool
	}{
		{
			name:        "Empty",
			list:        New[int](),
			position:    0,
			value:       1,
			successfull: false,
		},
		{
			name:        "One element, begin",
			list:        New[int](1),
			position:    -1,
			value:       1,
			successfull: false,
		},
		{
			name:        "One element, first",
			list:        New[int](1),
			position:    0,
			value:       1,
			successfull: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			it := test.list.Begin()

			successfull := it.SetAt(test.position, test.value)

			assert.Equalf(t, test.successfull, successfull, test.name)
		})
	}
}

// NOTE: Missing test case: other does not implement IndexedIterator
func TestArrayListIteratorDistanceTo(t *testing.T) {
	tests := []struct {
		name      string
		position1 int
		position2 int
		distance  int
	}{
		{
			name:      "Equal",
			position1: 0,
			position2: 0,
			distance:  0,
		},
		{
			name:      "First lower",
			position1: 0,
			position2: 1,
			distance:  -1,
		},
		{
			name:      "Second lower",
			position1: 1,
			position2: 0,
			distance:  1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			it1 := New[int]().Begin()
			it2 := New[int]().Begin()

			it1.MoveTo(test.position1)
			it2.MoveTo(test.position2)

			distance := it1.DistanceTo(it2)

			assert.Equalf(t, test.distance, distance, test.name)
		})
	}
}

func TestArrayListIteratorIsAfter(t *testing.T) {
	tests := []struct {
		name      string
		position1 int
		position2 int
		isAfter   bool
	}{
		{
			name:      "Equal",
			position1: 0,
			position2: 0,
			isAfter:   false,
		},
		{
			name:      "First lower",
			position1: 0,
			position2: 1,
			isAfter:   false,
		},
		{
			name:      "Second lower",
			position1: 1,
			position2: 0,
			isAfter:   true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			it1 := New[int]().Begin()
			it2 := New[int]().Begin()

			it1.MoveTo(test.position1)
			it2.MoveTo(test.position2)

			isAfter := it1.IsAfter(it2)

			assert.Equalf(t, test.isAfter, isAfter, test.name)
		})
	}
}

func TestArrayListIteratorIsBefore(t *testing.T) {
	tests := []struct {
		name      string
		position1 int
		position2 int
		isAfter   bool
	}{
		{
			name:      "Equal",
			position1: 0,
			position2: 0,
			isAfter:   false,
		},
		{
			name:      "First lower",
			position1: 0,
			position2: 1,
			isAfter:   true,
		},
		{
			name:      "Second lower",
			position1: 1,
			position2: 0,
			isAfter:   false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			it1 := New[int]().Begin()
			it2 := New[int]().Begin()

			it1.MoveTo(test.position1)
			it2.MoveTo(test.position2)

			isAfter := it1.IsBefore(it2)

			assert.Equalf(t, test.isAfter, isAfter, test.name)
		})
	}
}

func TestArrayListIteratorIsEqual(t *testing.T) {
	tests := []struct {
		name      string
		position1 int
		position2 int
		isAfter   bool
	}{
		{
			name:      "Equal",
			position1: 0,
			position2: 0,
			isAfter:   true,
		},
		{
			name:      "First lower",
			position1: 0,
			position2: 1,
			isAfter:   false,
		},
		{
			name:      "Second lower",
			position1: 1,
			position2: 0,
			isAfter:   false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			it1 := New[int]().Begin()
			it2 := New[int]().Begin()

			it1.MoveTo(test.position1)
			it2.MoveTo(test.position2)

			isAfter := it1.IsEqual(it2)

			assert.Equalf(t, test.isAfter, isAfter, test.name)
		})
	}
}

func TestArrayListIteratorIsBeginEndFirstLast(t *testing.T) {
	tests := []struct {
		name          string
		iteratorInit  func(*List[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
		iteratorCheck func(ds.ReadWriteOrdCompBidRandCollIterator[int, int]) bool
	}{
		{
			name:          "Begin",
			iteratorInit:  (*List[int]).Begin,
			iteratorCheck: (ds.ReadWriteOrdCompBidRandCollIterator[int, int]).IsBegin,
		},
		{
			name:          "End",
			iteratorInit:  (*List[int]).End,
			iteratorCheck: (ds.ReadWriteOrdCompBidRandCollIterator[int, int]).IsEnd,
		},
		{
			name:          "First",
			iteratorInit:  (*List[int]).First,
			iteratorCheck: (ds.ReadWriteOrdCompBidRandCollIterator[int, int]).IsFirst,
		},
		{
			name:          "Last",
			iteratorInit:  (*List[int]).Last,
			iteratorCheck: (ds.ReadWriteOrdCompBidRandCollIterator[int, int]).IsLast,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			it := test.iteratorInit(New[int](1, 2, 4, 5))
			assert.Truef(t, test.iteratorCheck(it), test.name)
		})
	}
}
