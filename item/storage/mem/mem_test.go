// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package mem_test

import (
	"testing"

	"github.com/vizigoto/vizigoto/item"
	"github.com/vizigoto/vizigoto/item/storage/mem"
	node "github.com/vizigoto/vizigoto/node/storage/mem"
)

func TestPutGetFolder(t *testing.T) {
	nodes := node.NewRepository()
	repo := mem.NewRepository(nodes)

	name, parent := "Home", item.ID("")
	root := item.NewFolder(name, parent)
	rootID, err := repo.Put(root)
	if err != nil {
		t.Fatal(err)
	}

	f, err := repo.Get(rootID)
	if err != nil {
		t.Fatal(err)
	}

	folder, ok := f.(*item.Folder)
	if !ok {
		t.Fatal("type match fail")
	}

	if folder.Name != name ||
		folder.Parent != parent {
		t.Fatal("folder properties error")
	}
}

func TestPutGetReport(t *testing.T) {
	nodes := node.NewRepository()
	var repo item.Repository = mem.NewRepository(nodes)

	name, parent, content := "Report", item.ID(""), "<h1>report"
	r := item.NewReport(name, parent, content)
	id, err := repo.Put(r)
	if err != nil {
		t.Fatal(err)
	}

	f, err := repo.Get(id)
	if err != nil {
		t.Fatal(err)
	}

	report, ok := f.(*item.Report)
	if !ok {
		t.Fatal("type match fail")
	}

	if report.Name != name ||
		report.Parent != parent ||
		report.Content != content {
		t.Fatal("report properties error")
	}
}

func TestChildren(t *testing.T) {
	nodes := node.NewRepository()
	var repo item.Repository = mem.NewRepository(nodes)
	root := item.NewFolder("Home", "")
	rootID, err := repo.Put(root)
	if err != nil {
		t.Fatal(err)
	}
	a := item.NewFolder("IT", rootID)
	aID, err := repo.Put(a)
	if err != nil {
		t.Fatal(err)
	}

	b := item.NewReport("Report", rootID, "<h1>report")
	bID, err := repo.Put(b)
	if err != nil {
		t.Fatal(err)
	}

	f, err := repo.Get(rootID)
	if err != nil {
		t.Fatal(err)
	}

	folder, ok := f.(*item.Folder)
	if !ok {
		t.Fatal("type match fail")
	}

	childrenIDs := []item.ID{aID, bID}

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
