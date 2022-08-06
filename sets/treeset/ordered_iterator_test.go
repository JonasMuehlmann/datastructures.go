package treeset

import (
	"testing"

	"github.com/JonasMuehlmann/datastructures.go/ds"
	testCommon "github.com/JonasMuehlmann/datastructures.go/tests"
	"github.com/JonasMuehlmann/datastructures.go/utils"
	"github.com/stretchr/testify/assert"
)

const (
	NoMoveMagicPosition = 7869543205234798
)

func TestArrayListIteratorIsValid(t *testing.T) {
	tests := []struct {
		name         string
		list         *Set[int]
		position     int
		isValid      bool
		iteratorInit func(*Set[int], utils.Comparator[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
	}{
		{
			name:         "Empty",
			list:         New[int](utils.BasicComparator[int]),
			position:     NoMoveMagicPosition,
			isValid:      false,
			iteratorInit: (*Set[int]).OrderedBegin,
		},
		{
			name:         "One element, begin",
			list:         New[int](utils.BasicComparator[int], 1),
			position:     NoMoveMagicPosition,
			isValid:      false,
			iteratorInit: (*Set[int]).OrderedBegin,
		},
		{
			name:         "One element, end",
			list:         New[int](utils.BasicComparator[int], 1),
			position:     NoMoveMagicPosition,
			isValid:      false,
			iteratorInit: (*Set[int]).OrderedEnd,
		},
		{
			name:         "One element, first",
			list:         New[int](utils.BasicComparator[int], 1),
			position:     NoMoveMagicPosition,
			isValid:      true,
			iteratorInit: (*Set[int]).OrderedFirst,
		},
		{
			name:         "One element, last",
			list:         New[int](utils.BasicComparator[int], 1),
			position:     NoMoveMagicPosition,
			isValid:      true,
			iteratorInit: (*Set[int]).OrderedLast,
		},
		{
			name:         "3 elements, middle",
			list:         New[int](utils.BasicComparator[int], 1, 2, 3),
			position:     1,
			isValid:      true,
			iteratorInit: (*Set[int]).OrderedBegin,
		},
	}

	for _, test := range tests {
test := test

		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.list, utils.BasicComparator[int])

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
		list         *Set[int]
		position     int
		isValid      bool
		iteratorInit func(*Set[int], utils.Comparator[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
	}{
		{
			name:         "Empty",
			list:         New[int](utils.BasicComparator[int]),
			position:     -1,
			isValid:      false,
			iteratorInit: (*Set[int]).OrderedBegin,
		},
		{
			name:         "One element, begin",
			list:         New[int](utils.BasicComparator[int], 1),
			position:     -1,
			isValid:      false,
			iteratorInit: (*Set[int]).OrderedBegin,
		},
		{
			name:         "One element, end",
			list:         New[int](utils.BasicComparator[int], 1),
			position:     1,
			isValid:      false,
			iteratorInit: (*Set[int]).OrderedEnd,
		},
		{
			name:         "One element, first",
			list:         New[int](utils.BasicComparator[int], 1),
			position:     0,
			isValid:      true,
			iteratorInit: (*Set[int]).OrderedFirst,
		},
		{
			name:         "One element, last",
			list:         New[int](utils.BasicComparator[int], 1),
			position:     0,
			isValid:      true,
			iteratorInit: (*Set[int]).OrderedLast,
		},
		{
			name:         "3 elements, middle",
			list:         New[int](utils.BasicComparator[int], 1, 2, 3),
			position:     -1,
			isValid:      false,
			iteratorInit: (*Set[int]).OrderedBegin,
		},
	}

	for _, test := range tests {
test := test

		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.list, utils.BasicComparator[int])

			position, valid := it.Index()

			assert.Equalf(t, test.isValid, valid, test.name)
			if test.isValid {
				assert.Equalf(t, test.position, position, test.name)
			}
		})
	}
}

func TestArrayListIteratorNext(t *testing.T) {
	tests := []struct {
		name          string
		list          *Set[int]
		position      int
		isValidBefore bool
		isValidAfter  bool
		iteratorInit  func(*Set[int], utils.Comparator[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
	}{
		{
			name:          "Empty",
			list:          New[int](utils.BasicComparator[int]),
			position:      NoMoveMagicPosition,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Set[int]).OrderedBegin,
		},
		{
			name:          "One element, begin",
			list:          New[int](utils.BasicComparator[int], 1),
			position:      NoMoveMagicPosition,
			isValidBefore: false,
			isValidAfter:  true,
			iteratorInit:  (*Set[int]).OrderedBegin,
		},
		{
			name:          "One element, end",
			list:          New[int](utils.BasicComparator[int], 1),
			position:      NoMoveMagicPosition,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Set[int]).OrderedEnd,
		},
		{
			name:          "One element, first",
			list:          New[int](utils.BasicComparator[int], 1),
			position:      NoMoveMagicPosition,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Set[int]).OrderedFirst,
		},
		{
			name:          "One element, last",
			list:          New[int](utils.BasicComparator[int], 1),
			position:      NoMoveMagicPosition,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Set[int]).OrderedLast,
		},
		{
			name:          "3 elements, middle",
			list:          New[int](utils.BasicComparator[int], 1, 2, 3),
			position:      1,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Set[int]).OrderedBegin,
		},
	}

	for _, test := range tests {
test := test

		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.list, utils.BasicComparator[int])

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
		list          *Set[int]
		position      int
		n             int
		isValidBefore bool
		isValidAfter  bool
		iteratorInit  func(*Set[int], utils.Comparator[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
	}{
		{
			name:          "Empty",
			list:          New[int](utils.BasicComparator[int]),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Set[int]).OrderedBegin,
		},
		{
			name:          "One element, begin",
			list:          New[int](utils.BasicComparator[int], 1),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: false,
			isValidAfter:  true,
			iteratorInit:  (*Set[int]).OrderedBegin,
		},
		{
			name:          "One element, end",
			list:          New[int](utils.BasicComparator[int], 1),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Set[int]).OrderedEnd,
		},
		{
			name:          "One element, first",
			list:          New[int](utils.BasicComparator[int], 1),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Set[int]).OrderedFirst,
		},
		{
			name:          "One element, last",
			list:          New[int](utils.BasicComparator[int], 1),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Set[int]).OrderedLast,
		},
		{
			name:          "3 elements, middle",
			list:          New[int](utils.BasicComparator[int], 1, 2, 3),
			position:      1,
			n:             1,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Set[int]).OrderedBegin,
		},
		{
			name:          "3 elements, middle, move out of bounds",
			list:          New[int](utils.BasicComparator[int], 1, 2, 3),
			position:      1,
			n:             5,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Set[int]).OrderedBegin,
		},
	}

	for _, test := range tests {
test := test

		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.list, utils.BasicComparator[int])

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
		list          *Set[int]
		position      int
		isValidBefore bool
		isValidAfter  bool
		iteratorInit  func(*Set[int], utils.Comparator[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
	}{
		{
			name:          "Empty",
			list:          New[int](utils.BasicComparator[int]),
			position:      NoMoveMagicPosition,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Set[int]).OrderedBegin,
		},
		{
			name:          "One element, begin",
			list:          New[int](utils.BasicComparator[int], 1),
			position:      NoMoveMagicPosition,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Set[int]).OrderedBegin,
		},
		{
			name:          "One element, end",
			list:          New[int](utils.BasicComparator[int], 1),
			position:      NoMoveMagicPosition,
			isValidBefore: false,
			isValidAfter:  true,
			iteratorInit:  (*Set[int]).OrderedEnd,
		},
		{
			name:          "One element, first",
			list:          New[int](utils.BasicComparator[int], 1),
			position:      NoMoveMagicPosition,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Set[int]).OrderedFirst,
		},
		{
			name:          "One element, last",
			list:          New[int](utils.BasicComparator[int], 1),
			position:      NoMoveMagicPosition,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Set[int]).OrderedLast,
		},
		{
			name:          "3 elements, middle",
			list:          New[int](utils.BasicComparator[int], 1, 2, 3),
			position:      1,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Set[int]).OrderedBegin,
		},
	}

	for _, test := range tests {
test := test

		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.list, utils.BasicComparator[int])

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
		list          *Set[int]
		position      int
		n             int
		isValidBefore bool
		isValidAfter  bool
		iteratorInit  func(*Set[int], utils.Comparator[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
	}{
		{
			name:          "Empty",
			list:          New[int](utils.BasicComparator[int]),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Set[int]).OrderedBegin,
		},
		{
			name:          "One element, begin",
			list:          New[int](utils.BasicComparator[int], 1),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Set[int]).OrderedBegin,
		},
		{
			name:          "One element, end",
			list:          New[int](utils.BasicComparator[int], 1),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: false,
			isValidAfter:  true,
			iteratorInit:  (*Set[int]).OrderedEnd,
		},
		{
			name:          "One element, first",
			list:          New[int](utils.BasicComparator[int], 1),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Set[int]).OrderedFirst,
		},
		{
			name:          "One element, last",
			list:          New[int](utils.BasicComparator[int], 1),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Set[int]).OrderedLast,
		},
		{
			name:          "3 elements, middle",
			list:          New[int](utils.BasicComparator[int], 1, 2, 3),
			position:      1,
			n:             1,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Set[int]).OrderedBegin,
		},
		{
			name:          "3 elements, middle, move out of bounds",
			list:          New[int](utils.BasicComparator[int], 1, 2, 3),
			position:      1,
			n:             5,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Set[int]).OrderedBegin,
		},
	}

	for _, test := range tests {
test := test

		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.list, utils.BasicComparator[int])

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
		list          *Set[int]
		position      int
		n             int
		isValidBefore bool
		isValidAfter  bool
		iteratorInit  func(*Set[int], utils.Comparator[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
	}{
		{
			name:          "Empty",
			list:          New[int](utils.BasicComparator[int]),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Set[int]).OrderedBegin,
		},
		{
			name:          "One element, begin",
			list:          New[int](utils.BasicComparator[int], 1),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: false,
			isValidAfter:  true,
			iteratorInit:  (*Set[int]).OrderedBegin,
		},
		{
			name:          "One element, end",
			list:          New[int](utils.BasicComparator[int], 1),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Set[int]).OrderedEnd,
		},
		{
			name:          "One element, first",
			list:          New[int](utils.BasicComparator[int], 1),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Set[int]).OrderedFirst,
		},
		{
			name:          "One element, last",
			list:          New[int](utils.BasicComparator[int], 1),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Set[int]).OrderedLast,
		},
		{
			name:          "3 elements, middle",
			list:          New[int](utils.BasicComparator[int], 1, 2, 3),
			position:      1,
			n:             1,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Set[int]).OrderedBegin,
		},
		{
			name:          "5 elements, middle, forward by 2",
			list:          New[int](utils.BasicComparator[int], 1, 2, 3, 4, 5),
			position:      2,
			n:             2,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Set[int]).OrderedBegin,
		}, {
			name:          "5 elements, middle, backward by 2",
			list:          New[int](utils.BasicComparator[int], 1, 2, 3, 4, 5),
			position:      2,
			n:             -2,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Set[int]).OrderedBegin,
		},
		{
			name:          "3 elements, middle, move out of bounds",
			list:          New[int](utils.BasicComparator[int], 1, 2, 3),
			position:      1,
			n:             5,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Set[int]).OrderedBegin,
		},
	}

	for _, test := range tests {
test := test

		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.list, utils.BasicComparator[int])

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
		list     *Set[int]
		position int
		value    int
		found    bool
	}{
		{
			name:     "Empty",
			list:     New[int](utils.BasicComparator[int]),
			position: NoMoveMagicPosition,
			found:    false,
		},
		{
			name:     "One element, begin",
			list:     New[int](utils.BasicComparator[int], 1),
			position: NoMoveMagicPosition,
			found:    false,
		},
		{
			name:     "One element, first",
			list:     New[int](utils.BasicComparator[int], 1),
			position: 0,
			value:    1,
			found:    true,
		},
	}

	for _, test := range tests {
test := test

		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.list.OrderedBegin(utils.BasicComparator[int])

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
		list        *Set[int]
		position    int
		value       int
		successfull bool
	}{
		{
			name:        "Empty",
			list:        New[int](utils.BasicComparator[int]),
			position:    NoMoveMagicPosition,
			value:       1,
			successfull: false,
		},
		{
			name:        "One element, begin",
			list:        New[int](utils.BasicComparator[int], 1),
			position:    NoMoveMagicPosition,
			value:       1,
			successfull: false,
		},
		{
			name:        "One element, first",
			list:        New[int](utils.BasicComparator[int], 1),
			position:    0,
			value:       1,
			successfull: true,
		},
	}

	for _, test := range tests {
test := test

		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.list.OrderedBegin(utils.BasicComparator[int])

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
		list     *Set[int]
		position int
		value    int
		found    bool
	}{
		{
			name:     "Empty",
			list:     New[int](utils.BasicComparator[int]),
			position: 0,
			found:    false,
		},
		{
			name:     "One element, begin",
			list:     New[int](utils.BasicComparator[int], 1),
			position: -1,
			found:    false,
		},
		{
			name:     "One element, first",
			list:     New[int](utils.BasicComparator[int], 1),
			position: 0,
			value:    1,
			found:    true,
		},
	}

	for _, test := range tests {
test := test

		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.list.OrderedBegin(utils.BasicComparator[int])

			value, found := it.GetAt(test.position)

			assert.Equalf(t, test.found, found, test.name)
			assert.Equalf(t, test.value, value, test.name)
		})
	}
}

func TestArrayListIteratorSetAt(t *testing.T) {
	tests := []struct {
		name        string
		list        *Set[int]
		position    int
		value       int
		successfull bool
	}{
		{
			name:        "Empty",
			list:        New[int](utils.BasicComparator[int]),
			position:    0,
			value:       1,
			successfull: false,
		},
		{
			name:        "One element, begin",
			list:        New[int](utils.BasicComparator[int], 1),
			position:    -1,
			value:       1,
			successfull: false,
		},
		{
			name:        "One element, first",
			list:        New[int](utils.BasicComparator[int], 1),
			position:    0,
			value:       1,
			successfull: true,
		},
	}

	for _, test := range tests {
test := test

		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.list.OrderedBegin(utils.BasicComparator[int])

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
test := test

		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it1 := New[int](utils.BasicComparator[int], 1, 2, 3, 4, 5).OrderedBegin(utils.BasicComparator[int])
			it2 := New[int](utils.BasicComparator[int], 1, 2, 3, 4, 5).OrderedBegin(utils.BasicComparator[int])

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
test := test

		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it1 := New[int](utils.BasicComparator[int], 1, 2, 3, 4, 5).OrderedBegin(utils.BasicComparator[int])
			it2 := New[int](utils.BasicComparator[int], 1, 2, 3, 4, 5).OrderedBegin(utils.BasicComparator[int])

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
test := test

		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it1 := New[int](utils.BasicComparator[int], 1, 2, 3, 4, 5).OrderedBegin(utils.BasicComparator[int])
			it2 := New[int](utils.BasicComparator[int], 1, 2, 3, 4, 5).OrderedBegin(utils.BasicComparator[int])

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
test := test

		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it1 := New[int](utils.BasicComparator[int], 1, 2, 3, 4, 5).OrderedBegin(utils.BasicComparator[int])
			it2 := New[int](utils.BasicComparator[int], 1, 2, 3, 4, 5).OrderedBegin(utils.BasicComparator[int])

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
		iteratorInit  func(*Set[int], utils.Comparator[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
		iteratorCheck func(ds.ReadWriteOrdCompBidRandCollIterator[int, int]) bool
	}{
		{
			name:          "Begin",
			iteratorInit:  (*Set[int]).OrderedBegin,
			iteratorCheck: (ds.ReadWriteOrdCompBidRandCollIterator[int, int]).IsBegin,
		},
		{
			name:          "End",
			iteratorInit:  (*Set[int]).OrderedEnd,
			iteratorCheck: (ds.ReadWriteOrdCompBidRandCollIterator[int, int]).IsEnd,
		},
		{
			name:          "First",
			iteratorInit:  (*Set[int]).OrderedFirst,
			iteratorCheck: (ds.ReadWriteOrdCompBidRandCollIterator[int, int]).IsFirst,
		},
		{
			name:          "Last",
			iteratorInit:  (*Set[int]).OrderedLast,
			iteratorCheck: (ds.ReadWriteOrdCompBidRandCollIterator[int, int]).IsLast,
		},
	}

	for _, test := range tests {
test := test

		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(New[int](utils.BasicComparator[int], 1, 2, 4, 5), utils.BasicComparator[int])
			assert.Truef(t, test.iteratorCheck(it), test.name)
		})
	}
}

func TestArrayListIteratorSize(t *testing.T) {
	tests := []struct {
		name         string
		list         *Set[int]
		iteratorInit func(*Set[int], utils.Comparator[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
		size         int
	}{
		{
			name:         "Empty",
			list:         New[int](utils.BasicComparator[int]),
			size:         0,
			iteratorInit: (*Set[int]).OrderedFirst,
		},

		{
			name:         "One element, first",
			list:         New[int](utils.BasicComparator[int], 1),
			size:         1,
			iteratorInit: (*Set[int]).OrderedFirst,
		},

		{
			name:         "3 elements, middle",
			list:         New[int](utils.BasicComparator[int], 1, 2, 3),
			size:         3,
			iteratorInit: (*Set[int]).OrderedFirst,
		},
	}

	for _, test := range tests {
test := test

		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.list, utils.BasicComparator[int])

			size := it.Size()

			assert.Equalf(t, test.size, size, test.name)
		})
	}
}
