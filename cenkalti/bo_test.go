package cenkalti

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	bo "github.com/cenkalti/backoff"
	"github.com/pkg/errors"
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

func validate(s string) (string, error) {

	rand.Seed(time.Now().UnixNano())

	operation := func() error {
		n := as[rand.Intn(len(as))]
		if s != n {
			return fmt.Errorf("%s is not %s", s, n)
		}
		return nil
	}

	backoff := bo.NewExponentialBackOff()
	// backoff.MaxElapsedTime = 10 * time.Second

	// setting max retries. is not thread-safe.
	b := bo.WithMaxRetries(backoff, 10)

	err := bo.Retry(operation, b)
	if err != nil {
		return "", errors.WithStack(err)
	}

	return "", nil
}

func TestCenkalti(t *testing.T) {

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			_, err := validate(tt.a)
			if err != nil {
				t.Error(err)
			}
		})
	}
}
