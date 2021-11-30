package utils

import (
	"gopkg.in/night-codes/types.v1"
	"time"
)

func GenerateToken() string {
	return types.String(time.Now().UnixNano())
}
