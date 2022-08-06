package doublylinkedlist

import (
	"testing"

	testCommon "github.com/JonasMuehlmann/datastructures.go/tests"

	"github.com/JonasMuehlmann/datastructures.go/ds"

	"github.com/stretchr/testify/assert"
)

const (
	NoMoveMagicPosition = 7869543205234798
)

func TestDoublyLinkedlistIteratorIsValid(t *testing.T) {
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
			iteratorInit: (*List[int]).First,
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
			iteratorInit: (*List[int]).First,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.list)

			if test.position != NoMoveMagicPosition {
				it.MoveTo(test.position)
			}

			isValid := it.IsValid()

			assert.Equalf(t, test.isValid, isValid, test.name)
		})
	}
}

func TestDoublyLinkedlistIteratorIndex(t *testing.T) {
	tests := []struct {
		name         string
		list         *List[int]
		position     int
		iteratorInit func(*List[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
	}{
		{
			name:         "Empty",
			list:         New[int](),
			position:     0,
			iteratorInit: (*List[int]).First,
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
			name:         "3 elements, first",
			list:         New[int](1, 2, 3),
			position:     0,
			iteratorInit: (*List[int]).First,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.list)

			position, valid := it.Index()

			assert.Equalf(t, test.position, position, test.name)
			assert.Truef(t, valid, test.name)
		})
	}
}

func TestDoublyLinkedlistIteratorNext(t *testing.T) {
	tests := []struct {
		name          string
		list          *List[int]
		position      int
		isValidBefore bool
		isValidAfter  bool
		value         int
		iteratorInit  func(*List[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
	}{
		{
			name:          "Empty",
			list:          New[int](),
			position:      NoMoveMagicPosition,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*List[int]).First,
		},
		{
			name:          "One element, begin",
			list:          New[int](1),
			position:      NoMoveMagicPosition,
			isValidBefore: false,
			isValidAfter:  true,
			value:         1,
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
			value:         3,
			iteratorInit:  (*List[int]).Begin,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.list)

			if test.position != NoMoveMagicPosition {
				it.MoveTo(test.position)
			}

			isValidBefore := it.IsValid()
			assert.Equalf(t, test.isValidBefore, isValidBefore, test.name)

			it.Next()
			value, _ := it.Get()

			isValidAfter := it.IsValid()
			assert.Equalf(t, test.isValidAfter, isValidAfter, test.name)
			assert.Equalf(t, test.value, value, test.name)
		})
	}
}

func TestDoublyLinkedlistIteratorNextN(t *testing.T) {
	tests := []struct {
		name          string
		list          *List[int]
		position      int
		n             int
		isValidBefore bool
		isValidAfter  bool
		value         int
		iteratorInit  func(*List[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
	}{
		{
			name:          "Empty",
			list:          New[int](),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*List[int]).First,
		},
		{
			name:          "One element, begin",
			list:          New[int](1),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: false,
			isValidAfter:  true,
			value:         1,
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
			value:         3,
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
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.list)

			if test.position != NoMoveMagicPosition {
				it.MoveTo(test.position)
			}

			isValidBefore := it.IsValid()
			assert.Equalf(t, test.isValidBefore, isValidBefore, test.name)

			it.NextN(test.n)
			value, _ := it.Get()

			isValidAfter := it.IsValid()
			assert.Equalf(t, test.isValidAfter, isValidAfter, test.name)
			assert.Equalf(t, test.value, value, test.name)
		})
	}
}

func TestDoublyLinkedlistIteratorPrevious(t *testing.T) {
	tests := []struct {
		name          string
		list          *List[int]
		position      int
		isValidBefore bool
		isValidAfter  bool
		value         int
		iteratorInit  func(*List[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
	}{
		{
			name:          "Empty",
			list:          New[int](),
			position:      NoMoveMagicPosition,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*List[int]).First,
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
			value:         1,
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
			value:         1,
			iteratorInit:  (*List[int]).Begin,
		},
		{
			name:          "5 elements, middle",
			list:          New[int](1, 2, 3, 4, 5),
			position:      2,
			isValidBefore: true,
			isValidAfter:  true,
			value:         2,
			iteratorInit:  (*List[int]).Begin,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.list)

			if test.position != NoMoveMagicPosition {
				it.MoveTo(test.position)
			}

			isValidBefore := it.IsValid()
			assert.Equalf(t, test.isValidBefore, isValidBefore, test.name+" valid before")

			it.Previous()
			value, _ := it.Get()

			isValidAfter := it.IsValid()
			assert.Equalf(t, test.isValidAfter, isValidAfter, test.name+" valid after")
			assert.Equalf(t, test.value, value, test.name+" value")
		})
	}
}

func TestDoublyLinkedlistIteratorPreviousN(t *testing.T) {
	tests := []struct {
		name          string
		list          *List[int]
		position      int
		n             int
		isValidBefore bool
		isValidAfter  bool
		value         int
		iteratorInit  func(*List[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
	}{
		{
			name:          "Empty",
			list:          New[int](),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*List[int]).First,
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
			value:         1,
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
			value:         1,
			iteratorInit:  (*List[int]).Begin,
		},
		{
			name:          "5 elements, middle",
			list:          New[int](1, 2, 3, 4, 5),
			position:      2,
			n:             1,
			isValidBefore: true,
			isValidAfter:  true,
			value:         2,
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
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.list)

			if test.position != NoMoveMagicPosition {
				it.MoveTo(test.position)
			}

			isValidBefore := it.IsValid()
			assert.Equalf(t, test.isValidBefore, isValidBefore, test.name)

			it.PreviousN(test.n)
			value, _ := it.Get()

			isValidAfter := it.IsValid()
			assert.Equalf(t, test.isValidAfter, isValidAfter, test.name)
			assert.Equalf(t, test.value, value, test.name)
		})
	}
}

func TestDoublyLinkedlistIteratorGet(t *testing.T) {
	tests := []struct {
		name         string
		list         *List[int]
		position     int
		value        int
		found        bool
		iteratorInit func(*List[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
	}{
		{
			name:         "Empty",
			list:         New[int](),
			position:     NoMoveMagicPosition,
			found:        false,
			iteratorInit: (*List[int]).First,
		},
		{
			name:         "One element, first",
			list:         New[int](1),
			position:     0,
			value:        1,
			found:        true,
			iteratorInit: (*List[int]).First,
		},
		{
			name:         "Three elements, begin",
			list:         New[int](1, 2, 3),
			position:     1,
			value:        2,
			found:        true,
			iteratorInit: (*List[int]).Begin,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it := test.list.First()

			if test.position != NoMoveMagicPosition {
				it.MoveTo(test.position)
			}

			value, found := it.Get()

			assert.Equalf(t, test.found, found, test.name)
			if test.found {
				assert.Equalf(t, test.value, value, test.name)
			}
		})
	}
}

func TestDoublyLinkedlistIteratorGetAt(t *testing.T) {
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
			name:     "One element, first",
			list:     New[int](1),
			position: 0,
			value:    1,
			found:    true,
		},
		{
			name:     "Three elements, first",
			list:     New[int](1, 2, 3),
			position: 0,
			value:    1,
			found:    true,
		},
		{
			name:     "Three elements, middle",
			list:     New[int](1, 2, 3),
			position: 1,
			value:    2,
			found:    true,
		},
		{
			name:     "Three elements, last",
			list:     New[int](1, 2, 3),
			position: 2,
			value:    3,
			found:    true,
		},
		{
			name:     "Three elements, out of bounds left",
			list:     New[int](1, 2, 3),
			position: -3,
			found:    false,
		},
		{
			name:     "Three elements, out of bounds right",
			list:     New[int](1, 2, 3),
			position: 5,
			found:    false,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it := test.list.First()

			value, found := it.GetAt(test.position)

			assert.Equalf(t, test.found, found, test.name)
			if test.found {
				assert.Equalf(t, test.value, value, test.name)
			}
		})
	}
}

func TestDoublyLinkedlistIteratorSet(t *testing.T) {
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
			name:        "One element, first",
			list:        New[int](1),
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
			it := test.list.First()
			successfull := it.Set(test.value)

			assert.Equalf(t, test.successfull, successfull, test.name)
		})
	}
}

func TestDoublyLinkedlistIteratorSetAt(t *testing.T) {
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
			name:     "One element, first",
			list:     New[int](1),
			position: 0,
			value:    1,
			found:    true,
		},
		{
			name:     "Three elements, first",
			list:     New[int](1, 2, 3),
			position: 0,
			value:    1,
			found:    true,
		},
		{
			name:     "Three elements, middle",
			list:     New[int](1, 2, 3),
			position: 1,
			value:    2,
			found:    true,
		},
		{
			name:     "Three elements, last",
			list:     New[int](1, 2, 3),
			position: 2,
			value:    3,
			found:    true,
		},
		{
			name:     "Three elements, out of bounds left",
			list:     New[int](1, 2, 3),
			position: -3,
			found:    false,
		},
		{
			name:     "Three elements, out of bounds right",
			list:     New[int](1, 2, 3),
			position: 5,
			found:    false,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it := test.list.First()

			value, found := it.GetAt(test.position)

			assert.Equalf(t, test.found, found, test.name)
			if test.found {
				assert.Equalf(t, test.value, value, test.name)
			}
		})
	}
}

// NOTE: Missing test case: other does not implement IndexedIterator
func TestDoublyLinkedlistIteratorDistanceTo(t *testing.T) {
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
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it1 := New[int](1, 2, 3, 4, 5).First()
			it2 := New[int](1, 2, 3, 4, 5).First()

			it1.MoveTo(test.position1)
			it2.MoveTo(test.position2)

			distance := it1.DistanceTo(it2)

			assert.Equalf(t, test.distance, distance, test.name)
		})
	}
}

func TestDoublyLinkedlistIteratorIsAfter(t *testing.T) {
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
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it1 := New[int](1, 2, 3, 4, 5).First()
			it2 := New[int](1, 2, 3, 4, 5).First()

			it1.MoveTo(test.position1)
			it2.MoveTo(test.position2)

			isAfter := it1.IsAfter(it2)

			assert.Equalf(t, test.isAfter, isAfter, test.name)
		})
	}
}

func TestDoublyLinkedlistIteratorIsBefore(t *testing.T) {
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
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it1 := New[int](1, 2, 3, 4, 5).First()
			it2 := New[int](1, 2, 3, 4, 5).First()

			it1.MoveTo(test.position1)
			it2.MoveTo(test.position2)

			isAfter := it1.IsBefore(it2)

			assert.Equalf(t, test.isAfter, isAfter, test.name)
		})
	}
}

func TestDoublyLinkedlistIteratorIsEqual(t *testing.T) {
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
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it1 := New[int](1, 2, 3, 4, 5).First()
			it2 := New[int](1, 2, 3, 4, 5).First()

			it1.MoveTo(test.position1)
			it2.MoveTo(test.position2)

			isAfter := it1.IsEqual(it2)

			assert.Equalf(t, test.isAfter, isAfter, test.name)
		})
	}
}

func TestDoublylInkedlistIteratorIsBeginEndFirstLast(t *testing.T) {
	tests := []struct {
		name          string
		iteratorInit  func(*List[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
		iteratorCheck func(ds.ReadWriteOrdCompBidRandCollIterator[int, int]) bool
	}{
		{
			name:          "Last",
			iteratorInit:  (*List[int]).Last,
			iteratorCheck: (ds.ReadWriteOrdCompBidRandCollIterator[int, int]).IsLast,
		},
		{
			name:          "First",
			iteratorInit:  (*List[int]).First,
			iteratorCheck: (ds.ReadWriteOrdCompBidRandCollIterator[int, int]).IsFirst,
		},
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
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(New[int](1, 2, 4, 5))
			assert.Truef(t, test.iteratorCheck(it), test.name)
		})
	}
}

func TestDoublyLinkedlistIteratorSize(t *testing.T) {
	tests := []struct {
		name         string
		list         *List[int]
		iteratorInit func(*List[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
		size         int
	}{
		{
			name:         "Empty",
			list:         New[int](),
			size:         0,
			iteratorInit: (*List[int]).First,
		},

		{
			name:         "One element, first",
			list:         New[int](1),
			size:         1,
			iteratorInit: (*List[int]).First,
		},

		{
			name:         "3 elements, middle",
			list:         New[int](1, 2, 3),
			size:         3,
			iteratorInit: (*List[int]).First,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.list)

			size := it.Size()

			assert.Equalf(t, test.size, size, test.name)
		})
	}
}
