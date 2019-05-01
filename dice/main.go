package main

import (
	r "math/rand"
	o "dice/output"
)

const (
	tryCount = 100000
	maxValue = 6
)

func main(){
	var t []int
	t = make([]int, 11)
	
	s := [] int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	
	for i:=1; i<=tryCount; i++ {
		t[diceValue() + diceValue() - 2] ++
	}

	o.PrintRow(s)
	o.PrintRow(t)
}

func diceValue() int {
	return (r.Intn(maxValue) + 1)
}
