package btree

import (
	"testing"

	"github.com/JonasMuehlmann/datastructures.go/ds"
	testCommon "github.com/JonasMuehlmann/datastructures.go/tests"
	"github.com/JonasMuehlmann/datastructures.go/utils"
	"github.com/stretchr/testify/assert"
)

func TestBTreeOrderedIteratorIsValid(t *testing.T) {
	tests := []struct {
		name         string
		map_         *Tree[string, int]
		position     string
		isValid      bool
		iteratorInit func(*Tree[string, int]) ds.ReadWriteOrdCompBidRandCollMapIterator[string, int]
	}{
		{
			name:         "Empty",
			map_:         New[string, int](3, utils.BasicComparator[string]),
			isValid:      false,
			iteratorInit: (*Tree[string, int]).OrderedFirst,
		},

		{
			name:         "One element, first",
			map_:         NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1}),
			position:     "",
			isValid:      true,
			iteratorInit: (*Tree[string, int]).OrderedFirst,
		},
		{
			name:         "3 elements, middle",
			map_:         NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1, "b": 2, "c": 3}),
			position:     "b",
			isValid:      true,
			iteratorInit: (*Tree[string, int]).OrderedFirst,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.map_)

			if test.position != "" {
				it.MoveToKey(test.position)
			}

			isValid := it.IsValid()

			assert.Equalf(t, test.isValid, isValid, test.name)
		})
	}
}

func TestBTreeOrderedIteratorGet(t *testing.T) {
	tests := []struct {
		name     string
		map_     *Tree[string, int]
		position string
		value    int
		found    bool
	}{
		{
			name:     "Empty",
			map_:     New[string, int](3, utils.BasicComparator[string]),
			position: "",
			found:    false,
		},
		{
			name:     "One element, first",
			map_:     NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1}),
			position: "a",
			value:    1,
			found:    true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.map_.OrderedFirst()

			if test.position != "" {
				it.MoveToKey(test.position)
			}

			value, found := it.Get()

			assert.Equalf(t, test.found, found, test.name)
			assert.Equalf(t, test.value, value, test.name)
		})
	}
}

func TestBTreeOrderedIteratorSet(t *testing.T) {
	tests := []struct {
		name        string
		map_        *Tree[string, int]
		position    string
		value       int
		successfull bool
	}{
		{
			name:        "Empty",
			map_:        New[string, int](3, utils.BasicComparator[string]),
			position:    "a",
			value:       1,
			successfull: false,
		},
		{
			name:        "One element, first",
			map_:        NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1}),
			position:    "a",
			value:       1,
			successfull: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.map_.OrderedFirst()

			if test.position != "" {
				it.MoveToKey(test.position)
			}

			successfull := it.Set(test.value)

			assert.Equalf(t, test.successfull, successfull, test.name)
		})
	}
}

func TestBTreeOrderedIteratorGetAtKey(t *testing.T) {
	tests := []struct {
		name     string
		map_     *Tree[string, int]
		position string
		value    int
		found    bool
	}{
		{
			name:     "Empty",
			map_:     New[string, int](3, utils.BasicComparator[string]),
			position: "",
			found:    false,
		},

		{
			name:     "One element, first",
			map_:     NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1}),
			position: "a",
			value:    1,
			found:    true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.map_.OrderedFirst()

			value, found := it.GetAtKey(test.position)

			assert.Equalf(t, test.found, found, test.name)
			assert.Equalf(t, test.value, value, test.name)
		})
	}
}

func TestBTreeOrderedIteratorSetAtKey(t *testing.T) {
	tests := []struct {
		name        string
		map_        *Tree[string, int]
		position    string
		value       int
		successfull bool
	}{
		{
			name:        "Empty",
			map_:        New[string, int](3, utils.BasicComparator[string]),
			position:    "a",
			value:       1,
			successfull: true,
		},

		{
			name:        "One element, first",
			map_:        NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1}),
			position:    "a",
			value:       1,
			successfull: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.map_.OrderedFirst()

			successfull := it.SetAtKey(test.position, test.value)

			assert.Equalf(t, test.successfull, successfull, test.name)
		})
	}
}

func TestBTreeOrderedIteratorDistanceTo(t *testing.T) {
	tests := []struct {
		name     string
		map_     *Tree[string, int]
		key1     string
		key2     string
		distance int
	}{
		{
			name:     "Equal",
			map_:     NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1, "b": 2, "c": 3}),
			key1:     "a",
			key2:     "a",
			distance: 0,
		},
		{
			name:     "First lower",
			map_:     NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1, "b": 2, "c": 3}),
			key1:     "a",
			key2:     "b",
			distance: -1,
		},
		{
			name:     "Second lower",
			map_:     NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1, "b": 2, "c": 3}),
			key1:     "b",
			key2:     "a",
			distance: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it1 := test.map_.OrderedBegin()
			it2 := test.map_.OrderedBegin()

			it1.MoveToKey(test.key1)
			it2.MoveToKey(test.key2)

			distance := it1.DistanceTo(it2)

			assert.Equalf(t, test.distance, distance, test.name)
		})
	}
}

func TestBTreeOrderedIteratorIsAfter(t *testing.T) {
	tests := []struct {
		name    string
		map_    *Tree[string, int]
		key1    string
		key2    string
		isAfter bool
	}{
		{
			name:    "Equal",
			map_:    NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1, "b": 2, "c": 3}),
			key1:    "a",
			key2:    "a",
			isAfter: false,
		},
		{
			name:    "First lower",
			map_:    NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1, "b": 2, "c": 3}),
			key1:    "a",
			key2:    "b",
			isAfter: false,
		},
		{
			name:    "Second lower",
			map_:    NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1, "b": 2, "c": 3}),
			key1:    "b",
			key2:    "a",
			isAfter: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it1 := test.map_.OrderedBegin()
			it2 := test.map_.OrderedBegin()

			it1.MoveToKey(test.key1)
			it2.MoveToKey(test.key2)

			isAfter := it1.IsAfter(it2)

			assert.Equalf(t, test.isAfter, isAfter, test.name)
		})
	}
}

func TestBTreeOrderedIteratorIsBefore(t *testing.T) {
	tests := []struct {
		name    string
		map_    *Tree[string, int]
		key1    string
		key2    string
		isAfter bool
	}{
		{
			name:    "Equal",
			map_:    NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1, "b": 2, "c": 3}),
			key1:    "a",
			key2:    "a",
			isAfter: false,
		},
		{
			name:    "First lower",
			map_:    NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1, "b": 2, "c": 3}),
			key1:    "a",
			key2:    "b",
			isAfter: true,
		},
		{
			name:    "Second lower",
			map_:    NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1, "b": 2, "c": 3}),
			key1:    "b",
			key2:    "a",
			isAfter: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it1 := test.map_.OrderedBegin()
			it2 := test.map_.OrderedBegin()

			it1.MoveToKey(test.key1)
			it2.MoveToKey(test.key2)

			isAfter := it1.IsBefore(it2)

			assert.Equalf(t, test.isAfter, isAfter, test.name)
		})
	}
}

func TestBTreeOrderedIteratorIsEqual(t *testing.T) {
	tests := []struct {
		name      string
		position1 string
		position2 string
		isAfter   bool
	}{
		{
			name:      "Equal",
			position1: "a",
			position2: "a",
			isAfter:   true,
		},
		{
			name:      "First lower",
			position1: "a",
			position2: "b",
			isAfter:   false,
		},
		{
			name:      "Second lower",
			position1: "b",
			position2: "a",
			isAfter:   false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			m := NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1, "b": 2, "c": 4, "d": 5})

			it1 := m.OrderedFirst()
			it2 := m.OrderedFirst()

			it1.MoveToKey(test.position1)
			it2.MoveToKey(test.position2)

			isAfter := it1.IsEqual(it2)

			assert.Equalf(t, test.isAfter, isAfter, test.name+m.ToString())
		})
	}
}

func TestHashmapOrderedIteratorIndex(t *testing.T) {
	tests := []struct {
		name         string
		map_         *Tree[string, int]
		key          string
		valid        bool
		iteratorInit func(*Tree[string, int]) ds.ReadWriteOrdCompBidRandCollMapIterator[string, int]
	}{
		{
			name:         "Empty",
			map_:         New[string, int](3, utils.BasicComparator[string]),
			valid:        false,
			iteratorInit: (*Tree[string, int]).OrderedBegin,
		},
		{
			name:         "One element, begin",
			map_:         NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1}),
			valid:        false,
			iteratorInit: (*Tree[string, int]).OrderedBegin,
		},
		{
			name:         "One element, end",
			map_:         NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1}),
			valid:        false,
			iteratorInit: (*Tree[string, int]).OrderedEnd,
		},
		{
			name:         "One element, first",
			map_:         NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1}),
			key:          "a",
			valid:        true,
			iteratorInit: (*Tree[string, int]).OrderedFirst,
		},
		{
			name:         "One element, last",
			map_:         NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1}),
			key:          "a",
			valid:        true,
			iteratorInit: (*Tree[string, int]).OrderedLast,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.map_)

			position, valid := it.GetKey()

			assert.Equalf(t, test.valid, valid, test.name)
			if test.valid {
				assert.Equalf(t, test.key, position, test.name)
			}
		})
	}
}

func TestHashmapOrderedIteratorSize(t *testing.T) {
	tests := []struct {
		name string
		map_ *Tree[string, int]
		size int
	}{
		{
			name: "Empty",
			map_: New[string, int](3, utils.BasicComparator[string]),
			size: 0,
		},
		{
			name: "3 elements, middle",
			map_: NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1, "b": 2, "c": 3}),
			size: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.map_.OrderedBegin()

			size := it.Size()

			assert.Equalf(t, test.size, size, test.name)
		})
	}
}

func TestHashmapOrderedIteratorNext(t *testing.T) {
	tests := []struct {
		name          string
		map_          *Tree[string, int]
		position      string
		isValidBefore bool
		isValidAfter  bool
		iteratorInit  func(*Tree[string, int]) ds.ReadWriteOrdCompBidRandCollMapIterator[string, int]
	}{
		{
			name:          "Empty",
			map_:          New[string, int](3, utils.BasicComparator[string]),
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Tree[string, int]).OrderedBegin,
		},
		{
			name:          "One element, begin",
			map_:          NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1}),
			isValidBefore: false,
			isValidAfter:  true,
			iteratorInit:  (*Tree[string, int]).OrderedBegin,
		},
		{
			name:          "One element, end",
			map_:          NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1}),
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Tree[string, int]).OrderedEnd,
		},
		{
			name:          "One element, first",
			map_:          NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1}),
			position:      "a",
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Tree[string, int]).OrderedFirst,
		},
		{
			name:          "One element, last",
			map_:          NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1}),
			position:      "a",
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Tree[string, int]).OrderedLast,
		},
		{
			name:          "3 elements, middle",
			map_:          NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1, "b": 2, "c": 3}),
			position:      "b",
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Tree[string, int]).OrderedBegin,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.map_)

			if test.position != "" {
				it.MoveToKey(test.position)
			}

			isValidBefore := it.IsValid()
			assert.Equalf(t, test.isValidBefore, isValidBefore, test.name)

			it.Next()

			isValidAfter := it.IsValid()
			assert.Equalf(t, test.isValidAfter, isValidAfter, test.name)
		})
	}
}

func TestHashmapOrderedIteratorNextN(t *testing.T) {
	tests := []struct {
		name          string
		map_          *Tree[string, int]
		position      string
		n             int
		isValidBefore bool
		isValidAfter  bool
		iteratorInit  func(*Tree[string, int]) ds.ReadWriteOrdCompBidRandCollMapIterator[string, int]
	}{
		{
			name:          "Empty",
			map_:          New[string, int](3, utils.BasicComparator[string]),
			position:      "",
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Tree[string, int]).OrderedBegin,
		},
		{
			name:          "One element, begin",
			map_:          NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1}),
			position:      "",
			n:             1,
			isValidBefore: false,
			isValidAfter:  true,
			iteratorInit:  (*Tree[string, int]).OrderedBegin,
		},
		{
			name:          "One element, end",
			map_:          NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1}),
			position:      "",
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Tree[string, int]).OrderedEnd,
		},
		{
			name:          "One element, first",
			map_:          NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1}),
			position:      "a",
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Tree[string, int]).OrderedFirst,
		},
		{
			name:          "One element, last",
			map_:          NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1}),
			position:      "a",
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Tree[string, int]).OrderedLast,
		},
		{
			name:          "3 elements, middle",
			map_:          NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1, "b": 2, "c": 3}),
			position:      "b",
			n:             1,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Tree[string, int]).OrderedBegin,
		},
		{
			name:          "3 elements, middle, move out of bounds",
			map_:          NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1, "b": 2, "c": 3}),
			position:      "b",
			n:             5,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Tree[string, int]).OrderedBegin,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.map_)

			if test.position != "" {
				it.MoveToKey(test.position)
			}

			isValidBefore := it.IsValid()
			assert.Equalf(t, test.isValidBefore, isValidBefore, test.name)

			it.NextN(test.n)

			isValidAfter := it.IsValid()
			assert.Equalf(t, test.isValidAfter, isValidAfter, test.name)
		})
	}
}

func TestHashmapOrderedIteratorPrevious(t *testing.T) {
	tests := []struct {
		name          string
		map_          *Tree[string, int]
		position      string
		isValidBefore bool
		isValidAfter  bool
		iteratorInit  func(*Tree[string, int]) ds.ReadWriteOrdCompBidRandCollMapIterator[string, int]
	}{
		{
			name:          "Empty",
			map_:          New[string, int](3, utils.BasicComparator[string]),
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Tree[string, int]).OrderedBegin,
		},
		{
			name:          "One element, begin",
			map_:          NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1}),
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Tree[string, int]).OrderedBegin,
		},
		{
			name:          "One element, end",
			map_:          NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1}),
			isValidBefore: false,
			isValidAfter:  true,
			iteratorInit:  (*Tree[string, int]).OrderedEnd,
		},
		{
			name:          "One element, first",
			map_:          NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1}),
			position:      "a",
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Tree[string, int]).OrderedFirst,
		},
		{
			name:          "One element, last",
			map_:          NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1}),
			position:      "a",
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Tree[string, int]).OrderedLast,
		},
		{
			name:          "3 elements, middle",
			map_:          NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1, "b": 2, "c": 3}),
			position:      "b",
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Tree[string, int]).OrderedBegin,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.map_)

			if test.position != "" {
				it.MoveToKey(test.position)
			}

			isValidBefore := it.IsValid()
			assert.Equalf(t, test.isValidBefore, isValidBefore, test.name)

			it.Previous()

			isValidAfter := it.IsValid()
			assert.Equalf(t, test.isValidAfter, isValidAfter, test.name)
		})
	}
}

func TestHashmapOrderedIteratorPreviousN(t *testing.T) {
	tests := []struct {
		name          string
		map_          *Tree[string, int]
		position      string
		n             int
		isValidBefore bool
		isValidAfter  bool
		iteratorInit  func(*Tree[string, int]) ds.ReadWriteOrdCompBidRandCollMapIterator[string, int]
	}{
		{
			name:          "Empty",
			map_:          New[string, int](3, utils.BasicComparator[string]),
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Tree[string, int]).OrderedBegin,
		},
		{
			name:          "One element, begin",
			map_:          NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1}),
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Tree[string, int]).OrderedBegin,
		},
		{
			name:          "One element, end",
			map_:          NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1}),
			n:             1,
			isValidBefore: false,
			isValidAfter:  true,
			iteratorInit:  (*Tree[string, int]).OrderedEnd,
		},
		{
			name:          "One element, first",
			map_:          NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1}),
			position:      "a",
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Tree[string, int]).OrderedFirst,
		},
		{
			name:          "One element, last",
			map_:          NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1}),
			position:      "a",
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Tree[string, int]).OrderedLast,
		},
		{
			name:          "3 elements, middle",
			map_:          NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1, "b": 2, "c": 3}),
			position:      "b",
			n:             1,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Tree[string, int]).OrderedBegin,
		},
		{
			name:          "3 elements, middle, move out of bounds",
			map_:          NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1, "b": 2, "c": 3}),
			position:      "b",
			n:             5,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Tree[string, int]).OrderedBegin,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.map_)

			if test.position != "" {
				it.MoveToKey(test.position)
			}

			isValidBefore := it.IsValid()
			assert.Equalf(t, test.isValidBefore, isValidBefore, test.name)

			it.PreviousN(test.n)

			isValidAfter := it.IsValid()
			assert.Equalf(t, test.isValidAfter, isValidAfter, test.name)
		})
	}
}

func TestHashmapOrderedIteratorMoveBy(t *testing.T) {
	tests := []struct {
		name          string
		map_          *Tree[string, int]
		position      string
		n             int
		isValidBefore bool
		isValidAfter  bool
		iteratorInit  func(*Tree[string, int]) ds.ReadWriteOrdCompBidRandCollMapIterator[string, int]
	}{
		{
			name:          "Empty",
			map_:          New[string, int](3, utils.BasicComparator[string]),
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Tree[string, int]).OrderedBegin,
		},
		{
			name:          "One element, begin",
			map_:          NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1}),
			n:             1,
			isValidBefore: false,
			isValidAfter:  true,
			iteratorInit:  (*Tree[string, int]).OrderedBegin,
		},
		{
			name:          "One element, end",
			map_:          NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1}),
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Tree[string, int]).OrderedEnd,
		},
		{
			name:          "One element, first",
			map_:          NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1}),
			position:      "a",
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Tree[string, int]).OrderedFirst,
		},
		{
			name:          "One element, last",
			map_:          NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1}),
			position:      "a",
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Tree[string, int]).OrderedLast,
		},
		{
			name:          "3 elements, middle",
			map_:          NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1, "b": 2, "c": 3}),
			position:      "b",
			n:             1,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Tree[string, int]).OrderedBegin,
		},
		{
			name:          "5 elements, middle, forward by 2",
			map_:          NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}),
			position:      "c",
			n:             2,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Tree[string, int]).OrderedBegin,
		},
		{
			name:          "5 elements, middle, backward by 2",
			map_:          NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}),
			position:      "c",
			n:             -2,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Tree[string, int]).OrderedBegin,
		},
		{
			name:          "3 elements, middle, move out of bounds",
			map_:          NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1, "b": 2, "c": 3}),
			position:      "b",
			n:             5,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Tree[string, int]).OrderedBegin,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)

			it := test.iteratorInit(test.map_)

			if test.position != "" {
				it.MoveToKey(test.position)
			}

			isValidBefore := it.IsValid()
			assert.Equalf(t, test.isValidBefore, isValidBefore, test.name)

			it.MoveBy(test.n)

			isValidAfter := it.IsValid()
			assert.Equalf(t, test.isValidAfter, isValidAfter, test.name)
		})
	}
}

func TestHashmapOrderedIteratorMoveToKey(t *testing.T) {
	tests := []struct {
		name         string
		map_         *Tree[string, int]
		position     string
		isValidAfter bool
		index        int
	}{
		{
			name:         "Empty",
			map_:         New[string, int](3, utils.BasicComparator[string]),
			isValidAfter: false,
		},
		{
			name:         "3 elements, first item",
			map_:         NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1, "b": 2, "c": 3}),
			isValidAfter: true,
			position:     "a",
			index:        0,
		},
		{
			name:         "3 elements, middle item",
			map_:         NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1, "b": 2, "c": 3}),
			position:     "b",
			isValidAfter: true,
			index:        1,
		},
		{
			name:         "3 elements, last item",
			map_:         NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1, "b": 2, "c": 3}),
			position:     "c",
			isValidAfter: true,
			index:        2,
		},
	}

	for _, test := range tests {
		testNameOrig := test.name
		for _, iteratorInit := range []struct {
			name string
			f    func(*Tree[string, int]) ds.ReadWriteOrdCompBidRandCollMapIterator[string, int]
		}{
			{", from first,", (*Tree[string, int]).OrderedFirst},
			{", from last,", (*Tree[string, int]).OrderedLast},
			{", from begin,", (*Tree[string, int]).OrderedBegin},
			{", from end,", (*Tree[string, int]).OrderedEnd},
		} {
			test.name = testNameOrig + iteratorInit.name

			t.Run(test.name, func(t *testing.T) {
				defer testCommon.HandlePanic(t, test.name)

				repr := test.map_.ToString()
				assert.NotEmpty(t, repr)

				it := iteratorInit.f(test.map_)

				if test.position != "" {
					it.MoveToKey(test.position)
				}

				isValidAfter := it.IsValid()
				assert.Equalf(t, test.isValidAfter, isValidAfter, test.name+" valid after")

				if test.isValidAfter {
					index := it.DistanceTo(test.map_.OrderedFirst())
					assert.Equalf(t, test.index, index, test.name+" index")
				}
			})
		}
	}
}

func TestBTreeOrderedIteratorIsBeginEndFirstLast(t *testing.T) {
	tests := []struct {
		name          string
		iteratorInit  func(*Tree[string, int]) ds.ReadWriteOrdCompBidRandCollMapIterator[string, int]
		iteratorCheck func(ds.ReadWriteOrdCompBidRandCollMapIterator[string, int]) bool
	}{
		{
			name:          "OrderedFirst",
			iteratorInit:  (*Tree[string, int]).OrderedFirst,
			iteratorCheck: (ds.ReadWriteOrdCompBidRandCollMapIterator[string, int]).IsFirst,
		}, {
			name:          "OrderedLast",
			iteratorInit:  (*Tree[string, int]).OrderedLast,
			iteratorCheck: (ds.ReadWriteOrdCompBidRandCollMapIterator[string, int]).IsLast,
		},
		{
			name:          "OrderedBegin",
			iteratorInit:  (*Tree[string, int]).OrderedBegin,
			iteratorCheck: (ds.ReadWriteOrdCompBidRandCollMapIterator[string, int]).IsBegin,
		},
		{
			name:          "OrderedEnd",
			iteratorInit:  (*Tree[string, int]).OrderedEnd,
			iteratorCheck: (ds.ReadWriteOrdCompBidRandCollMapIterator[string, int]).IsEnd,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(NewFromMap[string, int](3, utils.BasicComparator[string], map[string]int{"a": 1, "b": 2, "c": 4, "5": 5}))
			assert.Truef(t, test.iteratorCheck(it), test.name)
		})
	}
}
