// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package item_test

import (
	"os"
	"testing"

	"github.com/vizigoto/vizigoto/item"
	"github.com/vizigoto/vizigoto/item/storage/mem"
	node "github.com/vizigoto/vizigoto/node/storage/mem"
	"github.com/vizigoto/vizigoto/pkg/log"
	"github.com/vizigoto/vizigoto/pkg/testutil"
)

func TestLoggingRepository(t *testing.T) {
	w := log.NewSyncWriter(os.Stdout)
	logger := log.NewLogfmtLogger(w)

	nodes := node.NewRepository()
	repo := mem.NewRepository(nodes)
	repo = item.NewLoggingRepository(logger, repo)

	name, parent := "Home", item.ID("")
	root := item.NewFolder(name, parent)
	id, err := repo.Put(root)
	testutil.FatalOnError(t, err)

	n, err := repo.Get(id)
	testutil.FatalOnError(t, err)

	if folder, ok := n.(*item.Folder); ok {
		if root.Name != folder.Name {
			t.Fatal("error")
		}
	}
}

func TestLoggingService(t *testing.T) {
	w := log.NewSyncWriter(os.Stdout)
	logger := log.NewLogfmtLogger(w)

	nodes := node.NewRepository()
	repo := mem.NewRepository(nodes)
	service := item.NewService(repo)
	service = item.NewLoggingService(logger, service)

	rootID, err := service.AddFolder("Home", item.ID(""))
	testutil.FatalOnError(t, err)

	reportID, err := service.AddReport("report", rootID, "report content")
	testutil.FatalOnError(t, err)

	r, err := service.Get(reportID)
	testutil.FatalOnError(t, err)

	rep, ok := r.(*item.Report)
	testutil.FatalNotOK(t, ok, "type error")

	if rep.Name != "report" {
		t.Fatal("report error")
	}
}
