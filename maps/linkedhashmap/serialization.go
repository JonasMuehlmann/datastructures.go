// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linkedhashmap

import (
	"bytes"
	"encoding/json"

	"github.com/JonasMuehlmann/datastructures.go/ds"
	"github.com/JonasMuehlmann/datastructures.go/utils"
)

// Assert Serialization implementation
var _ ds.JSONSerializer = (*Map[string, any])(nil)
var _ ds.JSONDeserializer = (*Map[string, any])(nil)

// ToJSON outputs the JSON representation of map.
func (m *Map[TKey, TValue]) ToJSON() ([]byte, error) {
	var b []byte
	buf := bytes.NewBuffer(b)

	buf.WriteRune('{')

	it := m.Begin()
	lastIndex := m.Size() - 1
	index := 0

	for it.Next() {
		key, _ := it.Index()
		km, err := json.Marshal(key)
		if err != nil {
			return nil, err
		}
		buf.Write(km)

		buf.WriteRune(':')

		value, _ := it.Get()
		vm, err := json.Marshal(value)
		if err != nil {
			return nil, err
		}
		buf.Write(vm)

		if index != lastIndex {
			buf.WriteRune(',')
		}

		index++
	}

	buf.WriteRune('}')

	return buf.Bytes(), nil
}

// FromJSON populates map from the input JSON representation.
//func (m *Map) FromJSON(data []byte) error {
//	elements := make(map[string]interface{})
//	err := json.Unmarshal(data, &elements)
//	if err == nil {
//		m.Clear()
//		for key, value := range elements {
//			m.Put(key, value)
//		}
//	}
//	return err
//}

// FromJSON populates map from the input JSON representation.
func (m *Map[TKey, TValue]) FromJSON(data []byte) error {
	elements := make(map[TKey]TValue)
	err := json.Unmarshal(data, &elements)
	if err != nil {
		return err
	}

	index := make(map[TKey]int)
	var keys []TKey
	for key := range elements {
		keys = append(keys, key)
		esc, _ := json.Marshal(key)
		index[key] = bytes.Index(data, esc)
	}

	byIndex := func(a, b TKey) int {
		index1 := index[a]
		index2 := index[b]
		return index1 - index2
	}

	utils.Sort(keys, byIndex)

	m.Clear()

	for _, key := range keys {
		m.Put(key, elements[key])
	}

	return nil
}

// UnmarshalJSON @implements json.Unmarshaler
func (m *Map[Tkey, TValue]) UnmarshalJSON(bytes []byte) error {
	return m.FromJSON(bytes)
}

// MarshalJSON @implements json.Marshaler
func (m *Map[TKey, TValue]) MarshalJSON() ([]byte, error) {
	return m.ToJSON()
}
