// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package mem_test

import (
	"testing"

	"github.com/vizigoto/vizigoto/pkg/testutil"
	"github.com/vizigoto/vizigoto/user"
	"github.com/vizigoto/vizigoto/user/storage/mem"
)

func TestItemNotFound(t *testing.T) {
	repo := mem.NewRepository()

	_, err := repo.Get("abc")
	if err == nil {
		t.Fatal("error expected")
	}
}

func TestSimplePutGetUser(t *testing.T) {
	name := "Maria"
	userMaria := user.New(name)

	repo := mem.NewRepository()

	userID, err := repo.Put(userMaria)
	testutil.FatalOnError(t, err)

	u, err := repo.Get(userID)
	testutil.FatalOnError(t, err)

	if u.Name != userMaria.Name {
		t.Fatal("user name error")
	}
}
