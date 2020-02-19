package main

import (
	"github.com/go-vgo/robotgo"
	"math/rand"
	"time"
)
	
const (
	FIVE_SECONDS = 5
	MAX_X = 400
	MAX_Y = 400
)

var (
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
)


func init() {
	RandInt(0, time.Now().Nanosecond())
}

func RandInt(min, max int) int {
	return r.Intn(max-min) + min
}

func main() {
	t := time.NewTicker(time.Duration(time.Second * FIVE_SECONDS))
	defer t.Stop()

	for {
		<- t.C
		robotgo.Move(RandInt(0, MAX_X), RandInt(0, MAX_Y))
	}
} 