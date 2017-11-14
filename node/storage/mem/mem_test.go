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

func TestItemNotFound(t *testing.T) {
	var repo node.Repository = mem.NewRepository()
	_, err := repo.Get("abc")
	if err == nil {
		t.Fatal("error expected")
	}
}

func TestSimplePutGetFolder(t *testing.T) {
	var repo node.Repository = mem.NewRepository()
	folder := node.New("")
	folderID, err := repo.Put(folder)
	testutil.FatalOnError(t, err)

	i, err := repo.Get(folderID)
	testutil.FatalOnError(t, err)

	if folder.ID != i.ID {
		t.Fatal("id error")
	}
}

func TestSimplePutGetReport(t *testing.T) {
	var repo node.Repository = mem.NewRepository()
	report := node.New("")
	reportID, err := repo.Put(report)
	testutil.FatalOnError(t, err)

	i, err := repo.Get(reportID)
	testutil.FatalOnError(t, err)

	if report.ID != i.ID {
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
