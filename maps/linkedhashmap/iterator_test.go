package linkedhashmap

import (
	"testing"

	"github.com/JonasMuehlmann/datastructures.go/ds"
	"github.com/JonasMuehlmann/datastructures.go/maps/hashmap"
	testCommon "github.com/JonasMuehlmann/datastructures.go/tests"
	"github.com/JonasMuehlmann/datastructures.go/utils"
	"github.com/stretchr/testify/assert"
)

func TestHashMapIteratorIsValid(t *testing.T) {
	tests := []struct {
		name         string
		map_         *Map[string, int]
		position     string
		isValid      bool
		iteratorInit func(*Map[string, int]) ds.ReadWriteOrdCompBidRandCollIterator[string, int]
	}{
		{
			name:         "One element, first",
			map_:         NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			position:     "",
			isValid:      true,
			iteratorInit: (*Map[string, int]).First,
		},
		{
			name:         "3 elements, middle",
			map_:         NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3}).OrderedBegin(utils.BasicComparator[string])),
			position:     "b",
			isValid:      true,
			iteratorInit: (*Map[string, int]).First,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.map_)

			if test.position != "" {
				it.MoveTo(test.position)
			}

			isValid := it.IsValid()

			assert.Equalf(t, test.isValid, isValid, test.name)
		})
	}
}

func TestHashMapIteratorGet(t *testing.T) {
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
			map_:     NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			position: "a",
			value:    1,
			found:    true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.map_.First()

			if test.position != "" {
				it.MoveTo(test.position)
			}

			value, found := it.Get()

			assert.Equalf(t, test.found, found, test.name)
			assert.Equalf(t, test.value, value, test.name)
		})
	}
}

func TestHashMapIteratorSet(t *testing.T) {
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
			map_:        NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			position:    "a",
			value:       1,
			successfull: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.map_.First()

			if test.position != "" {
				it.MoveTo(test.position)
			}

			successfull := it.Set(test.value)

			assert.Equalf(t, test.successfull, successfull, test.name)
		})
	}
}

func TestHashMapIteratorGetAt(t *testing.T) {
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
			map_:     NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			position: "a",
			value:    1,
			found:    true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.map_.First()

			value, found := it.GetAt(test.position)

			assert.Equalf(t, test.found, found, test.name)
			assert.Equalf(t, test.value, value, test.name)
		})
	}
}

func TestHashMapIteratorSetAt(t *testing.T) {
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
			map_:        NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			position:    "a",
			value:       1,
			successfull: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.map_.First()

			successfull := it.SetAt(test.position, test.value)

			assert.Equalf(t, test.successfull, successfull, test.name)
		})
	}
}

func TestHashMapIteratorDistanceTo(t *testing.T) {
	tests := []struct {
		name     string
		map_     *Map[string, int]
		key1     string
		key2     string
		distance int
	}{
		{
			name:     "Equal",
			map_:     NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3}).OrderedBegin(utils.BasicComparator[string])),
			key1:     "a",
			key2:     "a",
			distance: 0,
		},
		{
			name:     "First lower",
			map_:     NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3}).OrderedBegin(utils.BasicComparator[string])),
			key1:     "a",
			key2:     "b",
			distance: -1,
		},
		{
			name:     "Second lower",
			map_:     NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3}).OrderedBegin(utils.BasicComparator[string])),
			key1:     "b",
			key2:     "a",
			distance: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it1 := test.map_.Begin()
			it2 := test.map_.Begin()

			it1.MoveTo(test.key1)
			it2.MoveTo(test.key2)

			distance := it1.DistanceTo(it2)

			assert.Equalf(t, test.distance, distance, test.name)
		})
	}
}

func TestHashMapIteratorIsAfter(t *testing.T) {
	tests := []struct {
		name    string
		map_    *Map[string, int]
		key1    string
		key2    string
		isAfter bool
	}{
		{
			name:    "Equal",
			map_:    NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3}).OrderedBegin(utils.BasicComparator[string])),
			key1:    "a",
			key2:    "a",
			isAfter: false,
		},
		{
			name:    "First lower",
			map_:    NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3}).OrderedBegin(utils.BasicComparator[string])),
			key1:    "a",
			key2:    "b",
			isAfter: false,
		},
		{
			name:    "Second lower",
			map_:    NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3}).OrderedBegin(utils.BasicComparator[string])),
			key1:    "b",
			key2:    "a",
			isAfter: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it1 := test.map_.Begin()
			it2 := test.map_.Begin()

			it1.MoveTo(test.key1)
			it2.MoveTo(test.key2)

			isAfter := it1.IsAfter(it2)

			assert.Equalf(t, test.isAfter, isAfter, test.name)
		})
	}
}

func TestHashMapIteratorIsBefore(t *testing.T) {
	tests := []struct {
		name    string
		map_    *Map[string, int]
		key1    string
		key2    string
		isAfter bool
	}{
		{
			name:    "Equal",
			map_:    NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3}).OrderedBegin(utils.BasicComparator[string])),
			key1:    "a",
			key2:    "a",
			isAfter: false,
		},
		{
			name:    "First lower",
			map_:    NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3}).OrderedBegin(utils.BasicComparator[string])),
			key1:    "a",
			key2:    "b",
			isAfter: true,
		},
		{
			name:    "Second lower",
			map_:    NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3}).OrderedBegin(utils.BasicComparator[string])),
			key1:    "b",
			key2:    "a",
			isAfter: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it1 := test.map_.Begin()
			it2 := test.map_.Begin()

			it1.MoveTo(test.key1)
			it2.MoveTo(test.key2)

			isAfter := it1.IsBefore(it2)

			assert.Equalf(t, test.isAfter, isAfter, test.name)
		})
	}
}

func TestHashMapIteratorIsEqual(t *testing.T) {
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
			m := NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 4, "5": 5}).OrderedBegin(utils.BasicComparator[string]))

			it1 := m.First()
			it2 := m.First()

			it1.MoveTo(test.position1)
			it2.MoveTo(test.position2)

			isAfter := it1.IsEqual(it2)

			assert.Equalf(t, test.isAfter, isAfter, test.name)
		})
	}
}

func TestHashmapIteratorIndex(t *testing.T) {
	tests := []struct {
		name         string
		map_         *Map[string, int]
		key          string
		valid        bool
		iteratorInit func(*Map[string, int]) ds.ReadWriteOrdCompBidRandCollIterator[string, int]
	}{
		{
			name:         "Empty",
			map_:         New[string, int](),
			valid:        false,
			iteratorInit: (*Map[string, int]).Begin,
		},
		{
			name:         "One element, begin",
			map_:         NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			valid:        false,
			iteratorInit: (*Map[string, int]).Begin,
		},
		{
			name:         "One element, end",
			map_:         NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			valid:        false,
			iteratorInit: (*Map[string, int]).End,
		},
		{
			name:         "One element, first",
			map_:         NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			key:          "a",
			valid:        true,
			iteratorInit: (*Map[string, int]).First,
		},
		{
			name:         "One element, last",
			map_:         NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			key:          "a",
			valid:        true,
			iteratorInit: (*Map[string, int]).Last,
		},
		{
			name:         "3 elements, middle",
			map_:         NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3}).OrderedBegin(utils.BasicComparator[string])),
			valid:        false,
			iteratorInit: (*Map[string, int]).Begin,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.map_)

			position, valid := it.Index()

			assert.Equalf(t, test.key, position, test.name)
			assert.Equalf(t, test.valid, valid, test.name)
		})
	}
}

func TestHashmapIteratorSize(t *testing.T) {
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
			map_: NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3}).OrderedBegin(utils.BasicComparator[string])),
			size: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.map_.Begin()

			size := it.Size()

			assert.Equalf(t, test.size, size, test.name)
		})
	}
}

func TestHashmapIteratorNext(t *testing.T) {
	tests := []struct {
		name          string
		map_          *Map[string, int]
		position      string
		isValidBefore bool
		isValidAfter  bool
		iteratorInit  func(*Map[string, int]) ds.ReadWriteOrdCompBidRandCollIterator[string, int]
	}{
		{
			name:          "Empty",
			map_:          New[string, int](),
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).Begin,
		},
		{
			name:          "One element, begin",
			map_:          NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			isValidBefore: false,
			isValidAfter:  true,
			iteratorInit:  (*Map[string, int]).Begin,
		},
		{
			name:          "One element, end",
			map_:          NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).End,
		},
		{
			name:          "One element, first",
			map_:          NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			position:      "a",
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).First,
		},
		{
			name:          "One element, last",
			map_:          NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			position:      "a",
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).Last,
		},
		{
			name:          "3 elements, middle",
			map_:          NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3}).OrderedBegin(utils.BasicComparator[string])),
			position:      "b",
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Map[string, int]).Begin,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.map_)

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

func TestHashmapIteratorNextN(t *testing.T) {
	tests := []struct {
		name          string
		map_          *Map[string, int]
		position      string
		n             int
		isValidBefore bool
		isValidAfter  bool
		iteratorInit  func(*Map[string, int]) ds.ReadWriteOrdCompBidRandCollIterator[string, int]
	}{
		{
			name:          "Empty",
			map_:          New[string, int](),
			position:      "",
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).Begin,
		},
		{
			name:          "One element, begin",
			map_:          NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			position:      "",
			n:             1,
			isValidBefore: false,
			isValidAfter:  true,
			iteratorInit:  (*Map[string, int]).Begin,
		},
		{
			name:          "One element, end",
			map_:          NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			position:      "",
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).End,
		},
		{
			name:          "One element, first",
			map_:          NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			position:      "a",
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).First,
		},
		{
			name:          "One element, last",
			map_:          NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			position:      "a",
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).Last,
		},
		{
			name:          "3 elements, middle",
			map_:          NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3}).OrderedBegin(utils.BasicComparator[string])),
			position:      "a",
			n:             1,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Map[string, int]).Begin,
		},
		{
			name:          "3 elements, middle, move out of bounds",
			map_:          NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3}).OrderedBegin(utils.BasicComparator[string])),
			position:      "a",
			n:             5,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).Begin,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.map_)

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

func TestHashmapIteratorPrevious(t *testing.T) {
	tests := []struct {
		name          string
		map_          *Map[string, int]
		position      string
		isValidBefore bool
		isValidAfter  bool
		iteratorInit  func(*Map[string, int]) ds.ReadWriteOrdCompBidRandCollIterator[string, int]
	}{
		{
			name:          "Empty",
			map_:          New[string, int](),
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).Begin,
		},
		{
			name:          "One element, begin",
			map_:          NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).Begin,
		},
		{
			name:          "One element, end",
			map_:          NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			isValidBefore: false,
			isValidAfter:  true,
			iteratorInit:  (*Map[string, int]).End,
		},
		{
			name:          "One element, first",
			map_:          NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			position:      "a",
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).First,
		},
		{
			name:          "One element, last",
			map_:          NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			position:      "a",
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).Last,
		},
		{
			name:          "3 elements, middle",
			map_:          NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3}).OrderedBegin(utils.BasicComparator[string])),
			position:      "b",
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Map[string, int]).Begin,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.map_)

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

func TestHashmapIteratorPreviousN(t *testing.T) {
	tests := []struct {
		name          string
		map_          *Map[string, int]
		position      string
		n             int
		isValidBefore bool
		isValidAfter  bool
		iteratorInit  func(*Map[string, int]) ds.ReadWriteOrdCompBidRandCollIterator[string, int]
	}{
		{
			name:          "Empty",
			map_:          New[string, int](),
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).Begin,
		},
		{
			name:          "One element, begin",
			map_:          NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).Begin,
		},
		{
			name:          "One element, end",
			map_:          NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			n:             1,
			isValidBefore: false,
			isValidAfter:  true,
			iteratorInit:  (*Map[string, int]).End,
		},
		{
			name:          "One element, first",
			map_:          NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			position:      "a",
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).First,
		},
		{
			name:          "One element, last",
			map_:          NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			position:      "a",
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).Last,
		},
		{
			name:          "3 elements, middle",
			map_:          NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3}).OrderedBegin(utils.BasicComparator[string])),
			position:      "b",
			n:             1,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Map[string, int]).Begin,
		},
		{
			name:          "3 elements, middle, move out of bounds",
			map_:          NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3}).OrderedBegin(utils.BasicComparator[string])),
			position:      "b",
			n:             5,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).Begin,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.map_)

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

func TestHashmapIteratorMoveBy(t *testing.T) {
	tests := []struct {
		name          string
		map_          *Map[string, int]
		position      string
		n             int
		isValidBefore bool
		isValidAfter  bool
		iteratorInit  func(*Map[string, int]) ds.ReadWriteOrdCompBidRandCollIterator[string, int]
	}{
		{
			name:          "Empty",
			map_:          New[string, int](),
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).Begin,
		},
		{
			name:          "One element, begin",
			map_:          NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			n:             1,
			isValidBefore: false,
			isValidAfter:  true,
			iteratorInit:  (*Map[string, int]).Begin,
		},
		{
			name:          "One element, end",
			map_:          NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			n:             1,
			isValidBefore: false,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).End,
		},
		{
			name:          "One element, first",
			map_:          NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			position:      "a",
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).First,
		},
		{
			name:          "One element, last",
			map_:          NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1}).OrderedBegin(utils.BasicComparator[string])),
			position:      "a",
			n:             1,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).Last,
		},
		{
			name:          "3 elements, middle",
			map_:          NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3}).OrderedBegin(utils.BasicComparator[string])),
			position:      "b",
			n:             1,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Map[string, int]).Begin,
		},
		{
			name:          "5 elements, middle, forward by 2",
			map_:          NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}).OrderedBegin(utils.BasicComparator[string])),
			position:      "c",
			n:             2,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Map[string, int]).Begin,
		}, {
			name:          "5 elements, middle, backward by 2",
			map_:          NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}).OrderedBegin(utils.BasicComparator[string])),
			position:      "c",
			n:             -2,
			isValidBefore: true,
			isValidAfter:  true,
			iteratorInit:  (*Map[string, int]).Begin,
		},
		{
			name:          "3 elements, middle, move out of bounds",
			map_:          NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 3}).OrderedBegin(utils.BasicComparator[string])),
			position:      "b",
			n:             5,
			isValidBefore: true,
			isValidAfter:  false,
			iteratorInit:  (*Map[string, int]).Begin,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(test.map_)

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

func TestHashMapIteratorIsBeginEndFirstLast(t *testing.T) {
	tests := []struct {
		name          string
		iteratorInit  func(*Map[string, int]) ds.ReadWriteOrdCompBidRandCollIterator[string, int]
		iteratorCheck func(ds.ReadWriteOrdCompBidRandCollIterator[string, int]) bool
	}{
		{
			name:          "First",
			iteratorInit:  (*Map[string, int]).First,
			iteratorCheck: (ds.ReadWriteOrdCompBidRandCollIterator[string, int]).IsFirst,
		}, {
			name:          "Last",
			iteratorInit:  (*Map[string, int]).Last,
			iteratorCheck: (ds.ReadWriteOrdCompBidRandCollIterator[string, int]).IsLast,
		},
		{
			name:          "Begin",
			iteratorInit:  (*Map[string, int]).Begin,
			iteratorCheck: (ds.ReadWriteOrdCompBidRandCollIterator[string, int]).IsBegin,
		},
		{
			name:          "End",
			iteratorInit:  (*Map[string, int]).End,
			iteratorCheck: (ds.ReadWriteOrdCompBidRandCollIterator[string, int]).IsEnd,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer testCommon.HandlePanic(t, test.name)
			it := test.iteratorInit(NewFromIterator[string, int](hashmap.NewFromMap(map[string]int{"a": 1, "b": 2, "c": 4, "5": 5}).OrderedBegin(utils.BasicComparator[string])))
			assert.Truef(t, test.iteratorCheck(it), test.name)
		})
	}
}
