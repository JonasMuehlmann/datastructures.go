package circularbuffer

import (
	"testing"

	testCommon "github.com/JonasMuehlmann/datastructures.go/tests"

	"github.com/JonasMuehlmann/datastructures.go/ds"

	"github.com/stretchr/testify/assert"
)

const (
	NoMoveMagicPosition = 7869543205234798
)

func TestCircularBufferIteratorIsValid(t *testing.T) {
	tests := []struct {
		name         string
		list         *Queue[int]
		position     int
		isValid      bool
		iteratorInit func(*Queue[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
	}{
		{
			name:         "Empty",
			list:         New[int](5),
			position:     NoMoveMagicPosition,
			isValid:      false,
			iteratorInit: (*Queue[int]).Begin,
		},
		{
			name:         "One element, begin",
			list:         NewFromSlice[int](5, []int{1}),
			position:     NoMoveMagicPosition,
			isValid:      false,
			iteratorInit: (*Queue[int]).Begin,
		},
		{
			name:         "One element, end",
			list:         NewFromSlice[int](5, []int{1}),
			position:     NoMoveMagicPosition,
			isValid:      false,
			iteratorInit: (*Queue[int]).End,
		},
		{
			name:         "One element, first",
			list:         NewFromSlice[int](5, []int{1}),
			position:     NoMoveMagicPosition,
			isValid:      true,
			iteratorInit: (*Queue[int]).First,
		},
		{
			name:         "One element, last",
			list:         NewFromSlice[int](5, []int{1}),
			position:     NoMoveMagicPosition,
			isValid:      true,
			iteratorInit: (*Queue[int]).Last,
		},
		{
			name:         "3 elements, middle",
			list:         NewFromSlice[int](5, []int{1, 2, 3}),
			position:     1,
			isValid:      true,
			iteratorInit: (*Queue[int]).Begin,
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

func TestCircularBufferIteratorIndex(t *testing.T) {
	tests := []struct {
		name         string
		list         *Queue[int]
		position     int
		iteratorInit func(*Queue[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
	}{
		{
			name:         "Empty",
			list:         New[int](5),
			position:     -1,
			iteratorInit: (*Queue[int]).Begin,
		},
		{
			name:         "One element, begin",
			list:         NewFromSlice[int](5, []int{1}),
			position:     -1,
			iteratorInit: (*Queue[int]).Begin,
		},
		{
			name:         "One element, end",
			list:         NewFromSlice[int](5, []int{1}),
			position:     1,
			iteratorInit: (*Queue[int]).End,
		},
		{
			name:         "One element, first",
			list:         NewFromSlice[int](5, []int{1}),
			position:     0,
			iteratorInit: (*Queue[int]).First,
		},
		{
			name:         "One element, last",
			list:         NewFromSlice[int](5, []int{1}),
			position:     0,
			iteratorInit: (*Queue[int]).Last,
		},
		{
			name:         "3 elements, middle",
			list:         NewFromSlice[int](5, []int{1, 2, 3}),
			position:     -1,
			iteratorInit: (*Queue[int]).Begin,
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

func TestCircularBufferIteratorNext(t *testing.T) {
	tests := []struct {
		name          string
		list          *Queue[int]
		position      int
		isValidBefore bool
		isValidAfter  bool
		iteratorInit  func(*Queue[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
	}{
		{
			name:          "Empty",
			list:          New[int](5),
			position:      NoMoveMagicPosition,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Queue[int]).Begin,
		},
		{
			name:          "One element, begin",
			list:          NewFromSlice[int](5, []int{1}),
			position:      NoMoveMagicPosition,
			isValidBefore: false,
			isValidAfter:  true,
			iteratorInit:  (*Queue[int]).Begin,
		},
		{
			name:          "One element, end",
			list:          NewFromSlice[int](5, []int{1}),
			position:      NoMoveMagicPosition,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Queue[int]).End,
		},
		{
			name:          "One element, first",
			list:          NewFromSlice[int](5, []int{1}),
			position:      NoMoveMagicPosition,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Queue[int]).First,
		},
		{
			name:          "One element, last",
			list:          NewFromSlice[int](5, []int{1}),
			position:      NoMoveMagicPosition,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Queue[int]).Last,
		},
		{
			name:          "3 elements, middle",
			list:          NewFromSlice[int](5, []int{1, 2, 3}),
			position:      1,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Queue[int]).Begin,
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

			isValidAfter := it.IsValid()
			assert.Equalf(t, test.isValidAfter, isValidAfter, test.name)
		})
	}
}

func TestCircularBufferIteratorNextN(t *testing.T) {
	tests := []struct {
		name          string
		list          *Queue[int]
		position      int
		n             int
		isValidBefore bool
		isValidAfter  bool
		iteratorInit  func(*Queue[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
	}{
		{
			name:          "Empty",
			list:          New[int](5),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Queue[int]).Begin,
		},
		{
			name:          "One element, begin",
			list:          NewFromSlice[int](5, []int{1}),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: false,
			isValidAfter:  true,
			iteratorInit:  (*Queue[int]).Begin,
		},
		{
			name:          "One element, end",
			list:          NewFromSlice[int](5, []int{1}),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Queue[int]).End,
		},
		{
			name:          "One element, first",
			list:          NewFromSlice[int](5, []int{1}),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Queue[int]).First,
		},
		{
			name:          "One element, last",
			list:          NewFromSlice[int](5, []int{1}),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Queue[int]).Last,
		},
		{
			name:          "3 elements, middle",
			list:          NewFromSlice[int](5, []int{1, 2, 3}),
			position:      1,
			n:             1,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Queue[int]).Begin,
		},
		{
			name:          "3 elements, middle, move out of bounds",
			list:          NewFromSlice[int](5, []int{1, 2, 3}),
			position:      1,
			n:             5,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Queue[int]).Begin,
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

			isValidAfter := it.IsValid()
			assert.Equalf(t, test.isValidAfter, isValidAfter, test.name)
		})
	}
}

func TestCircularBufferIteratorPrevious(t *testing.T) {
	tests := []struct {
		name          string
		list          *Queue[int]
		position      int
		isValidBefore bool
		isValidAfter  bool
		iteratorInit  func(*Queue[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
	}{
		{
			name:          "Empty",
			list:          New[int](5),
			position:      NoMoveMagicPosition,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Queue[int]).Begin,
		},
		{
			name:          "One element, begin",
			list:          NewFromSlice[int](5, []int{1}),
			position:      NoMoveMagicPosition,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Queue[int]).Begin,
		},
		{
			name:          "One element, end",
			list:          NewFromSlice[int](5, []int{1}),
			position:      NoMoveMagicPosition,
			isValidBefore: false,
			isValidAfter:  true,
			iteratorInit:  (*Queue[int]).End,
		},
		{
			name:          "One element, first",
			list:          NewFromSlice[int](5, []int{1}),
			position:      NoMoveMagicPosition,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Queue[int]).First,
		},
		{
			name:          "One element, last",
			list:          NewFromSlice[int](5, []int{1}),
			position:      NoMoveMagicPosition,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Queue[int]).Last,
		},
		{
			name:          "3 elements, middle",
			list:          NewFromSlice[int](5, []int{1, 2, 3}),
			position:      1,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Queue[int]).Begin,
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

			it.Previous()

			isValidAfter := it.IsValid()
			assert.Equalf(t, test.isValidAfter, isValidAfter, test.name)
		})
	}
}

func TestCircularBufferIteratorPreviousN(t *testing.T) {
	tests := []struct {
		name          string
		list          *Queue[int]
		position      int
		n             int
		isValidBefore bool
		isValidAfter  bool
		iteratorInit  func(*Queue[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
	}{
		{
			name:          "Empty",
			list:          New[int](5),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Queue[int]).Begin,
		},
		{
			name:          "One element, begin",
			list:          NewFromSlice[int](5, []int{1}),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Queue[int]).Begin,
		},
		{
			name:          "One element, end",
			list:          NewFromSlice[int](5, []int{1}),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: false,
			isValidAfter:  true,
			iteratorInit:  (*Queue[int]).End,
		},
		{
			name:          "One element, first",
			list:          NewFromSlice[int](5, []int{1}),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Queue[int]).First,
		},
		{
			name:          "One element, last",
			list:          NewFromSlice[int](5, []int{1}),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Queue[int]).Last,
		},
		{
			name:          "3 elements, middle",
			list:          NewFromSlice[int](5, []int{1, 2, 3}),
			position:      1,
			n:             1,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Queue[int]).Begin,
		},
		{
			name:          "3 elements, middle, move out of bounds",
			list:          NewFromSlice[int](5, []int{1, 2, 3}),
			position:      1,
			n:             5,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Queue[int]).Begin,
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

			isValidAfter := it.IsValid()
			assert.Equalf(t, test.isValidAfter, isValidAfter, test.name)
		})
	}
}

func TestCircularBufferIteratorMoveBy(t *testing.T) {
	tests := []struct {
		name          string
		list          *Queue[int]
		position      int
		n             int
		isValidBefore bool
		isValidAfter  bool
		iteratorInit  func(*Queue[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
	}{
		{
			name:          "Empty",
			list:          New[int](5),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Queue[int]).Begin,
		},
		{
			name:          "One element, begin",
			list:          NewFromSlice[int](5, []int{1}),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: false,
			isValidAfter:  true,
			iteratorInit:  (*Queue[int]).Begin,
		},
		{
			name:          "One element, end",
			list:          NewFromSlice[int](5, []int{1}),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Queue[int]).End,
		},
		{
			name:          "One element, first",
			list:          NewFromSlice[int](5, []int{1}),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Queue[int]).First,
		},
		{
			name:          "One element, last",
			list:          NewFromSlice[int](5, []int{1}),
			position:      NoMoveMagicPosition,
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Queue[int]).Last,
		},
		{
			name:          "3 elements, middle",
			list:          NewFromSlice[int](5, []int{1, 2, 3}),
			position:      1,
			n:             1,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Queue[int]).Begin,
		},
		{
			name:          "5 elements, middle, forward by 2",
			list:          NewFromSlice[int](5, []int{1, 2, 3, 4, 5}),
			position:      2,
			n:             2,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Queue[int]).Begin,
		}, {
			name:          "5 elements, middle, backward by 2",
			list:          NewFromSlice[int](5, []int{1, 2, 3, 4, 5}),
			position:      2,
			n:             -2,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Queue[int]).Begin,
		},
		{
			name:          "3 elements, middle, move out of bounds",
			list:          NewFromSlice[int](5, []int{1, 2, 3}),
			position:      1,
			n:             5,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Queue[int]).Begin,
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

			it.MoveBy(test.n)

			isValidAfter := it.IsValid()
			assert.Equalf(t, test.isValidAfter, isValidAfter, test.name)
		})
	}
}

func TestCircularBufferIteratorGet(t *testing.T) {
	tests := []struct {
		name     string
		list     *Queue[int]
		position int
		value    int
		found    bool
	}{
		{
			name:     "Empty",
			list:     New[int](5),
			position: NoMoveMagicPosition,
			found:    false,
		},
		{
			name:     "One element, begin",
			list:     NewFromSlice[int](5, []int{1}),
			position: NoMoveMagicPosition,
			found:    false,
		},
		{
			name:     "One element, first",
			list:     NewFromSlice[int](5, []int{1}),
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

func TestCircularBufferIteratorSet(t *testing.T) {
	tests := []struct {
		name        string
		list        *Queue[int]
		position    int
		value       int
		successfull bool
	}{
		{
			name:        "Empty",
			list:        New[int](5),
			position:    NoMoveMagicPosition,
			value:       1,
			successfull: false,
		},
		{
			name:        "One element, begin",
			list:        NewFromSlice[int](5, []int{1}),
			position:    NoMoveMagicPosition,
			value:       1,
			successfull: false,
		},
		{
			name:        "One element, first",
			list:        NewFromSlice[int](5, []int{1}),
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
			it := test.list.Begin()

			if test.position != NoMoveMagicPosition {
				it.MoveTo(test.position)
			}

			successfull := it.Set(test.value)

			assert.Equalf(t, test.successfull, successfull, test.name)
		})
	}
}

func TestCircularBufferIteratorGetAt(t *testing.T) {
	tests := []struct {
		name     string
		list     *Queue[int]
		position int
		value    int
		found    bool
	}{
		{
			name:     "Empty",
			list:     New[int](5),
			position: 0,
			found:    false,
		},
		{
			name:     "One element, begin",
			list:     NewFromSlice[int](5, []int{1}),
			position: -1,
			found:    false,
		},
		{
			name:     "One element, first",
			list:     NewFromSlice[int](5, []int{1}),
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
			it := test.list.Begin()

			value, found := it.GetAt(test.position)

			assert.Equalf(t, test.found, found, test.name)
			assert.Equalf(t, test.value, value, test.name)
		})
	}
}

func TestCircularBufferIteratorSetAt(t *testing.T) {
	tests := []struct {
		name        string
		list        *Queue[int]
		position    int
		value       int
		successfull bool
	}{
		{
			name:        "Empty",
			list:        New[int](5),
			position:    0,
			value:       1,
			successfull: false,
		},
		{
			name:        "One element, begin",
			list:        NewFromSlice[int](5, []int{1}),
			position:    -1,
			value:       1,
			successfull: false,
		},
		{
			name:        "One element, first",
			list:        NewFromSlice[int](5, []int{1}),
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
			it := test.list.Begin()

			successfull := it.SetAt(test.position, test.value)

			assert.Equalf(t, test.successfull, successfull, test.name)
		})
	}
}

// NOTE: Missing test case: other does not implement IndexedIterator
func TestCircularBufferIteratorDistanceTo(t *testing.T) {
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
			it1 := NewFromSlice[int](5, []int{1, 2, 3, 4, 5}).Begin()
			it2 := NewFromSlice[int](5, []int{1, 2, 3, 4, 5}).Begin()

			it1.MoveTo(test.position1)
			it2.MoveTo(test.position2)

			distance := it1.DistanceTo(it2)

			assert.Equalf(t, test.distance, distance, test.name)
		})
	}
}

func TestCircularBufferIteratorIsAfter(t *testing.T) {
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
			it1 := NewFromSlice[int](5, []int{1, 2, 3, 4, 5}).Begin()
			it2 := NewFromSlice[int](5, []int{1, 2, 3, 4, 5}).Begin()

			it1.MoveTo(test.position1)
			it2.MoveTo(test.position2)

			isAfter := it1.IsAfter(it2)

			assert.Equalf(t, test.isAfter, isAfter, test.name)
		})
	}
}

func TestCircularBufferIteratorIsBefore(t *testing.T) {
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
			it1 := NewFromSlice[int](5, []int{1, 2, 3, 4, 5}).Begin()
			it2 := NewFromSlice[int](5, []int{1, 2, 3, 4, 5}).Begin()

			it1.MoveTo(test.position1)
			it2.MoveTo(test.position2)

			isAfter := it1.IsBefore(it2)

			assert.Equalf(t, test.isAfter, isAfter, test.name)
		})
	}
}

func TestCircularBufferIteratorIsEqual(t *testing.T) {
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
			it1 := NewFromSlice[int](5, []int{1, 2, 3, 4, 5}).Begin()
			it2 := NewFromSlice[int](5, []int{1, 2, 3, 4, 5}).Begin()

			it1.MoveTo(test.position1)
			it2.MoveTo(test.position2)

			isAfter := it1.IsEqual(it2)

			assert.Equalf(t, test.isAfter, isAfter, test.name)
		})
	}
}

func TestCircularBufferIteratorIsBeginEndFirstLast(t *testing.T) {
	tests := []struct {
		name          string
		iteratorInit  func(*Queue[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
		iteratorCheck func(ds.ReadWriteOrdCompBidRandCollIterator[int, int]) bool
	}{
		{
			name:          "Begin",
			iteratorInit:  (*Queue[int]).Begin,
			iteratorCheck: (ds.ReadWriteOrdCompBidRandCollIterator[int, int]).IsBegin,
		},
		{
			name:          "End",
			iteratorInit:  (*Queue[int]).End,
			iteratorCheck: (ds.ReadWriteOrdCompBidRandCollIterator[int, int]).IsEnd,
		},
		{
			name:          "First",
			iteratorInit:  (*Queue[int]).First,
			iteratorCheck: (ds.ReadWriteOrdCompBidRandCollIterator[int, int]).IsFirst,
		},
		{
			name:          "Last",
			iteratorInit:  (*Queue[int]).Last,
			iteratorCheck: (ds.ReadWriteOrdCompBidRandCollIterator[int, int]).IsLast,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(NewFromSlice[int](5, []int{1, 2, 4, 5}))
			assert.Truef(t, test.iteratorCheck(it), test.name)
		})
	}
}

func TestCircularBufferIteratorSize(t *testing.T) {
	tests := []struct {
		name         string
		list         *Queue[int]
		iteratorInit func(*Queue[int]) ds.ReadWriteOrdCompBidRandCollIterator[int, int]
		size         int
	}{
		{
			name:         "Empty",
			list:         New[int](5),
			size:         0,
			iteratorInit: (*Queue[int]).First,
		},

		{
			name:         "One element, first",
			list:         NewFromSlice[int](5, []int{1}),
			size:         1,
			iteratorInit: (*Queue[int]).First,
		},

		{
			name:         "3 elements, middle",
			list:         NewFromSlice[int](5, []int{1, 2, 3}),
			size:         3,
			iteratorInit: (*Queue[int]).First,
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
