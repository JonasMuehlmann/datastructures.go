package hashmap

import (
	"testing"

	"github.com/JonasMuehlmann/datastructures.go/ds"
	"github.com/stretchr/testify/assert"
)

func TestHashMapIteratorIsValid(t *testing.T) {
	tests := []struct {
		name         string
		list         *Map[string, int]
		position     string
		isValid      bool
		iteratorInit func(*Map[string, int]) ds.ReadWriteCompForRandCollIterator[string, int]
	}{
		{
			name:         "One element, first",
			list:         NewFromMap[string, int](map[string]int{"a": 1}),
			position:     "",
			isValid:      true,
			iteratorInit: (*Map[string, int]).First,
		},
		{
			name:         "3 elements, middle",
			list:         NewFromMap[string, int](map[string]int{"a": 1, "b": 2, "c": 3}),
			position:     "b",
			isValid:      true,
			iteratorInit: (*Map[string, int]).First,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			it := test.iteratorInit(test.list)

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
			it := test.list.First()

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
			it := test.list.First()

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
			it := test.list.First()

			value, found := it.GetAt(test.position)

			assert.Equalf(t, test.found, found, test.name)
			assert.Equalf(t, test.value, value, test.name)
		})
	}
}

func TestHashMapIteratorSetAt(t *testing.T) {
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
			it := test.list.First()

			successfull := it.SetAt(test.position, test.value)

			assert.Equalf(t, test.successfull, successfull, test.name)
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
			it1 := New[string, int]().First()
			it2 := New[string, int]().First()

			it1.MoveTo(test.position1)
			it2.MoveTo(test.position2)

			isAfter := it1.IsEqual(it2)

			assert.Equalf(t, test.isAfter, isAfter, test.name)
		})
	}
}

func TestHashMapIteratorIsBeginEndFirstLast(t *testing.T) {
	tests := []struct {
		name          string
		iteratorInit  func(*Map[string, int]) ds.ReadWriteCompForRandCollIterator[string, int]
		iteratorCheck func(ds.ReadWriteCompForRandCollIterator[string, int]) bool
	}{
		{
			name:          "First",
			iteratorInit:  (*Map[string, int]).First,
			iteratorCheck: (ds.ReadWriteCompForRandCollIterator[string, int]).IsFirst,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			it := test.iteratorInit(NewFromMap[string, int](map[string]int{"a": 1, "b": 2, "c": 4, "5": 5}))
			assert.Truef(t, test.iteratorCheck(it), test.name)
		})
	}
}
