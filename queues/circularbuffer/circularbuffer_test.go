// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package circularbuffer

import (
	"testing"

	testCommon "github.com/JonasMuehlmann/datastructures.go/tests"

	"github.com/JonasMuehlmann/datastructures.go/ds"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// func TestCircularBufferContains(t *testing.T) {
// 	tests := []struct {
// 		name         string
// 		originalList *Queue[string]
// 		value        string
// 		found        bool
// 	}{
// 		{
// 			name:         "empty list",
// 			originalList: New[string](5),
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

// func TestCircularBufferIndexOf(t *testing.T) {
// 	tests := []struct {
// 		name         string
// 		originalList *Queue[string]
// 		value        string
// 		position     int
// 	}{
// 		{
// 			name:         "empty list",
// 			originalList: New[string](5),
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

func TestCircularBufferGetValues(t *testing.T) {
	tests := []struct {
		name         string
		originalList *Queue[string]
	}{
		{
			name:         "empty list",
			originalList: New[string](5),
		},
		{
			name:         "3 items, not found",
			originalList: NewFromSlice[string](10, []string{"foo", "bar", "baz"}),
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			values := test.originalList.GetValues()

			assert.Equalf(t, test.originalList.GetValues(), values, test.name)
		})
	}
}

func TestCircularBufferIsEmpty(t *testing.T) {
	tests := []struct {
		name         string
		originalList *Queue[string]
		isEmpty      bool
	}{
		{
			name:         "empty list",
			originalList: New[string](5),
			isEmpty:      true,
		},
		{
			name:         "3 items, found",
			originalList: NewFromSlice[string](10, []string{"foo", "bar", "baz"}),
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

func TestCircularBufferClear(t *testing.T) {
	tests := []struct {
		name         string
		originalList *Queue[string]
	}{
		{
			name:         "empty list",
			originalList: New[string](5),
		},
		{
			name:         "3 items, found",
			originalList: NewFromSlice[string](10, []string{"foo", "bar", "baz"}),
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			isEmpty := test.originalList.IsEmpty()
			assert.Equalf(t, test.originalList.Size() == 0, isEmpty, test.name)

			test.originalList.Clear()

			isEmpty = test.originalList.IsEmpty()
			assert.Truef(t, isEmpty, test.name)
		})
	}
}

func TestCircularBufferEnqueue(t *testing.T) {
	tests := []struct {
		name         string
		originalList *Queue[string]
		valueToAdd   string
		newItems     []string
	}{
		{
			name:         "empty list",
			originalList: New[string](5),
			valueToAdd:   "foo",
			newItems:     []string{"foo"},
		},
		{
			name:         "1 item",
			originalList: NewFromSlice[string](10, []string{"foo"}),
			valueToAdd:   "bar",
			newItems:     []string{"foo", "bar"},
		},

		{
			name:         "list with 4 items, remove 1",
			originalList: NewFromSlice[string](10, []string{"foo", "bar", "baz"}),
			valueToAdd:   "foo",
			newItems:     []string{"foo", "bar", "baz", "foo"},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			test.originalList.Enqueue(test.valueToAdd)

			assert.Equalf(t, test.originalList.GetValues(), test.newItems, test.name)
		})
	}
}

func TestCircularBufferDequeue(t *testing.T) {
	tests := []struct {
		name         string
		originalList *Queue[string]
		newItems     []string
	}{
		{
			name:         "empty list",
			originalList: New[string](5),
			newItems:     []string{},
		},
		{
			name:         "1 item",
			originalList: NewFromSlice[string](10, []string{"foo"}),
			newItems:     []string{},
		},

		{
			name:         "list with 4 items, remove 1",
			originalList: NewFromSlice[string](10, []string{"foo", "bar", "baz", "foo"}),
			newItems:     []string{"bar", "baz", "foo"},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			test.originalList.Dequeue()

			assert.Equalf(t, test.originalList.GetValues(), test.newItems, test.name)
		})
	}
}

func TestCircularBufferPeek(t *testing.T) {
	tests := []struct {
		name         string
		originalList *Queue[string]
		found        bool
		value        string
	}{
		{
			name:         "empty list",
			originalList: New[string](5),
			found:        false,
		},
		{
			name:         "1 item",
			originalList: NewFromSlice[string](10, []string{"foo"}),
			found:        true,
			value:        "foo",
		},

		{
			name:         "list with 4 items",
			originalList: NewFromSlice[string](10, []string{"foo", "bar", "baz", "foo"}),
			found:        true,
			value:        "foo",
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
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
			originalList: New[string](5),
		},
		{
			name:         "single item",
			originalList: NewFromSlice[string](10, []string{"foo"}),
		},
		{
			name:         "3 items",
			originalList: NewFromSlice[string](10, []string{"foo", "bar", "baz"}),
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			newList := NewFromSlice[string](10, test.originalList.GetValues())

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
			originalList: New[string](5),
		},
		{
			name:         "single item",
			originalList: NewFromSlice[string](10, []string{"foo"}),
		},
		{
			name:         "3 items",
			originalList: NewFromSlice[string](10, []string{"foo", "bar", "baz"}),
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it := test.originalList.Begin()
			newList := NewFromIterator[string](10, it)

			assert.Equalf(t, test.originalList.GetValues(), newList.GetValues(), test.name)
		})
	}

}

// NOTE: Missing test case: unordered iterator, which prevents preallocation
func TestNewFromIterators(t *testing.T) {
	tests := []struct {
		name              string
		originalList      *Queue[string]
		iteratorInitFirst func(*Queue[string]) ds.ReadWriteOrdCompBidRandCollIterator[int, string]
		iteratorInitEnd   func(*Queue[string]) ds.ReadWriteOrdCompBidRandCollIterator[int, string]
	}{
		{
			name:              "empty list",
			originalList:      New[string](10),
			iteratorInitFirst: (*Queue[string]).Begin,
			iteratorInitEnd:   (*Queue[string]).End,
		},
		{
			name:              "single item",
			originalList:      NewFromSlice[string](10, []string{"foo"}),
			iteratorInitFirst: (*Queue[string]).Begin,
			iteratorInitEnd:   (*Queue[string]).End,
		},
		{
			name:              "3 items",
			originalList:      NewFromSlice[string](10, []string{"foo", "bar", "baz"}),
			iteratorInitFirst: (*Queue[string]).Begin,
			iteratorInitEnd:   (*Queue[string]).End,
		},
		{
			name:              "3 items, end and first swapped",
			originalList:      NewFromSlice[string](10, []string{"foo", "bar", "baz"}),
			iteratorInitFirst: (*Queue[string]).End,
			iteratorInitEnd:   (*Queue[string]).Begin,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			first := test.originalList.Begin()
			end := test.originalList.End()
			newList := NewFromIterators[string](10, first, end)

			assert.Equalf(t, test.originalList.GetValues(), newList.GetValues(), test.name)
		})
	}

}

func BenchmarkCircularBufferDequeue(b *testing.B) {
	b.StopTimer()
	variants := []struct {
		name string
		f    func(n int, name string)
	}{
		{
			name: "Ours",
			f: func(n int, name string) {
				m := New[string](5)
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
		testCommon.RunBenchmarkWithDefualtInputSizes(b, variant.name, variant.f)
	}
}
