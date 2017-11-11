// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package auth_test

import (
	"testing"

	"github.com/vizigoto/vizigoto/authz"
	"github.com/vizigoto/vizigoto/authz/storage/mem"
	"github.com/vizigoto/vizigoto/pkg/testutil"
)

func TestAuthServiceManager(t *testing.T) {
	repo := mem.NewRepository()
	service := auth.NewAuthServiceManager(repo)

	rootName, rootParent := "root", ""
	rootID, err := service.AddGroup(rootName, rootParent)
	testutil.FatalOnError(t, err)

	fooName := "report"
	fooID, err := service.AddGroup(fooName, rootID)
	testutil.FatalOnError(t, err)

	root, err := service.Get(rootID)
	testutil.FatalOnError(t, err)

	foo, err := service.Get(fooID)
	testutil.FatalOnError(t, err)

	if root.Name != rootName {
		t.Fatal("root name fail")
	}
	if foo.Name != fooName {
		t.Fatal("foo name fail")
	}

}
