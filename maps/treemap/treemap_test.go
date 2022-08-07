// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package treemap

import (
	"testing"

	testCommon "github.com/JonasMuehlmann/datastructures.go/tests"

	"github.com/JonasMuehlmann/datastructures.go/utils"
	"github.com/stretchr/testify/assert"
)

func TestTreeMapRemove(t *testing.T) {
	tests := []struct {
		name        string
		originalMap *Map[string, int]
		newMap      *Map[string, int]
		toRemove    string
	}{

		{
			name:        "empty list",
			originalMap: New[string, int](utils.BasicComparator[string]),
			newMap:      New[string, int](utils.BasicComparator[string]),
			toRemove:    "foo",
		},
		{
			name:        "single item",
			toRemove:    "foo",
			originalMap: NewFromMap[string, int](utils.BasicComparator[string], map[string]int{"foo": 1}),
			newMap:      New[string, int](utils.BasicComparator[string]),
		},
		{
			name:        "single item, target does not exist",
			toRemove:    "bar",
			originalMap: NewFromMap[string, int](utils.BasicComparator[string], map[string]int{"foo": 1}),
			newMap:      NewFromMap[string, int](utils.BasicComparator[string], map[string]int{"foo": 1}),
		},
		{
			name:        "3 items",
			toRemove:    "bar",
			originalMap: NewFromMap[string, int](utils.BasicComparator[string], map[string]int{"foo": 1, "bar": 2, "baz": 3}),
			newMap:      NewFromMap[string, int](utils.BasicComparator[string], map[string]int{"foo": 1, "baz": 3}),
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			test.originalMap.Remove(utils.BasicComparator[string], test.toRemove)

			assert.ElementsMatchf(t, test.originalMap.GetKeys(), test.newMap.GetKeys(), test.name)
		})
	}
}

func TestTreeMapPut(t *testing.T) {
	tests := []struct {
		name        string
		originalMap *Map[string, int]
		newMap      *Map[string, int]
		keyToAdd    string
		valueToAdd  int
	}{

		{
			name:        "empty list",
			originalMap: New[string, int](utils.BasicComparator[string]),
			newMap:      NewFromMap[string, int](utils.BasicComparator[string], map[string]int{"foo": 1}),
			keyToAdd:    "foo",
			valueToAdd:  1,
		},
		{
			name:        "single item",
			keyToAdd:    "foo",
			valueToAdd:  1,
			newMap:      NewFromMap[string, int](utils.BasicComparator[string], map[string]int{"foo": 1}),
			originalMap: New[string, int](utils.BasicComparator[string]),
		},
		{
			name:        "single item, overwrite",
			keyToAdd:    "foo",
			valueToAdd:  2,
			originalMap: NewFromMap[string, int](utils.BasicComparator[string], map[string]int{"foo": 1}),
			newMap:      NewFromMap[string, int](utils.BasicComparator[string], map[string]int{"foo": 2}),
		},
		{
			name:        "3 items",
			keyToAdd:    "bar",
			valueToAdd:  2,
			originalMap: NewFromMap[string, int](utils.BasicComparator[string], map[string]int{"foo": 1, "baz": 3}),
			newMap:      NewFromMap[string, int](utils.BasicComparator[string], map[string]int{"foo": 1, "bar": 2, "baz": 3}),
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			test.originalMap.Put(test.keyToAdd, test.valueToAdd)

			assert.ElementsMatchf(t, test.originalMap.GetKeys(), test.newMap.GetKeys(), test.name)
		})
	}
}

func TestTreeMapGet(t *testing.T) {
	tests := []struct {
		name        string
		originalMap *Map[string, int]
		keyToGet    string
		value       int
		found       bool
	}{

		{
			name:        "empty list",
			originalMap: New[string, int](utils.BasicComparator[string]),
			keyToGet:    "foo",
			found:       false,
		},
		{
			name:        "single item",
			keyToGet:    "foo",
			originalMap: NewFromMap[string, int](utils.BasicComparator[string], map[string]int{"foo": 1}),
			value:       1,
			found:       true,
		},
		{
			name:        "single item, target does not exist",
			keyToGet:    "bar",
			originalMap: NewFromMap[string, int](utils.BasicComparator[string], map[string]int{"foo": 1}),
			found:       false,
		},
		{
			name:        "3 items",
			keyToGet:    "bar",
			originalMap: NewFromMap[string, int](utils.BasicComparator[string], map[string]int{"foo": 1, "bar": 2, "baz": 3}),
			value:       2,
			found:       true,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			value, found := test.originalMap.Get(test.keyToGet)

			assert.Equalf(t, test.value, value, test.name)
			assert.Equalf(t, test.found, found, test.name)
		})
	}
}

func TestTreeMapGetKeys(t *testing.T) {
	tests := []struct {
		name        string
		originalMap *Map[string, int]
		keys        []string
	}{

		{
			name:        "empty list",
			originalMap: New[string, int](utils.BasicComparator[string]),
			keys:        []string{},
		},
		{
			name:        "single item",
			originalMap: NewFromMap[string, int](utils.BasicComparator[string], map[string]int{"foo": 1}),
			keys:        []string{"foo"},
		},
		{
			name:        "3 items",
			originalMap: NewFromMap[string, int](utils.BasicComparator[string], map[string]int{"foo": 1, "bar": 2, "baz": 3}),
			keys:        []string{"foo", "bar", "baz"},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			keys := test.originalMap.GetKeys()

			assert.ElementsMatchf(t, test.keys, keys, test.name)
		})
	}
}

func TestTreeMapGetValues(t *testing.T) {
	tests := []struct {
		name        string
		originalMap *Map[string, int]
		values      []int
	}{

		{
			name:        "empty list",
			originalMap: New[string, int](utils.BasicComparator[string]),
			values:      []int{},
		},
		{
			name:        "single item",
			originalMap: NewFromMap[string, int](utils.BasicComparator[string], map[string]int{"foo": 1}),
			values:      []int{1},
		},
		{
			name:        "3 items",
			originalMap: NewFromMap[string, int](utils.BasicComparator[string], map[string]int{"foo": 1, "bar": 2, "baz": 3}),
			values:      []int{1, 2, 3},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			values := test.originalMap.GetValues()

			assert.ElementsMatchf(t, test.values, values, test.name)
		})
	}
}

func TestTreeMapIsEmpty(t *testing.T) {
	tests := []struct {
		name        string
		originalMap *Map[string, int]
		isEmpty     bool
	}{

		{
			name:        "empty list",
			originalMap: New[string, int](utils.BasicComparator[string]),
			isEmpty:     true,
		},
		{
			name:        "single item",
			originalMap: NewFromMap[string, int](utils.BasicComparator[string], map[string]int{"foo": 1}),
			isEmpty:     false,
		},
		{
			name:        "3 items",
			originalMap: NewFromMap[string, int](utils.BasicComparator[string], map[string]int{"foo": 1, "bar": 2, "baz": 3}),
			isEmpty:     false,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			isEmpty := test.originalMap.IsEmpty()

			assert.Equal(t, test.isEmpty, isEmpty, test.name)
		})
	}
}

func TestTreeMapClear(t *testing.T) {
	tests := []struct {
		name          string
		originalMap   *Map[string, int]
		isEmptyBefore bool
		isEmptyAfter  bool
	}{

		{
			name:          "empty list",
			originalMap:   New[string, int](utils.BasicComparator[string]),
			isEmptyBefore: true,
			isEmptyAfter:  true,
		},
		{
			name:          "single item",
			originalMap:   NewFromMap[string, int](utils.BasicComparator[string], map[string]int{"foo": 1}),
			isEmptyBefore: false,
			isEmptyAfter:  true,
		},
		{
			name:          "3 items",
			originalMap:   NewFromMap[string, int](utils.BasicComparator[string], map[string]int{"foo": 1, "bar": 2, "baz": 3}),
			isEmptyBefore: false,
			isEmptyAfter:  true,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			isEmptyBefore := test.originalMap.IsEmpty()
			assert.Equal(t, test.isEmptyBefore, isEmptyBefore, test.name)

			test.originalMap.Clear()

			isEmptAfter := test.originalMap.IsEmpty()
			assert.Equal(t, test.isEmptyAfter, isEmptAfter, test.name)
		})
	}
}

func TestTreeMapNewFromIterator(t *testing.T) {
	tests := []struct {
		name        string
		originalMap *Map[string, int]
	}{

		{
			name:        "empty list",
			originalMap: New[string, int](utils.BasicComparator[string]),
		},
		{
			name:        "single item",
			originalMap: NewFromMap[string, int](utils.BasicComparator[string], map[string]int{"foo": 1}),
		},
		{
			name:        "3 items",
			originalMap: NewFromMap[string, int](utils.BasicComparator[string], map[string]int{"foo": 1, "bar": 2, "baz": 3}),
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it := test.originalMap.OrderedBegin(utils.BasicComparator[string])

			newMap := NewFromIterator[string, int](utils.BasicComparator[string], it)

			assert.ElementsMatchf(t, test.originalMap.GetKeys(), newMap.GetKeys(), test.name)
		})
	}

}

func TestTreeMapNewFromIterators(t *testing.T) {
	tests := []struct {
		name        string
		originalMap *Map[string, int]
	}{
		{
			name:        "empty list",
			originalMap: New[string, int](utils.BasicComparator[string]),
		},
		{
			name:        "single item",
			originalMap: NewFromMap[string, int](utils.BasicComparator[string], map[string]int{"foo": 1}),
		},
		{
			name:        "3 items",
			originalMap: NewFromMap[string, int](utils.BasicComparator[string], map[string]int{"foo": 1, "bar": 2, "baz": 3}),
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			first := test.originalMap.OrderedBegin(utils.BasicComparator[string])
			end := test.originalMap.OrderedEnd(utils.BasicComparator[string])

			newMap := NewFromIterators[string, int](utils.BasicComparator[string], first, end)

			assert.ElementsMatchf(t, test.originalMap.GetKeys(), newMap.GetKeys(), test.name)
		})
	}

}

// TODO: Compare lists after operations, to require correctnes
func BenchmarkHashMapRemove(b *testing.B) {
	b.StopTimer()
	variants := []struct {
		name string
		f    func(n int, name string)
	}{
		{
			name: "Ours",
			f: func(n int, name string) {
				m := New[int, string](utils.BasicComparator[int])
				for i := 0; i < n; i++ {
					m.Put(i, "foo")
				}
				b.StartTimer()
				for i := 0; i < n; i++ {
					m.Remove(utils.BasicComparator[int], i)
				}
				b.StopTimer()
			},
		},
		{
			name: "Raw",
			f: func(n int, name string) {
				m := make(map[int]string)
				for i := 0; i < n; i++ {
					m[i] = "foo"
				}
				b.StartTimer()
				for i := 0; i < n; i++ {
					delete(m, i)
				}
				b.StopTimer()
			},
		},
	}
	for _, variant := range variants {
		testCommon.RunBenchmarkWithDefualtInputSizes(b, variant.name, variant.f)
	}
}

func BenchmarkHashMapGet(b *testing.B) {
	b.StopTimer()
	variants := []struct {
		name string
		f    func(n int, name string)
	}{
		{
			name: "Ours",
			f: func(n int, name string) {
				m := New[int, string](utils.BasicComparator[int])
				for i := 0; i < n; i++ {
					m.Put(i, "foo")
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
				m := make(map[int]string)
				for i := 0; i < n; i++ {
					m[i] = "foo"
				}
				b.StartTimer()
				for i := 0; i < n; i++ {
					_, _ = m[i]
				}
				b.StopTimer()
			},
		},
	}

	for _, variant := range variants {
		testCommon.RunBenchmarkWithDefualtInputSizes(b, variant.name, variant.f)
	}
}

func BenchmarkHashMapPut(b *testing.B) {
	b.StopTimer()
	variants := []struct {
		name string
		f    func(n int, name string)
	}{
		{
			name: "Ours",
			f: func(n int, name string) {
				m := New[int, string](utils.BasicComparator[int])
				b.StartTimer()
				for i := 0; i < n; i++ {
					m.Put(i, "foo")
				}
				b.StopTimer()
			},
		},
		{
			name: "Raw",
			f: func(n int, name string) {
				m := make(map[int]string)
				b.StartTimer()
				for i := 0; i < n; i++ {
					m[i] = "foo"
				}
				b.StopTimer()
			},
		},
	}
	for _, variant := range variants {
		testCommon.RunBenchmarkWithDefualtInputSizes(b, variant.name, variant.f)
	}
}
