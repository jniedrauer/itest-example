package pkg1

import (
	"testing"

	"github.com/jniedrauer/itest-example/testutil"
)

func TestHelloWorld(t *testing.T) {
	testutil.IntegrationTest(t)

	t.Log("integration test ran")
}
