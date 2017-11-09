// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package postgres_test

import (
	"testing"

	_ "github.com/lib/pq"
	"github.com/vizigoto/vizigoto/item"
	"github.com/vizigoto/vizigoto/item/storage/postgres"
	node "github.com/vizigoto/vizigoto/node/storage/postgres"
	"github.com/vizigoto/vizigoto/pkg/testutil"
)

func TestPutGetFolder(t *testing.T) {
	db := testutil.GetDB()
	defer db.Close()

	nodes := node.NewRepository(db)
	repo := postgres.NewRepository(db, nodes)

	root := item.NewFolder("Home", "")
	id, err := repo.Put(root)
	testutil.FatalOnError(t, err)

	f, err := repo.Get(id)
	testutil.FatalOnError(t, err)

	folder, ok := f.(*item.Folder)
	if !ok {
		t.Fatal("type match fail")
	}

	if folder.Name != root.Name ||
		folder.Parent != root.Parent {
		t.Fatal("properties error")
	}
}

func TestPutGetReport(t *testing.T) {
	db := testutil.GetDB()
	defer db.Close()

	nodes := node.NewRepository(db)
	repo := postgres.NewRepository(db, nodes)

	root := item.NewFolder("Home", "")
	rootID, err := repo.Put(root)
	testutil.FatalOnError(t, err)

	r := item.NewReport("hr report", rootID, "content")
	id, err := repo.Put(r)
	testutil.FatalOnError(t, err)

	result, err := repo.Get(id)
	testutil.FatalOnError(t, err)

	report, ok := result.(*item.Report)
	if !ok {
		t.Fatal("type error")
	}

	if r.Name != report.Name ||
		r.Parent != report.Parent ||
		r.Content != report.Content {
		t.Fatal("properties error")
	}

}
