package utils

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateSimpleRandomString() string {
	rand.Seed(time.Now().UnixNano())
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	randomChar := rand.Intn(100)

	return timestamp + "-" + strconv.Itoa(randomChar)
}
