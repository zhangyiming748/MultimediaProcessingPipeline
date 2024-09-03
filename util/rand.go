package util

import (
	"math/rand"
	"time"
)

var Seed *rand.Rand

func SetSeed() {
	Seed = rand.New(rand.NewSource(time.Now().Unix()))
}
func GetSeed() *rand.Rand {
	return Seed
}
