package util

import (
	"math/rand"
	"time"
)

func GetSeed() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().Unix()))
}
