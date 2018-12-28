package cenkalti

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	bo "github.com/lestrrat-go/backoff"
)

var (
	cases = []struct {
		name string
		a    string
	}{
		{"", "a"},
		{"", "b"},
		{"", "c"},
		{"", "d"},
		{"", "e"},
	}
	as     = []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	policy = bo.NewExponential()
)

func validate(s string) (string, error) {
	b, cancel := policy.Start(context.Background())
	defer cancel()
	rand.Seed(time.Now().UnixNano())

	for bo.Continue(b) {
		n := as[rand.Intn(len(as))]
		if s == n {
			return s, nil
		}
	}

	return "", fmt.Errorf("%s is invalid", s)
}

func TestLestrrat(t *testing.T) {

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			_, err := validate(tt.a)
			if err != nil {
				t.Error(err)
			}
		})
	}
}
