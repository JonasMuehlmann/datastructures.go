// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linkedhashset

import (
	"testing"

	testCommon "github.com/JonasMuehlmann/datastructures.go/tests"

	"github.com/JonasMuehlmann/datastructures.go/utils"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/maps"
)

func TestLinkedHashSetRemove(t *testing.T) {
	tests := []struct {
		name        string
		originalSet *Set[string]
		newSet      *Set[string]
		toRemove    string
	}{

		{
			name:        "empty list",
			originalSet: New[string](),
			newSet:      New[string](),
			toRemove:    "foo",
		},
		{
			name:        "single item",
			toRemove:    "foo",
			originalSet: NewFromSlice[string]([]string{"foo"}),
			newSet:      New[string](),
		},
		{
			name:        "single item, target does not exist",
			toRemove:    "bar",
			originalSet: NewFromSlice[string]([]string{"foo"}),
			newSet:      NewFromSlice[string]([]string{"foo"}),
		},
		{
			name:        "3 items",
			toRemove:    "bar",
			originalSet: NewFromSlice[string]([]string{"foo", "bar", "baz"}),
			newSet:      NewFromSlice[string]([]string{"foo", "baz"}),
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			test.originalSet.Remove(utils.BasicComparator[string], test.toRemove)

			assert.Equalf(t, test.originalSet, test.newSet, test.name)
		})
	}
}

func TestLinkedHashSetAdd(t *testing.T) {
	tests := []struct {
		name        string
		originalSet *Set[string]
		newSet      *Set[string]
		keyToAdd    string
		valueToAdd  int
	}{

		{
			name:        "empty list",
			originalSet: New[string](),
			newSet:      NewFromSlice[string]([]string{"foo"}),
			keyToAdd:    "foo",
		},
		{
			name:        "single item",
			keyToAdd:    "foo",
			newSet:      NewFromSlice[string]([]string{"foo"}),
			originalSet: New[string](),
		},
		{
			name:        "single item, overwrite",
			keyToAdd:    "foo",
			originalSet: NewFromSlice[string]([]string{"foo"}),
			newSet:      NewFromSlice[string]([]string{"foo"}),
		},
		{
			name:        "3 items",
			keyToAdd:    "bar",
			originalSet: NewFromSlice[string]([]string{"foo", "baz"}),
			newSet:      NewFromSlice[string]([]string{"foo", "baz", "bar"}),
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			test.originalSet.Add(test.keyToAdd)

			assert.Equalf(t, test.originalSet, test.newSet, test.name)
		})
	}
}

func TestLinkedHashSetGetValues(t *testing.T) {
	tests := []struct {
		name        string
		originalSet *Set[string]
		values      []string
	}{

		{
			name:        "empty list",
			originalSet: New[string](),
			values:      []string{},
		},
		{
			name:        "single item",
			originalSet: NewFromSlice[string]([]string{"foo"}),
			values:      []string{"foo"},
		},
		{
			name:        "3 items",
			originalSet: NewFromSlice[string]([]string{"foo", "bar", "baz"}),
			values:      []string{"foo", "bar", "baz"},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			values := test.originalSet.GetValues()

			assert.Equalf(t, test.values, values, test.name)
		})
	}
}

func TestLinkedHashSetContains(t *testing.T) {
	tests := []struct {
		name        string
		originalSet *Set[string]
		value       string
		doesContain bool
	}{

		{
			name:        "empty list",
			originalSet: New[string](),
			value:       "foo",
			doesContain: false,
		},
		{
			name:        "single item",
			originalSet: NewFromSlice[string]([]string{"foo"}),
			value:       "foo",
			doesContain: true,
		},
		{
			name:        "3 items",
			originalSet: NewFromSlice[string]([]string{"foo", "bar", "baz"}),
			value:       "foo",
			doesContain: true,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			assert.Equalf(t, test.doesContain, test.originalSet.Contains(test.value), test.name)
		})
	}
}

func TestLinkedHashSetIsEmpty(t *testing.T) {
	tests := []struct {
		name        string
		originalSet *Set[string]
		isEmpty     bool
	}{

		{
			name:        "empty list",
			originalSet: New[string](),
			isEmpty:     true,
		},
		{
			name:        "single item",
			originalSet: NewFromSlice[string]([]string{"foo"}),
			isEmpty:     false,
		},
		{
			name:        "3 items",
			originalSet: NewFromSlice[string]([]string{"foo", "bar", "baz"}),
			isEmpty:     false,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			isEmpty := test.originalSet.IsEmpty()

			assert.Equal(t, test.isEmpty, isEmpty, test.name)
		})
	}
}

func TestLinkedHashSetClear(t *testing.T) {
	tests := []struct {
		name          string
		originalSet   *Set[string]
		isEmptyBefore bool
		isEmptyAfter  bool
	}{

		{
			name:          "empty list",
			originalSet:   New[string](),
			isEmptyBefore: true,
			isEmptyAfter:  true,
		},
		{
			name:          "single item",
			originalSet:   NewFromSlice[string]([]string{"foo"}),
			isEmptyBefore: false,
			isEmptyAfter:  true,
		},
		{
			name:          "3 items",
			originalSet:   NewFromSlice[string]([]string{"foo", "bar", "baz"}),
			isEmptyBefore: false,
			isEmptyAfter:  true,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			isEmptyBefore := test.originalSet.IsEmpty()
			assert.Equal(t, test.isEmptyBefore, isEmptyBefore, test.name)

			test.originalSet.Clear()

			isEmptAfter := test.originalSet.IsEmpty()
			assert.Equal(t, test.isEmptyAfter, isEmptAfter, test.name)
		})
	}
}

func TestLinkedHashSetNewFromIterator(t *testing.T) {
	tests := []struct {
		name        string
		originalSet *Set[string]
	}{

		{
			name:        "empty list",
			originalSet: New[string](),
		},
		{
			name:        "single item",
			originalSet: NewFromSlice[string]([]string{"foo"}),
		},
		{
			name:        "3 items",
			originalSet: NewFromSlice[string]([]string{"foo", "bar", "baz"}),
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it := test.originalSet.Begin(utils.BasicComparator[string])

			newSet := NewFromIterator[string](it)

			assert.Equalf(t, test.originalSet.GetValues(), newSet.GetValues(), test.name)
		})
	}

}

func TestLinkedHashSetNewFromIterators(t *testing.T) {
	tests := []struct {
		name        string
		originalSet *Set[string]
	}{
		{
			name:        "empty list",
			originalSet: New[string](),
		},
		{
			name:        "single item",
			originalSet: NewFromSlice[string]([]string{"foo"}),
		},
		{
			name:        "3 items",
			originalSet: NewFromSlice[string]([]string{"foo", "bar", "baz"}),
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			first := test.originalSet.Begin(utils.BasicComparator[string])
			end := test.originalSet.End(utils.BasicComparator[string])

			newSet := NewFromIterators[string](first, end)

			assert.Equalf(t, test.originalSet.GetValues(), newSet.GetValues(), test.name)
		})
	}
}

func TestLinkedHashSetMakeIntersectionWith(t *testing.T) {
	tests := []struct {
		name         string
		a            *Set[string]
		b            *Set[string]
		intersection *Set[string]
	}{
		{
			name:         "first empty",
			a:            New[string](),
			b:            New[string]("foo", "bar", "baz"),
			intersection: New[string](),
		},
		{
			name:         "Second empty",
			a:            New[string]("foo", "bar", "baz"),
			b:            New[string](),
			intersection: New[string](),
		},
		{
			name:         "equal",
			a:            New[string]("foo", "bar", "baz"),
			b:            New[string]("foo", "bar", "baz"),
			intersection: New[string]("foo", "bar", "baz"),
		},
		{
			name:         "first shorter",
			a:            New[string]("bar", "baz"),
			b:            New[string]("foo", "bar", "baz"),
			intersection: New[string]("bar", "baz"),
		},
		{
			name:         "second shorter",
			a:            New[string]("foo", "bar", "baz"),
			b:            New[string]("bar", "baz"),
			intersection: New[string]("bar", "baz"),
		},
		{
			name:         "No overlap",
			a:            New[string]("foo", "bar", "baz"),
			b:            New[string]("1", "2"),
			intersection: New[string](),
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			newSet := test.a.MakeIntersectionWith(test.b)

			assert.ElementsMatchf(t, test.intersection.GetValues(), newSet.GetValues(), test.name)
		})
	}
}

func TestLinkedHashSetMakeUnionWith(t *testing.T) {
	tests := []struct {
		name         string
		a            *Set[string]
		b            *Set[string]
		intersection *Set[string]
	}{
		{
			name:         "first empty",
			a:            New[string](),
			b:            New[string]("foo", "bar", "baz"),
			intersection: New[string]("foo", "bar", "baz"),
		},
		{
			name:         "Second empty",
			a:            New[string]("foo", "bar", "baz"),
			b:            New[string](),
			intersection: New[string]("foo", "bar", "baz"),
		},
		{
			name:         "equal",
			a:            New[string]("foo", "bar", "baz"),
			b:            New[string]("foo", "bar", "baz"),
			intersection: New[string]("foo", "bar", "baz"),
		},
		{
			name:         "first shorter",
			a:            New[string]("bar", "baz"),
			b:            New[string]("foo", "bar", "baz"),
			intersection: New[string]("foo", "bar", "baz"),
		},
		{
			name:         "second shorter",
			a:            New[string]("foo", "bar", "baz"),
			b:            New[string]("bar", "baz"),
			intersection: New[string]("foo", "bar", "baz"),
		},
		{
			name:         "No overlap",
			a:            New[string]("foo", "bar", "baz"),
			b:            New[string]("1", "2"),
			intersection: New[string]("foo", "bar", "baz", "1", "2"),
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			newSet := test.a.MakeUnionWith(test.b)

			assert.ElementsMatchf(t, test.intersection.GetValues(), newSet.GetValues(), test.name)
		})
	}
}

func TestLinkedHashSetMakeDifferenceWith(t *testing.T) {
	tests := []struct {
		name         string
		a            *Set[string]
		b            *Set[string]
		intersection *Set[string]
	}{
		{
			name:         "first empty",
			a:            New[string](),
			b:            New[string]("foo", "bar", "baz"),
			intersection: New[string](),
		},
		{
			name:         "Second empty",
			a:            New[string]("foo", "bar", "baz"),
			b:            New[string](),
			intersection: New[string]("foo", "bar", "baz"),
		},
		{
			name:         "equal",
			a:            New[string]("foo", "bar", "baz"),
			b:            New[string]("foo", "bar", "baz"),
			intersection: New[string](),
		},
		{
			name:         "first shorter",
			a:            New[string]("bar", "baz"),
			b:            New[string]("foo", "bar", "baz"),
			intersection: New[string](),
		},
		{
			name:         "second shorter",
			a:            New[string]("foo", "bar", "baz"),
			b:            New[string]("bar", "baz"),
			intersection: New[string]("foo"),
		},
		{
			name:         "No overlap",
			a:            New[string]("foo", "bar", "baz"),
			b:            New[string]("1", "2"),
			intersection: New[string]("foo", "bar", "baz"),
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			newSet := test.a.MakeDifferenceWith(test.b)

			assert.ElementsMatchf(t, test.intersection.GetValues(), newSet.GetValues(), test.name)
		})
	}
}

// TODO: Compare lists after operations, to require correctnes
func BenchmarkLinkedHashSetRemove(b *testing.B) {
	b.StopTimer()
	variants := []struct {
		name string
		f    func(n int, name string)
	}{
		{
			name: "Ours",
			f: func(n int, name string) {
				m := New[int]()
				for i := 0; i < n; i++ {
					m.Add(i)
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
				m := make(map[int]struct{})
				for i := 0; i < n; i++ {
					m[i] = struct{}{}
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

func BenchmarkLinkedHashSetAdd(b *testing.B) {
	b.StopTimer()
	variants := []struct {
		name string
		f    func(n int, name string)
	}{
		{
			name: "Ours",
			f: func(n int, name string) {
				m := New[int]()
				b.StartTimer()
				for i := 0; i < n; i++ {
					m.Add(i)
				}
				b.StopTimer()
			},
		},
		{
			name: "Raw",
			f: func(n int, name string) {
				m := make(map[int]struct{})
				b.StartTimer()
				for i := 0; i < n; i++ {
					m[i] = struct{}{}
				}
				b.StopTimer()
			},
		},
	}
	for _, variant := range variants {
		testCommon.RunBenchmarkWithDefualtInputSizes(b, variant.name, variant.f)
	}
}

func BenchmarkLinkedHashSetGetValues(b *testing.B) {
	b.StopTimer()
	variants := []struct {
		name string
		f    func(n int, name string)
	}{
		{
			name: "Ours",
			f: func(n int, name string) {
				m := New[int]()
				for i := 0; i < n; i++ {
					m.Add(i)
				}
				b.StartTimer()
				_ = m.GetValues()
				b.StopTimer()
			},
		},
		{
			name: "golang.org_x_exp",
			f: func(n int, name string) {
				m := New[int]()
				for i := 0; i < n; i++ {
					m.Add(i)
				}
				b.StartTimer()
				_ = maps.Keys(m.table)
				b.StopTimer()
			},
		},
	}

	for _, variant := range variants {
		testCommon.RunBenchmarkWithDefualtInputSizes(b, variant.name, variant.f)
	}
}
