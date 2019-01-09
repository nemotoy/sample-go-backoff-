package cenkalti

import (
	"fmt"
	"log"
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
var as = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n"}

func validate(s string) (string, error) {

	rand.Seed(time.Now().UnixNano())

	operation := func() error {
		fmt.Println("***")
		n := as[rand.Intn(len(as))]
		if s != n {
			return fmt.Errorf("%s is not %s", s, n)
		}
		return nil
	}

	backoff := bo.NewExponentialBackOff()
	backoff.InitialInterval = 500 * time.Millisecond
	backoff.RandomizationFactor = 0.5
	backoff.Multiplier = 1.5
	backoff.MaxInterval = 5 * time.Second
	backoff.MaxElapsedTime = 30 * time.Second

	// not included initial request. so, amount of requests is max 4 times.
	// b := bo.WithMaxRetries(backoff, 5)

	notify := func(err error, d time.Duration) {
		log.Printf("err: %v, duration: %v", err, d)
	}

	err := bo.RetryNotify(operation, backoff, notify)
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
