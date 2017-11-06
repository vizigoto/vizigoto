package testutil_test

import (
	"os"
	"testing"

	"github.com/vizigoto/vizigoto/pkg/testutil"
)

func TestDB(t *testing.T) {
	param := testutil.GetParams()

	if param.String() != "host=localhost dbname=vizi user=vizi password=vizi" {
		t.Fatal("local param db error")
	}

	os.Setenv("SEMAPHORE", "true")
	os.Setenv("DATABASE_POSTGRESQL_USERNAME", "semaphore")
	os.Setenv("DATABASE_POSTGRESQL_PASSWORD", "semaphore")

	param = testutil.GetParams()

	if param.String() != "host=localhost dbname=vizigoto user=semaphore password=semaphore" {
		t.Fatal("local param db error")
	}
}
