package main

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/util/conn"
)

// create sample exponential backoff duration
func main() {
	var d time.Duration = 1 * time.Second

	for i := 0; i < 10; i++ {
		fmt.Println("base: ", d)
		d = conn.Exponential(d)
		fmt.Println("dur: ", d)
		fmt.Println("-----------------------------------")
		if d == time.Minute {
			break
		}
	}
}
