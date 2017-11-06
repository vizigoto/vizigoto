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
	"github.com/vizigoto/vizigoto/item"
	"github.com/vizigoto/vizigoto/item/storage/postgres"
	node "github.com/vizigoto/vizigoto/node/storage/postgres"
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
	_, err = db.Exec("truncate viitems.folders")
	_, err = db.Exec("truncate viitems.reports")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func TestPutGetFolder(t *testing.T) {
	db, err := getDB()
	if err != nil {
		t.Fatal(err)
	}

	nodes := node.NewRepository(db)
	var repo item.Repository = postgres.NewRepository(db, nodes)
	root := item.NewFolder("Home", "", "")
	id, err := repo.Put(root)
	if err != nil {
		t.Fatal(err)
	}

	if id == "" {
		t.Fatal("id error")
	}

	f, err := repo.Get(id)
	if err != nil {
		t.Fatal(err)
	}

	folder, ok := f.(*item.Folder)
	if !ok {
		t.Fatal("type match fail")
	}

	if folder.Name != "Home" {
		t.Fatal("name error")
	}
}

func TestPutGetReport(t *testing.T) {
	db, err := getDB()
	if err != nil {
		t.Fatal(err)
	}

	nodes := node.NewRepository(db)
	var repo item.Repository = postgres.NewRepository(db, nodes)
	root := item.NewFolder("Home", "", "")
	_, err = repo.Put(root)
	if err != nil {
		t.Fatal(err)
	}

	r := item.NewReport("hr report", root.ID, "", "content")
	id, err := repo.Put(r)
	if err != nil {
		t.Fatal(err)
	}

	if id == "" {
		t.Fatal("id error")
	}

	result, err := repo.Get(id)
	if err != nil {
		t.Fatal(err)
	}

	report, ok := result.(*item.Report)
	if !ok {
		t.Fatal("type error")
	}

	t.Log(r)
	t.Log(report)

	if r.Name != report.Name ||
		r.Owner != report.Owner ||
		r.Parent != report.Parent ||
		r.Content != report.Content {
		t.Fatal("properties error")
	}

}
