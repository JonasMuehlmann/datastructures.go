// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package arraylist

import (
	"testing"

	testCommon "github.com/JonasMuehlmann/datastructures.go/tests"

	"github.com/JonasMuehlmann/datastructures.go/ds"

	"github.com/JonasMuehlmann/datastructures.go/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestArrayListGet(t *testing.T) {
	tests := []struct {
		name         string
		originalList *List[string]
		position     int
		value        string
		found        bool
	}{
		{
			name:         "empty list",
			originalList: New[string](),
			position:     0,
			found:        false,
		},
		{
			name:         "3 items, position out of bounds",
			originalList: NewFromSlice[string]([]string{"foo", "bar", "baz"}),
			position:     5,
			found:        false,
		},
		{
			name:         "3 items, position in middle",
			originalList: NewFromSlice[string]([]string{"foo", "bar", "baz"}),
			value:        "bar",
			position:     1,
			found:        true,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			value, found := test.originalList.Get(test.position)

			assert.Equalf(t, test.value, value, test.name)
			assert.Equalf(t, test.found, found, test.name)
		})
	}
}

func TestArrayListContains(t *testing.T) {
	tests := []struct {
		name         string
		originalList *List[string]
		value        string
		found        bool
	}{
		{
			name:         "empty list",
			originalList: New[string](),
			value:        "foo",
			found:        false,
		},
		{
			name:         "3 items, not found",
			originalList: NewFromSlice[string]([]string{"foo", "bar", "baz"}),
			value:        "golang",
			found:        false,
		},
		{
			name:         "3 items, found",
			originalList: NewFromSlice[string]([]string{"foo", "bar", "baz"}),
			value:        "bar",
			found:        true,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			found := test.originalList.Contains(utils.BasicComparator[string], test.value)

			assert.Equalf(t, test.found, found, test.name)
		})
	}
}

func TestArrayListIndexOf(t *testing.T) {
	tests := []struct {
		name         string
		originalList *List[string]
		value        string
		position     int
	}{
		{
			name:         "empty list",
			originalList: New[string](),
			value:        "foo",
			position:     -1,
		},
		{
			name:         "3 items, not found",
			originalList: NewFromSlice[string]([]string{"foo", "bar", "baz"}),
			value:        "golang",
			position:     -1,
		},
		{
			name:         "3 items, found",
			originalList: NewFromSlice[string]([]string{"foo", "bar", "baz"}),
			value:        "bar",
			position:     1,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			position := test.originalList.IndexOf(utils.BasicComparator[string], test.value)

			assert.Equalf(t, test.position, position, test.name)
		})
	}
}

func TestArrayListGetValues(t *testing.T) {
	tests := []struct {
		name         string
		originalList *List[string]
	}{
		{
			name:         "empty list",
			originalList: New[string](),
		},
		{
			name:         "3 items, not found",
			originalList: NewFromSlice[string]([]string{"foo", "bar", "baz"}),
		},
		{
			name:         "3 items, found",
			originalList: NewFromSlice[string]([]string{"foo", "bar", "baz"}),
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			values := test.originalList.GetValues()

			assert.Equalf(t, test.originalList.elements, values, test.name)
		})
	}
}

func TestArrayListGetSlice(t *testing.T) {
	tests := []struct {
		name         string
		originalList *List[string]
	}{
		{
			name:         "empty list",
			originalList: New[string](),
		},
		{
			name:         "3 items, found",
			originalList: NewFromSlice[string]([]string{"foo", "bar", "baz"}),
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			slice := test.originalList.GetSlice()

			assert.Equalf(t, test.originalList.elements, slice, test.name)
		})
	}
}

func TestArrayListIsEmpty(t *testing.T) {
	tests := []struct {
		name         string
		originalList *List[string]
		isEmpty      bool
	}{
		{
			name:         "empty list",
			originalList: New[string](),
			isEmpty:      true,
		},
		{
			name:         "3 items, found",
			originalList: NewFromSlice[string]([]string{"foo", "bar", "baz"}),
			isEmpty:      false,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			isEmpty := test.originalList.IsEmpty()

			assert.Equalf(t, test.isEmpty, isEmpty, test.name)
		})
	}
}

func TestArrayListClear(t *testing.T) {
	tests := []struct {
		name         string
		originalList *List[string]
	}{
		{
			name:         "empty list",
			originalList: New[string](),
		},
		{
			name:         "3 items, found",
			originalList: NewFromSlice[string]([]string{"foo", "bar", "baz"}),
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			isEmpty := test.originalList.IsEmpty()
			assert.Equalf(t, len(test.originalList.elements) == 0, isEmpty, test.name)

			test.originalList.Clear()

			isEmpty = test.originalList.IsEmpty()
			assert.Truef(t, isEmpty, test.name)
		})
	}
}

func TestArrayListSet(t *testing.T) {
	tests := []struct {
		name         string
		originalList *List[string]
		value        string
		position     int
		successfull  bool
	}{
		{
			name:         "empty list, set position 0",
			originalList: New[string](),
			value:        "foo",
			position:     0,
			successfull:  true,
		},
		{
			name:         "empty list, set position 5",
			originalList: New[string](),
			value:        "foo",
			position:     5,
			successfull:  false,
		},
		{
			name:         "position out of bounds",
			originalList: New[string]("foo", "bar", "baz"),
			value:        "foo",
			position:     -1,
			successfull:  false,
		},

		{
			name:         "3 items, set middle",
			originalList: NewFromSlice[string]([]string{"foo", "bar", "baz"}),
			value:        "golang",
			position:     1,
			successfull:  true,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			test.originalList.Set(test.position, test.value)

			index := test.originalList.IndexOf(utils.BasicComparator[string], test.value)

			assert.Equalf(t, test.successfull, index == test.position, test.name)
		})
	}
}

func TestArrayListInsert(t *testing.T) {
	tests := []struct {
		name         string
		originalList *List[string]
		newList      *List[string]
		value        string
		position     int
	}{
		{
			name:         "empty list, set position 0",
			originalList: New[string](),
			newList:      New[string]("foo"),
			value:        "foo",
			position:     0,
		},
		{
			name:         "empty list, set position 5",
			originalList: New[string](),
			newList:      New[string](),
			value:        "foo",
			position:     5,
		},
		{
			name:         "position out of bounds",
			originalList: New[string]("foo", "bar", "baz"),
			newList:      New[string]("foo", "bar", "baz"),
			value:        "foo",
			position:     -1,
		},

		{
			name:         "3 items, insert in middle",
			originalList: NewFromSlice[string]([]string{"foo", "bar", "baz"}),
			newList:      NewFromSlice[string]([]string{"foo", "golang", "bar", "baz"}),
			value:        "golang",
			position:     1,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			test.originalList.Insert(test.position, test.value)

			assert.ElementsMatch(t, test.newList.elements, test.originalList.elements, test.name)
		})
	}
}

func TestArrayListSwap(t *testing.T) {
	tests := []struct {
		name         string
		originalList *List[string]
		newList      *List[string]
		position1    int
		position2    int
	}{
		{
			name:         "empty list",
			originalList: New[string](),
			newList:      New[string](),
			position1:    1,
			position2:    2,
		},
		{
			name:         "3 items, position 1 out of bounds",
			originalList: NewFromSlice[string]([]string{"foo", "golang", "bar", "baz"}),
			newList:      NewFromSlice[string]([]string{"foo", "golang", "bar", "baz"}),
			position1:    5,
			position2:    1,
		}, {
			name:         "3 items, position 2 out of bounds",
			originalList: NewFromSlice[string]([]string{"foo", "golang", "bar", "baz"}),
			newList:      NewFromSlice[string]([]string{"foo", "golang", "bar", "baz"}),
			position1:    1,
			position2:    5,
		},
		{
			name:         "3 items, swap first two",
			originalList: NewFromSlice[string]([]string{"foo", "golang", "bar", "baz"}),
			newList:      NewFromSlice[string]([]string{"foo", "golang", "bar", "baz"}),
			position1:    0,
			position2:    1,
		}, {
			name:         "3 items, swap first and last",
			originalList: NewFromSlice[string]([]string{"foo", "golang", "bar", "baz"}),
			newList:      NewFromSlice[string]([]string{"foo", "golang", "bar", "baz"}),
			position1:    0,
			position2:    2,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			test.originalList.Swap(test.position1, test.position2)

			assert.ElementsMatch(t, test.newList.elements, test.originalList.elements, test.name)
		})
	}
}

func TestArrayListSort(t *testing.T) {
	tests := []struct {
		name         string
		originalList *List[string]
		newList      *List[string]
	}{
		{
			name:         "empty list",
			originalList: New[string](),
			newList:      New[string](),
		},
		{
			name:         "single item",
			originalList: New[string]("foo"),
			newList:      New[string]("foo"),
		},
		{
			name:         "two items",
			originalList: New[string]("foo", "bar"),
			newList:      New[string]("bar", "foo"),
		},

		{
			name:         "3 items, unchanged",
			originalList: NewFromSlice[string]([]string{"bar", "baz", "foo"}),
			newList:      NewFromSlice[string]([]string{"bar", "baz", "foo"}),
		},
		{
			name:         "3 items, reverse",
			originalList: NewFromSlice[string]([]string{"foo", "baz", "bar"}),
			newList:      NewFromSlice[string]([]string{"bar", "baz", "foo"}),
		},
		{
			name:         "3 items, random order",
			originalList: NewFromSlice[string]([]string{"baz", "bar", "foo"}),
			newList:      NewFromSlice[string]([]string{"bar", "baz", "foo"}),
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			test.originalList.Sort(utils.BasicComparator[string])

			assert.ElementsMatch(t, test.newList.elements, test.originalList.elements, test.name)
		})
	}
}

func TestArrayListPushFront(t *testing.T) {
	tests := []struct {
		name         string
		originalList *List[string]
		itemsToAdd   []string
		newItems     []string
	}{
		{
			name:         "empty list, add nothing",
			originalList: New[string](),
			itemsToAdd:   []string{},
			newItems:     []string{},
		},
		{
			name:         "empty list, add 2",
			originalList: New[string](),
			itemsToAdd:   []string{"foo", "bar"},
			newItems:     []string{"foo", "bar"},
		},
		{
			name:         "list with 2 items, add nothing",
			originalList: New[string]("foo", "bar"),
			itemsToAdd:   []string{},
			newItems:     []string{"foo", "bar"},
		},
		{
			name:         "list with 2 items, add 2",
			originalList: New[string]("foo", "bar"),
			itemsToAdd:   []string{"foo2", "bar2"},
			newItems:     []string{"foo2", "bar2", "foo", "bar"},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			test.originalList.PushFront(test.itemsToAdd...)

			assert.Equalf(t, test.originalList.elements, test.newItems, test.name)
		})
	}
}

func TestArrayListPopBack(t *testing.T) {
	tests := []struct {
		name         string
		originalList *List[string]
		n            int
		newItems     []string
	}{
		{
			name:         "empty list, remove nothing",
			originalList: New[string](),
			newItems:     []string{},
		},
		{
			name:         "empty list, remove 2",
			originalList: New[string](),
			n:            2,
			newItems:     []string{},
		},
		{
			name:         "list with 2 items, remove nothing",
			originalList: New[string]("foo", "bar"),
			n:            0,
			newItems:     []string{"foo", "bar"},
		},
		{
			name:         "list with 4 items, remove 2",
			originalList: New[string]("foo", "bar", "baz", "foo"),
			n:            2,
			newItems:     []string{"foo", "bar"},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			test.originalList.PopBack(test.n)

			assert.Equalf(t, test.originalList.elements, test.newItems, test.name)
		})
	}
}

func TestArrayListPopFront(t *testing.T) {
	tests := []struct {
		name         string
		originalList *List[string]
		n            int
		newItems     []string
	}{
		{
			name:         "empty list, remove nothing",
			originalList: New[string](),
			newItems:     []string{},
		},
		{
			name:         "empty list, remove 2",
			originalList: New[string](),
			n:            2,
			newItems:     []string{},
		},
		{
			name:         "list with 2 items, remove nothing",
			originalList: New[string]("foo", "bar"),
			n:            0,
			newItems:     []string{"foo", "bar"},
		},
		{
			name:         "list with 4 items, remove 2",
			originalList: New[string]("foo", "bar", "baz", "foo"),
			n:            2,
			newItems:     []string{"baz", "foo"},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			test.originalList.PopFront(test.n)

			assert.Equalf(t, test.originalList.elements, test.newItems, test.name)
		})
	}
}

func TestArrayListShrinkToFit(t *testing.T) {
	tests := []struct {
		name         string
		originalList *List[string]
		n            int
		newLen       int
		newCap       int
	}{
		{
			name:         "empty list, remove nothing",
			originalList: New[string](),
			newLen:       0,
			newCap:       0,
		},
		{
			name:         "empty list, remove 2",
			originalList: New[string](),
			n:            2,
			newLen:       0,
			newCap:       0,
		},
		{
			name:         "list with 2 items, remove nothing",
			originalList: New[string]("foo", "bar"),
			n:            0,
			newLen:       2,
			newCap:       2,
		},
		{
			name:         "list with 4 items, remove 2",
			originalList: New[string]("foo", "bar", "baz", "foo"),
			n:            2,
			newLen:       2,
			newCap:       2,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			test.originalList.PopBack(test.n)
			test.originalList.ShrinkToFit()

			assert.Equal(t, test.newLen, len(test.originalList.elements), test.name)
			assert.Equal(t, test.newCap, cap(test.originalList.elements), test.name)
		})
	}
}

func TestArrayListRemoveStable(t *testing.T) {
	tests := []struct {
		name         string
		originalList *List[string]
		i            int
		newItems     []string
	}{
		{
			name:         "empty list, remove nothing",
			originalList: New[string](),
			i:            -1,
			newItems:     []string{},
		},
		{
			name:         "empty list, remove 2",
			originalList: New[string](),
			i:            2,
			newItems:     []string{},
		},
		{
			name:         "list with 2 items, remove nothing",
			originalList: New[string]("foo", "bar"),
			i:            -1,
			newItems:     []string{"foo", "bar"},
		},
		{
			name:         "list with 4 items, remove 2",
			originalList: New[string]("foo", "bar", "baz", "foo"),
			i:            2,
			newItems:     []string{"foo", "bar", "foo"},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			test.originalList.RemoveStable(test.i)

			assert.Equalf(t, test.originalList.elements, test.newItems, test.name)
		})
	}
}

func TestArrayListRemove(t *testing.T) {
	tests := []struct {
		name         string
		originalList *List[string]
		i            int
		newItems     []string
	}{
		{
			name:         "empty list, remove nothing",
			originalList: New[string](),
			i:            -1,
			newItems:     []string{},
		},
		{
			name:         "empty list, remove 2",
			originalList: New[string](),
			i:            2,
			newItems:     []string{},
		},
		{
			name:         "list with 2 items, remove nothing",
			originalList: New[string]("foo", "bar"),
			i:            -1,
			newItems:     []string{"foo", "bar"},
		},
		{
			name:         "list with 4 items, remove 2",
			originalList: New[string]("foo", "bar", "baz", "foo"),
			i:            2,
			newItems:     []string{"foo", "bar", "foo"},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			test.originalList.Remove(test.i)

			assert.Equalf(t, test.originalList.elements, test.newItems, test.name)
		})
	}
}

func TestArrayListNewFromSlice(t *testing.T) {
	tests := []struct {
		name         string
		originalList *List[string]
	}{
		{
			name:         "empty list",
			originalList: New[string](),
		},
		{
			name:         "single item",
			originalList: New[string]("foo"),
		},
		{
			name:         "3 items",
			originalList: New[string]("foo", "bar", "baz"),
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			newList := NewFromSlice[string](test.originalList.elements)

			assert.Equalf(t, test.originalList.elements, newList.elements, test.name)
		})
	}

}

func TestArrayListNewFromIterator(t *testing.T) {
	tests := []struct {
		name         string
		originalList *List[string]
	}{
		{
			name:         "empty list",
			originalList: New[string](),
		},
		{
			name:         "single item",
			originalList: New[string]("foo"),
		},
		{
			name:         "3 items",
			originalList: New[string]("foo", "bar", "baz"),
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it := test.originalList.Begin()
			newList := NewFromIterator[string](it)

			assert.Equalf(t, test.originalList.elements, newList.elements, test.name)
		})
	}

}

// NOTE: Missing test case: unordered iterator, which prevents preallocation
func TestArrayListNewFromIterators(t *testing.T) {
	tests := []struct {
		name              string
		originalList      *List[string]
		newList           *List[string]
		iteratorInitFirst func(*List[string]) ds.ReadWriteOrdCompBidRandCollIterator[int, string]
		iteratorInitEnd   func(*List[string]) ds.ReadWriteOrdCompBidRandCollIterator[int, string]
	}{
		{
			name:              "empty list",
			originalList:      New[string](),
			newList:           New[string](),
			iteratorInitFirst: (*List[string]).First,
			iteratorInitEnd:   (*List[string]).End,
		},
		{
			name:              "single item",
			originalList:      New[string]("foo"),
			iteratorInitFirst: (*List[string]).First,
			iteratorInitEnd:   (*List[string]).End,
		},
		{
			name:              "3 items",
			originalList:      New[string]("foo", "bar", "baz"),
			newList:           New[string]("foo", "bar", "baz"),
			iteratorInitFirst: (*List[string]).First,
			iteratorInitEnd:   (*List[string]).End,
		},
		{
			name:              "3 items, end and first swapped",
			originalList:      New[string]("foo", "bar", "baz"),
			newList:           New[string](),
			iteratorInitFirst: (*List[string]).End,
			iteratorInitEnd:   (*List[string]).First,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			first := test.originalList.Begin()
			end := test.originalList.End()
			newList := NewFromIterators[string](first, end)

			assert.Equalf(t, test.originalList.elements, newList.elements, test.name)
		})
	}

}

func BenchmarkArrayListGet(b *testing.B) {
	b.StopTimer()
	variants := []struct {
		name string
		f    func(n int, name string)
	}{
		{
			name: "Ours",
			f: func(n int, name string) {
				m := New[string]()
				for i := 0; i < n; i++ {
					m.Set(i, "foo")
				}
				b.StartTimer()
				for i := 0; i < n; i++ {
					_, _ = m.Get(i)
				}
				b.StopTimer()
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
					_ = m[i]
				}
				b.StopTimer()
			},
		},
	}

	for _, variant := range variants {
		testCommon.RunBenchmarkWithDefualtInputSizes(b, variant.name, variant.f)
	}
}

func BenchmarkArrayListPushBack(b *testing.B) {
	b.StopTimer()
	variants := []struct {
		name string
		f    func(n int, name string)
	}{
		{
			name: "Ours",
			f: func(n int, name string) {
				m := New[string]()
				b.StartTimer()
				for i := 0; i < n; i++ {
					m.PushBack("foo")
				}
				b.StopTimer()
				require.Equalf(b, n, len(m.elements), name)
			},
		},
		{
			name: "Raw",
			f: func(n int, name string) {
				m := make([]string, 0)
				b.StartTimer()
				for i := 0; i < n; i++ {
					m = append(m, "foo")
				}
				b.StopTimer()
				require.Equalf(b, n, len(m), name)
			},
		},
	}

	for _, variant := range variants {
		testCommon.RunBenchmarkWithDefualtInputSizes(b, variant.name, variant.f)
	}
}

func BenchmarkArrayListPushFront(b *testing.B) {
	b.StopTimer()
	variants := []struct {
		name string
		f    func(n int, name string)
	}{
		{
			name: "Ours",
			f: func(n int, name string) {
				m := New[string]()
				b.StartTimer()
				for i := 0; i < n; i++ {
					m.PushFront("foo")
				}
				b.StopTimer()
				require.Equalf(b, n, len(m.elements), name)
			},
		},
		{
			name: "Raw",
			f: func(n int, name string) {
				m := make([]string, 0)
				b.StartTimer()
				for i := 0; i < n; i++ {
					m = append([]string{"foo"}, m...)
				}
				b.StopTimer()
				require.Equalf(b, n, len(m), name)
			},
		},
	}

	for _, variant := range variants {
		testCommon.RunBenchmarkWithDefualtInputSizes(b, variant.name, variant.f)
	}
}

func BenchmarkArrayListRemoveStable(b *testing.B) {
	b.StopTimer()
	variants := []struct {
		name string
		f    func(n int, name string)
	}{
		{
			name: "Ours",
			f: func(n int, name string) {
				m := New[string]()
				for i := 0; i < n; i++ {
					m.PushBack("foo")
				}
				b.StartTimer()
				for i := 0; i < n-1; i++ {
					m.RemoveStable(1)
				}
				b.StopTimer()
				require.Equalf(b, 1, len(m.elements), name)
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
				for i := 0; i < n-1; i++ {
					m = append(m[:1], m[2:]...)
				}
				b.StopTimer()
				require.Equalf(b, 1, len(m), name)
			},
		},
	}

	for _, variant := range variants {
		testCommon.RunBenchmarkWithDefualtInputSizes(b, variant.name, variant.f)
	}
}

func BenchmarkArrayListRemove(b *testing.B) {
	b.StopTimer()
	variants := []struct {
		name string
		f    func(n int, name string)
	}{
		{
			name: "Ours",
			f: func(n int, name string) {
				m := New[string]()
				for i := 0; i < n; i++ {
					m.PushBack("foo")
				}
				b.StartTimer()
				for i := 0; i < n-1; i++ {
					m.Remove(1)
				}
				b.StopTimer()
				require.Equalf(b, 1, len(m.elements), name)
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
				for i := 0; i < n-1; i++ {

					m[1] = m[len(m)-1]
					m = m[:len(m)-1]
				}
				b.StopTimer()
				require.Equalf(b, 1, len(m), name)
			},
		},
	}

	for _, variant := range variants {
		testCommon.RunBenchmarkWithDefualtInputSizes(b, variant.name, variant.f)
	}
}

func BenchmarkArrayListPopBack(b *testing.B) {
	b.StopTimer()
	variants := []struct {
		name string
		f    func(n int, name string)
	}{
		{
			name: "Ours",
			f: func(n int, name string) {
				m := New[string]()
				for i := 0; i < n; i++ {
					m.PushBack("foo")
				}
				b.StartTimer()
				for i := 0; i < n; i++ {
					m.PopBack(1)
				}
				b.StopTimer()
				require.Equalf(b, 0, len(m.elements), name)
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

func BenchmarkArrayListPopFront(b *testing.B) {
	b.StopTimer()
	variants := []struct {
		name string
		f    func(n int, name string)
	}{
		{
			name: "Ours",
			f: func(n int, name string) {
				m := New[string]()
				for i := 0; i < n; i++ {
					m.PushBack("foo")
				}
				b.StartTimer()
				for i := 0; i < n; i++ {
					m.PopFront(1)
				}
				b.StopTimer()
				require.Equalf(b, 0, len(m.elements), name)
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
					m = m[1:]
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
