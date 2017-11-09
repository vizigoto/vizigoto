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
	folder := node.New(node.Folder, "")
	folderID, err := repo.Put(folder)
	testutil.FatalOnError(t, err)

	i, err := repo.Get(folderID)
	if err != nil {
		t.Fatal(err)
	}
	if folder.ID != i.ID {
		t.Fatal("id error")
	}
	if folder.Kind != i.Kind {
		t.Fatal("kind error")
	}
}

func TestSimplePutGetReport(t *testing.T) {
	var repo node.Repository = mem.NewRepository()
	report := node.New(node.Report, "")
	reportID, err := repo.Put(report)
	testutil.FatalOnError(t, err)

	i, err := repo.Get(reportID)
	testutil.FatalOnError(t, err)

	if report.ID != i.ID {
		t.Fatal("id error")
	}
	if report.Kind != i.Kind {
		t.Fatal("kind error")
	}
}

func TestPutGet(t *testing.T) {
	var repo node.Repository = mem.NewRepository()

	root := node.New(node.Folder, "")
	rootID, err := repo.Put(root)
	testutil.FatalOnError(t, err)

	children := []string{"a", "b", "c", "d"}
	childrenIDs := []node.ID{}
	for range children {
		id, err := repo.Put(node.New(node.Folder, rootID))
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
