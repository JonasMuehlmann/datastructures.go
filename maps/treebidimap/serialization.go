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
var _ ds.JSONSerializer = (*Map)(nil)
var _ ds.JSONDeserializer = (*Map)(nil)

// ToJSON outputs the JSON representation of the map.
func (m *Map) ToJSON() ([]byte, error) {
	elements := make(map[string]interface{})
	it := m.Iterator()
	for it.Next() {
		elements[utils.ToString(it.Key())] = it.Value()
	}
	return json.Marshal(&elements)
}

// FromJSON populates the map from the input JSON representation.
func (m *Map) FromJSON(data []byte) error {
	elements := make(map[string]interface{})
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
func (m *Map) UnmarshalJSON(bytes []byte) error {
	return m.FromJSON(bytes)
}

// MarshalJSON @implements json.Marshaler
func (m *Map) MarshalJSON() ([]byte, error) {
	return m.ToJSON()
}
