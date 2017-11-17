// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package mem_test

import (
	"context"
	"testing"

	"github.com/vizigoto/vizigoto/mem"
	"github.com/vizigoto/vizigoto/node"
	"github.com/vizigoto/vizigoto/pkg/testutil"
)

func TestPath(t *testing.T) {
	ctx := context.Background()
	repo := mem.NewNodeRepository()
	root := node.New("")
	rootID, err := repo.Put(ctx, root)
	testutil.FatalOnError(t, err)

	a := node.New(rootID)
	aID, err := repo.Put(ctx, a)
	testutil.FatalOnError(t, err)

	b := node.New(aID)
	bID, err := repo.Put(ctx, b)
	testutil.FatalOnError(t, err)

	c := node.New(bID)
	cID, err := repo.Put(ctx, c)
	testutil.FatalOnError(t, err)

	path := []string{rootID, aID, bID, cID}

	c, err = repo.Get(ctx, cID)
	testutil.FatalOnError(t, err)

	for k, v := range path {
		if v != c.Path[k].ID {
			t.Fatal("path error")
		}
	}

	path = []string{rootID, aID, bID}

	c, err = repo.Get(ctx, bID)
	testutil.FatalOnError(t, err)

	for k, v := range path {
		if v != c.Path[k].ID {
			t.Fatal("path error")
		}
	}
}

func TestNodeNotFound(t *testing.T) {
	ctx := context.Background()
	repo := mem.NewNodeRepository()
	_, err := repo.Get(ctx, "abc")
	if err == nil {
		t.Fatal("error expected")
	}
}

func TestPutGet(t *testing.T) {
	ctx := context.Background()
	repo := mem.NewNodeRepository()

	root := node.New("")
	rootID, err := repo.Put(ctx, root)
	testutil.FatalOnError(t, err)

	children := []string{"a", "b", "c", "d"}
	childrenIDs := []string{}
	for range children {
		id, err := repo.Put(ctx, node.New(rootID))
		testutil.FatalOnError(t, err)

		childrenIDs = append(childrenIDs, id)
	}

	i, err := repo.Get(ctx, rootID)
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

func TestMove(t *testing.T) {
	ctx := context.Background()
	repo := mem.NewNodeRepository()

	root := node.New("")
	rootID, err := repo.Put(ctx, root)
	testutil.FatalOnError(t, err)

	a := node.New(rootID)
	aID, err := repo.Put(ctx, a)
	testutil.FatalOnError(t, err)

	b := node.New(rootID)
	bID, err := repo.Put(ctx, b)
	testutil.FatalOnError(t, err)

	c := node.New(aID)
	cID, err := repo.Put(ctx, c)
	testutil.FatalOnError(t, err)
	c, err = repo.Get(ctx, cID)
	testutil.FatalOnError(t, err)

	pathToC := []string{rootID, aID, cID}
	for k, v := range pathToC {
		if v != c.Path[k].ID {
			t.Fatal("path error")
		}
	}

	err = repo.Move(ctx, c, bID)
	testutil.FatalOnError(t, err)

	pathToC = []string{rootID, bID, cID}
	c, err = repo.Get(ctx, cID)
	testutil.FatalOnError(t, err)

	for k, v := range pathToC {
		if v != c.Path[k].ID {
			t.Fatal("path error")
		}
	}
}
