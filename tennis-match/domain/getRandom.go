package domain

import (
	"math/rand"
	"time"
)

func GetRandom(max int) int {
	t := time.Now().UnixNano()
	s := rand.NewSource(t)
	r := rand.New(s)

	return r.Intn(max + 1)
}
