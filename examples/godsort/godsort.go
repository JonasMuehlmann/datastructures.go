// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "github.com/JonasMuehlmann/datastructures.go/utils"

// SortExample to demonstrate basic usage of basic sort
func main() {
	strings := []string{}                              // []
	strings = append(strings, "d")                     // ["d"]
	strings = append(strings, "a")                     // ["d","a"]
	strings = append(strings, "b")                     // ["d","a",b"
	strings = append(strings, "c")                     // ["d","a",b","c"]
	utils.Sort(strings, utils.BasicComparator[string]) // ["a","b","c","d"]
}
