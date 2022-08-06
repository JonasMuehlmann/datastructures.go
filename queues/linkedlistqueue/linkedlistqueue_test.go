// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linkedlistqueue

import (
	"testing"

	"github.com/JonasMuehlmann/datastructures.go/ds"
	"github.com/JonasMuehlmann/datastructures.go/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// func TestArrayQueueContains(t *testing.T) {
// 	tests := []struct {
// 		name         string
// 		originalList *Queue[string]
// 		value        string
// 		found        bool
// 	}{
// 		{
// 			name:         "empty list",
// 			originalList: New[string](),
// 			value:        "foo",
// 			found:        false,
// 		},
// 		{
// 			name:         "3 items, not found",
// 			originalList: NewFromSlice[string]([]string{"foo", "bar", "baz"}),
// 			value:        "golang",
// 			found:        false,
// 		},
// 		{
// 			name:         "3 items, found",
// 			originalList: NewFromSlice[string]([]string{"foo", "bar", "baz"}),
// 			value:        "bar",
// 			found:        true,
// 		},
// 	}

// 	for _, test := range tests {

// 		assert.Equalf(t, test.found, found, test.name)
// 	}
// }

// func TestArrayQueueIndexOf(t *testing.T) {
// 	tests := []struct {
// 		name         string
// 		originalList *Queue[string]
// 		value        string
// 		position     int
// 	}{
// 		{
// 			name:         "empty list",
// 			originalList: New[string](),
// 			value:        "foo",
// 			position:     -1,
// 		},
// 		{
// 			name:         "3 items, not found",
// 			originalList: NewFromSlice[string]([]string{"foo", "bar", "baz"}),
// 			value:        "golang",
// 			position:     -1,
// 		},
// 		{
// 			name:         "3 items, found",
// 			originalList: NewFromSlice[string]([]string{"foo", "bar", "baz"}),
// 			value:        "bar",
// 			position:     1,
// 		},
// 	}

// 	for _, test := range tests {

// 		position := test.originalList.IndexOf(utils.BasicComparator[string], test.value)

// 		assert.Equalf(t, test.position, position, test.name)
// 	}
// }

func TestArrayQueueGetValues(t *testing.T) {
	tests := []struct {
		name         string
		originalList *Queue[string]
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
			values := test.originalList.GetValues()

			assert.Equalf(t, test.originalList.list.GetValues(), values, test.name)
		})
	}
}

func TestArrayQueueIsEmpty(t *testing.T) {
	tests := []struct {
		name         string
		originalList *Queue[string]
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
			isEmpty := test.originalList.IsEmpty()

			assert.Equalf(t, test.isEmpty, isEmpty, test.name)
		})
	}
}

func TestArrayQueueClear(t *testing.T) {
	tests := []struct {
		name         string
		originalList *Queue[string]
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
			isEmpty := test.originalList.IsEmpty()
			assert.Equalf(t, test.originalList.Size() == 0, isEmpty, test.name)

			test.originalList.Clear()

			isEmpty = test.originalList.IsEmpty()
			assert.Truef(t, isEmpty, test.name)
		})
	}
}

func TestArrayQueueEnqueue(t *testing.T) {
	tests := []struct {
		name         string
		originalList *Queue[string]
		valueToAdd   string
		newItems     []string
	}{
		{
			name:         "empty list",
			originalList: New[string](),
			valueToAdd:   "foo",
			newItems:     []string{"foo"},
		},
		{
			name:         "1 item",
			originalList: New[string]("foo"),
			valueToAdd:   "bar",
			newItems:     []string{"foo", "bar"},
		},

		{
			name:         "list with 4 items, remove 1",
			originalList: New[string]("foo", "bar", "baz"),
			valueToAdd:   "foo",
			newItems:     []string{"foo", "bar", "baz", "foo"},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {

			t.Parallel()
			test.originalList.Enqueue(test.valueToAdd)

			assert.Equalf(t, test.originalList.GetValues(), test.newItems, test.name)
		})
	}
}

func TestArrayQueueDequeue(t *testing.T) {
	tests := []struct {
		name         string
		originalList *Queue[string]
		newItems     []string
	}{
		{
			name:         "empty list",
			originalList: New[string](),
			newItems:     []string{},
		},
		{
			name:         "1 item",
			originalList: New[string]("foo"),
			newItems:     []string{},
		},

		{
			name:         "list with 4 items, remove 1",
			originalList: New[string]("foo", "bar", "baz", "foo"),
			newItems:     []string{"bar", "baz", "foo"},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {

			t.Parallel()
			test.originalList.Dequeue()

			assert.Equalf(t, test.originalList.GetValues(), test.newItems, test.name)
		})
	}
}

func TestArrayQueuePeek(t *testing.T) {
	tests := []struct {
		name         string
		originalList *Queue[string]
		found        bool
		value        string
	}{
		{
			name:         "empty list",
			originalList: New[string](),
			found:        false,
		},
		{
			name:         "1 item",
			originalList: New[string]("foo"),
			found:        true,
			value:        "foo",
		},

		{
			name:         "list with 4 items",
			originalList: New[string]("foo", "bar", "baz", "foo"),
			found:        true,
			value:        "foo",
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {

			t.Parallel()
			value, found := test.originalList.Peek()

			assert.Equalf(t, test.found, found, test.name)

			if test.found {
				assert.Equalf(t, test.value, value, test.name)
			}

		})
	}
}

func TestNewFromSlice(t *testing.T) {
	tests := []struct {
		name         string
		originalList *Queue[string]
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
			newList := NewFromSlice[string](test.originalList.GetValues())

			assert.Equalf(t, test.originalList.GetValues(), newList.GetValues(), test.name)
		})
	}

}

func TestNewFromIterator(t *testing.T) {
	tests := []struct {
		name         string
		originalList *Queue[string]
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
			it := test.originalList.Begin()
			newList := NewFromIterator[string](it)

			assert.Equalf(t, test.originalList.GetValues(), newList.GetValues(), test.name)
		})
	}

}

// NOTE: Missing test case: unordered iterator, which prevents preallocation
func TestNewFromIterators(t *testing.T) {
	tests := []struct {
		name              string
		originalList      *Queue[string]
		newList           *Queue[string]
		iteratorInitFirst func(*Queue[string]) ds.ReadWriteOrdCompForRandCollIterator[int, string]
		iteratorInitEnd   func(*Queue[string]) ds.ReadWriteOrdCompForRandCollIterator[int, string]
	}{
		{
			name:              "empty list",
			originalList:      New[string](),
			newList:           New[string](),
			iteratorInitFirst: (*Queue[string]).Begin,
			iteratorInitEnd:   (*Queue[string]).End,
		},
		{
			name:              "single item",
			originalList:      New[string]("foo"),
			iteratorInitFirst: (*Queue[string]).Begin,
			iteratorInitEnd:   (*Queue[string]).End,
		},
		{
			name:              "3 items",
			originalList:      New[string]("foo", "bar", "baz"),
			newList:           New[string]("foo", "bar", "baz"),
			iteratorInitFirst: (*Queue[string]).Begin,
			iteratorInitEnd:   (*Queue[string]).End,
		},
		{
			name:              "3 items, end and first swapped",
			originalList:      New[string]("foo", "bar", "baz"),
			newList:           New[string](),
			iteratorInitFirst: (*Queue[string]).End,
			iteratorInitEnd:   (*Queue[string]).First,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {

			t.Parallel()
			first := test.originalList.Begin()
			end := test.originalList.End()
			newList := NewFromIterators[string](first, end)

			assert.Equalf(t, test.originalList.GetValues(), newList.GetValues(), test.name)
		})
	}

}

func BenchmarkArrayQueueDequeue(b *testing.B) {
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
					m.Enqueue("foo")
				}
				b.StartTimer()
				for i := 0; i < n; i++ {
					m.Dequeue()
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
		tests.RunBenchmarkWithDefualtInputSizes(b, variant.name, variant.f)
	}
}
