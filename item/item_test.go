// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package item_test

import (
	"testing"

	"github.com/vizigoto/vizigoto/item"
	"github.com/vizigoto/vizigoto/item/storage/mem"
	node "github.com/vizigoto/vizigoto/node/storage/mem"
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
	nodes := node.NewRepository()
	repo := mem.NewRepository(nodes)
	service := item.NewService(repo)

	rootName, rootParent := "Home", ""
	rootID, err := service.AddFolder(rootName, rootParent)
	testutil.FatalOnError(t, err)

	reportName, reportContent := "report", "<h1>content"
	reportID, err := service.AddReport(reportName, rootID, reportContent)
	testutil.FatalOnError(t, err)

	f, err := service.Get(rootID)
	testutil.FatalOnError(t, err)

	r, err := service.Get(reportID)
	testutil.FatalOnError(t, err)

	folder, ok := f.(*item.Folder)
	testutil.FatalNotOK(t, ok, "type error")

	if folder.Name != rootName ||
		folder.Parent != rootParent {
		t.Fatal("folder error")
	}

	report, ok := r.(*item.Report)
	testutil.FatalNotOK(t, ok, "type error")

	if report.Name != reportName ||
		report.Content != reportContent ||
		report.Parent != rootID {
		t.Fatal("report error")
	}
}
