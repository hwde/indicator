// Copyright (c) 2021-2024 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package helper

import "slices"

// Highest returns a channel that emits the highest value
// within a sliding window of size w from the input channel c.
func Highest[T Number](c <-chan T, w int) <-chan T {
	r := make(chan T)

	go func() {
		defer close(r)
		h := make([]T, w)
		n, cnt := 0, 0
		ok := true
		for ok {
			if h[n], ok = <-c; ok {
				if cnt < w {
					cnt++
					r <- slices.Max(h[:cnt])
				} else {
					r <- slices.Max(h)
				}
			}
			n = (n + 1) % w
		}
	}()

	return r
}
