package hashmap

import (
	"testing"

	"github.com/JonasMuehlmann/datastructures.go/ds"
	"github.com/JonasMuehlmann/datastructures.go/utils"
	"github.com/stretchr/testify/assert"
)

func TestHashMapOrderedIteratorIsValid(t *testing.T) {
	tests := []struct {
		name         string
		list         *Map[string, int]
		position     string
		isValid      bool
		iteratorInit func(*Map[string, int], utils.Comparator[string]) ds.ReadWriteUnordCompBidRandCollIterator[string, int]
	}{
		{
			name:         "One element, first",
			list:         NewFromMap[string, int](map[string]int{"a": 1}),
			position:     "",
			isValid:      true,
			iteratorInit: (*Map[string, int]).OrderedFirst,
		},
		{
			name:         "3 elements, middle",
			list:         NewFromMap[string, int](map[string]int{"a": 1, "b": 2, "c": 3}),
			position:     "b",
			isValid:      true,
			iteratorInit: (*Map[string, int]).OrderedFirst,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			it := test.iteratorInit(test.list, utils.BasicComparator[string])

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
		list     *Map[string, int]
		position string
		value    int
		found    bool
	}{
		{
			name:     "Empty",
			list:     New[string, int](),
			position: "",
			found:    false,
		},
		{
			name:     "One element, first",
			list:     NewFromMap[string, int](map[string]int{"a": 1}),
			position: "a",
			value:    1,
			found:    true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			it := test.list.OrderedFirst(utils.BasicComparator[string])

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
		list        *Map[string, int]
		position    string
		value       int
		successfull bool
	}{
		{
			name:        "Empty",
			list:        New[string, int](),
			position:    "a",
			value:       1,
			successfull: false,
		},
		{
			name:        "One element, first",
			list:        NewFromMap[string, int](map[string]int{"a": 1}),
			position:    "a",
			value:       1,
			successfull: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			it := test.list.OrderedFirst(utils.BasicComparator[string])

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
		list     *Map[string, int]
		position string
		value    int
		found    bool
	}{
		{
			name:     "Empty",
			list:     New[string, int](),
			position: "",
			found:    false,
		},

		{
			name:     "One element, first",
			list:     NewFromMap[string, int](map[string]int{"a": 1}),
			position: "a",
			value:    1,
			found:    true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			it := test.list.OrderedFirst(utils.BasicComparator[string])

			value, found := it.GetAt(test.position)

			assert.Equalf(t, test.found, found, test.name)
			assert.Equalf(t, test.value, value, test.name)
		})
	}
}

func TestHashMapOrderedIteratorSetAt(t *testing.T) {
	tests := []struct {
		name        string
		list        *Map[string, int]
		position    string
		value       int
		successfull bool
	}{
		{
			name:        "Empty",
			list:        New[string, int](),
			position:    "a",
			value:       1,
			successfull: true,
		},

		{
			name:        "One element, first",
			list:        NewFromMap[string, int](map[string]int{"a": 1}),
			position:    "a",
			value:       1,
			successfull: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			it := test.list.OrderedFirst(utils.BasicComparator[string])

			successfull := it.SetAt(test.position, test.value)

			assert.Equalf(t, test.successfull, successfull, test.name)
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

func TestHashMapOrderedIteratorIsBeginEndOrderedFirstLast(t *testing.T) {
	tests := []struct {
		name          string
		iteratorInit  func(*Map[string, int], utils.Comparator[string]) ds.ReadWriteUnordCompBidRandCollIterator[string, int]
		iteratorCheck func(ds.ReadWriteUnordCompBidRandCollIterator[string, int]) bool
	}{
		{
			name:          "OrderedFirst",
			iteratorInit:  (*Map[string, int]).OrderedFirst,
			iteratorCheck: (ds.ReadWriteUnordCompBidRandCollIterator[string, int]).IsFirst,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			it := test.iteratorInit(NewFromMap[string, int](map[string]int{"a": 1, "b": 2, "c": 4, "5": 5}), utils.BasicComparator[string])
			assert.Truef(t, test.iteratorCheck(it), test.name)
		})
	}
}
