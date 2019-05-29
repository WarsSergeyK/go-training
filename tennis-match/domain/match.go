package domain

import (
	"fmt"
	"sync"
)

type Match struct {
	Score [2]int
}

func (m Match) Start() {

	p1 := Player{Name: "Vasya", Skill: 80}
	p2 := Player{Name: "Petya", Skill: 81}

	ch := make(chan struct{})

	wg := &sync.WaitGroup{}
	wg.Add(2)

	var winner string

	go func(ch chan struct{}) {
		point := p1.Play(ch, wg)
		m.Score[0] += point
		if point == 1 {
			winner = p1.Name
		}
	}(ch)

	go func(ch chan struct{}) {
		point := p2.Play(ch, wg)
		m.Score[1] += point
		if point == 1 {
			winner = p2.Name
		}
	}(ch)

	ch <- struct{}{}
	wg.Wait()

	fmt.Println(winner, "won!")

	fmt.Println("Score:", m.Score)
}
