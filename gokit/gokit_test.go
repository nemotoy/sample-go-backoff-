package gokit_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-kit/kit/util/conn"
)

var (
	maxAttemt int = 10
)

func BenchmarkGokit(b *testing.B) {

	for i := 0; i < b.N; i++ {
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
}
