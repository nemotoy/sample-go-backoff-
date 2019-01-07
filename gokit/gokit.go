package gokit

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/util/conn"
)

var (
	maxAttemt int = 10
)

// create sample exponential backoff duration
func main() {

	var (
		d time.Duration = 1 * time.Millisecond
	)
	for i := 1; i <= maxAttemt; i++ {
		d = conn.Exponential(d)

		// do something...

		fmt.Printf("dur: %s, attempt: %v\n", d, i)
		if d == time.Minute {
			break
		}
	}
}
