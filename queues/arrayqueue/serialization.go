// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package arrayqueue

import (
	"github.com/JonasMuehlmann/datastructures.go/ds"
)

// Assert Serialization implementation
var _ ds.JSONSerializer = (*Queue[any])(nil)
var _ ds.JSONDeserializer = (*Queue[any])(nil)

// ToJSON outputs the JSON representation of the queue.
func (queue *Queue[T]) ToJSON() ([]byte, error) {
	return queue.list.ToJSON()
}

// FromJSON populates the queue from the input JSON representation.
func (queue *Queue[T]) FromJSON(data []byte) error {
	return queue.list.FromJSON(data)
}

// UnmarshalJSON @implements json.Unmarshaler
func (queue *Queue[T]) UnmarshalJSON(bytes []byte) error {
	return queue.FromJSON(bytes)
}

// MarshalJSON @implements json.Marshaler
func (queue *Queue[T]) MarshalJSON() ([]byte, error) {
	return queue.ToJSON()
}
