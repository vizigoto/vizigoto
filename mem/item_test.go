// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package mem_test

import (
	"testing"

	"github.com/vizigoto/vizigoto/item"
	"github.com/vizigoto/vizigoto/mem"
	"github.com/vizigoto/vizigoto/pkg/testutil"
)

func TestItemPath(t *testing.T) {
	nodes := mem.NewNodeRepository()
	repo := mem.NewItemRepository(nodes)

	name, parent := "Home", ""
	root := item.NewFolder(name, parent)
	rootID, err := repo.Put(root)
	testutil.FatalOnError(t, err)

	a := item.NewFolder("a", rootID)
	aID, err := repo.Put(a)
	testutil.FatalOnError(t, err)

	b := item.NewFolder("b", aID)
	bID, err := repo.Put(b)
	testutil.FatalOnError(t, err)

	c := item.NewFolder("c", bID)
	cID, err := repo.Put(c)
	testutil.FatalOnError(t, err)

	pathIDs := []string{rootID, aID, bID, cID}
	pathNames := []string{root.Name, a.Name, b.Name, c.Name}

	cg, err := repo.Get(cID)
	testutil.FatalOnError(t, err)
	cf := cg.(item.Folder)

	for k, _ := range pathIDs {
		if pathIDs[k] != cf.Path[k].PathID() {
			t.Fatal("path id error")
		}
		if pathNames[k] != cf.Path[k].PathName() {
			t.Fatal("path name error")
		}
	}

	pathIDs = []string{rootID, aID, bID}
	pathNames = []string{root.Name, a.Name, b.Name}

	cg, err = repo.Get(cID)
	testutil.FatalOnError(t, err)
	cf = cg.(item.Folder)

	for k, _ := range pathIDs {
		if pathIDs[k] != cf.Path[k].PathID() {
			t.Fatal("path id error")
		}
		if pathNames[k] != cf.Path[k].PathName() {
			t.Fatal("path name error")
		}
	}
}

func TestPutGetFolder(t *testing.T) {
	nodes := mem.NewNodeRepository()
	repo := mem.NewItemRepository(nodes)

	name, parent := "Home", ""
	root := item.NewFolder(name, parent)
	rootID, err := repo.Put(root)
	testutil.FatalOnError(t, err)

	f, err := repo.Get(rootID)
	testutil.FatalOnError(t, err)

	folder, ok := f.(item.Folder)
	testutil.FatalNotOK(t, ok, "type error")

	if folder.Name != name ||
		folder.Parent != parent {
		t.Fatal("folder properties error")
	}
}

func TestPutGetReport(t *testing.T) {
	nodes := mem.NewNodeRepository()
	repo := mem.NewItemRepository(nodes)

	name, parent, content := "Report", "", "<h1>report"
	r := item.NewReport(name, parent, content)
	id, err := repo.Put(r)
	testutil.FatalOnError(t, err)

	f, err := repo.Get(id)
	testutil.FatalOnError(t, err)

	report, ok := f.(item.Report)
	testutil.FatalNotOK(t, ok, "type error")

	if report.Name != name ||
		report.Parent != parent ||
		report.Content != content {
		t.Fatal("report properties error")
	}
}

func TestChildren(t *testing.T) {
	nodes := mem.NewNodeRepository()
	repo := mem.NewItemRepository(nodes)

	root := item.NewFolder("Home", "")
	rootID, err := repo.Put(root)
	testutil.FatalOnError(t, err)

	a := item.NewFolder("IT", rootID)
	aID, err := repo.Put(a)
	testutil.FatalOnError(t, err)

	b := item.NewReport("Report", rootID, "<h1>report")
	bID, err := repo.Put(b)
	testutil.FatalOnError(t, err)

	f, err := repo.Get(rootID)
	testutil.FatalOnError(t, err)

	folder, ok := f.(item.Folder)
	testutil.FatalNotOK(t, ok, "type error")

	childrenIDs := []string{aID, bID}

	for _, j := range folder.Children {
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
