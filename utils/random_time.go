package utils

import (
	"math/rand"
	"time"
)

func generateRandomTime() time.Time {
	min := time.Date(2011, 3, 1, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(),
		time.Now().Minute(), time.Now().Second(), 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}
