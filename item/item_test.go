// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package item_test

import (
	"context"
	"testing"

	"github.com/vizigoto/vizigoto/item"
	"github.com/vizigoto/vizigoto/mem"
	"github.com/vizigoto/vizigoto/pkg/testutil"
)

func TestNewFolder(t *testing.T) {
	name, parent := "Home", ""
	i := item.NewFolder(name, parent)
	if i.Name != name ||
		i.Parent != parent {
		t.Fatal("folder error")
	}
}

func TestNewReport(t *testing.T) {
	name, parent, content := "Home", "", "<h1>report"
	i := item.NewReport(name, parent, content)
	if i.Name != name ||
		i.Parent != parent {
		t.Fatal("report error")
	}
}

func TestService(t *testing.T) {
	ctx := context.Background()
	nodes := mem.NewNodeRepository()
	repo := mem.NewItemRepository(nodes)
	service := item.NewService(repo)

	rootName, rootParent := "Home", ""
	rootID, err := service.AddFolder(ctx, rootName, rootParent)
	testutil.FatalOnError(t, err)

	reportName, reportContent := "report", "<h1>content"
	reportID, err := service.AddReport(ctx, reportName, rootID, reportContent)
	testutil.FatalOnError(t, err)

	f, err := service.Get(ctx, rootID)
	testutil.FatalOnError(t, err)

	r, err := service.Get(ctx, reportID)
	testutil.FatalOnError(t, err)

	folder, ok := f.(item.Folder)
	testutil.FatalNotOK(t, ok, "type error")

	if folder.Name != rootName ||
		folder.Parent != rootParent {
		t.Fatal("folder error")
	}

	report, ok := r.(item.Report)
	testutil.FatalNotOK(t, ok, "type error")

	if report.Name != reportName ||
		report.Content != reportContent ||
		report.Parent != rootID {
		t.Fatal("report error")
	}
}

func TestPathInterface(t *testing.T) {
	ctx := context.Background()
	nodes := mem.NewNodeRepository()
	repo := mem.NewItemRepository(nodes)
	service := item.NewService(repo)

	rootName, rootParent := "Home", ""
	rootID, err := service.AddFolder(ctx, rootName, rootParent)
	testutil.FatalOnError(t, err)
	rootItem, err := service.Get(ctx, rootID)
	testutil.FatalOnError(t, err)
	root := rootItem.(item.Folder)

	aName, aParent := "a", rootID
	aID, err := service.AddFolder(ctx, aName, aParent)
	testutil.FatalOnError(t, err)
	aItem, err := service.Get(ctx, aID)
	testutil.FatalOnError(t, err)
	a := aItem.(item.Folder)

	bName, bParent := "b", aID
	bID, err := service.AddFolder(ctx, bName, bParent)
	testutil.FatalOnError(t, err)
	bItem, err := service.Get(ctx, bID)
	testutil.FatalOnError(t, err)
	b := bItem.(item.Folder)

	cName, cParent, cContent := "c", bID, "content"
	cID, err := service.AddReport(ctx, cName, cParent, cContent)
	testutil.FatalOnError(t, err)
	cItem, err := service.Get(ctx, cID)
	testutil.FatalOnError(t, err)
	c := cItem.(item.Report)

	pathToReport := []item.Path{root, a, b, c}

	for k := range pathToReport {
		if pathToReport[k].PathID() != c.Path[k].PathID() {
			t.Fatal("path error")
		}
		if pathToReport[k].PathName() != c.Path[k].PathName() {
			t.Fatal("path error")
		}
	}
}
