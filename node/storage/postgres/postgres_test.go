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
	repo := postgres.NewRepository(db)
	root := node.New(node.Folder, "")
	rootID, err := repo.Put(root)
	if err != nil {
		t.Fatal(err)
	}
	childNode := node.New(node.Folder, rootID)
	childID, err := repo.Put(childNode)
	if err != nil {
		t.Fatal(err)
	}

	c, err := repo.Get(childID)
	if err != nil {
		t.Fatal(err)
	}

	if c.ID != childID ||
		c.Kind != childNode.Kind ||
		c.Parent != childNode.Parent {
		t.Fatal("child error")
	}
}

func TestPutRoot(t *testing.T) {
	db := testutil.GetDB()
	repo := postgres.NewRepository(db)
	root := node.New(node.Folder, "")
	id, err := repo.Put(root)
	if err != nil {
		t.Fatal(err)
	}

	var n node.Node
	var p sql.NullString
	var lft, rgt int
	query := "select id, parent, lft, rgt, kind from vinodes.nodes where id = $1"
	err = db.QueryRow(query, id).Scan(&n.ID, &p, &lft, &rgt, &n.Kind)

	if err != nil {
		t.Fatal(err)
	}
	if root.Kind != n.Kind {
		t.Fatal("kind error")
	}
	if p.Valid {
		t.Fatal("null expected")
	}
	if rgt-1-lft != 0 {
		t.Fatal("lft or rgt error")
	}
}

func TestPutFirstChild(t *testing.T) {
	db := testutil.GetDB()
	repo := postgres.NewRepository(db)
	rootNode := node.New(node.Folder, "")
	rootID, err := repo.Put(rootNode)
	if err != nil {
		t.Fatal(err)
	}
	childNode := node.New(node.Folder, rootID)
	childID, err := repo.Put(childNode)
	if err != nil {
		t.Fatal(err)
	}

	var n node.Node
	var p sql.NullString
	var lft, rgt int
	query := "select id, parent, lft, rgt, kind from vinodes.nodes where id = $1"
	err = db.QueryRow(query, childID).Scan(&n.ID, &p, &lft, &rgt, &n.Kind)

	if err != nil {
		t.Fatal(err)
	}
	if childNode.Kind != n.Kind {
		t.Fatal("kind error")
	}
	if !p.Valid {
		t.Fatal("valid parent expected")
	}
	if rgt-1-lft != 0 {
		t.Fatal("lft or rgt error")
	}
}

func TestPutSecondChild(t *testing.T) {
	db := testutil.GetDB()
	repo := postgres.NewRepository(db)
	rootNode := node.New(node.Folder, "")
	rootID, err := repo.Put(rootNode)
	if err != nil {
		t.Fatal(err)
	}
	firstChildNode := node.New(node.Folder, rootID)
	_, err = repo.Put(firstChildNode)
	if err != nil {
		t.Fatal(err)
	}
	secondChildNode := node.New(node.Folder, rootID)
	secondID, err := repo.Put(secondChildNode)
	if err != nil {
		t.Fatal(err)
	}

	var n node.Node
	var p sql.NullString
	var lft, rgt int
	query := "select id, parent, lft, rgt, kind from vinodes.nodes where id = $1"
	err = db.QueryRow(query, secondID).Scan(&n.ID, &p, &lft, &rgt, &n.Kind)

	if err != nil {
		t.Fatal(err)
	}
	if secondChildNode.Kind != n.Kind {
		t.Fatal("kind error")
	}
	if !p.Valid {
		t.Fatal("valid parent expected")
	}
	if rgt-lft-1 != 0 {
		t.Fatal("lft or rgt error")
	}
}

func TestGet(t *testing.T) {
	db := testutil.GetDB()
	repo := postgres.NewRepository(db)
	rootNode := node.New(node.Folder, "")
	rootID, err := repo.Put(rootNode)
	if err != nil {
		t.Fatal(err)
	}
	firstChildNode := node.New(node.Folder, rootID)
	firstID, err := repo.Put(firstChildNode)
	if err != nil {
		t.Fatal(err)
	}
	secondChildNode := node.New(node.Folder, rootID)
	secondID, err := repo.Put(secondChildNode)
	if err != nil {
		t.Fatal(err)
	}

	n, err := repo.Get(rootID)
	if err != nil {
		t.Fatal(err)
	}

	if rootNode.Parent != n.Parent ||
		rootNode.Kind != n.Kind {
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
