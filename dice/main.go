package main

import (
	o "dice/output"
	r "math/rand"
)

const (
	tryCount = 100000
	maxValue = 6
)

func main() {
	var t []int
	t = make([]int, 11)

	s := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	frequency := make([]int, 11)

	for i := 1; i <= tryCount; i++ {
		t[diceValue()+diceValue()-2]++
	}

	o.PrintRow(s)
	o.PrintRow(t)

	for i, v := range t {
		frequency[i] = v * 100 / tryCount
	}

	o.PrintRow(frequency)
}

func diceValue() int {
	return (r.Intn(maxValue) + 1)
}
