package hashmap

import (
	"testing"

	"github.com/JonasMuehlmann/datastructures.go/ds"
	testCommon "github.com/JonasMuehlmann/datastructures.go/tests"
	"github.com/JonasMuehlmann/datastructures.go/utils"
	"github.com/stretchr/testify/assert"
)

func TestHashMapOrderedIteratorIsValid(t *testing.T) {
	tests := []struct {
		name         string
		map_         *Map[string, int]
		position     string
		isValid      bool
		iteratorInit func(*Map[string, int], utils.Comparator[string]) ds.ReadWriteOrdCompBidRandCollIterator[string, int]
	}{
		{
			name:         "One element, first",
			map_:         NewFromMap[string, int](map[string]int{"a": 1}),
			position:     "",
			isValid:      true,
			iteratorInit: (*Map[string, int]).OrderedFirst,
		},
		{
			name:         "3 elements, middle",
			map_:         NewFromMap[string, int](map[string]int{"a": 1, "b": 2, "c": 3}),
			position:     "b",
			isValid:      true,
			iteratorInit: (*Map[string, int]).OrderedFirst,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.map_, utils.BasicComparator[string])

			if test.position != "" {
				it.MoveTo(test.position)
			}

			isValid := it.IsValid()

			assert.Equalf(t, test.isValid, isValid, test.name)
		})
	}
}

func TestHashMapOrderedIteratorGet(t *testing.T) {
	tests := []struct {
		name     string
		map_     *Map[string, int]
		position string
		value    int
		found    bool
	}{
		{
			name:     "Empty",
			map_:     New[string, int](),
			position: "",
			found:    false,
		},
		{
			name:     "One element, first",
			map_:     NewFromMap[string, int](map[string]int{"a": 1}),
			position: "a",
			value:    1,
			found:    true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.map_.OrderedFirst(utils.BasicComparator[string])

			if test.position != "" {
				it.MoveTo(test.position)
			}

			value, found := it.Get()

			assert.Equalf(t, test.found, found, test.name)
			assert.Equalf(t, test.value, value, test.name)
		})
	}
}

func TestHashMapOrderedIteratorSet(t *testing.T) {
	tests := []struct {
		name        string
		map_        *Map[string, int]
		position    string
		value       int
		successfull bool
	}{
		{
			name:        "Empty",
			map_:        New[string, int](),
			position:    "a",
			value:       1,
			successfull: false,
		},
		{
			name:        "One element, first",
			map_:        NewFromMap[string, int](map[string]int{"a": 1}),
			position:    "a",
			value:       1,
			successfull: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.map_.OrderedFirst(utils.BasicComparator[string])

			if test.position != "" {
				it.MoveTo(test.position)
			}

			successfull := it.Set(test.value)

			assert.Equalf(t, test.successfull, successfull, test.name)
		})
	}
}

func TestHashMapOrderedIteratorGetAt(t *testing.T) {
	tests := []struct {
		name     string
		map_     *Map[string, int]
		position string
		value    int
		found    bool
	}{
		{
			name:     "Empty",
			map_:     New[string, int](),
			position: "",
			found:    false,
		},

		{
			name:     "One element, first",
			map_:     NewFromMap[string, int](map[string]int{"a": 1}),
			position: "a",
			value:    1,
			found:    true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.map_.OrderedFirst(utils.BasicComparator[string])

			value, found := it.GetAt(test.position)

			assert.Equalf(t, test.found, found, test.name)
			assert.Equalf(t, test.value, value, test.name)
		})
	}
}

func TestHashMapOrderedIteratorSetAt(t *testing.T) {
	tests := []struct {
		name        string
		map_        *Map[string, int]
		position    string
		value       int
		successfull bool
	}{
		{
			name:        "Empty",
			map_:        New[string, int](),
			position:    "a",
			value:       1,
			successfull: true,
		},

		{
			name:        "One element, first",
			map_:        NewFromMap[string, int](map[string]int{"a": 1}),
			position:    "a",
			value:       1,
			successfull: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.map_.OrderedFirst(utils.BasicComparator[string])

			successfull := it.SetAt(test.position, test.value)

			assert.Equalf(t, test.successfull, successfull, test.name)
		})
	}
}

func TestHashMapOrderedIteratorDistanceTo(t *testing.T) {
	tests := []struct {
		name     string
		map_     *Map[string, int]
		key1     string
		key2     string
		distance int
	}{
		{
			name:     "Equal",
			map_:     NewFromMap[string, int](map[string]int{"a": 1, "b": 2, "c": 3}),
			key1:     "a",
			key2:     "a",
			distance: 0,
		},
		{
			name:     "First lower",
			map_:     NewFromMap[string, int](map[string]int{"a": 1, "b": 2, "c": 3}),
			key1:     "a",
			key2:     "b",
			distance: -1,
		},
		{
			name:     "Second lower",
			map_:     NewFromMap[string, int](map[string]int{"a": 1, "b": 2, "c": 3}),
			key1:     "b",
			key2:     "a",
			distance: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it1 := test.map_.OrderedBegin(utils.BasicComparator[string])
			it2 := test.map_.OrderedBegin(utils.BasicComparator[string])

			it1.MoveTo(test.key1)
			it2.MoveTo(test.key2)

			distance := it1.DistanceTo(it2)

			assert.Equalf(t, test.distance, distance, test.name)
		})
	}
}

func TestHashMapOrderedIteratorIsAfter(t *testing.T) {
	tests := []struct {
		name    string
		map_    *Map[string, int]
		key1    string
		key2    string
		isAfter bool
	}{
		{
			name:    "Equal",
			map_:    NewFromMap[string, int](map[string]int{"a": 1, "b": 2, "c": 3}),
			key1:    "a",
			key2:    "a",
			isAfter: false,
		},
		{
			name:    "First lower",
			map_:    NewFromMap[string, int](map[string]int{"a": 1, "b": 2, "c": 3}),
			key1:    "a",
			key2:    "b",
			isAfter: false,
		},
		{
			name:    "Second lower",
			map_:    NewFromMap[string, int](map[string]int{"a": 1, "b": 2, "c": 3}),
			key1:    "b",
			key2:    "a",
			isAfter: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it1 := test.map_.OrderedBegin(utils.BasicComparator[string])
			it2 := test.map_.OrderedBegin(utils.BasicComparator[string])

			it1.MoveTo(test.key1)
			it2.MoveTo(test.key2)

			isAfter := it1.IsAfter(it2)

			assert.Equalf(t, test.isAfter, isAfter, test.name)
		})
	}
}

func TestHashMapOrderedIteratorIsBefore(t *testing.T) {
	tests := []struct {
		name    string
		map_    *Map[string, int]
		key1    string
		key2    string
		isAfter bool
	}{
		{
			name:    "Equal",
			map_:    NewFromMap[string, int](map[string]int{"a": 1, "b": 2, "c": 3}),
			key1:    "a",
			key2:    "a",
			isAfter: false,
		},
		{
			name:    "First lower",
			map_:    NewFromMap[string, int](map[string]int{"a": 1, "b": 2, "c": 3}),
			key1:    "a",
			key2:    "b",
			isAfter: true,
		},
		{
			name:    "Second lower",
			map_:    NewFromMap[string, int](map[string]int{"a": 1, "b": 2, "c": 3}),
			key1:    "b",
			key2:    "a",
			isAfter: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it1 := test.map_.OrderedBegin(utils.BasicComparator[string])
			it2 := test.map_.OrderedBegin(utils.BasicComparator[string])

			it1.MoveTo(test.key1)
			it2.MoveTo(test.key2)

			isAfter := it1.IsBefore(it2)

			assert.Equalf(t, test.isAfter, isAfter, test.name)
		})
	}
}

func TestHashMapOrderedIteratorIsEqual(t *testing.T) {
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
			name:      "OrderedFirst lower",
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
			m := NewFromMap[string, int](map[string]int{"a": 1, "b": 2, "c": 4, "5": 5})

			it1 := m.OrderedFirst(utils.BasicComparator[string])
			it2 := m.OrderedFirst(utils.BasicComparator[string])

			it1.MoveTo(test.position1)
			it2.MoveTo(test.position2)

			isAfter := it1.IsEqual(it2)

			assert.Equalf(t, test.isAfter, isAfter, test.name)
		})
	}
}

func TestHashmapOrderedIteratorIndex(t *testing.T) {
	tests := []struct {
		name         string
		map_         *Map[string, int]
		key          string
		valid        bool
		iteratorInit func(*Map[string, int], utils.Comparator[string]) ds.ReadWriteOrdCompBidRandCollIterator[string, int]
	}{
		{
			name:         "Empty",
			map_:         New[string, int](),
			valid:        false,
			iteratorInit: (*Map[string, int]).OrderedBegin,
		},
		{
			name:         "One element, begin",
			map_:         NewFromMap[string, int](map[string]int{"a": 1}),
			valid:        false,
			iteratorInit: (*Map[string, int]).OrderedBegin,
		},
		{
			name:         "One element, end",
			map_:         NewFromMap[string, int](map[string]int{"a": 1}),
			valid:        false,
			iteratorInit: (*Map[string, int]).OrderedEnd,
		},
		{
			name:         "One element, first",
			map_:         NewFromMap[string, int](map[string]int{"a": 1}),
			key:          "a",
			valid:        true,
			iteratorInit: (*Map[string, int]).OrderedFirst,
		},
		{
			name:         "One element, last",
			map_:         NewFromMap[string, int](map[string]int{"a": 1}),
			key:          "a",
			valid:        true,
			iteratorInit: (*Map[string, int]).OrderedLast,
		},
		{
			name:         "3 elements, middle",
			map_:         NewFromMap[string, int](map[string]int{"a": 1, "b": 2, "c": 3}),
			valid:        false,
			iteratorInit: (*Map[string, int]).OrderedBegin,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.map_, utils.BasicComparator[string])

			position, valid := it.Index()

			assert.Equalf(t, test.key, position, test.name)
			assert.Equalf(t, test.valid, valid, test.name)
		})
	}
}

func TestHashmapOrderedIteratorSize(t *testing.T) {
	tests := []struct {
		name string
		map_ *Map[string, int]
		size int
	}{
		{
			name: "Empty",
			map_: New[string, int](),
			size: 0,
		},
		{
			name: "3 elements, middle",
			map_: NewFromMap[string, int](map[string]int{"a": 1, "b": 2, "c": 3}),
			size: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.map_.OrderedBegin(utils.BasicComparator[string])

			size := it.Size()

			assert.Equalf(t, test.size, size, test.name)
		})
	}
}

func TestHashmapOrderedIteratorNext(t *testing.T) {
	tests := []struct {
		name          string
		map_          *Map[string, int]
		position      string
		isValidBefore bool
		isValidAfter  bool
		iteratorInit  func(*Map[string, int], utils.Comparator[string]) ds.ReadWriteOrdCompBidRandCollIterator[string, int]
	}{
		{
			name:          "Empty",
			map_:          New[string, int](),
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).OrderedBegin,
		},
		{
			name:          "One element, begin",
			map_:          NewFromMap[string, int](map[string]int{"a": 1}),
			isValidBefore: false,
			isValidAfter:  true,
			iteratorInit:  (*Map[string, int]).OrderedBegin,
		},
		{
			name:          "One element, end",
			map_:          NewFromMap[string, int](map[string]int{"a": 1}),
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).OrderedEnd,
		},
		{
			name:          "One element, first",
			map_:          NewFromMap[string, int](map[string]int{"a": 1}),
			position:      "a",
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).OrderedFirst,
		},
		{
			name:          "One element, last",
			map_:          NewFromMap[string, int](map[string]int{"a": 1}),
			position:      "a",
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).OrderedLast,
		},
		{
			name:          "3 elements, middle",
			map_:          NewFromMap[string, int](map[string]int{"a": 1, "b": 2, "c": 3}),
			position:      "b",
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Map[string, int]).OrderedBegin,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.map_, utils.BasicComparator[string])

			if test.position != "" {
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

func TestHashmapOrderedIteratorNextN(t *testing.T) {
	tests := []struct {
		name          string
		map_          *Map[string, int]
		position      string
		n             int
		isValidBefore bool
		isValidAfter  bool
		iteratorInit  func(*Map[string, int], utils.Comparator[string]) ds.ReadWriteOrdCompBidRandCollIterator[string, int]
	}{
		{
			name:          "Empty",
			map_:          New[string, int](),
			position:      "",
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).OrderedBegin,
		},
		{
			name:          "One element, begin",
			map_:          NewFromMap[string, int](map[string]int{"a": 1}),
			position:      "",
			n:             1,
			isValidBefore: false,
			isValidAfter:  true,
			iteratorInit:  (*Map[string, int]).OrderedBegin,
		},
		{
			name:          "One element, end",
			map_:          NewFromMap[string, int](map[string]int{"a": 1}),
			position:      "",
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).OrderedEnd,
		},
		{
			name:          "One element, first",
			map_:          NewFromMap[string, int](map[string]int{"a": 1}),
			position:      "a",
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).OrderedFirst,
		},
		{
			name:          "One element, last",
			map_:          NewFromMap[string, int](map[string]int{"a": 1}),
			position:      "a",
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).OrderedLast,
		},
		{
			name:          "3 elements, middle",
			map_:          NewFromMap[string, int](map[string]int{"a": 1, "b": 2, "c": 3}),
			position:      "a",
			n:             1,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Map[string, int]).OrderedBegin,
		},
		{
			name:          "3 elements, middle, move out of bounds",
			map_:          NewFromMap[string, int](map[string]int{"a": 1, "b": 2, "c": 3}),
			position:      "a",
			n:             5,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).OrderedBegin,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.map_, utils.BasicComparator[string])

			if test.position != "" {
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

func TestHashmapOrderedIteratorPrevious(t *testing.T) {
	tests := []struct {
		name          string
		map_          *Map[string, int]
		position      string
		isValidBefore bool
		isValidAfter  bool
		iteratorInit  func(*Map[string, int], utils.Comparator[string]) ds.ReadWriteOrdCompBidRandCollIterator[string, int]
	}{
		{
			name:          "Empty",
			map_:          New[string, int](),
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).OrderedBegin,
		},
		{
			name:          "One element, begin",
			map_:          NewFromMap[string, int](map[string]int{"a": 1}),
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).OrderedBegin,
		},
		{
			name:          "One element, end",
			map_:          NewFromMap[string, int](map[string]int{"a": 1}),
			isValidBefore: false,
			isValidAfter:  true,
			iteratorInit:  (*Map[string, int]).OrderedEnd,
		},
		{
			name:          "One element, first",
			map_:          NewFromMap[string, int](map[string]int{"a": 1}),
			position:      "a",
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).OrderedFirst,
		},
		{
			name:          "One element, last",
			map_:          NewFromMap[string, int](map[string]int{"a": 1}),
			position:      "a",
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).OrderedLast,
		},
		{
			name:          "3 elements, middle",
			map_:          NewFromMap[string, int](map[string]int{"a": 1, "b": 2, "c": 3}),
			position:      "b",
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Map[string, int]).OrderedBegin,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.map_, utils.BasicComparator[string])

			if test.position != "" {
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

func TestHashmapOrderedIteratorPreviousN(t *testing.T) {
	tests := []struct {
		name          string
		map_          *Map[string, int]
		position      string
		n             int
		isValidBefore bool
		isValidAfter  bool
		iteratorInit  func(*Map[string, int], utils.Comparator[string]) ds.ReadWriteOrdCompBidRandCollIterator[string, int]
	}{
		{
			name:          "Empty",
			map_:          New[string, int](),
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).OrderedBegin,
		},
		{
			name:          "One element, begin",
			map_:          NewFromMap[string, int](map[string]int{"a": 1}),
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).OrderedBegin,
		},
		{
			name:          "One element, end",
			map_:          NewFromMap[string, int](map[string]int{"a": 1}),
			n:             1,
			isValidBefore: false,
			isValidAfter:  true,
			iteratorInit:  (*Map[string, int]).OrderedEnd,
		},
		{
			name:          "One element, first",
			map_:          NewFromMap[string, int](map[string]int{"a": 1}),
			position:      "a",
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).OrderedFirst,
		},
		{
			name:          "One element, last",
			map_:          NewFromMap[string, int](map[string]int{"a": 1}),
			position:      "a",
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).OrderedLast,
		},
		{
			name:          "3 elements, middle",
			map_:          NewFromMap[string, int](map[string]int{"a": 1, "b": 2, "c": 3}),
			position:      "b",
			n:             1,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Map[string, int]).OrderedBegin,
		},
		{
			name:          "3 elements, middle, move out of bounds",
			map_:          NewFromMap[string, int](map[string]int{"a": 1, "b": 2, "c": 3}),
			position:      "b",
			n:             5,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).OrderedBegin,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.map_, utils.BasicComparator[string])

			if test.position != "" {
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

func TestHashmapOrderedIteratorMoveBy(t *testing.T) {
	tests := []struct {
		name          string
		map_          *Map[string, int]
		position      string
		n             int
		isValidBefore bool
		isValidAfter  bool
		iteratorInit  func(*Map[string, int], utils.Comparator[string]) ds.ReadWriteOrdCompBidRandCollIterator[string, int]
	}{
		{
			name:          "Empty",
			map_:          New[string, int](),
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).OrderedBegin,
		},
		{
			name:          "One element, begin",
			map_:          NewFromMap[string, int](map[string]int{"a": 1}),
			n:             1,
			isValidBefore: false,
			isValidAfter:  true,
			iteratorInit:  (*Map[string, int]).OrderedBegin,
		},
		{
			name:          "One element, end",
			map_:          NewFromMap[string, int](map[string]int{"a": 1}),
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).OrderedEnd,
		},
		{
			name:          "One element, first",
			map_:          NewFromMap[string, int](map[string]int{"a": 1}),
			position:      "a",
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).OrderedFirst,
		},
		{
			name:          "One element, last",
			map_:          NewFromMap[string, int](map[string]int{"a": 1}),
			position:      "a",
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).OrderedLast,
		},
		{
			name:          "3 elements, middle",
			map_:          NewFromMap[string, int](map[string]int{"a": 1, "b": 2, "c": 3}),
			position:      "b",
			n:             1,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Map[string, int]).OrderedBegin,
		},
		{
			name:          "5 elements, middle, forward by 2",
			map_:          NewFromMap[string, int](map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}),
			position:      "c",
			n:             2,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Map[string, int]).OrderedBegin,
		}, {
			name:          "5 elements, middle, backward by 2",
			map_:          NewFromMap[string, int](map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}),
			position:      "c",
			n:             -2,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Map[string, int]).OrderedBegin,
		},
		{
			name:          "3 elements, middle, move out of bounds",
			map_:          NewFromMap[string, int](map[string]int{"a": 1, "b": 2, "c": 3}),
			position:      "b",
			n:             5,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).OrderedBegin,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.map_, utils.BasicComparator[string])

			if test.position != "" {
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

func TestHashMapOrderedIteratorIsBeginEndFirstLast(t *testing.T) {
	tests := []struct {
		name          string
		iteratorInit  func(*Map[string, int], utils.Comparator[string]) ds.ReadWriteOrdCompBidRandCollIterator[string, int]
		iteratorCheck func(ds.ReadWriteOrdCompBidRandCollIterator[string, int]) bool
	}{
		{
			name:          "OrderedFirst",
			iteratorInit:  (*Map[string, int]).OrderedFirst,
			iteratorCheck: (ds.ReadWriteOrdCompBidRandCollIterator[string, int]).IsFirst,
		}, {
			name:          "OrderedLast",
			iteratorInit:  (*Map[string, int]).OrderedLast,
			iteratorCheck: (ds.ReadWriteOrdCompBidRandCollIterator[string, int]).IsLast,
		},
		{
			name:          "OrderedBegin",
			iteratorInit:  (*Map[string, int]).OrderedBegin,
			iteratorCheck: (ds.ReadWriteOrdCompBidRandCollIterator[string, int]).IsBegin,
		},
		{
			name:          "OrderedEnd",
			iteratorInit:  (*Map[string, int]).OrderedEnd,
			iteratorCheck: (ds.ReadWriteOrdCompBidRandCollIterator[string, int]).IsEnd,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(NewFromMap[string, int](map[string]int{"a": 1, "b": 2, "c": 4, "5": 5}), utils.BasicComparator[string])
			assert.Truef(t, test.iteratorCheck(it), test.name)
		})
	}
}
