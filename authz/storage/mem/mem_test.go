// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package mem_test

import (
	"testing"

	"github.com/vizigoto/vizigoto/authz"
	"github.com/vizigoto/vizigoto/authz/storage/mem"
	"github.com/vizigoto/vizigoto/pkg/testutil"
)

func TestGroupNotFound(t *testing.T) {
	var repo auth.Repository = mem.NewRepository()
	_, err := repo.Get("abc")
	if err == nil {
		t.Fatal("error expected")
	}
}

func TestSimplePutGet(t *testing.T) {
	repo := mem.NewRepository()
	group := auth.NewGroup("name", "")
	groupID, err := repo.Put(group)
	testutil.FatalOnError(t, err)

	i, err := repo.Get(groupID)
	testutil.FatalOnError(t, err)

	if groupID != i.ID {
		t.Fatal("id error")
	}

	if group.Name != i.Name {
		t.Fatal("name error")
	}
}
