// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package postgres_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/vizigoto/vizigoto/node"
	"github.com/vizigoto/vizigoto/node/storage/postgres"
)

func getDB() (*sql.DB, error) {
	hostname := os.Getenv("PGHOSTNAME")
	database := os.Getenv("PGDATABASE")
	username := os.Getenv("PGUSERNAME")
	password := os.Getenv("PGPASSWORD")

	conInfo := fmt.Sprintf("host=%s dbname=%s user=%s password=%s",
		hostname, database, username, password)

	db, err := sql.Open("postgres", conInfo)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("truncate vinodes.nodes")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func TestPutRoot(t *testing.T) {
	db, err := getDB()
	if err != nil {
		t.Fatal(err)
	}

	repo := postgres.NewRepository(db)

	n := node.New("Home", node.Folder, "", "x")

	id, err := repo.Put(n)
	if err != nil {
		t.Fatal(err)
	}

	if id == "" {
		t.Fatal("id not returned")
	}

	var dbNode node.Node

	err = db.QueryRow("select id, name, kind, owner from vinodes.nodes where id = $1", id).Scan(&dbNode.ID, &dbNode.Name, &dbNode.Kind, &dbNode.Owner)

	if err != nil {
		t.Fatal(err)
	}

	if n.ID != dbNode.ID ||
		n.Name != dbNode.Name ||
		n.Kind != dbNode.Kind ||
		n.Owner != dbNode.Owner {
		t.Fatal("propertie error")
	}
}

func TestPutFirstChild(t *testing.T) {
	db, err := getDB()
	if err != nil {
		t.Fatal(err)
	}

	repo := postgres.NewRepository(db)

	rootNode := node.New("Home", node.Folder, "", "x")

	id, err := repo.Put(rootNode)
	if err != nil {
		t.Fatal(err)
	}

	if id == "" {
		t.Fatal("id not returned")
	}

	childNode := node.New("IT", node.Folder, rootNode.ID, "x")

	id, err = repo.Put(childNode)
	if err != nil {
		t.Fatal(err)
	}

	if id == "" {
		t.Fatal("id not returned")
	}

	var dbNode node.Node

	err = db.QueryRow("select id, name, kind, owner from vinodes.nodes where id = $1", id).Scan(&dbNode.ID, &dbNode.Name, &dbNode.Kind, &dbNode.Owner)

	if err != nil {
		t.Fatal(err)
	}

	if childNode.ID != dbNode.ID ||
		childNode.Name != dbNode.Name ||
		childNode.Kind != dbNode.Kind ||
		childNode.Owner != dbNode.Owner {
		t.Fatal("propertie error")
	}
}

func TestPutSecondChild(t *testing.T) {
	db, err := getDB()
	if err != nil {
		t.Fatal(err)
	}

	repo := postgres.NewRepository(db)

	rootNode := node.New("Home", node.Folder, "", "x")

	id, err := repo.Put(rootNode)
	if err != nil {
		t.Fatal(err)
	}

	if id == "" {
		t.Fatal("id not returned")
	}

	firstChildNode := node.New("IT", node.Folder, rootNode.ID, "x")

	id, err = repo.Put(firstChildNode)
	if err != nil {
		t.Fatal(err)
	}

	if id == "" {
		t.Fatal("id not returned")
	}

	secondChildNode := node.New("HR", node.Folder, rootNode.ID, "x")

	id, err = repo.Put(secondChildNode)
	if err != nil {
		t.Fatal(err)
	}

	if id == "" {
		t.Fatal("id not returned")
	}

	var dbNode node.Node

	err = db.QueryRow("select id, name, kind, owner from vinodes.nodes where id = $1", id).Scan(&dbNode.ID, &dbNode.Name, &dbNode.Kind, &dbNode.Owner)

	if err != nil {
		t.Fatal(err)
	}

	if secondChildNode.ID != dbNode.ID ||
		secondChildNode.Name != dbNode.Name ||
		secondChildNode.Kind != dbNode.Kind ||
		secondChildNode.Owner != dbNode.Owner {
		t.Fatal("propertie error")
	}
}

func TestGet(t *testing.T) {
	db, err := getDB()
	if err != nil {
		t.Fatal(err)
	}

	repo := postgres.NewRepository(db)

	rootNode := node.New("Home", node.Folder, "", "x")

	id, err := repo.Put(rootNode)
	if err != nil {
		t.Fatal(err)
	}

	if id == "" {
		t.Fatal("id not returned")
	}

	firstChildNode := node.New("IT", node.Folder, rootNode.ID, "x")

	id, err = repo.Put(firstChildNode)
	if err != nil {
		t.Fatal(err)
	}

	if id == "" {
		t.Fatal("id not returned")
	}

	a, err := repo.Get(rootNode.ID)
	if err != nil {
		t.Fatal(err)
	}

	if a.ID != rootNode.ID ||
		a.Name != rootNode.Name ||
		a.Owner != rootNode.Owner ||
		a.Parent != rootNode.Parent {
		t.Fatal("properties differ")
	}

	a, err = repo.Get(firstChildNode.ID)
	if err != nil {
		t.Fatal(err)
	}

	if a.ID != firstChildNode.ID ||
		a.Name != firstChildNode.Name ||
		a.Owner != firstChildNode.Owner ||
		a.Parent != firstChildNode.Parent {
		t.Fatal("properties differ")
	}
}
