package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	var randomValue int

	randomValue = randomWithRange(2, 8)
	fmt.Println("number random :", randomValue)

}
func randomWithRange(min, max int) int {
	var value = rand.Int()%(max-min+1) + min
	return value
}
