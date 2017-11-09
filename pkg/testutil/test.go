package testutil

import "testing"

//FatalOnError marks a test as failed if there is an error.
func FatalOnError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}
