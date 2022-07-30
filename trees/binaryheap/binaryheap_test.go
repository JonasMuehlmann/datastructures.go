// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package binaryheap

import (
	"testing"

	"github.com/JonasMuehlmann/datastructures.go/ds"
	testCommon "github.com/JonasMuehlmann/datastructures.go/tests"
	"github.com/JonasMuehlmann/datastructures.go/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestArrayHeapGetValues(t *testing.T) {
	tests := []struct {
		name         string
		originalList *Heap[string]
	}{
		{
			name:         "empty list",
			originalList: New[string](utils.BasicComparator[string]),
		},
		{
			name:         "3 items, not found",
			originalList: NewFromSlice[string](utils.BasicComparator[string], []string{"foo", "bar", "baz"}),
		},
		{
			name:         "3 items, found",
			originalList: NewFromSlice[string](utils.BasicComparator[string], []string{"foo", "bar", "baz"}),
		},
	}

	for _, test := range tests {
		values := test.originalList.GetValues()

		assert.ElementsMatchf(t, test.originalList.list.GetValues(), values, test.name)
	}
}

func TestArrayHeapIsEmpty(t *testing.T) {
	tests := []struct {
		name         string
		originalList *Heap[string]
		isEmpty      bool
	}{
		{
			name:         "empty list",
			originalList: New[string](utils.BasicComparator[string]),
			isEmpty:      true,
		},
		{
			name:         "3 items, found",
			originalList: NewFromSlice[string](utils.BasicComparator[string], []string{"foo", "bar", "baz"}),
			isEmpty:      false,
		},
	}

	for _, test := range tests {
		isEmpty := test.originalList.IsEmpty()

		assert.Equalf(t, test.isEmpty, isEmpty, test.name)
	}
}

func TestArrayHeapClear(t *testing.T) {
	tests := []struct {
		name         string
		originalList *Heap[string]
	}{
		{
			name:         "empty list",
			originalList: New[string](utils.BasicComparator[string]),
		},
		{
			name:         "3 items, found",
			originalList: NewFromSlice[string](utils.BasicComparator[string], []string{"foo", "bar", "baz"}),
		},
	}

	for _, test := range tests {
		isEmpty := test.originalList.IsEmpty()
		assert.Equalf(t, test.originalList.Size() == 0, isEmpty, test.name)

		test.originalList.Clear()

		isEmpty = test.originalList.IsEmpty()
		assert.Truef(t, isEmpty, test.name)
	}
}

func TestArrayHeapPush(t *testing.T) {
	tests := []struct {
		name         string
		originalList *Heap[string]
		valueToAdd   string
		newItems     []string
	}{
		{
			name:         "empty list",
			originalList: New[string](utils.BasicComparator[string]),
			valueToAdd:   "foo",
			newItems:     []string{"foo"},
		},
		{
			name:         "1 item",
			originalList: New[string](utils.BasicComparator[string], "foo"),
			valueToAdd:   "bar",
			newItems:     []string{"foo", "bar"},
		},

		{
			name:         "list with 4 items, remove 1",
			originalList: New[string](utils.BasicComparator[string], "foo", "bar", "baz"),
			valueToAdd:   "foo",
			newItems:     []string{"foo", "bar", "baz", "foo"},
		},
	}

	for _, test := range tests {
		test.originalList.Push(test.valueToAdd)

		assert.ElementsMatchf(t, test.originalList.GetValues(), test.newItems, test.name)
	}
}

func TestArrayHeapPop(t *testing.T) {
	tests := []struct {
		name         string
		originalList *Heap[string]
		newItems     []string
	}{
		{
			name:         "empty list",
			originalList: New[string](utils.BasicComparator[string]),
			newItems:     []string{},
		},
		{
			name:         "1 item",
			originalList: New[string](utils.BasicComparator[string], "foo"),
			newItems:     []string{},
		},

		{
			name:         "list with 4 items, remove 1",
			originalList: New[string](utils.BasicComparator[string], "foo", "bar", "baz"),
			newItems:     []string{"foo", "baz"},
		},
	}

	for _, test := range tests {
		test.originalList.Pop()

		assert.ElementsMatchf(t, test.originalList.GetValues(), test.newItems, test.name)
	}
}

func TestArrayHeapPeek(t *testing.T) {
	tests := []struct {
		name         string
		originalList *Heap[string]
		found        bool
		value        string
	}{
		{
			name:         "empty list",
			originalList: New[string](utils.BasicComparator[string]),
			found:        false,
		},
		{
			name:         "1 item",
			originalList: New[string](utils.BasicComparator[string], "foo"),
			found:        true,
			value:        "foo",
		},

		{
			name:         "list with 4 items",
			originalList: New[string](utils.BasicComparator[string], "foo", "bar", "baz"),
			found:        true,
			value:        "bar",
		},
	}

	for _, test := range tests {
		value, found := test.originalList.Peek()

		assert.Equalf(t, test.found, found, test.name)

		if test.found {
			assert.Equalf(t, test.value, value, test.name)
		}

	}
}

func TestNewFromSlice(t *testing.T) {
	tests := []struct {
		name         string
		originalList *Heap[string]
	}{
		{
			name:         "empty list",
			originalList: New[string](utils.BasicComparator[string]),
		},
		{
			name:         "single item",
			originalList: New[string](utils.BasicComparator[string], "foo"),
		},
		{
			name:         "3 items",
			originalList: New[string](utils.BasicComparator[string], "foo", "bar", "baz"),
		},
	}

	for _, test := range tests {
		newList := NewFromSlice[string](utils.BasicComparator[string], test.originalList.GetValues())

		assert.ElementsMatchf(t, test.originalList.GetValues(), newList.GetValues(), test.name)
	}

}

func TestNewFromIterator(t *testing.T) {
	tests := []struct {
		name         string
		originalList *Heap[string]
	}{
		{
			name:         "empty list",
			originalList: New[string](utils.BasicComparator[string]),
		},
		{
			name:         "single item",
			originalList: New[string](utils.BasicComparator[string], "foo"),
		},
		{
			name:         "3 items",
			originalList: New[string](utils.BasicComparator[string], "foo", "bar", "baz"),
		},
	}

	for _, test := range tests {
		it := test.originalList.OrderedBegin()
		newList := NewFromIterator[string](utils.BasicComparator[string], it)

		assert.ElementsMatchf(t, test.originalList.GetValues(), newList.GetValues(), test.name)
	}

}

// NOTE: Missing test case: unordered iterator, which prevents preallocation
func TestNewFromIterators(t *testing.T) {
	tests := []struct {
		name                     string
		originalList             *Heap[string]
		newList                  *Heap[string]
		iteratorInitOrderedFirst func(*Heap[string]) ds.ReadWriteOrdCompBidRandCollIterator[int, string]
		iteratorInitOrderedEnd   func(*Heap[string]) ds.ReadWriteOrdCompBidRandCollIterator[int, string]
	}{
		{
			name:                     "empty list",
			originalList:             New[string](utils.BasicComparator[string]),
			newList:                  New[string](utils.BasicComparator[string]),
			iteratorInitOrderedFirst: (*Heap[string]).OrderedFirst,
			iteratorInitOrderedEnd:   (*Heap[string]).OrderedEnd,
		},
		{
			name:                     "single item",
			originalList:             New[string](utils.BasicComparator[string], "foo"),
			iteratorInitOrderedFirst: (*Heap[string]).OrderedFirst,
			iteratorInitOrderedEnd:   (*Heap[string]).OrderedEnd,
		},
		{
			name:                     "3 items",
			originalList:             New[string](utils.BasicComparator[string], "foo", "bar", "baz"),
			newList:                  New[string](utils.BasicComparator[string], "foo", "bar", "baz"),
			iteratorInitOrderedFirst: (*Heap[string]).OrderedFirst,
			iteratorInitOrderedEnd:   (*Heap[string]).OrderedEnd,
		},
		{
			name:                     "3 items, end and OrderedFirst swapped",
			originalList:             New[string](utils.BasicComparator[string], "foo", "bar", "baz"),
			newList:                  New[string](utils.BasicComparator[string]),
			iteratorInitOrderedFirst: (*Heap[string]).OrderedEnd,
			iteratorInitOrderedEnd:   (*Heap[string]).OrderedFirst,
		},
	}

	for _, test := range tests {
		OrderedFirst := test.originalList.OrderedBegin()
		end := test.originalList.OrderedEnd()
		newList := NewFromIterators[string](utils.BasicComparator[string], OrderedFirst, end)

		assert.ElementsMatchf(t, test.originalList.GetValues(), newList.GetValues(), test.name)
	}

}

func BenchmarkArrayHeapPop(b *testing.B) {
	b.StopTimer()
	variants := []struct {
		name string
		f    func(n int, name string)
	}{
		{
			name: "Ours",
			f: func(n int, name string) {
				m := New[string](utils.BasicComparator[string])
				for i := 0; i < n; i++ {
					m.Push("foo")
				}
				b.StartTimer()
				for i := 0; i < n; i++ {
					m.Pop()
				}
				b.StopTimer()
				require.Equalf(b, 0, m.Size(), name)
			},
		},
		{
			name: "Raw",
			f: func(n int, name string) {
				m := make([]string, 0)
				for i := 0; i < n; i++ {
					m = append(m, "foo")
				}
				b.StartTimer()
				for i := 0; i < n; i++ {
					m = m[:len(m)-1]
				}
				b.StopTimer()
				require.Equalf(b, 0, len(m), name)
			},
		},
	}

	for _, variant := range variants {
		testCommon.RunBenchmarkWithDefualtInputSizes(b, variant.name, variant.f)
	}
}
