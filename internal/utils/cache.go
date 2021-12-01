package utils

import (
	"gopkg.in/night-codes/types.v1"
	"math/rand"
	"time"
)

func GenerateToken() string {
	rand.Seed(time.Now().UnixNano())
	return types.String(rand.Int())
}
