// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package postgres

import (
	"database/sql"

	"github.com/vizigoto/vizigoto/node"
	"github.com/vizigoto/vizigoto/pkg/uuid"
)

type repository struct {
	db *sql.DB
}

// NewRepository returns an instance of a node repository.
func NewRepository(db *sql.DB) node.Repository {
	return &repository{db}
}

func (repo *repository) Get(id string) (*node.Node, error) {
	n := &node.Node{ID: id, Children: []string{}}

	rows, err := repo.db.Query(sqlGet, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var r struct {
			parent  sql.NullString
			childID sql.NullString
		}
		err = rows.Scan(&r.parent, &r.childID)
		if err != nil {
			return nil, err
		}
		if r.parent.Valid {
			n.Parent = r.parent.String
		}
		if r.childID.Valid {
			n.Children = append(n.Children, r.childID.String)
		}
	}

	path, err := repo.path(id)
	if err != nil {
		return nil, err
	}

	n.Path = path

	return n, nil
}

func (repo *repository) Put(n *node.Node) (string, error) {
	if n.Parent == "" {
		return repo.putRoot(n)
	}
	return repo.putChild(n)
}

func (repo *repository) putRoot(n *node.Node) (id string, err error) {
	id = uuid.New()
	if _, err = repo.db.Exec(sqlInsert, id, nil, 0, 1); err != nil {
		return "", err
	}
	return
}

func (repo *repository) putChild(n *node.Node) (id string, err error) {
	var lft, rgt int
	if err = repo.db.QueryRow(sqlPos, n.Parent).Scan(&lft, &rgt); err != nil {
		return
	}
	if rgt-1-lft == 0 {
		return repo.putFirstChild(n, lft)
	}
	return repo.putSecondChild(n, rgt)
}

func (repo *repository) putFirstChild(n *node.Node, lft int) (id string, err error) {
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
	if _, err = tx.Exec("update vinodes.nodes set rgt = rgt + 2 where rgt > $1", lft); err != nil {
		return
	}
	if _, err = tx.Exec("update vinodes.nodes set lft = lft + 2 where lft > $1", lft); err != nil {
		return
	}
	id = uuid.New()
	if _, err = tx.Exec(sqlInsert, id, n.Parent, lft+1, lft+2); err != nil {
		return "", err
	}
	return
}

func (repo *repository) putSecondChild(n *node.Node, lft int) (id string, err error) {
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
	if _, err = repo.db.Exec("update vinodes.nodes set rgt = rgt + 2 where rgt >= $1", lft); err != nil {
		return
	}
	if _, err = repo.db.Exec("update vinodes.nodes set lft = lft + 2 where lft >= $1", lft); err != nil {
		return
	}
	id = uuid.New()
	if _, err = repo.db.Exec(sqlInsert, id, n.Parent, lft, lft+1); err != nil {
		return "", err
	}
	return
}

func (repo *repository) path(id string) ([]string, error) {
	rows, err := repo.db.Query(sqlPath, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	paths := []string{}

	for rows.Next() {
		var id string
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}

		paths = append(paths, id)
	}

	return paths, nil
}
