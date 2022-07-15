// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package treemap

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/JonasMuehlmann/datastructures.go/utils"
	"github.com/stretchr/testify/assert"
)

func TestMapPut(t *testing.T) {
	m := NewWith[int, string](utils.BasicComparator[int])
	m.Put(5, "e")
	m.Put(6, "f")
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(1, "x")
	m.Put(2, "b")
	m.Put(1, "a") //overwrite

	if actualValue := m.Size(); actualValue != 7 {
		t.Errorf("Got %v expected %v", actualValue, 7)
	}
	assert.ElementsMatch(t, []int{1, 2, 3, 4, 5, 6, 7}, m.GetKeys())
	assert.ElementsMatch(t, []string{}, m.GetValues())
	assert.ElementsMatch(t, []int{1, 2, 3, 4, 5, 6, 7}, m.GetValues())

	// key,expectedValue,expectedFound
	tests1 := [][]interface{}{
		{1, "a", true},
		{2, "b", true},
		{3, "c", true},
		{4, "d", true},
		{5, "e", true},
		{6, "f", true},
		{7, "g", true},
		{8, nil, false},
	}

	for _, test := range tests1 {
		// retrievals
		actualValue, actualFound := m.Get(test.key)
		if actualValue != test.value || actualFound != test.retrieved {
			t.Errorf("Got %v expected %v", actualValue, test.value)
		}
	}
}

func TestMapMin(t *testing.T) {
	m := NewWith[int, string](utils.BasicComparator[int])

	if k, v := m.Min(); k != nil || v != nil {
		t.Errorf("Got %v->%v expected %v->%v", k, v, nil, nil)
	}

	m.Put(5, "e")
	m.Put(6, "f")
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(1, "x")
	m.Put(2, "b")
	m.Put(1, "a") //overwrite

	actualKey, actualValue := m.Min()
	expectedKey, expectedValue := 1, "a"
	if actualKey != expectedKey {
		t.Errorf("Got %v expected %v", actualKey, expectedKey)
	}
	if actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestMapMax(t *testing.T) {
	m := NewWith[int, string](utils.BasicComparator[int])

	if k, v := m.Max(); k != nil || v != nil {
		t.Errorf("Got %v->%v expected %v->%v", k, v, nil, nil)
	}

	m.Put(5, "e")
	m.Put(6, "f")
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(1, "x")
	m.Put(2, "b")
	m.Put(1, "a") //overwrite

	actualKey, actualValue := m.Max()
	expectedKey, expectedValue := 7, "g"
	if actualKey != expectedKey {
		t.Errorf("Got %v expected %v", actualKey, expectedKey)
	}
	if actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestMapClear(t *testing.T) {
	m := NewWith[int, string](utils.BasicComparator[int])
	m.Put(5, "e")
	m.Put(6, "f")
	m.Put(7, "g")
	m.Put(3, "c")
	if actualValue, expectedValue := m.Size(), 4; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	m.Clear()
	if actualValue, expectedValue := m.Size(), 0; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestMapRemove(t *testing.T) {
	m := NewWith[int, string](utils.BasicComparator[int])
	m.Put(5, "e")
	m.Put(6, "f")
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(1, "x")
	m.Put(2, "b")
	m.Put(1, "a") //overwrite

	m.Remove(5)
	m.Remove(6)
	m.Remove(7)
	m.Remove(8)
	m.Remove(5)

	if actualValue, expectedValue := m.GetKeys(), []interface{}{1, 2, 3, 4}; !sameElements(actualValue, expectedValue) {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	if actualValue, expectedValue := m.GetValues(), []interface{}{"a", "b", "c", "d"}; !sameElements(actualValue, expectedValue) {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue := m.Size(); actualValue != 4 {
		t.Errorf("Got %v expected %v", actualValue, 4)
	}

	tests2 := [][]interface{}{
		{1, "a", true},
		{2, "b", true},
		{3, "c", true},
		{4, "d", true},
		{5, nil, false},
		{6, nil, false},
		{7, nil, false},
		{8, nil, false},
	}

	for _, test := range tests2 {
		actualValue, actualFound := m.Get(test[0])
		if actualValue != test[1] || actualFound != test[2] {
			t.Errorf("Got %v expected %v", actualValue, test[1])
		}
	}

	m.Remove(1)
	m.Remove(4)
	m.Remove(2)
	m.Remove(3)
	m.Remove(2)
	m.Remove(2)

	if actualValue, expectedValue := fmt.Sprintf("%s", m.GetKeys()), "[]"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue, expectedValue := fmt.Sprintf("%s", m.GetValues()), "[]"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue := m.Size(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
	if actualValue := m.IsEmpty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
}

func TestMapFloor(t *testing.T) {
	m := NewWith[int, string](utils.BasicComparator[int])
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(1, "a")

	// key,expectedKey,expectedValue,expectedFound
	tests1 := [][]interface{}{
		{-1, nil, nil, false},
		{0, nil, nil, false},
		{1, 1, "a", true},
		{2, 1, "a", true},
		{3, 3, "c", true},
		{4, 3, "c", true},
		{7, 7, "g", true},
		{8, 7, "g", true},
	}

	for _, test := range tests1 {
		// retrievals
		actualKey, actualValue := m.Floor(test[0])
		actualFound := actualKey != nil && actualValue != nil
		if actualKey != test[1] || actualValue != test[2] || actualFound != test[3] {
			t.Errorf("Got %v, %v, %v, expected %v, %v, %v", actualKey, actualValue, actualFound, test[1], test[2], test[3])
		}
	}
}

func TestMapCeiling(t *testing.T) {
	m := NewWith[int, string](utils.BasicComparator[int])
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(1, "a")

	// key,expectedKey,expectedValue,expectedFound
	tests1 := [][]interface{}{
		{-1, 1, "a", true},
		{0, 1, "a", true},
		{1, 1, "a", true},
		{2, 3, "c", true},
		{3, 3, "c", true},
		{4, 7, "g", true},
		{7, 7, "g", true},
		{8, nil, nil, false},
	}

	for _, test := range tests1 {
		// retrievals
		actualKey, actualValue := m.Ceiling(test[0])
		actualFound := actualKey != nil && actualValue != nil
		if actualKey != test[1] || actualValue != test[2] || actualFound != test[3] {
			t.Errorf("Got %v, %v, %v, expected %v, %v, %v", actualKey, actualValue, actualFound, test[1], test[2], test[3])
		}
	}
}

func TestMapIteratorNextOnEmpty(t *testing.T) {
	m := NewWith[string, string](utils.BasicComparator[string])
	it := m.Iterator()
	it = m.Iterator()
	for it.Next() {
		t.Errorf("Shouldn't iterate on empty map")
	}
}

func TestMapIteratorPrevOnEmpty(t *testing.T) {
	m := NewWith[string, string](utils.BasicComparator[string])
	it := m.Iterator()
	it = m.Iterator()
	for it.Prev() {
		t.Errorf("Shouldn't iterate on empty map")
	}
}

func TestMapIteratorNext(t *testing.T) {
	m := NewWith[string, string](utils.BasicComparator[string])
	m.Put("c", 3)
	m.Put("a", 1)
	m.Put("b", 2)

	it := m.Iterator()
	count := 0
	for it.Next() {
		count++
		key := it.Key()
		value := it.Value()
		switch key {
		case "a":
			if actualValue, expectedValue := value, 1; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case "b":
			if actualValue, expectedValue := value, 2; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case "c":
			if actualValue, expectedValue := value, 3; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			t.Errorf("Too many")
		}
		if actualValue, expectedValue := value, count; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
	}
	if actualValue, expectedValue := count, 3; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestMapIteratorPrev(t *testing.T) {
	m := NewWith[string, string](utils.BasicComparator[string])
	m.Put("c", 3)
	m.Put("a", 1)
	m.Put("b", 2)

	it := m.Iterator()
	for it.Next() {
	}
	countDown := m.Size()
	for it.Prev() {
		key := it.Key()
		value := it.Value()
		switch key {
		case "a":
			if actualValue, expectedValue := value, 1; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case "b":
			if actualValue, expectedValue := value, 2; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case "c":
			if actualValue, expectedValue := value, 3; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			t.Errorf("Too many")
		}
		if actualValue, expectedValue := value, countDown; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
		countDown--
	}
	if actualValue, expectedValue := countDown, 0; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestMapIteratorBegin(t *testing.T) {
	m := NewWith[int, string](utils.BasicComparator[int])
	it := m.Iterator()
	it.Begin()
	m.Put(3, "c")
	m.Put(1, "a")
	m.Put(2, "b")
	for it.Next() {
	}
	it.Begin()
	it.Next()
	if key, value := it.Key(), it.Value(); key != 1 || value != "a" {
		t.Errorf("Got %v,%v expected %v,%v", key, value, 1, "a")
	}
}

func TestMapIteratorEnd(t *testing.T) {
	m := NewWith[int, string](utils.BasicComparator[int])
	it := m.Iterator()
	m.Put(3, "c")
	m.Put(1, "a")
	m.Put(2, "b")
	it.End()
	it.Prev()
	if key, value := it.Key(), it.Value(); key != 3 || value != "c" {
		t.Errorf("Got %v,%v expected %v,%v", key, value, 3, "c")
	}
}

func TestMapIteratorFirst(t *testing.T) {
	m := NewWith[int, string](utils.BasicComparator[int])
	m.Put(3, "c")
	m.Put(1, "a")
	m.Put(2, "b")
	it := m.Iterator()
	if actualValue, expectedValue := it.First(), true; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if key, value := it.Key(), it.Value(); key != 1 || value != "a" {
		t.Errorf("Got %v,%v expected %v,%v", key, value, 1, "a")
	}
}

func TestMapIteratorLast(t *testing.T) {
	m := NewWith[int, string](utils.BasicComparator[int])
	m.Put(3, "c")
	m.Put(1, "a")
	m.Put(2, "b")
	it := m.Iterator()
	if actualValue, expectedValue := it.Last(), true; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if key, value := it.Key(), it.Value(); key != 3 || value != "c" {
		t.Errorf("Got %v,%v expected %v,%v", key, value, 3, "c")
	}
}

func TestMapIteratorNextTo(t *testing.T) {
	// Sample seek function, i.e. string starting with "b"
	seek := func(index interface{}, value interface{}) bool {
		return strings.HasSuffix(value, "b")
	}

	// NextTo (empty)
	{
		m := NewWith[int, string](utils.BasicComparator[int])
		it := m.Iterator()
		for it.NextTo(seek) {
			t.Errorf("Shouldn't iterate on empty map")
		}
	}

	// NextTo (not found)
	{
		m := NewWith[int, string](utils.BasicComparator[int])
		m.Put(0, "xx")
		m.Put(1, "yy")
		it := m.Iterator()
		for it.NextTo(seek) {
			t.Errorf("Shouldn't iterate on empty map")
		}
	}

	// NextTo (found)
	{
		m := NewWith[int, string](utils.BasicComparator[int])
		m.Put(0, "aa")
		m.Put(1, "bb")
		m.Put(2, "cc")
		it := m.Iterator()
		it.Begin()
		if !it.NextTo(seek) {
			t.Errorf("Shouldn't iterate on empty map")
		}
		if index, value := it.Key(), it.Value(); index != 1 || value != "bb" {
			t.Errorf("Got %v,%v expected %v,%v", index, value, 1, "bb")
		}
		if !it.Next() {
			t.Errorf("Should go to first element")
		}
		if index, value := it.Key(), it.Value(); index != 2 || value != "cc" {
			t.Errorf("Got %v,%v expected %v,%v", index, value, 2, "cc")
		}
		if it.Next() {
			t.Errorf("Should not go past last element")
		}
	}
}

func TestMapIteratorPrevTo(t *testing.T) {
	// Sample seek function, i.e. string starting with "b"
	seek := func(index interface{}, value interface{}) bool {
		return strings.HasSuffix(value, "b")
	}

	// PrevTo (empty)
	{
		m := NewWith[int, string](utils.BasicComparator[int])
		it := m.Iterator()
		it.End()
		for it.PrevTo(seek) {
			t.Errorf("Shouldn't iterate on empty map")
		}
	}

	// PrevTo (not found)
	{
		m := NewWith[int, string](utils.BasicComparator[int])
		m.Put(0, "xx")
		m.Put(1, "yy")
		it := m.Iterator()
		it.End()
		for it.PrevTo(seek) {
			t.Errorf("Shouldn't iterate on empty map")
		}
	}

	// PrevTo (found)
	{
		m := NewWith[int, string](utils.BasicComparator[int])
		m.Put(0, "aa")
		m.Put(1, "bb")
		m.Put(2, "cc")
		it := m.Iterator()
		it.End()
		if !it.PrevTo(seek) {
			t.Errorf("Shouldn't iterate on empty map")
		}
		if index, value := it.Key(), it.Value(); index != 1 || value != "bb" {
			t.Errorf("Got %v,%v expected %v,%v", index, value, 1, "bb")
		}
		if !it.Prev() {
			t.Errorf("Should go to first element")
		}
		if index, value := it.Key(), it.Value(); index != 0 || value != "aa" {
			t.Errorf("Got %v,%v expected %v,%v", index, value, 0, "aa")
		}
		if it.Prev() {
			t.Errorf("Should not go before first element")
		}
	}
}

func TestMapSerialization(t *testing.T) {
	for i := 0; i < 10; i++ {
		original := NewWith[string, string](utils.BasicComparator[string])
		original.Put("d", "4")
		original.Put("e", "5")
		original.Put("c", "3")
		original.Put("b", "2")
		original.Put("a", "1")

		assertSerialization(original, "A", t)

		serialized, err := original.ToJSON()
		if err != nil {
			t.Errorf("Got error %v", err)
		}
		assertSerialization(original, "B", t)

		deserialized := NewWith[string, string](utils.BasicComparator[string])
		err = deserialized.FromJSON(serialized)
		if err != nil {
			t.Errorf("Got error %v", err)
		}
		assertSerialization(deserialized, "C", t)
	}

	m := NewWith[string, string](utils.BasicComparator[string])
	m.Put("a", 1.0)
	m.Put("b", 2.0)
	m.Put("c", 3.0)

	_, err := json.Marshal([]interface{}{"a", "b", "c", m})
	if err != nil {
		t.Errorf("Got error %v", err)
	}

	err = json.Unmarshal([]byte(`{"a":1,"b":2}`), &m)
	if err != nil {
		t.Errorf("Got error %v", err)
	}
}

func TestMapString(t *testing.T) {
	c := NewWith[string, string](utils.BasicComparator[string])
	c.Put("a", 1)
	if !strings.HasPrefix(c.ToString(), "TreeMap") {
		t.Errorf("ToString should start with container name")
	}
}

//noinspection GoBoolExpressions
func assertSerialization(m *Map[int, struct{}], txt string, t *testing.T) {
	if actualValue := m.GetKeys(); false ||
		actualValue[0] != "a" ||
		actualValue[1] != "b" ||
		actualValue[2] != "c" ||
		actualValue[3] != "d" ||
		actualValue[4] != "e" {
		t.Errorf("[%s] Got %v expected %v", txt, actualValue, "[a,b,c,d,e]")
	}
	if actualValue := m.GetValues(); false ||
		actualValue[0] != "1" ||
		actualValue[1] != "2" ||
		actualValue[2] != "3" ||
		actualValue[3] != "4" ||
		actualValue[4] != "5" {
		t.Errorf("[%s] Got %v expected %v", txt, actualValue, "[1,2,3,4,5]")
	}
	if actualValue, expectedValue := m.Size(), 5; actualValue != expectedValue {
		t.Errorf("[%s] Got %v expected %v", txt, actualValue, expectedValue)
	}
}

func benchmarkGet(b *testing.B, m *Map[int, struct{}], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			m.Get(n)
		}
	}
}

func benchmarkPut(b *testing.B, m *Map[int, struct{}], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			m.Put(n, struct{}{})
		}
	}
}

func benchmarkRemove(b *testing.B, m *Map[int, struct{}], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			m.Remove(n)
		}
	}
}

func BenchmarkTreeMapGet100(b *testing.B) {
	b.StopTimer()
	size := 100
	m := NewWith[int, struct{}](utils.BasicComparator[int])
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkTreeMapGet1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	m := NewWith[int, struct{}](utils.BasicComparator[int])
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkTreeMapGet10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	m := NewWith[int, struct{}](utils.BasicComparator[int])
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkTreeMapGet100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	m := NewWith[int, struct{}](utils.BasicComparator[int])
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkTreeMapPut100(b *testing.B) {
	b.StopTimer()
	size := 100
	m := NewWith[int, struct{}](utils.BasicComparator[int])
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkTreeMapPut1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	m := NewWith[int, struct{}](utils.BasicComparator[int])
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkTreeMapPut10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	m := NewWith[int, struct{}](utils.BasicComparator[int])
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkTreeMapPut100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	m := NewWith[int, struct{}](utils.BasicComparator[int])
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkTreeMapRemove100(b *testing.B) {
	b.StopTimer()
	size := 100
	m := NewWith[int, struct{}](utils.BasicComparator[int])
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, m, size)
}

func BenchmarkTreeMapRemove1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	m := NewWith[int, struct{}](utils.BasicComparator[int])
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, m, size)
}

func BenchmarkTreeMapRemove10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	m := NewWith[int, struct{}](utils.BasicComparator[int])
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, m, size)
}

func BenchmarkTreeMapRemove100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	m := NewWith[int, struct{}](utils.BasicComparator[int])
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, m, size)
}
