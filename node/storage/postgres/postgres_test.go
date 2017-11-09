// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package postgres_test

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	"github.com/vizigoto/vizigoto/node"
	"github.com/vizigoto/vizigoto/node/storage/postgres"
	"github.com/vizigoto/vizigoto/pkg/testutil"
)

func TestGetChild(t *testing.T) {
	db := testutil.GetDB()
	defer db.Close()

	repo := postgres.NewRepository(db)
	root := node.New("")
	rootID, err := repo.Put(root)
	testutil.FatalOnError(t, err)

	childNode := node.New(rootID)
	childID, err := repo.Put(childNode)
	testutil.FatalOnError(t, err)

	c, err := repo.Get(childID)
	testutil.FatalOnError(t, err)

	if c.ID != childID ||
		c.Parent != childNode.Parent {
		t.Fatal("child error")
	}
}

func TestPutRoot(t *testing.T) {
	db := testutil.GetDB()
	defer db.Close()
	repo := postgres.NewRepository(db)
	root := node.New("")
	id, err := repo.Put(root)
	testutil.FatalOnError(t, err)

	var n node.Node
	var p sql.NullString
	var lft, rgt int

	query := "select id, parent, lft, rgt from vinodes.nodes where id = $1"
	err = db.QueryRow(query, id).Scan(&n.ID, &p, &lft, &rgt)
	testutil.FatalOnError(t, err)

	if p.Valid {
		t.Fatal("null expected")
	}
	if rgt-1-lft != 0 {
		t.Fatal("lft or rgt error")
	}
}

func TestPutFirstChild(t *testing.T) {
	db := testutil.GetDB()
	defer db.Close()

	repo := postgres.NewRepository(db)
	rootNode := node.New("")
	rootID, err := repo.Put(rootNode)
	testutil.FatalOnError(t, err)

	childNode := node.New(rootID)
	childID, err := repo.Put(childNode)
	testutil.FatalOnError(t, err)

	var n node.Node
	var p sql.NullString
	var lft, rgt int
	query := "select id, parent, lft, rgt from vinodes.nodes where id = $1"
	err = db.QueryRow(query, childID).Scan(&n.ID, &p, &lft, &rgt)
	testutil.FatalOnError(t, err)

	if !p.Valid {
		t.Fatal("valid parent expected")
	}
	if rgt-1-lft != 0 {
		t.Fatal("lft or rgt error")
	}
}

func TestPutSecondChild(t *testing.T) {
	db := testutil.GetDB()
	defer db.Close()

	repo := postgres.NewRepository(db)
	rootNode := node.New("")
	rootID, err := repo.Put(rootNode)
	testutil.FatalOnError(t, err)

	firstChildNode := node.New(rootID)
	_, err = repo.Put(firstChildNode)
	testutil.FatalOnError(t, err)

	secondChildNode := node.New(rootID)
	secondID, err := repo.Put(secondChildNode)
	testutil.FatalOnError(t, err)

	var n node.Node
	var p sql.NullString
	var lft, rgt int
	query := "select id, parent, lft, rgt from vinodes.nodes where id = $1"
	err = db.QueryRow(query, secondID).Scan(&n.ID, &p, &lft, &rgt)
	testutil.FatalOnError(t, err)

	if !p.Valid {
		t.Fatal("valid parent expected")
	}
	if rgt-lft-1 != 0 {
		t.Fatal("lft or rgt error")
	}
}

func TestGet(t *testing.T) {
	db := testutil.GetDB()
	defer db.Close()

	repo := postgres.NewRepository(db)
	rootNode := node.New("")
	rootID, err := repo.Put(rootNode)
	testutil.FatalOnError(t, err)

	firstChildNode := node.New(rootID)
	firstID, err := repo.Put(firstChildNode)
	testutil.FatalOnError(t, err)

	secondChildNode := node.New(rootID)
	secondID, err := repo.Put(secondChildNode)
	testutil.FatalOnError(t, err)

	n, err := repo.Get(rootID)
	testutil.FatalOnError(t, err)

	if rootNode.Parent != n.Parent {
		t.Fatal("root properties error")
	}

	ids := []node.ID{firstID, secondID}

	for _, j := range n.Children {
		fail := true
		for _, c := range ids {
			if j == c {
				fail = false
			}
		}
		if fail {
			t.Fatal("children not found")
		}
	}
}
