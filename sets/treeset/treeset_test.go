// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package treeset

import (
	"testing"

	testCommon "github.com/JonasMuehlmann/datastructures.go/tests"

	"github.com/JonasMuehlmann/datastructures.go/utils"
	"github.com/stretchr/testify/assert"
)

func TestTreeSetRemove(t *testing.T) {
	tests := []struct {
		name        string
		originalSet *Set[string]
		newSet      *Set[string]
		toRemove    string
	}{

		{
			name:        "empty list",
			originalSet: New[string](utils.BasicComparator[string]),
			newSet:      New[string](utils.BasicComparator[string]),
			toRemove:    "foo",
		},
		{
			name:        "single item",
			toRemove:    "foo",
			originalSet: NewFromSlice[string](utils.BasicComparator[string], []string{"foo"}),
			newSet:      New[string](utils.BasicComparator[string]),
		},
		{
			name:        "single item, target does not exist",
			toRemove:    "bar",
			originalSet: NewFromSlice[string](utils.BasicComparator[string], []string{"foo"}),
			newSet:      NewFromSlice[string](utils.BasicComparator[string], []string{"foo"}),
		},
		{
			name:        "3 items",
			toRemove:    "bar",
			originalSet: NewFromSlice[string](utils.BasicComparator[string], []string{"foo", "bar", "baz"}),
			newSet:      NewFromSlice[string](utils.BasicComparator[string], []string{"foo", "baz"}),
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			test.originalSet.Remove(utils.BasicComparator[string], test.toRemove)

			assert.Equalf(t, test.originalSet.GetValues(), test.newSet.GetValues(), test.name)
		})
	}
}

func TestTreeSetAdd(t *testing.T) {
	tests := []struct {
		name        string
		originalSet *Set[string]
		newSet      *Set[string]
		keyToAdd    string
		valueToAdd  int
	}{

		{
			name:        "empty list",
			originalSet: New[string](utils.BasicComparator[string]),
			newSet:      NewFromSlice[string](utils.BasicComparator[string], []string{"foo"}),
			keyToAdd:    "foo",
		},
		{
			name:        "single item",
			keyToAdd:    "foo",
			newSet:      NewFromSlice[string](utils.BasicComparator[string], []string{"foo"}),
			originalSet: New[string](utils.BasicComparator[string]),
		},
		{
			name:        "single item, overwrite",
			keyToAdd:    "foo",
			originalSet: NewFromSlice[string](utils.BasicComparator[string], []string{"foo"}),
			newSet:      NewFromSlice[string](utils.BasicComparator[string], []string{"foo"}),
		},
		{
			name:        "3 items",
			keyToAdd:    "bar",
			originalSet: NewFromSlice[string](utils.BasicComparator[string], []string{"foo", "baz"}),
			newSet:      NewFromSlice[string](utils.BasicComparator[string], []string{"foo", "bar", "baz"}),
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			test.originalSet.Add(test.keyToAdd)

			assert.Equalf(t, test.originalSet.GetValues(), test.newSet.GetValues(), test.name)
		})
	}
}

func TestTreeSetGetValues(t *testing.T) {
	tests := []struct {
		name        string
		originalSet *Set[string]
		values      []string
	}{

		{
			name:        "empty list",
			originalSet: New[string](utils.BasicComparator[string]),
			values:      []string{},
		},
		{
			name:        "single item",
			originalSet: NewFromSlice[string](utils.BasicComparator[string], []string{"foo"}),
			values:      []string{"foo"},
		},
		{
			name:        "3 items",
			originalSet: NewFromSlice[string](utils.BasicComparator[string], []string{"foo", "bar", "baz"}),
			values:      []string{"bar", "baz", "foo"},
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

func TestTreeSetContains(t *testing.T) {
	tests := []struct {
		name        string
		originalSet *Set[string]
		value       string
		doesContain bool
	}{

		{
			name:        "empty list",
			originalSet: New[string](utils.BasicComparator[string]),
			value:       "foo",
			doesContain: false,
		},
		{
			name:        "single item",
			originalSet: NewFromSlice[string](utils.BasicComparator[string], []string{"foo"}),
			value:       "foo",
			doesContain: true,
		},
		{
			name:        "3 items",
			originalSet: NewFromSlice[string](utils.BasicComparator[string], []string{"foo", "bar", "baz"}),
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

func TestTreeSetIsEmpty(t *testing.T) {
	tests := []struct {
		name        string
		originalSet *Set[string]
		isEmpty     bool
	}{

		{
			name:        "empty list",
			originalSet: New[string](utils.BasicComparator[string]),
			isEmpty:     true,
		},
		{
			name:        "single item",
			originalSet: NewFromSlice[string](utils.BasicComparator[string], []string{"foo"}),
			isEmpty:     false,
		},
		{
			name:        "3 items",
			originalSet: NewFromSlice[string](utils.BasicComparator[string], []string{"foo", "bar", "baz"}),
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

func TestTreeSetClear(t *testing.T) {
	tests := []struct {
		name          string
		originalSet   *Set[string]
		isEmptyBefore bool
		isEmptyAfter  bool
	}{

		{
			name:          "empty list",
			originalSet:   New[string](utils.BasicComparator[string]),
			isEmptyBefore: true,
			isEmptyAfter:  true,
		},
		{
			name:          "single item",
			originalSet:   NewFromSlice[string](utils.BasicComparator[string], []string{"foo"}),
			isEmptyBefore: false,
			isEmptyAfter:  true,
		},
		{
			name:          "3 items",
			originalSet:   NewFromSlice[string](utils.BasicComparator[string], []string{"foo", "bar", "baz"}),
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

func TestTreeSetNewFromIterator(t *testing.T) {
	tests := []struct {
		name        string
		originalSet *Set[string]
	}{

		{
			name:        "empty list",
			originalSet: New[string](utils.BasicComparator[string]),
		},
		{
			name:        "single item",
			originalSet: NewFromSlice[string](utils.BasicComparator[string], []string{"foo"}),
		},
		{
			name:        "3 items",
			originalSet: NewFromSlice[string](utils.BasicComparator[string], []string{"foo", "bar", "baz"}),
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			it := test.originalSet.OrderedBegin(utils.BasicComparator[string])

			newSet := NewFromIterator[string](utils.BasicComparator[string], it)

			assert.Equalf(t, test.originalSet.GetValues(), newSet.GetValues(), test.name)
		})
	}

}

func TestTreeSetNewFromIterators(t *testing.T) {
	tests := []struct {
		name        string
		originalSet *Set[string]
	}{
		{
			name:        "empty list",
			originalSet: New[string](utils.BasicComparator[string]),
		},
		{
			name:        "single item",
			originalSet: NewFromSlice[string](utils.BasicComparator[string], []string{"foo"}),
		},
		{
			name:        "3 items",
			originalSet: NewFromSlice[string](utils.BasicComparator[string], []string{"foo", "bar", "baz"}),
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			first := test.originalSet.OrderedBegin(utils.BasicComparator[string])
			end := test.originalSet.OrderedEnd(utils.BasicComparator[string])

			newSet := NewFromIterators[string](utils.BasicComparator[string], first, end)

			assert.Equalf(t, test.originalSet.GetValues(), newSet.GetValues(), test.name)
		})
	}
}

func TestTreeSetMakeIntersectionWith(t *testing.T) {
	tests := []struct {
		name         string
		a            *Set[string]
		b            *Set[string]
		intersection *Set[string]
	}{
		{
			name:         "first empty",
			a:            New[string](utils.BasicComparator[string]),
			b:            New[string](utils.BasicComparator[string], "foo", "bar", "baz"),
			intersection: New[string](utils.BasicComparator[string]),
		},
		{
			name:         "Second empty",
			a:            New[string](utils.BasicComparator[string], "foo", "bar", "baz"),
			b:            New[string](utils.BasicComparator[string]),
			intersection: New[string](utils.BasicComparator[string]),
		},
		{
			name:         "equal",
			a:            New[string](utils.BasicComparator[string], "foo", "bar", "baz"),
			b:            New[string](utils.BasicComparator[string], "foo", "bar", "baz"),
			intersection: New[string](utils.BasicComparator[string], "foo", "bar", "baz"),
		},
		{
			name:         "first shorter",
			a:            New[string](utils.BasicComparator[string], "bar", "baz"),
			b:            New[string](utils.BasicComparator[string], "foo", "bar", "baz"),
			intersection: New[string](utils.BasicComparator[string], "bar", "baz"),
		},
		{
			name:         "second shorter",
			a:            New[string](utils.BasicComparator[string], "foo", "bar", "baz"),
			b:            New[string](utils.BasicComparator[string], "bar", "baz"),
			intersection: New[string](utils.BasicComparator[string], "bar", "baz"),
		},
		{
			name:         "No overlap",
			a:            New[string](utils.BasicComparator[string], "foo", "bar", "baz"),
			b:            New[string](utils.BasicComparator[string], "1", "2"),
			intersection: New[string](utils.BasicComparator[string]),
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

func TestTreeSetMakeUnionWith(t *testing.T) {
	tests := []struct {
		name         string
		a            *Set[string]
		b            *Set[string]
		intersection *Set[string]
	}{
		{
			name:         "first empty",
			a:            New[string](utils.BasicComparator[string]),
			b:            New[string](utils.BasicComparator[string], "foo", "bar", "baz"),
			intersection: New[string](utils.BasicComparator[string], "foo", "bar", "baz"),
		},
		{
			name:         "Second empty",
			a:            New[string](utils.BasicComparator[string], "foo", "bar", "baz"),
			b:            New[string](utils.BasicComparator[string]),
			intersection: New[string](utils.BasicComparator[string], "foo", "bar", "baz"),
		},
		{
			name:         "equal",
			a:            New[string](utils.BasicComparator[string], "foo", "bar", "baz"),
			b:            New[string](utils.BasicComparator[string], "foo", "bar", "baz"),
			intersection: New[string](utils.BasicComparator[string], "foo", "bar", "baz"),
		},
		{
			name:         "first shorter",
			a:            New[string](utils.BasicComparator[string], "bar", "baz"),
			b:            New[string](utils.BasicComparator[string], "foo", "bar", "baz"),
			intersection: New[string](utils.BasicComparator[string], "foo", "bar", "baz"),
		},
		{
			name:         "second shorter",
			a:            New[string](utils.BasicComparator[string], "foo", "bar", "baz"),
			b:            New[string](utils.BasicComparator[string], "bar", "baz"),
			intersection: New[string](utils.BasicComparator[string], "foo", "bar", "baz"),
		},
		{
			name:         "No overlap",
			a:            New[string](utils.BasicComparator[string], "foo", "bar", "baz"),
			b:            New[string](utils.BasicComparator[string], "1", "2"),
			intersection: New[string](utils.BasicComparator[string], "foo", "bar", "baz", "1", "2"),
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

func TestTreeSetMakeDifferenceWith(t *testing.T) {
	tests := []struct {
		name         string
		a            *Set[string]
		b            *Set[string]
		intersection *Set[string]
	}{
		{
			name:         "first empty",
			a:            New[string](utils.BasicComparator[string]),
			b:            New[string](utils.BasicComparator[string], "foo", "bar", "baz"),
			intersection: New[string](utils.BasicComparator[string]),
		},
		{
			name:         "Second empty",
			a:            New[string](utils.BasicComparator[string], "foo", "bar", "baz"),
			b:            New[string](utils.BasicComparator[string]),
			intersection: New[string](utils.BasicComparator[string], "foo", "bar", "baz"),
		},
		{
			name:         "equal",
			a:            New[string](utils.BasicComparator[string], "foo", "bar", "baz"),
			b:            New[string](utils.BasicComparator[string], "foo", "bar", "baz"),
			intersection: New[string](utils.BasicComparator[string]),
		},
		{
			name:         "first shorter",
			a:            New[string](utils.BasicComparator[string], "bar", "baz"),
			b:            New[string](utils.BasicComparator[string], "foo", "bar", "baz"),
			intersection: New[string](utils.BasicComparator[string]),
		},
		{
			name:         "second shorter",
			a:            New[string](utils.BasicComparator[string], "foo", "bar", "baz"),
			b:            New[string](utils.BasicComparator[string], "bar", "baz"),
			intersection: New[string](utils.BasicComparator[string], "foo"),
		},
		{
			name:         "No overlap",
			a:            New[string](utils.BasicComparator[string], "foo", "bar", "baz"),
			b:            New[string](utils.BasicComparator[string], "1", "2"),
			intersection: New[string](utils.BasicComparator[string], "foo", "bar", "baz"),
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
