package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// 必定发财，双色球随机号码

func main() {

	rand.Seed(time.Now().Unix())
	exit := make(map[int]bool)
	var number []int
	for {
		n := rand.Intn(33)
		if _, ok := exit[n]; ok || n == 0 {
			continue
		}

		exit[n] = true
		number = append(number, n)

		if len(number) == 6 {
			break
		}
	}

	sort.Ints(number)
	fmt.Println(number)
	fmt.Println(rand.Intn(16))
}
