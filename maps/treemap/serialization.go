// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package treemap

import (
	"github.com/JonasMuehlmann/datastructures.go/ds"
)

// Assert Serialization implementation
var _ ds.JSONSerializer = (*Map[string, any])(nil)
var _ ds.JSONDeserializer = (*Map[string, any])(nil)

// ToJSON outputs the JSON representation of the map.
func (m *Map[Tkey, TValue]) ToJSON() ([]byte, error) {
	return m.tree.ToJSON()
}

// FromJSON populates the map from the input JSON representation.
func (m *Map[Tkey, TValue]) FromJSON(data []byte) error {
	return m.tree.FromJSON(data)
}

// UnmarshalJSON @implements json.Unmarshaler
func (m *Map[Tkey, TValue]) UnmarshalJSON(bytes []byte) error {
	return m.FromJSON(bytes)
}

// MarshalJSON @implements json.Marshaler
func (m *Map[Tkey, TValue]) MarshalJSON() ([]byte, error) {
	return m.ToJSON()
}
