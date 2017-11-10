// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package testutil_test

import (
	"os"
	"testing"

	"github.com/vizigoto/vizigoto/pkg/testutil"
)

func TestGetDB(t *testing.T) {
	db := testutil.GetDB()
	defer db.Close()

	var version string
	db.QueryRow("select version()").Scan(&version)
	t.Log(version)
}

func TestSemaphoreDB(t *testing.T) {
	os.Setenv("TRAVIS", "")
	os.Setenv("SEMAPHORE", "true")
	os.Setenv("DATABASE_POSTGRESQL_USERNAME", "semaphore")
	os.Setenv("DATABASE_POSTGRESQL_PASSWORD", "semaphore")
	param := testutil.GetParams()
	if param.String() != "host=localhost dbname=vizigoto user=semaphore password=semaphore" {
		t.Fatal("semaphore param db error")
	}
}

func TestTravisDB(t *testing.T) {
	os.Setenv("SEMAPHORE", "")
	os.Setenv("TRAVIS", "true")
	param := testutil.GetParams()
	if param.String() != "host=localhost dbname=travis_ci_test user=postgres password=" {
		t.Fatal("travis param db error")
	}
}

func TestLocalDB(t *testing.T) {
	os.Setenv("SEMAPHORE", "")
	os.Setenv("TRAVIS", "")
	os.Setenv("PGHOSTNAME", "local")
	os.Setenv("PGDATABASE", "local")
	os.Setenv("PGUSERNAME", "local")
	os.Setenv("PGPASSWORD", "local")

	param := testutil.GetParams()
	if param.String() != "host=local dbname=local user=local password=local" {
		t.Log(param)
		t.Fatal("travis param db error")
	}
}
