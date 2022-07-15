// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package singlylinkedlist

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/JonasMuehlmann/datastructures.go/utils"
	"github.com/stretchr/testify/assert"
)

func TestListNew(t *testing.T) {
	list1 := New[string]()

	if actualValue := list1.IsEmpty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	list2 := New("a", "b")

	if actualValue := list2.Size(); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}

	if actualValue, ok := list2.Get(0); actualValue != "a" || !ok {
		t.Errorf("Got %v expected %v", actualValue, 1)
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

func TestListAppendAndPushFront(t *testing.T) {
	list := New[string]()
	list.PushBack("b")
	list.PushFront("a")
	list.Append("c")
	if actualValue := list.IsEmpty(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	if actualValue := list.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, ok := list.Get(0); actualValue != "a" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "c")
	}
	if actualValue, ok := list.Get(1); actualValue != "b" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "c")
	}
	if actualValue, ok := list.Get(2); actualValue != "c" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "c")
	}
}

func TestListRemove(t *testing.T) {
	list := New[string]()
	list.PushBack("a")
	list.PushBack("b", "c")
	list.Remove(2)
	if _, ok := list.Get(2); ok {
		t.Errorf("Got %v expected %v", ok, false)
	}
	list.Remove(1)
	list.Remove(0)
	list.Remove(0) // no effect
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
	if _, ok := list.Get(3); ok {
		t.Errorf("Got %v expected %v", ok, false)
	}
	list.Remove(0)
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
		t.Errorf("Got %v expected %v", actualValue, "c")
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

	assert.ElementsMatch(t, []string{"a", "b", "c"}, list.GetValues())
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
	assert.ElementsMatch(t, []string{"a", "b", "c", "d"}, list.GetValues())
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
	assert.ElementsMatch(t, []string{"a", "bb", "c"}, list.GetValues())
}

func TestListIteratorNextOnEmpty(t *testing.T) {
	list := New[string]()
	it := list.Iterator()
	for it.Next() {
		t.Errorf("Shouldn't iterate on empty list")
	}
}

func TestListIteratorNext(t *testing.T) {
	list := New[string]()
	list.PushBack("a", "b", "c")
	it := list.Iterator()
	count := 0
	for it.Next() {
		count++
		index := it.Index()
		value := it.Value()
		switch index {
		case 0:
			if actualValue, expectedValue := value, "a"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case 1:
			if actualValue, expectedValue := value, "b"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case 2:
			if actualValue, expectedValue := value, "c"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			t.Errorf("Too many")
		}
	}
	if actualValue, expectedValue := count, 3; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestListIteratorBegin(t *testing.T) {
	list := New[string]()
	it := list.Iterator()
	it.Begin()
	list.PushBack("a", "b", "c")
	for it.Next() {
	}
	it.Begin()
	it.Next()
	if index, value := it.Index(), it.Value(); index != 0 || value != "a" {
		t.Errorf("Got %v,%v expected %v,%v", index, value, 0, "a")
	}
}

func TestListIteratorFirst(t *testing.T) {
	list := New[string]()
	it := list.Iterator()
	if actualValue, expectedValue := it.First(), false; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	list.PushBack("a", "b", "c")
	if actualValue, expectedValue := it.First(), true; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if index, value := it.Index(), it.Value(); index != 0 || value != "a" {
		t.Errorf("Got %v,%v expected %v,%v", index, value, 0, "a")
	}
}

func TestListIteratorNextTo(t *testing.T) {
	// Sample seek function, i.e. string starting with "b"
	seek := func(index int, value interface{}) bool {
		return strings.HasSuffix(value.(string), "b")
	}

	// NextTo (empty)
	{
		list := New[string]()
		it := list.Iterator()
		for it.NextTo(seek) {
			t.Errorf("Shouldn't iterate on empty list")
		}
	}

	// NextTo (not found)
	{
		list := New[string]()
		list.PushBack("xx", "yy")
		it := list.Iterator()
		for it.NextTo(seek) {
			t.Errorf("Shouldn't iterate on empty list")
		}
	}

	// NextTo (found)
	{
		list := New[string]()
		list.PushBack("aa", "bb", "cc")
		it := list.Iterator()
		it.Begin()
		if !it.NextTo(seek) {
			t.Errorf("Shouldn't iterate on empty list")
		}
		if index, value := it.Index(), it.Value(); index != 1 || value.(string) != "bb" {
			t.Errorf("Got %v,%v expected %v,%v", index, value, 1, "bb")
		}
		if !it.Next() {
			t.Errorf("Should go to first element")
		}
		if index, value := it.Index(), it.Value(); index != 2 || value.(string) != "cc" {
			t.Errorf("Got %v,%v expected %v,%v", index, value, 2, "cc")
		}
		if it.Next() {
			t.Errorf("Should not go past last element")
		}
	}
}

func TestListSerialization(t *testing.T) {
	list := New[string]()
	list.PushBack("a", "b", "c")

	var err error
	assert := func() {
		assert.ElementsMatch(t, []string{"a", "b", "c"}, list.GetValues())
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

	err = json.Unmarshal([]byte(`["a", "b","c"]`), &list)
	if err != nil {
		t.Errorf("Got error %v", err)
	}
}

func TestListString(t *testing.T) {
	c := New[int]()
	c.PushBack(1)
	if !strings.HasPrefix(c.ToString(), "SinglyLinkedList") {
		t.Errorf("ToString should start with container name")
	}
}

func benchmarkGet(b *testing.B, list *List[int], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			list.Get(n)
		}
	}
}

func benchmarkPushBack(b *testing.B, list *List[int], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			list.PushBack(n)
		}
	}
}

func benchmarkRemove(b *testing.B, list *List[int], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			list.Remove(n)
		}
	}
}

func BenchmarkSinglyLinkedListGet100(b *testing.B) {
	b.StopTimer()
	size := 100
	list := New[int]()
	for n := 0; n < size; n++ {
		list.PushBack(n)
	}
	b.StartTimer()
	benchmarkGet(b, list, size)
}

func BenchmarkSinglyLinkedListGet1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	list := New[int]()
	for n := 0; n < size; n++ {
		list.PushBack(n)
	}
	b.StartTimer()
	benchmarkGet(b, list, size)
}

func BenchmarkSinglyLinkedListGet10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	list := New[int]()
	for n := 0; n < size; n++ {
		list.PushBack(n)
	}
	b.StartTimer()
	benchmarkGet(b, list, size)
}

func BenchmarkSinglyLinkedListGet100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	list := New[int]()
	for n := 0; n < size; n++ {
		list.PushBack(n)
	}
	b.StartTimer()
	benchmarkGet(b, list, size)
}

func BenchmarkSinglyLinkedListAdd100(b *testing.B) {
	b.StopTimer()
	size := 100
	list := New[int]()
	b.StartTimer()
	benchmarkPushBack(b, list, size)
}

func BenchmarkSinglyLinkedListAdd1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	list := New[int]()
	for n := 0; n < size; n++ {
		list.PushBack(n)
	}
	b.StartTimer()
	benchmarkPushBack(b, list, size)
}

func BenchmarkSinglyLinkedListAdd10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	list := New[int]()
	for n := 0; n < size; n++ {
		list.PushBack(n)
	}
	b.StartTimer()
	benchmarkPushBack(b, list, size)
}

func BenchmarkSinglyLinkedListAdd100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	list := New[int]()
	for n := 0; n < size; n++ {
		list.PushBack(n)
	}
	b.StartTimer()
	benchmarkPushBack(b, list, size)
}

func BenchmarkSinglyLinkedListRemove100(b *testing.B) {
	b.StopTimer()
	size := 100
	list := New[int]()
	for n := 0; n < size; n++ {
		list.PushBack(n)
	}
	b.StartTimer()
	benchmarkRemove(b, list, size)
}

func BenchmarkSinglyLinkedListRemove1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	list := New[int]()
	for n := 0; n < size; n++ {
		list.PushBack(n)
	}
	b.StartTimer()
	benchmarkRemove(b, list, size)
}

func BenchmarkSinglyLinkedListRemove10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	list := New[int]()
	for n := 0; n < size; n++ {
		list.PushBack(n)
	}
	b.StartTimer()
	benchmarkRemove(b, list, size)
}

func BenchmarkSinglyLinkedListRemove100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	list := New[int]()
	for n := 0; n < size; n++ {
		list.PushBack(n)
	}
	b.StartTimer()
	benchmarkRemove(b, list, size)
}
