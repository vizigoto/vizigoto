// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package mem_test

import (
	"testing"

	"github.com/vizigoto/vizigoto/node"
	"github.com/vizigoto/vizigoto/node/storage/mem"
)

func TestPutGet(t *testing.T) {
	var repo node.Repository = mem.NewRepository()
	root := node.New("Home", node.Folder, node.ID(""), "x")
	rootID, err := repo.Put(root)
	if err != nil {
		t.Fatal(err)
	}

	i, err := repo.Get(rootID)
	if err != nil {
		t.Fatal(err)
	}

	if root.ID != i.ID {
		t.Fatal("id error")
	}

	_, err = repo.Get("unknow")
	if err == nil {
		t.Fatal("error expected")
	}

	children := []string{"a", "b", "c", "d"}
	childrenIDs := []node.ID{}
	for _, c := range children {
		id, err := repo.Put(node.New(c, node.Folder, rootID, "x"))
		if err != nil {
			t.Fatal(err)
		}
		childrenIDs = append(childrenIDs, id)
	}

	i, err = repo.Get(rootID)
	if err != nil {
		t.Fatal(err)
	}
	ok := true
	for _, j := range i.Children {
		for _, c := range childrenIDs {
			if j == c {
				ok = false
			}
		}
	}
	if ok {
		t.Fatal("children fail")
	}
}
