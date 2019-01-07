package gokit

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/go-kit/kit/util/conn"
)

var (
	maxRetries   int           = 5
	baseInterval time.Duration = 10 * time.Millisecond
)

var cases = []struct {
	name string
	a    string
}{
	{"", "a"},
	{"", "b"},
	{"", "c"},
	{"", "d"},
	{"", "e"},
}
var as = []string{"a", "b", "c", "d", "e", "f", "g", "h"}

func validate(s string, retry int) (string, error) {

	rand.Seed(time.Now().UnixNano())
	d := conn.Exponential(baseInterval)

	n := as[rand.Intn(len(as))]
	retry++
	if retry == maxRetries {
		return "", fmt.Errorf("Reached maxRetries")
	}
	if s != n {
		time.Sleep(d)
		validate(s, retry)
	}

	return n, nil
}

func TestCenkalti(t *testing.T) {

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			_, err := validate(tt.a, 0)
			if err != nil {
				t.Error(err)
			}
		})
	}
}
