// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hashmap

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/JonasMuehlmann/datastructures.go/tests"
	"github.com/JonasMuehlmann/datastructures.go/utils"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/maps"
)

func TestMapPut(t *testing.T) {
	m := New[int, string]()
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

	actualValue, expectedValue := m.GetKeys(), []int{1, 2, 3, 4, 5, 6, 7}
	assert.ElementsMatch(t, actualValue, expectedValue)

	actualValue2, expectedValue2 := m.GetValues(), []string{"a", "b", "c", "d", "e", "f", "g"}
	assert.ElementsMatch(t, actualValue2, expectedValue2)

	// key,expectedValue,expectedFound
	tests1 := []struct {
		key   int
		value string
		found bool
	}{
		{1, "a", true},
		{2, "b", true},
		{3, "c", true},
		{4, "d", true},
		{5, "e", true},
		{6, "f", true},
		{7, "g", true},
	}

	for _, test := range tests1 {
		// retrievals
		actualValue, actualFound := m.Get(test.key)
		if actualValue != test.value || actualFound != test.found {
			t.Errorf("Got %v expected %v", actualValue, test.value)
		}
	}
}

func TestMapRemove(t *testing.T) {
	m := New[int, string]()
	m.Put(5, "e")
	m.Put(6, "f")
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(1, "x")
	m.Put(2, "b")
	m.Put(1, "a") //overwrite

	m.Remove(utils.BasicComparator[int], 5)
	m.Remove(utils.BasicComparator[int], 6)
	m.Remove(utils.BasicComparator[int], 7)
	m.Remove(utils.BasicComparator[int], 8)
	m.Remove(utils.BasicComparator[int], 5)

	actualValue, expectedValue := m.GetKeys(), []int{1, 2, 3, 4}
	assert.ElementsMatch(t, actualValue, expectedValue)

	actualValue2, expectedValue2 := m.GetValues(), []string{"a", "b", "c", "d"}
	assert.ElementsMatch(t, actualValue2, expectedValue2)

	if actualValue := m.Size(); actualValue != 4 {
		t.Errorf("Got %v expected %v", actualValue, 4)
	}

	tests2 := []struct {
		key   int
		value string
		found bool
	}{
		{1, "a", true},
		{2, "b", true},
		{3, "c", true},
		{4, "d", true},
	}

	for _, test := range tests2 {
		actualValue, actualFound := m.Get(test.key)
		if actualValue != test.value || actualFound != test.found {
			t.Errorf("Got %v expected %v", actualValue, test.value)
		}

	}

	m.Remove(utils.BasicComparator[int], 1)
	m.Remove(utils.BasicComparator[int], 4)
	m.Remove(utils.BasicComparator[int], 2)
	m.Remove(utils.BasicComparator[int], 3)
	m.Remove(utils.BasicComparator[int], 2)
	m.Remove(utils.BasicComparator[int], 2)

	assert.Empty(t, m.GetKeys())
	assert.Empty(t, m.GetValues())

	if actualValue := m.Size(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
	if actualValue := m.IsEmpty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
}

func TestMapSerialization(t *testing.T) {
	m := New[string, float32]()
	m.Put("a", 1.0)
	m.Put("b", 2.0)
	m.Put("c", 3.0)

	var err error
	assert := func() {
		actualValue, expectedValue := m.GetValues(), []float32{1.0, 2.0, 3.0}
		assert.ElementsMatch(t, actualValue, expectedValue)

		actualValue2, expectedValue2 := m.GetKeys(), []string{"a", "b", "c"}
		assert.ElementsMatch(t, actualValue2, expectedValue2)

		if actualValue, expectedValue := m.Size(), 3; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
		if err != nil {
			t.Errorf("Got error %v", err)
		}
	}

	assert()

	bytes, err := m.ToJSON()
	assert()

	err = m.FromJSON(bytes)
	assert()

	bytes, err = json.Marshal([]string{"a", "b", "c"})
	if err != nil {
		t.Errorf("Got error %v", err)
	}

	err = json.Unmarshal([]byte(`{"a":1,"b":2}`), &m)
	if err != nil {
		t.Errorf("Got error %v", err)
	}
}

func TestMapstring(t *testing.T) {
	c := New[string, int]()
	c.Put("a", 1)
	if !strings.HasPrefix(c.ToString(), "HashMap") {
		t.Errorf("Tostring should start with container name")
	}
}

func BenchmarkHashMapRemove(b *testing.B) {
	b.StopTimer()
	variants := []struct {
		name string
		f    func(n int)
	}{
		{
			name: "Ours",
			f: func(n int) {
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
			f: func(n int) {
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
		f    func(n int)
	}{
		{
			name: "Ours",
			f: func(n int) {
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
			f: func(n int) {
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
		f    func(n int)
	}{
		{
			name: "Ours",
			f: func(n int) {
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
			f: func(n int) {
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
		f    func(n int)
	}{
		{
			name: "Ours",
			f: func(n int) {
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
			f: func(n int) {
				m := New[int, string]()
				for i := 0; i < n; i++ {
					m.Put(i, "foo")
				}
				b.StartTimer()
				_ = maps.Keys(m.m)
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
		f    func(n int)
	}{
		{
			name: "Ours",
			f: func(n int) {
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
			f: func(n int) {
				m := New[int, string]()
				for i := 0; i < n; i++ {
					m.Put(i, "foo")
				}
				b.StartTimer()
				_ = maps.Values(m.m)
				b.StopTimer()
			},
		},
	}

	for _, variant := range variants {
		tests.RunBenchmarkWithDefualtInputSizes(b, variant.name, variant.f)
	}
}
