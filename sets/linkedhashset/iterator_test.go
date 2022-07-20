package linkedhashset

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

func TestLinkedHashSetIteratorIsValid(t *testing.T) {
	tests := []struct {
		name         string
		list         *Set[int]
		position     int
		isValid      bool
		iteratorInit func(*Set[int], utils.Comparator[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
	}{
		{
			name:         "Empty",
			list:         New[int](),
			position:     NoMoveMagicPosition,
			isValid:      false,
			iteratorInit: (*Set[int]).Begin,
		},
		{
			name:         "One element, begin",
			list:         New[int](1),
			position:     NoMoveMagicPosition,
			isValid:      false,
			iteratorInit: (*Set[int]).Begin,
		},
		{
			name:         "One element, end",
			list:         New[int](1),
			position:     NoMoveMagicPosition,
			isValid:      false,
			iteratorInit: (*Set[int]).End,
		},
		{
			name:         "One element, first",
			list:         New[int](1),
			position:     NoMoveMagicPosition,
			isValid:      true,
			iteratorInit: (*Set[int]).First,
		},
		{
			name:         "One element, last",
			list:         New[int](1),
			position:     NoMoveMagicPosition,
			isValid:      true,
			iteratorInit: (*Set[int]).Last,
		},
		{
			name:         "3 elements, middle",
			list:         New[int](1, 2, 3),
			position:     1,
			isValid:      true,
			iteratorInit: (*Set[int]).First,
		},
	}

	for _, test := range tests {
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

func TestLinkedHashSetIteratorIndex(t *testing.T) {
	tests := []struct {
		name         string
		list         *Set[int]
		position     int
		iteratorInit func(*Set[int], utils.Comparator[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
	}{
		{
			name:         "Empty",
			list:         New[int](),
			position:     0,
			iteratorInit: (*Set[int]).First,
		},
		{
			name:         "One element, begin",
			list:         New[int](1),
			position:     -1,
			iteratorInit: (*Set[int]).Begin,
		},
		{
			name:         "One element, end",
			list:         New[int](1),
			position:     1,
			iteratorInit: (*Set[int]).End,
		},
		{
			name:         "One element, first",
			list:         New[int](1),
			position:     0,
			iteratorInit: (*Set[int]).First,
		},
		{
			name:         "One element, last",
			list:         New[int](1),
			position:     0,
			iteratorInit: (*Set[int]).Last,
		},
		{
			name:         "3 elements, first",
			list:         New[int](1, 2, 3),
			position:     0,
			iteratorInit: (*Set[int]).First,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.list, utils.BasicComparator[int])

			position, valid := it.Index()

			assert.Equalf(t, test.position, position, test.name)
			assert.Truef(t, valid, test.name)
		})
	}
}

func TestLinkedHashSetIteratorNext(t *testing.T) {
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
			list:          New[int](),
			position:      NoMoveMagicPosition,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Set[int]).Begin,
		},
		{
			name:          "One element, begin",
			list:          New[int](1),
			position:      NoMoveMagicPosition,
			isValidBefore: false,
			isValidAfter:  true,
			iteratorInit:  (*Set[int]).Begin,
		},
		{
			name:          "One element, end",
			list:          New[int](1),
			position:      NoMoveMagicPosition,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Set[int]).End,
		},
		{
			name:          "One element, first",
			list:          New[int](1),
			position:      NoMoveMagicPosition,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Set[int]).First,
		},
		{
			name:          "One element, last",
			list:          New[int](1),
			position:      NoMoveMagicPosition,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Set[int]).Last,
		},
		{
			name:          "3 elements, middle",
			list:          New[int](1, 2, 3),
			position:      1,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Set[int]).Begin,
		},
	}

	for _, test := range tests {
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

func TestLinkedHashSetIteratorNextN(t *testing.T) {
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
			list:          New[int](),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Set[int]).Begin,
		},
		{
			name:          "One element, begin",
			list:          New[int](1),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: false,
			isValidAfter:  true,
			iteratorInit:  (*Set[int]).Begin,
		},
		{
			name:          "One element, end",
			list:          New[int](1),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Set[int]).End,
		},
		{
			name:          "One element, first",
			list:          New[int](1),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Set[int]).First,
		},
		{
			name:          "One element, last",
			list:          New[int](1),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Set[int]).Last,
		},
		{
			name:          "3 elements, middle",
			list:          New[int](1, 2, 3),
			position:      1,
			n:             1,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Set[int]).Begin,
		},
		{
			name:          "3 elements, middle, move out of bounds",
			list:          New[int](1, 2, 3),
			position:      1,
			n:             5,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Set[int]).Begin,
		},
	}

	for _, test := range tests {
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

func TestLinkedHashSetIteratorPrevious(t *testing.T) {
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
			list:          New[int](),
			position:      NoMoveMagicPosition,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Set[int]).Begin,
		},
		{
			name:          "One element, begin",
			list:          New[int](1),
			position:      NoMoveMagicPosition,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Set[int]).Begin,
		},
		{
			name:          "One element, end",
			list:          New[int](1),
			position:      NoMoveMagicPosition,
			isValidBefore: false,
			isValidAfter:  true,
			iteratorInit:  (*Set[int]).End,
		},
		{
			name:          "One element, first",
			list:          New[int](1),
			position:      NoMoveMagicPosition,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Set[int]).First,
		},
		{
			name:          "One element, last",
			list:          New[int](1),
			position:      NoMoveMagicPosition,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Set[int]).Last,
		},
		{
			name:          "3 elements, middle",
			list:          New[int](1, 2, 3),
			position:      1,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Set[int]).Begin,
		},
	}

	for _, test := range tests {
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

func TestLinkedHashSetIteratorPreviousN(t *testing.T) {
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
			list:          New[int](),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Set[int]).Begin,
		},
		{
			name:          "One element, begin",
			list:          New[int](1),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Set[int]).Begin,
		},
		{
			name:          "One element, end",
			list:          New[int](1),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: false,
			isValidAfter:  true,
			iteratorInit:  (*Set[int]).End,
		},
		{
			name:          "One element, first",
			list:          New[int](1),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Set[int]).First,
		},
		{
			name:          "One element, last",
			list:          New[int](1),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Set[int]).Last,
		},
		{
			name:          "3 elements, middle",
			list:          New[int](1, 2, 3),
			position:      1,
			n:             1,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Set[int]).Begin,
		},
		{
			name:          "3 elements, middle, move out of bounds",
			list:          New[int](1, 2, 3),
			position:      1,
			n:             5,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Set[int]).Begin,
		},
	}

	for _, test := range tests {
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

func TestLinkedHashSetIteratorMoveBy(t *testing.T) {
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
			list:          New[int](),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Set[int]).Begin,
		},
		{
			name:          "One element, begin",
			list:          New[int](1),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: false,
			isValidAfter:  true,
			iteratorInit:  (*Set[int]).Begin,
		},
		{
			name:          "One element, end",
			list:          New[int](1),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Set[int]).End,
		},
		{
			name:          "One element, first",
			list:          New[int](1),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Set[int]).First,
		},
		{
			name:          "One element, last",
			list:          New[int](1),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Set[int]).Last,
		},
		{
			name:          "3 elements, middle",
			list:          New[int](1, 2, 3),
			position:      1,
			n:             1,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Set[int]).Begin,
		},
		{
			name:          "5 elements, middle, forward by 2",
			list:          New[int](1, 2, 3, 4, 5),
			position:      2,
			n:             2,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Set[int]).Begin,
		}, {
			name:          "5 elements, middle, backward by 2",
			list:          New[int](1, 2, 3, 4, 5),
			position:      2,
			n:             -2,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Set[int]).Begin,
		},
		{
			name:          "3 elements, middle, move out of bounds",
			list:          New[int](1, 2, 3),
			position:      1,
			n:             5,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Set[int]).Begin,
		},
	}

	for _, test := range tests {
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

func TestLinkedHashSetIteratorGet(t *testing.T) {
	tests := []struct {
		name     string
		list     *Set[int]
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
			defer testCommon.HandlePanic(t, test.name)
			it := test.list.Begin(utils.BasicComparator[int])

			if test.position != NoMoveMagicPosition {
				it.MoveTo(test.position)
			}

			value, found := it.Get()

			assert.Equalf(t, test.found, found, test.name)
			assert.Equalf(t, test.value, value, test.name)
		})
	}
}

func TestLinkedHashSetIteratorSet(t *testing.T) {
	tests := []struct {
		name        string
		list        *Set[int]
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
			defer testCommon.HandlePanic(t, test.name)
			it := test.list.Begin(utils.BasicComparator[int])

			if test.position != NoMoveMagicPosition {
				it.MoveTo(test.position)
			}

			successfull := it.Set(test.value)

			assert.Equalf(t, test.successfull, successfull, test.name)
		})
	}
}

func TestLinkedHashSetIteratorGetAt(t *testing.T) {
	tests := []struct {
		name     string
		list     *Set[int]
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
			defer testCommon.HandlePanic(t, test.name)
			it := test.list.Begin(utils.BasicComparator[int])

			value, found := it.GetAt(test.position)

			assert.Equalf(t, test.found, found, test.name)
			assert.Equalf(t, test.value, value, test.name)
		})
	}
}

func TestLinkedHashSetIteratorSetAt(t *testing.T) {
	tests := []struct {
		name        string
		list        *Set[int]
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
			defer testCommon.HandlePanic(t, test.name)
			it := test.list.Begin(utils.BasicComparator[int])

			successfull := it.SetAt(test.position, test.value)

			assert.Equalf(t, test.successfull, successfull, test.name)
		})
	}
}

// NOTE: Missing test case: other does not implement IndexedIterator
func TestLinkedHashSetIteratorDistanceTo(t *testing.T) {
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
			defer testCommon.HandlePanic(t, test.name)
			it1 := New[int]().Begin(utils.BasicComparator[int])
			it2 := New[int]().Begin(utils.BasicComparator[int])

			it1.MoveTo(test.position1)
			it2.MoveTo(test.position2)

			distance := it1.DistanceTo(it2)

			assert.Equalf(t, test.distance, distance, test.name)
		})
	}
}

func TestLinkedHashSetIteratorIsAfter(t *testing.T) {
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
			defer testCommon.HandlePanic(t, test.name)
			it1 := New[int]().Begin(utils.BasicComparator[int])
			it2 := New[int]().Begin(utils.BasicComparator[int])

			it1.MoveTo(test.position1)
			it2.MoveTo(test.position2)

			isAfter := it1.IsAfter(it2)

			assert.Equalf(t, test.isAfter, isAfter, test.name)
		})
	}
}

func TestLinkedHashSetIteratorIsBefore(t *testing.T) {
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
			defer testCommon.HandlePanic(t, test.name)
			it1 := New[int]().Begin(utils.BasicComparator[int])
			it2 := New[int]().Begin(utils.BasicComparator[int])

			it1.MoveTo(test.position1)
			it2.MoveTo(test.position2)

			isAfter := it1.IsBefore(it2)

			assert.Equalf(t, test.isAfter, isAfter, test.name)
		})
	}
}

func TestLinkedHashSetIteratorIsEqual(t *testing.T) {
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
			defer testCommon.HandlePanic(t, test.name)
			it1 := New[int]().Begin(utils.BasicComparator[int])
			it2 := New[int]().Begin(utils.BasicComparator[int])

			it1.MoveTo(test.position1)
			it2.MoveTo(test.position2)

			isAfter := it1.IsEqual(it2)

			assert.Equalf(t, test.isAfter, isAfter, test.name)
		})
	}
}

func TestLinkedHashSetIteratorIsBeginEndFirstLast(t *testing.T) {
	tests := []struct {
		name          string
		iteratorInit  func(*Set[int], utils.Comparator[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
		iteratorCheck func(ds.ReadWriteOrdCompBidRandCollIterator[int, int]) bool
	}{
		{
			name:          "Begin",
			iteratorInit:  (*Set[int]).Begin,
			iteratorCheck: (ds.ReadWriteOrdCompBidRandCollIterator[int, int]).IsBegin,
		},
		{
			name:          "End",
			iteratorInit:  (*Set[int]).End,
			iteratorCheck: (ds.ReadWriteOrdCompBidRandCollIterator[int, int]).IsEnd,
		},
		{
			name:          "First",
			iteratorInit:  (*Set[int]).First,
			iteratorCheck: (ds.ReadWriteOrdCompBidRandCollIterator[int, int]).IsFirst,
		},
		{
			name:          "Last",
			iteratorInit:  (*Set[int]).Last,
			iteratorCheck: (ds.ReadWriteOrdCompBidRandCollIterator[int, int]).IsLast,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(New[int](1, 2, 4, 5), utils.BasicComparator[int])
			assert.Truef(t, test.iteratorCheck(it), test.name)
		})
	}
}

func TestLinkedHashSetIteratorSize(t *testing.T) {
	tests := []struct {
		name         string
		list         *Set[int]
		iteratorInit func(*Set[int], utils.Comparator[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
		size         int
	}{
		{
			name:         "Empty",
			list:         New[int](),
			size:         0,
			iteratorInit: (*Set[int]).First,
		},

		{
			name:         "One element, first",
			list:         New[int](1),
			size:         1,
			iteratorInit: (*Set[int]).First,
		},

		{
			name:         "3 elements, middle",
			list:         New[int](1, 2, 3),
			size:         3,
			iteratorInit: (*Set[int]).First,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.list, utils.BasicComparator[int])

			size := it.Size()

			assert.Equalf(t, test.size, size, test.name)
		})
	}
}
