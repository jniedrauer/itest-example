package testutil

import (
	"flag"
	"testing"
)

var integration bool

func init() {
	flag.BoolVar(&integration, "integration", false, "run integration tests")
}

// IntegrationTest filters integration tests if the flag is not set.
func IntegrationTest(t *testing.T) {
	if !integration {
		t.Skip("skip: integration flag not set")
	}
}
