package main

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/util/conn"
)

// create sample exponential backoff duration
func main() {

	var (
		d          time.Duration = 1 * time.Millisecond
		maxAttemot int           = 10
	)
	for i := 1; i <= maxAttemot; i++ {
		d = conn.Exponential(d)

		// do something...

		fmt.Printf("dur: %s, attempt: %v\n", d, i)
		if d == time.Minute {
			break
		}
	}
}
