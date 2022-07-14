// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package arraylist

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/JonasMuehlmann/datastructures.go/tests"
	"github.com/JonasMuehlmann/datastructures.go/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TODO: Refactor tests with testify and table tests
func TestListNew(t *testing.T) {
	list1 := New[int]()

	if actualValue := list1.IsEmpty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	list2 := New[string]("a", "b")

	if actualValue := list2.Size(); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}

	if actualValue, ok := list2.Get(0); actualValue != "a" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "a")
	}

	if actualValue, ok := list2.Get(1); actualValue != "b" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "b")
	}

	if _, ok := list2.Get(2); ok {
		t.Errorf("Got %v expected %v", ok, false)
	}
}

func TestListPushBack(t *testing.T) {
	list := New[string]()
	list.PushBack("a")
	list.PushBack("b", "c")
	if actualValue := list.IsEmpty(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	if actualValue := list.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, ok := list.Get(2); actualValue != "c" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "c")
	}
}

func TestListPushFront(t *testing.T) {
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
		test.originalList.PushFront(test.itemsToAdd...)

		assert.ElementsMatchf(t, test.originalList.elements, test.newItems, test.name)
	}
}

func TestListPopBack(t *testing.T) {
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
		test.originalList.PopBack(test.n)

		assert.ElementsMatchf(t, test.originalList.elements, test.newItems, test.name)
	}
}

func TestListPopFront(t *testing.T) {
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
		test.originalList.PopFront(test.n)

		assert.ElementsMatchf(t, test.originalList.elements, test.newItems, test.name)
	}
}

func TestListShrinkToFit(t *testing.T) {
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
		test.originalList.PopBack(test.n)
		test.originalList.ShrinkToFit()

		assert.Equal(t, test.newLen, len(test.originalList.elements), test.name)
		assert.Equal(t, test.newCap, cap(test.originalList.elements), test.name)
	}
}

func TestListIndexOf(t *testing.T) {
	list := New[string]()

	expectedIndex := -1
	if index := list.IndexOf(utils.BasicComparator[string], "a"); index != expectedIndex {
		t.Errorf("Got %v expected %v", index, expectedIndex)
	}

	list.PushBack("a")
	list.PushBack("b", "c")

	expectedIndex = 0
	if index := list.IndexOf(utils.BasicComparator[string], "a"); index != expectedIndex {
		t.Errorf("Got %v expected %v", index, expectedIndex)
	}

	expectedIndex = 1
	if index := list.IndexOf(utils.BasicComparator[string], "b"); index != expectedIndex {
		t.Errorf("Got %v expected %v", index, expectedIndex)
	}

	expectedIndex = 2
	if index := list.IndexOf(utils.BasicComparator[string], "c"); index != expectedIndex {
		t.Errorf("Got %v expected %v", index, expectedIndex)
	}
}

func TestListRemoveStable(t *testing.T) {
	list := New[string]()
	list.PushBack("a")
	list.PushBack("b", "c")
	list.RemoveStable(2)
	if actualValue, ok := list.Get(2); ok {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
	list.RemoveStable(1)
	list.RemoveStable(0)
	list.RemoveStable(0) // no effect
	if actualValue := list.IsEmpty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := list.Size(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
}

func TestListGet(t *testing.T) {
	list := New[string]()
	list.PushBack("a")
	list.PushBack("b", "c")
	if actualValue, ok := list.Get(0); actualValue != "a" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "a")
	}
	if actualValue, ok := list.Get(1); actualValue != "b" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "b")
	}
	if actualValue, ok := list.Get(2); actualValue != "c" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "c")
	}
	if actualValue, ok := list.Get(3); ok {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
	list.RemoveStable(0)
	if actualValue, ok := list.Get(0); actualValue != "b" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "b")
	}
}

func TestListSwap(t *testing.T) {
	list := New[string]()
	list.PushBack("a")
	list.PushBack("b", "c")
	list.Swap(0, 1)
	if actualValue, ok := list.Get(0); actualValue != "b" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "b")
	}
}

func TestListSort(t *testing.T) {
	list := New[string]()
	list.Sort(utils.BasicComparator[string])
	list.PushBack("e", "f", "g", "a", "b", "c", "d")
	list.Sort(utils.BasicComparator[string])
	for i := 1; i < list.Size(); i++ {
		a, _ := list.Get(i - 1)
		b, _ := list.Get(i)
		if a > b {
			t.Errorf("Not sorted! %s > %s", a, b)
		}
	}
}

func TestListClear(t *testing.T) {
	list := New[string]()
	list.PushBack("e", "f", "g", "a", "b", "c", "d")
	list.Clear()
	if actualValue := list.IsEmpty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := list.Size(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
}

func TestListContains(t *testing.T) {
	list := New[string]()
	list.PushBack("a")
	list.PushBack("b", "c")
	if actualValue := list.Contains(utils.BasicComparator[string], "a"); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := list.Contains(utils.BasicComparator[string], "a", "b", "c"); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := list.Contains(utils.BasicComparator[string], "a", "b", "c", "d"); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	list.Clear()
	if actualValue := list.Contains(utils.BasicComparator[string], "a"); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	if actualValue := list.Contains(utils.BasicComparator[string], "a", "b", "c"); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
}

func TestListValues(t *testing.T) {
	list := New[string]()
	list.PushBack("a")
	list.PushBack("b", "c")
	actualValue, expectedValue := list.GetValues(), []string{"a", "b", "c"}
	assert.Equal(t, actualValue, expectedValue)
}

func TestListInsert(t *testing.T) {
	list := New[string]()
	list.Insert(0, "b", "c")
	list.Insert(0, "a")
	list.Insert(10, "x") // ignore
	if actualValue := list.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	list.Insert(3, "d") // append
	if actualValue := list.Size(); actualValue != 4 {
		t.Errorf("Got %v expected %v", actualValue, 4)
	}
	actualValue, expectedValue := list.GetValues(), []string{"a", "b", "c", "d"}
	assert.Equal(t, actualValue, expectedValue)
}

func TestListSet(t *testing.T) {
	list := New[string]()
	list.Set(0, "a")
	list.Set(1, "b")
	if actualValue := list.Size(); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	list.Set(2, "c") // append
	if actualValue := list.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	list.Set(4, "d")  // ignore
	list.Set(1, "bb") // update
	if actualValue := list.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	actualValue, expectedValue := list.GetValues(), []string{"a", "bb", "c"}
	assert.Equal(t, actualValue, expectedValue)
}

func TestListSerialization(t *testing.T) {
	list := New[string]()
	list.PushBack("a", "b", "c")

	var err error
	assert := func() {
		actualValue, expectedValue := list.GetValues(), []string{"a", "b", "c"}
		assert.Equal(t, actualValue, expectedValue)
		if actualValue, expectedValue := list.Size(), 3; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
		if err != nil {
			t.Errorf("Got error %v", err)
		}
	}

	assert()

	bytes, err := list.ToJSON()
	assert()

	err = list.FromJSON(bytes)
	assert()

	bytes, err = json.Marshal([]interface{}{"a", "b", "c", list})
	if err != nil {
		t.Errorf("Got error %v", err)
	}

	err = json.Unmarshal([]byte(`["a", "b", "c"]`), &list)
	if err != nil {
		t.Errorf("Got error %v", err)
	}
}

func TestListString(t *testing.T) {
	c := New[int]()
	c.PushBack(1)
	if !strings.HasPrefix(c.ToString(), "ArrayList") {
		t.Errorf("ToString should start with container name")
	}
}

func BenchmarkArrayListGet(b *testing.B) {
	b.StopTimer()
	variants := []struct {
		name string
		f    func(n int)
	}{
		{
			name: "Ours",
			f: func(n int) {
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
			f: func(n int) {
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
		tests.RunBenchmarkWithDefualtInputSizes(b, variant.name, variant.f)
	}
}

func BenchmarkArrayListPushBack(b *testing.B) {
	b.StopTimer()
	variants := []struct {
		name string
		f    func(n int)
	}{
		{
			name: "Ours",
			f: func(n int) {
				m := New[string]()
				b.StartTimer()
				for i := 0; i < n; i++ {
					m.PushBack("foo")
				}
				b.StopTimer()
			},
		},
		{
			name: "Raw",
			f: func(n int) {
				m := make([]string, 0)
				b.StartTimer()
				for i := 0; i < n; i++ {
					m = append(m, "foo")
				}
				b.StopTimer()
			},
		},
	}

	for _, variant := range variants {
		tests.RunBenchmarkWithDefualtInputSizes(b, variant.name, variant.f)
	}
}

func BenchmarkArrayListRemoveStable(b *testing.B) {
	b.StopTimer()
	variants := []struct {
		name string
		f    func(n int)
	}{
		{
			name: "Ours",
			f: func(n int) {
				m := New[string]()
				for i := 0; i < n; i++ {
					m.PushBack("foo")
				}
				b.StartTimer()
				for i := 0; i < n; i++ {
					m.RemoveStable(i)
				}
				b.StopTimer()
			},
		},
		{
			name: "Raw",
			f: func(n int) {
				m := make([]string, 0)
				for i := 0; i < n; i++ {
					m = append(m, "foo")
				}
				b.StartTimer()
				for i := 0; i < n; i++ {
					if i > 0 && i < len(m) {
						copy(m[i:], m[i+1:])
					}
				}
				b.StopTimer()
			},
		},
	}

	for _, variant := range variants {
		tests.RunBenchmarkWithDefualtInputSizes(b, variant.name, variant.f)
	}
}
