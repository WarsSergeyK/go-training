package main

import (
	r "math/rand"
	o "dice/output"
)

func main(){
	const TryCount = 100000

	var t []int
	t = make([]int, 11)
	
	s := [] int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	
	for i:=1; i<=TryCount; i++ {
		t[diceValue() + diceValue() - 2] ++
	}

	o.PrintRow(s)
	o.PrintRow(t)
}

func diceValue() int {
	const MaxValue = 6
	
	return (r.Intn(MaxValue) + 1)
}
