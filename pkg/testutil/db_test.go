package testutil_test

import (
	"os"
	"testing"

	"github.com/vizigoto/vizigoto/pkg/testutil"
)

func TestDB(t *testing.T) {
	os.Setenv("SEMAPHORE", "true")
	os.Setenv("DATABASE_POSTGRESQL_USERNAME", "semaphore")
	os.Setenv("DATABASE_POSTGRESQL_PASSWORD", "semaphore")
	param := testutil.GetParams()
	if param.String() != "host=localhost dbname=vizigoto user=semaphore password=semaphore" {
		t.Fatal("semaphore param db error")
	}

	os.Setenv("SEMAPHORE", "")
	os.Setenv("TRAVIS", "true")
	param = testutil.GetParams()
	if param.String() != "host=localhost dbname=travis_ci_test user=postgres password=" {
		t.Fatal("travis param db error")
	}
}
