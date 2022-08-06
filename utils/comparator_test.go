// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utils

import (
	"testing"
	"time"
)

type comparisionInput[T any] struct {
	A           T
	B           T
	Comparision int
}

func TestBasicComparator(t *testing.T) {

	// i1,i2,expected
	tests := []comparisionInput[int]{
		{1, 1, 0},
		{1, 2, -1},
		{2, 1, 1},
		{11, 22, -1},
		{0, 0, 0},
		{1, 0, 1},
		{0, 1, -1},
	}

	for _, test := range tests {
		test := test

		actual := BasicComparator(test.A, test.B)
		expected := test.Comparision
		if actual != expected {
			t.Errorf("Got %v expected %v", actual, expected)
		}
	}
}

func TestTimeComparator(t *testing.T) {

	now := time.Now()

	// i1,i2,expected
	tests := []comparisionInput[time.Time]{
		{now, now, 0},
		{now.Add(24 * 7 * 2 * time.Hour), now, 1},
		{now, now.Add(24 * 7 * 2 * time.Hour), -1},
	}

	for _, test := range tests {
		test := test

		actual := TimeComparator(test.A, test.B)
		expected := test.Comparision
		if actual != expected {
			t.Errorf("Got %v expected %v", actual, expected)
		}
	}
}
