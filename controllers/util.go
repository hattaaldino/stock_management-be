package controllers

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateRandomNumber(length int) string {
	randSeed := time.Now().UnixNano()
	randGenerator := rand.New(rand.NewSource(randSeed))

	result := ""
	for i := 0; i < length; i++ {
		result += strconv.Itoa(randGenerator.Intn(10))
	}
	return result
}
