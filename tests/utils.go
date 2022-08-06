// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tests

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func RunBenchmarkWithDefualtInputSizes(b *testing.B, name string, f func(n int, name string)) {
	RunBenchmarkWithInputSizes(b, 10, 1, 4, name, f)
}

func RunBenchmarkWithInputSizes(b *testing.B, base int, lowExponent int, highExponent int, name string, f func(n int, name string)) {
	for i := lowExponent; i <= highExponent; i++ {
		n := math.Pow(float64(base), float64(i))

		fullName := fmt.Sprint(name, "/", n)
		b.Run(fullName, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				f(int(n), fullName)
			}
		})
	}
}

func HandlePanic(t *testing.T, name string) {
	if err := recover(); err != nil {
		assert.Fail(t, name, err)
	}
}
