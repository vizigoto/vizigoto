// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package user_test

import (
	"testing"

	"github.com/vizigoto/vizigoto/mem"
	"github.com/vizigoto/vizigoto/pkg/testutil"
	"github.com/vizigoto/vizigoto/user"
)

func TestNew(t *testing.T) {
	name := "Maria"
	u := user.New(name)

	if u.Name != name {
		t.Fatal("name error")
	}
}

func TestService(t *testing.T) {
	name := "Maria"

	repo := mem.NewUserRepository()
	service := user.NewService(repo)
	id, err := service.AddUser(name)
	testutil.FatalOnError(t, err)

	u, err := service.Get(id)
	testutil.FatalOnError(t, err)

	if u.Name != name {
		t.Fatal("user name error")
	}
}
