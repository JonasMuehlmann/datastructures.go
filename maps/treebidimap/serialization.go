// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package treebidimap

import (
	"encoding/json"

	"github.com/JonasMuehlmann/datastructures.go/ds"
	"github.com/JonasMuehlmann/datastructures.go/utils"
)

// Assert Serialization implementation
var _ ds.JSONSerializer = (*Map[string, string])(nil)
var _ ds.JSONDeserializer = (*Map[string, string])(nil)

// ToJSON outputs the JSON representation of the map.
func (m *Map[TKey, TValue]) ToJSON() ([]byte, error) {
	elements := make(map[string]TValue)
	it := m.OrderedBegin(m.forwardMap.Comparator)
	for it.Next() {
		key, _ := it.Index()
		value, _ := it.Get()
		elements[utils.ToString(key)] = value
	}
	return json.Marshal(&elements)
}

// FromJSON populates the map from the input JSON representation.
func (m *Map[TKey, TValue]) FromJSON(data []byte) error {
	elements := make(map[TKey]TValue)
	err := json.Unmarshal(data, &elements)
	if err == nil {
		m.Clear()
		for key, value := range elements {
			m.Put(key, value)
		}
	}
	return err
}

// UnmarshalJSON @implements json.Unmarshaler
func (m *Map[TKey, TValue]) UnmarshalJSON(bytes []byte) error {
	return m.FromJSON(bytes)
}

// MarshalJSON @implements json.Marshaler
func (m *Map[TKey, TValue]) MarshalJSON() ([]byte, error) {
	return m.ToJSON()
}
