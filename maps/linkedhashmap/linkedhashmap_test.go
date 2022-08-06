// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linkedhashmap

import (
	"testing"

	"github.com/JonasMuehlmann/datastructures.go/tests"
	"github.com/JonasMuehlmann/datastructures.go/utils"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/maps"
)

func TestRemove(t *testing.T) {
	tests := []struct {
		name        string
		originalMap *Map[string, int]
		newMap      *Map[string, int]
		toRemove    string
	}{

		{
			name:        "empty list",
			originalMap: New[string, int](),
			newMap:      New[string, int](),
			toRemove:    "foo",
		},
		{
			name:        "single item",
			toRemove:    "foo",
			originalMap: NewFromMap[string, int](map[string]int{"foo": 1}),
			newMap:      New[string, int](),
		},
		{
			name:        "single item, target does not exist",
			toRemove:    "bar",
			originalMap: NewFromMap[string, int](map[string]int{"foo": 1}),
			newMap:      NewFromMap[string, int](map[string]int{"foo": 1}),
		},
		{
			name:        "3 items",
			toRemove:    "bar",
			originalMap: NewFromMap[string, int](map[string]int{"foo": 1, "bar": 2, "baz": 3}),
			newMap:      NewFromMap[string, int](map[string]int{"foo": 1, "baz": 3}),
		},
	}

	for _, test := range tests {
test := test

		t.Run(test.name, func(t *testing.T) {

t.Parallel()
			test.originalMap.Remove(utils.BasicComparator[string], test.toRemove)

			assert.Equalf(t, test.originalMap.table, test.newMap.table, test.name)
		})
	}
}

func TestPut(t *testing.T) {
	tests := []struct {
		name        string
		originalMap *Map[string, int]
		newMap      *Map[string, int]
		keyToAdd    string
		valueToAdd  int
	}{

		{
			name:        "empty list",
			originalMap: New[string, int](),
			newMap:      NewFromMap[string, int](map[string]int{"foo": 1}),
			keyToAdd:    "foo",
			valueToAdd:  1,
		},
		{
			name:        "single item",
			keyToAdd:    "foo",
			valueToAdd:  1,
			newMap:      NewFromMap[string, int](map[string]int{"foo": 1}),
			originalMap: New[string, int](),
		},
		{
			name:        "single item, overwrite",
			keyToAdd:    "foo",
			valueToAdd:  2,
			originalMap: NewFromMap[string, int](map[string]int{"foo": 1}),
			newMap:      NewFromMap[string, int](map[string]int{"foo": 2}),
		},
		{
			name:        "3 items",
			keyToAdd:    "bar",
			valueToAdd:  2,
			originalMap: NewFromMap[string, int](map[string]int{"foo": 1, "baz": 3}),
			newMap:      NewFromMap[string, int](map[string]int{"foo": 1, "bar": 2, "baz": 3}),
		},
	}

	for _, test := range tests {
test := test

		t.Run(test.name, func(t *testing.T) {

t.Parallel()
			test.originalMap.Put(test.keyToAdd, test.valueToAdd)

			assert.Equalf(t, test.originalMap.table, test.newMap.table, test.name)
		})
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		name        string
		originalMap *Map[string, int]
		keyToGet    string
		value       int
		found       bool
	}{

		{
			name:        "empty list",
			originalMap: New[string, int](),
			keyToGet:    "foo",
			found:       false,
		},
		{
			name:        "single item",
			keyToGet:    "foo",
			originalMap: NewFromMap[string, int](map[string]int{"foo": 1}),
			value:       1,
			found:       true,
		},
		{
			name:        "single item, target does not exist",
			keyToGet:    "bar",
			originalMap: NewFromMap[string, int](map[string]int{"foo": 1}),
			found:       false,
		},
		{
			name:        "3 items",
			keyToGet:    "bar",
			originalMap: NewFromMap[string, int](map[string]int{"foo": 1, "bar": 2, "baz": 3}),
			value:       2,
			found:       true,
		},
	}

	for _, test := range tests {
test := test

		t.Run(test.name, func(t *testing.T) {

t.Parallel()
			value, found := test.originalMap.Get(test.keyToGet)

			assert.Equalf(t, test.value, value, test.name)
			assert.Equalf(t, test.found, found, test.name)
		})
	}
}

func TestGetKeys(t *testing.T) {
	tests := []struct {
		name        string
		originalMap *Map[string, int]
		keys        []string
	}{

		{
			name:        "empty list",
			originalMap: New[string, int](),
			keys:        []string{},
		},
		{
			name:        "single item",
			originalMap: NewFromMap[string, int](map[string]int{"foo": 1}),
			keys:        []string{"foo"},
		},
		{
			name:        "3 items",
			originalMap: NewFromMap[string, int](map[string]int{"foo": 1, "bar": 2, "baz": 3}),
			keys:        []string{"foo", "bar", "baz"},
		},
	}

	for _, test := range tests {
test := test

		t.Run(test.name, func(t *testing.T) {

t.Parallel()
			keys := test.originalMap.GetKeys()

			assert.ElementsMatch(t, test.keys, keys, test.name)
		})
	}
}

func TestGetValues(t *testing.T) {
	tests := []struct {
		name        string
		originalMap *Map[string, int]
		values      []int
	}{

		{
			name:        "empty list",
			originalMap: New[string, int](),
			values:      []int{},
		},
		{
			name:        "single item",
			originalMap: NewFromMap[string, int](map[string]int{"foo": 1}),
			values:      []int{1},
		},
		{
			name:        "3 items",
			originalMap: NewFromMap[string, int](map[string]int{"foo": 1, "bar": 2, "baz": 3}),
			values:      []int{1, 2, 3},
		},
	}

	for _, test := range tests {
test := test

		t.Run(test.name, func(t *testing.T) {

t.Parallel()
			values := test.originalMap.GetValues()

			assert.ElementsMatch(t, test.values, values, test.name)
		})
	}
}

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		name        string
		originalMap *Map[string, int]
		isEmpty     bool
	}{

		{
			name:        "empty list",
			originalMap: New[string, int](),
			isEmpty:     true,
		},
		{
			name:        "single item",
			originalMap: NewFromMap[string, int](map[string]int{"foo": 1}),
			isEmpty:     false,
		},
		{
			name:        "3 items",
			originalMap: NewFromMap[string, int](map[string]int{"foo": 1, "bar": 2, "baz": 3}),
			isEmpty:     false,
		},
	}

	for _, test := range tests {
test := test

		t.Run(test.name, func(t *testing.T) {

t.Parallel()
			isEmpty := test.originalMap.IsEmpty()

			assert.Equal(t, test.isEmpty, isEmpty, test.name)
		})
	}
}

func TestClear(t *testing.T) {
	tests := []struct {
		name          string
		originalMap   *Map[string, int]
		isEmptyBefore bool
		isEmptyAfter  bool
	}{

		{
			name:          "empty list",
			originalMap:   New[string, int](),
			isEmptyBefore: true,
			isEmptyAfter:  true,
		},
		{
			name:          "single item",
			originalMap:   NewFromMap[string, int](map[string]int{"foo": 1}),
			isEmptyBefore: false,
			isEmptyAfter:  true,
		},
		{
			name:          "3 items",
			originalMap:   NewFromMap[string, int](map[string]int{"foo": 1, "bar": 2, "baz": 3}),
			isEmptyBefore: false,
			isEmptyAfter:  true,
		},
	}

	for _, test := range tests {
test := test

		t.Run(test.name, func(t *testing.T) {

t.Parallel()
			isEmptyBefore := test.originalMap.IsEmpty()
			assert.Equal(t, test.isEmptyBefore, isEmptyBefore, test.name)

			test.originalMap.Clear()

			isEmptAfter := test.originalMap.IsEmpty()
			assert.Equal(t, test.isEmptyAfter, isEmptAfter, test.name)
		})
	}
}

func TestNewFromIterator(t *testing.T) {
	tests := []struct {
		name        string
		originalMap *Map[string, int]
	}{

		{
			name:        "empty list",
			originalMap: New[string, int](),
		},
		{
			name:        "single item",
			originalMap: NewFromMap[string, int](map[string]int{"foo": 1}),
		},
		{
			name:        "3 items",
			originalMap: NewFromMap[string, int](map[string]int{"foo": 1, "bar": 2, "baz": 3}),
		},
	}

	for _, test := range tests {
test := test

		t.Run(test.name, func(t *testing.T) {

t.Parallel()
			it := test.originalMap.Begin()

			newMap := NewFromIterator[string, int](it)

			assert.EqualValues(t, test.originalMap.table, newMap.table, test.name)
		})
	}

}

func TestNewFromIterators(t *testing.T) {
	tests := []struct {
		name        string
		originalMap *Map[string, int]
	}{
		{
			name:        "empty list",
			originalMap: New[string, int](),
		},
		{
			name:        "single item",
			originalMap: NewFromMap[string, int](map[string]int{"foo": 1}),
		},
		{
			name:        "3 items",
			originalMap: NewFromMap[string, int](map[string]int{"foo": 1, "bar": 2, "baz": 3}),
		},
	}

	for _, test := range tests {
test := test

		t.Run(test.name, func(t *testing.T) {

t.Parallel()
			first := test.originalMap.Begin()
			end := test.originalMap.End()

			newMap := NewFromIterators[string, int](first, end)

			assert.EqualValues(t, test.originalMap.table, newMap.table, test.name)
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
				m := New[int, string]()
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
		tests.RunBenchmarkWithDefualtInputSizes(b, variant.name, variant.f)
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
				m := New[int, string]()
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
		tests.RunBenchmarkWithDefualtInputSizes(b, variant.name, variant.f)
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
				m := New[int, string]()
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
		tests.RunBenchmarkWithDefualtInputSizes(b, variant.name, variant.f)
	}
}

func BenchmarkHashMapGetKeys(b *testing.B) {
	b.StopTimer()
	variants := []struct {
		name string
		f    func(n int, name string)
	}{
		{
			name: "Ours",
			f: func(n int, name string) {
				m := New[int, string]()
				for i := 0; i < n; i++ {
					m.Put(i, "foo")
				}
				b.StartTimer()
				_ = m.GetKeys()
				b.StopTimer()
			},
		},
		{
			name: "golang.org_x_exp",
			f: func(n int, name string) {
				m := New[int, string]()
				for i := 0; i < n; i++ {
					m.Put(i, "foo")
				}
				b.StartTimer()
				_ = maps.Keys(m.table)
				b.StopTimer()
			},
		},
	}

	for _, variant := range variants {
		tests.RunBenchmarkWithDefualtInputSizes(b, variant.name, variant.f)
	}
}

func BenchmarkHashMapGetValues(b *testing.B) {
	b.StopTimer()
	variants := []struct {
		name string
		f    func(n int, name string)
	}{
		{
			name: "Ours",
			f: func(n int, name string) {
				m := New[int, string]()
				for i := 0; i < n; i++ {
					m.Put(i, "foo")
				}
				b.StartTimer()
				_ = m.GetValues()
				b.StopTimer()
			},
		},
		{
			name: "golang.org_x_exp",
			f: func(n int, name string) {
				m := New[int, string]()
				for i := 0; i < n; i++ {
					m.Put(i, "foo")
				}
				b.StartTimer()
				_ = maps.Values(m.table)
				b.StopTimer()
			},
		},
	}

	for _, variant := range variants {
		tests.RunBenchmarkWithDefualtInputSizes(b, variant.name, variant.f)
	}
}
