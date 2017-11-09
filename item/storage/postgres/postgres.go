// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package postgres

import (
	"database/sql"
	"errors"

	"github.com/vizigoto/vizigoto/item"
	"github.com/vizigoto/vizigoto/node"
)

type repository struct {
	db    *sql.DB
	nodes node.Repository
}

// NewRepository returns an instance of a item repository.
func NewRepository(db *sql.DB, nodes node.Repository) item.Repository {
	return &repository{db, nodes}
}

func (repo *repository) Get(id item.ID) (interface{}, error) {
	n, err := repo.nodes.Get(node.ID(id))
	if err != nil {
		return nil, err
	}

	kind, err := repo.getItemKind(id)
	if err != nil {
		return nil, err
	}

	if item.Kind(kind) == item.KindFolder {
		return repo.getFolder(n)
	}
	if item.Kind(kind) == item.KindReport {
		return repo.getReport(n)
	}
	return nil, errors.New("item not found")
}

func (repo *repository) getItemKind(id item.ID) (kind string, err error) {
	err = repo.db.QueryRow("select kind from viitems.items where id = $1", id).Scan(&kind)
	if err != nil {
		return "", err
	}
	return
}

func (repo *repository) getFolder(n *node.Node) (interface{}, error) {
	folder := &item.Folder{}
	folder.ID = item.ID(n.ID)
	folder.Parent = item.ID(n.Parent)

	err := repo.db.QueryRow("select name from viitems.items where id = $1", n.ID).Scan(&folder.Name)
	if err != nil {
		return nil, err
	}

	return folder, nil
}

func (repo *repository) getReport(n *node.Node) (interface{}, error) {
	report := &item.Report{}
	report.ID = item.ID(n.ID)
	report.Parent = item.ID(n.Parent)

	if err := repo.db.QueryRow("select name from viitems.items where id = $1", n.ID).Scan(&report.Name); err != nil {
		return nil, err
	}
	if err := repo.db.QueryRow("select content from viitems.reports where id = $1", n.ID).Scan(&report.Content); err != nil {
		return nil, err
	}
	return report, nil
}

func (repo *repository) Put(i interface{}) (item.ID, error) {
	if folder, ok := i.(*item.Folder); ok {
		return repo.putFolder(folder)
	}
	if report, ok := i.(*item.Report); ok {
		return repo.putReport(report)
	}
	return "", nil
}

func (repo *repository) putFolder(folder *item.Folder) (id item.ID, err error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	n := node.New(node.ID(folder.Parent))
	nodeID, err := repo.nodes.Put(n)
	if err != nil {
		return
	}
	_, err = tx.Exec("insert into viitems.items (id, name, kind) values($1, $2, $3)", nodeID, folder.Name, item.KindFolder)
	if err != nil {
		return
	}
	_, err = tx.Exec("insert into viitems.folders (id) values($1)", nodeID)
	if err != nil {
		return
	}
	return item.ID(nodeID), nil
}

func (repo *repository) putReport(report *item.Report) (id item.ID, err error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	n := node.New(node.ID(report.Parent))
	nodeID, err := repo.nodes.Put(n)
	if err != nil {
		return
	}
	_, err = tx.Exec("insert into viitems.items(id, name, kind) values($1, $2, $3)", nodeID, report.Name, item.KindReport)
	if err != nil {
		return
	}
	_, err = tx.Exec("insert into viitems.reports(id, content) values($1, $2)", nodeID, report.Content)
	if err != nil {
		return
	}
	return item.ID(nodeID), nil
}
