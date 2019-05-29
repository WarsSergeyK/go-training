package domain

import (
	"sync"
)

type Player struct {
	Name  string
	Skill int
}

func (p Player) Play(c chan struct{}, wg *sync.WaitGroup) int {
	defer wg.Done()

	for {
		ball, ok := <-c
		if !ok {
			return 1 // return won point
		}

		tmp := GetRandom(maxBallСomplexity)
		if p.Skill < tmp {
			close(c)
			return 0
		}

		c <- ball
	}
}
