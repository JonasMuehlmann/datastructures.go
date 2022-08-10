package treebidimap

import (
	"testing"

	testCommon "github.com/JonasMuehlmann/datastructures.go/tests"

	"github.com/JonasMuehlmann/datastructures.go/ds"
	"github.com/JonasMuehlmann/datastructures.go/maps/hashmap"

	"github.com/JonasMuehlmann/datastructures.go/utils"
	"github.com/stretchr/testify/assert"
)

func TestTreeMapOrderedIteratorIsValid(t *testing.T) {
	tests := []struct {
		name         string
		map_         *Map[string, int]
		position     string
		isValid      bool
		iteratorInit func(*Map[string, int], utils.Comparator[string]) ds.ReadWriteOrdCompBidRandCollIterator[string, int]
	}{
		{
			name:         "One element, first",
			map_:         NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			position:     "",
			isValid:      true,
			iteratorInit: (*Map[string, int]).OrderedFirst,
		},
		{
			name:         "3 elements, middle",
			map_:         NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3}).OrderedBegin(utils.BasicComparator[string])),
			position:     "b",
			isValid:      true,
			iteratorInit: (*Map[string, int]).OrderedFirst,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.map_, utils.BasicComparator[string])

			if test.position != "" {
				it.MoveToKey(test.position)
			}

			isValid := it.IsValid()

			assert.Equalf(t, test.isValid, isValid, test.name)
		})
	}
}

func TestTreeMapOrderedIteratorGet(t *testing.T) {
	tests := []struct {
		name     string
		map_     *Map[string, int]
		position string
		value    int
		found    bool
	}{
		{
			name:     "Empty",
			map_:     New[string, int](utils.BasicComparator[string], utils.BasicComparator[int]),
			position: "",
			found:    false,
		},
		{
			name:     "One element, first",
			map_:     NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			position: "a",
			value:    1,
			found:    true,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it := test.map_.OrderedFirst(utils.BasicComparator[string])

			if test.position != "" {
				it.MoveToKey(test.position)
			}

			value, found := it.Get()

			assert.Equalf(t, test.found, found, test.name)
			assert.Equalf(t, test.value, value, test.name)
		})
	}
}

func TestTreeMapOrderedIteratorSet(t *testing.T) {
	tests := []struct {
		name        string
		map_        *Map[string, int]
		position    string
		value       int
		successfull bool
	}{
		{
			name:        "Empty",
			map_:        New[string, int](utils.BasicComparator[string], utils.BasicComparator[int]),
			position:    "a",
			value:       1,
			successfull: false,
		},
		{
			name:        "One element, first",
			map_:        NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			position:    "a",
			value:       1,
			successfull: true,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it := test.map_.OrderedFirst(utils.BasicComparator[string])

			if test.position != "" {
				it.MoveToKey(test.position)
			}

			successfull := it.Set(test.value)

			assert.Equalf(t, test.successfull, successfull, test.name)
		})
	}
}

func TestTreeMapOrderedIteratorGetAtKey(t *testing.T) {
	tests := []struct {
		name     string
		map_     *Map[string, int]
		position string
		value    int
		found    bool
	}{
		{
			name:     "Empty",
			map_:     New[string, int](utils.BasicComparator[string], utils.BasicComparator[int]),
			position: "",
			found:    false,
		},

		{
			name:     "One element, first",
			map_:     NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			position: "a",
			value:    1,
			found:    true,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it := test.map_.OrderedFirst(utils.BasicComparator[string])

			value, found := it.GetAtKey(test.position)

			assert.Equalf(t, test.found, found, test.name)
			assert.Equalf(t, test.value, value, test.name)
		})
	}
}

func TestTreeMapOrderedIteratorSetAtKey(t *testing.T) {
	tests := []struct {
		name        string
		map_        *Map[string, int]
		position    string
		value       int
		successfull bool
	}{
		{
			name:        "Empty",
			map_:        New[string, int](utils.BasicComparator[string], utils.BasicComparator[int]),
			position:    "a",
			value:       1,
			successfull: true,
		},

		{
			name:        "One element, first",
			map_:        NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			position:    "a",
			value:       1,
			successfull: true,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it := test.map_.OrderedFirst(utils.BasicComparator[string])

			successfull := it.SetAtKey(test.position, test.value)

			assert.Equalf(t, test.successfull, successfull, test.name)
		})
	}
}

func TestTreeMapOrderedIteratorDistanceTo(t *testing.T) {
	tests := []struct {
		name     string
		map_     *Map[string, int]
		key1     string
		key2     string
		distance int
	}{
		{
			name:     "Equal",
			map_:     NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3}).OrderedBegin(utils.BasicComparator[string])),
			key1:     "a",
			key2:     "a",
			distance: 0,
		},
		{
			name:     "First lower",
			map_:     NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3}).OrderedBegin(utils.BasicComparator[string])),
			key1:     "a",
			key2:     "b",
			distance: -1,
		},
		{
			name:     "Second lower",
			map_:     NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3}).OrderedBegin(utils.BasicComparator[string])),
			key1:     "b",
			key2:     "a",
			distance: 1,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it1 := test.map_.OrderedBegin(utils.BasicComparator[string])
			it2 := test.map_.OrderedBegin(utils.BasicComparator[string])

			it1.MoveToKey(test.key1)
			it2.MoveToKey(test.key2)

			distance := it1.DistanceTo(it2)

			assert.Equalf(t, test.distance, distance, test.name)
		})
	}
}

func TestTreeMapOrderedIteratorIsAfter(t *testing.T) {
	tests := []struct {
		name    string
		map_    *Map[string, int]
		key1    string
		key2    string
		isAfter bool
	}{
		{
			name:    "Equal",
			map_:    NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3}).OrderedBegin(utils.BasicComparator[string])),
			key1:    "a",
			key2:    "a",
			isAfter: false,
		},
		{
			name:    "First lower",
			map_:    NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3}).OrderedBegin(utils.BasicComparator[string])),
			key1:    "a",
			key2:    "b",
			isAfter: false,
		},
		{
			name:    "Second lower",
			map_:    NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3}).OrderedBegin(utils.BasicComparator[string])),
			key1:    "b",
			key2:    "a",
			isAfter: true,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it1 := test.map_.OrderedBegin(utils.BasicComparator[string])
			it2 := test.map_.OrderedBegin(utils.BasicComparator[string])

			it1.MoveToKey(test.key1)
			it2.MoveToKey(test.key2)

			isAfter := it1.IsAfter(it2)

			assert.Equalf(t, test.isAfter, isAfter, test.name)
		})
	}
}

func TestTreeMapOrderedIteratorIsBefore(t *testing.T) {
	tests := []struct {
		name    string
		map_    *Map[string, int]
		key1    string
		key2    string
		isAfter bool
	}{
		{
			name:    "Equal",
			map_:    NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3}).OrderedBegin(utils.BasicComparator[string])),
			key1:    "a",
			key2:    "a",
			isAfter: false,
		},
		{
			name:    "First lower",
			map_:    NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3}).OrderedBegin(utils.BasicComparator[string])),
			key1:    "a",
			key2:    "b",
			isAfter: true,
		},
		{
			name:    "Second lower",
			map_:    NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3}).OrderedBegin(utils.BasicComparator[string])),
			key1:    "b",
			key2:    "a",
			isAfter: false,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it1 := test.map_.OrderedBegin(utils.BasicComparator[string])
			it2 := test.map_.OrderedBegin(utils.BasicComparator[string])

			it1.MoveToKey(test.key1)
			it2.MoveToKey(test.key2)

			isAfter := it1.IsBefore(it2)

			assert.Equalf(t, test.isAfter, isAfter, test.name)
		})
	}
}

func TestTreeMapOrderedIteratorIsEqual(t *testing.T) {
	tests := []struct {
		name      string
		map_      *Map[string, int]
		position1 string
		position2 string
		isAfter   bool
	}{
		{
			name:      "Equal",
			map_:      NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3}).OrderedBegin(utils.BasicComparator[string])),
			position1: "a",
			position2: "a",
			isAfter:   true,
		},
		{
			name:      "OrderedFirst lower",
			map_:      NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3}).OrderedBegin(utils.BasicComparator[string])),
			position1: "a",
			position2: "b",
			isAfter:   false,
		},
		{
			name:      "Second lower",
			map_:      NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3}).OrderedBegin(utils.BasicComparator[string])),
			position1: "b",
			position2: "a",
			isAfter:   false,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)

			it1 := test.map_.OrderedFirst(utils.BasicComparator[string])
			it2 := test.map_.OrderedFirst(utils.BasicComparator[string])

			it1.MoveToKey(test.position1)
			it2.MoveToKey(test.position2)

			isAfter := it1.IsEqual(it2)

			assert.Equalf(t, test.isAfter, isAfter, test.name)
		})
	}
}

func TestTreeMapOrderedIteratorIndex(t *testing.T) {
	tests := []struct {
		name         string
		map_         *Map[string, int]
		key          string
		valid        bool
		iteratorInit func(*Map[string, int], utils.Comparator[string]) ds.ReadWriteOrdCompBidRandCollIterator[string, int]
	}{
		{
			name:         "Empty",
			map_:         New[string, int](utils.BasicComparator[string], utils.BasicComparator[int]),
			valid:        false,
			iteratorInit: (*Map[string, int]).OrderedBegin,
		},
		{
			name:         "One element, begin",
			map_:         NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			valid:        false,
			iteratorInit: (*Map[string, int]).OrderedBegin,
		},
		{
			name:         "One element, end",
			map_:         NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			valid:        false,
			iteratorInit: (*Map[string, int]).OrderedEnd,
		},
		{
			name:         "One element, first",
			map_:         NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			key:          "a",
			valid:        true,
			iteratorInit: (*Map[string, int]).OrderedFirst,
		},
		{
			name:         "One element, last",
			map_:         NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			key:          "a",
			valid:        true,
			iteratorInit: (*Map[string, int]).OrderedLast,
		},
		{
			name:         "3 elements, middle",
			map_:         NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3}).OrderedBegin(utils.BasicComparator[string])),
			valid:        false,
			iteratorInit: (*Map[string, int]).OrderedBegin,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.map_, utils.BasicComparator[string])

			position, valid := it.GetKey()

			assert.Equalf(t, test.key, position, test.name)
			assert.Equalf(t, test.valid, valid, test.name)
		})
	}
}

func TestTreeMapOrderedIteratorSize(t *testing.T) {
	tests := []struct {
		name string
		map_ *Map[string, int]
		size int
	}{
		{
			name: "Empty",
			map_: New[string, int](utils.BasicComparator[string], utils.BasicComparator[int]),
			size: 0,
		},
		{
			name: "3 elements, middle",
			map_: NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3}).OrderedBegin(utils.BasicComparator[string])),
			size: 3,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it := test.map_.OrderedBegin(utils.BasicComparator[string])

			size := it.Size()

			assert.Equalf(t, test.size, size, test.name)
		})
	}
}

func TestTreeMapOrderedIteratorNext(t *testing.T) {
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
			map_:          New[string, int](utils.BasicComparator[string], utils.BasicComparator[int]),
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).OrderedBegin,
		},
		{
			name:          "One element, begin",
			map_:          NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			isValidBefore: false,
			isValidAfter:  true,
			iteratorInit:  (*Map[string, int]).OrderedBegin,
		},
		{
			name:          "One element, end",
			map_:          NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).OrderedEnd,
		},
		{
			name:          "One element, first",
			map_:          NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			position:      "a",
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).OrderedFirst,
		},
		{
			name:          "One element, last",
			map_:          NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			position:      "a",
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).OrderedLast,
		},
		{
			name:          "3 elements, middle",
			map_:          NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3}).OrderedBegin(utils.BasicComparator[string])),
			position:      "b",
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Map[string, int]).OrderedBegin,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.map_, utils.BasicComparator[string])

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

func TestTreeMapOrderedIteratorNextN(t *testing.T) {
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
			map_:          New[string, int](utils.BasicComparator[string], utils.BasicComparator[int]),
			position:      "",
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).OrderedBegin,
		},
		{
			name:          "One element, begin",
			map_:          NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			position:      "",
			n:             1,
			isValidBefore: false,
			isValidAfter:  true,
			iteratorInit:  (*Map[string, int]).OrderedBegin,
		},
		{
			name:          "One element, end",
			map_:          NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			position:      "",
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).OrderedEnd,
		},
		{
			name:          "One element, first",
			map_:          NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			position:      "a",
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).OrderedFirst,
		},
		{
			name:          "One element, last",
			map_:          NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			position:      "a",
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).OrderedLast,
		},
		{
			name:          "3 elements, middle",
			map_:          NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3}).OrderedBegin(utils.BasicComparator[string])),
			position:      "a",
			n:             1,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Map[string, int]).OrderedBegin,
		},
		{
			name:          "3 elements, middle, move out of bounds",
			map_:          NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3}).OrderedBegin(utils.BasicComparator[string])),
			position:      "a",
			n:             5,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).OrderedBegin,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.map_, utils.BasicComparator[string])

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

func TestTreeMapOrderedIteratorPrevious(t *testing.T) {
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
			map_:          New[string, int](utils.BasicComparator[string], utils.BasicComparator[int]),
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).OrderedBegin,
		},
		{
			name:          "One element, begin",
			map_:          NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).OrderedBegin,
		},
		{
			name:          "One element, end",
			map_:          NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			isValidBefore: false,
			isValidAfter:  true,
			iteratorInit:  (*Map[string, int]).OrderedEnd,
		},
		{
			name:          "One element, first",
			map_:          NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			position:      "a",
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).OrderedFirst,
		},
		{
			name:          "One element, last",
			map_:          NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			position:      "a",
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).OrderedLast,
		},
		{
			name:          "3 elements, middle",
			map_:          NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3}).OrderedBegin(utils.BasicComparator[string])),
			position:      "b",
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Map[string, int]).OrderedBegin,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.map_, utils.BasicComparator[string])

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

func TestTreeMapOrderedIteratorPreviousN(t *testing.T) {
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
			map_:          New[string, int](utils.BasicComparator[string], utils.BasicComparator[int]),
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).OrderedBegin,
		},
		{
			name:          "One element, begin",
			map_:          NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).OrderedBegin,
		},
		{
			name:          "One element, end",
			map_:          NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			n:             1,
			isValidBefore: false,
			isValidAfter:  true,
			iteratorInit:  (*Map[string, int]).OrderedEnd,
		},
		{
			name:          "One element, first",
			map_:          NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			position:      "a",
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).OrderedFirst,
		},
		{
			name:          "One element, last",
			map_:          NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			position:      "a",
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).OrderedLast,
		},
		{
			name:          "3 elements, middle",
			map_:          NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3}).OrderedBegin(utils.BasicComparator[string])),
			position:      "b",
			n:             1,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Map[string, int]).OrderedBegin,
		},
		{
			name:          "3 elements, middle, move out of bounds",
			map_:          NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3}).OrderedBegin(utils.BasicComparator[string])),
			position:      "b",
			n:             5,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).OrderedBegin,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.map_, utils.BasicComparator[string])

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

func TestTreeMapOrderedIteratorMoveBy(t *testing.T) {
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
			map_:          New[string, int](utils.BasicComparator[string], utils.BasicComparator[int]),
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).OrderedBegin,
		},
		{
			name:          "One element, begin",
			map_:          NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			n:             1,
			isValidBefore: false,
			isValidAfter:  true,
			iteratorInit:  (*Map[string, int]).OrderedBegin,
		},
		{
			name:          "One element, end",
			map_:          NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).OrderedEnd,
		},
		{
			name:          "One element, first",
			map_:          NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			position:      "a",
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).OrderedFirst,
		},
		{
			name:          "One element, last",
			map_:          NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			position:      "a",
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).OrderedLast,
		},
		{
			name:          "3 elements, middle",
			map_:          NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3}).OrderedBegin(utils.BasicComparator[string])),
			position:      "b",
			n:             1,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Map[string, int]).OrderedBegin,
		},
		{
			name:          "5 elements, middle, forward by 2",
			map_:          NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}).OrderedBegin(utils.BasicComparator[string])),
			position:      "c",
			n:             2,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Map[string, int]).OrderedBegin,
		}, {
			name:          "5 elements, middle, backward by 2",
			map_:          NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}).OrderedBegin(utils.BasicComparator[string])),
			position:      "c",
			n:             -2,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Map[string, int]).OrderedBegin,
		},
		{
			name:          "3 elements, middle, move out of bounds",
			map_:          NewFromIterator[string, int](utils.BasicComparator[string], utils.BasicComparator[int], hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3}).OrderedBegin(utils.BasicComparator[string])),
			position:      "b",
			n:             5,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).OrderedBegin,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.map_, utils.BasicComparator[string])

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

func TestTreeMapOrderedIteratorIsBeginEndFirstLast(t *testing.T) {
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
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(NewFromMap[string, int](utils.BasicComparator[string], utils.BasicComparator[int], map[string]int{"a": 1, "b": 2, "c": 4, "5": 5}), utils.BasicComparator[string])
			assert.Truef(t, test.iteratorCheck(it), test.name)
		})
	}
}
