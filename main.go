package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start time:", time.Now().UTC())
	l := 10
	ticker := time.NewTicker(time.Second * 1)
	var curr, last time.Time = time.Now(), time.Now()
	diff := make([]time.Duration, l)
	for i := 1; i <= l; i++ {
		c := time.Now().UTC()
		last = curr
		curr = c
		fmt.Printf("\rtimer: %d\ttick time: %v", i, c)
		diff[i-1] = curr.Sub(last)
		<-ticker.C
	}
	ticker.Stop()
	fmt.Printf("\rtimer done: %ds\tend time: %v\n", l, time.Now().UTC())
	usedif := diff[1:]
	var totalt int64
	for _, ud := range usedif {
		diff := ud - (time.Second * 1)
		fmt.Println(diff)
		totalt += diff.Microseconds()
	}
	fmt.Println(totalt / 9)
}
