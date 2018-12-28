package cenkalti

import (
	"fmt"
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

	operation := func() error {
		if s != "a" {
			return fmt.Errorf("%s is not 'a'", s)
		}
		return nil
	}

	backoff := bo.NewExponentialBackOff()
	backoff.MaxElapsedTime = 3 * time.Second

	err := bo.Retry(operation, backoff)
	if err != nil {
		return "", errors.WithStack(err)
	}

	return "", nil
}

func TestCenkalti(t *testing.T) {

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			s, err := validate(tt.a)
			if err != nil {
				t.Error(err)
			}
			t.Log(s)
		})
	}
}
