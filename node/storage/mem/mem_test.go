// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package mem_test

import (
	"testing"

	"github.com/vizigoto/vizigoto/node"
	"github.com/vizigoto/vizigoto/node/storage/mem"
	"github.com/vizigoto/vizigoto/pkg/testutil"
)

func TestPath(t *testing.T) {
	repo := mem.NewRepository()
	root := node.New("")
	rootID, err := repo.Put(root)
	testutil.FatalOnError(t, err)

	a := node.New(rootID)
	aID, err := repo.Put(a)
	testutil.FatalOnError(t, err)

	b := node.New(aID)
	bID, err := repo.Put(b)
	testutil.FatalOnError(t, err)

	c := node.New(bID)
	cID, err := repo.Put(c)
	testutil.FatalOnError(t, err)

	path := []string{rootID, aID, bID, cID}

	c, err = repo.Get(cID)
	testutil.FatalOnError(t, err)

	for k, v := range path {
		if v != c.Path[k] {
			t.Fatal("path error")
		}
	}

	path = []string{rootID, aID, bID}

	c, err = repo.Get(bID)
	testutil.FatalOnError(t, err)

	for k, v := range path {
		if v != c.Path[k] {
			t.Fatal("path error")
		}
	}
}

func TestItemNotFound(t *testing.T) {
	repo := mem.NewRepository()
	_, err := repo.Get("abc")
	if err == nil {
		t.Fatal("error expected")
	}
}

func TestSimplePutGet(t *testing.T) {
	repo := mem.NewRepository()
	folder := node.New("")
	folderID, err := repo.Put(folder)
	testutil.FatalOnError(t, err)

	i, err := repo.Get(folderID)
	testutil.FatalOnError(t, err)

	if folder.ID != i.ID {
		t.Fatal("id error")
	}
}

func TestPutGet(t *testing.T) {
	var repo node.Repository = mem.NewRepository()

	root := node.New("")
	rootID, err := repo.Put(root)
	testutil.FatalOnError(t, err)

	children := []string{"a", "b", "c", "d"}
	childrenIDs := []string{}
	for range children {
		id, err := repo.Put(node.New(rootID))
		testutil.FatalOnError(t, err)

		childrenIDs = append(childrenIDs, id)
	}

	i, err := repo.Get(rootID)
	testutil.FatalOnError(t, err)

	for _, j := range i.Children {
		fail := true
		for _, c := range childrenIDs {
			if j == c {
				fail = false
			}
		}
		if fail {
			t.Fatal("children not found")
		}
	}
}
