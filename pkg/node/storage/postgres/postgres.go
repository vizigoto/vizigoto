package postgres

import (
	"database/sql"

	"github.com/vizigoto/vizigoto/pkg/node"
	"github.com/vizigoto/vizigoto/pkg/uuid"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) node.Repository {
	return &repository{db}
}

func (repo *repository) Get(id node.ID) (*node.Node, error) {
	n := &node.Node{}
	n.ID = id

	var p sql.NullString
	var kind string
	err := repo.db.QueryRow("select name, kind, owner, parent from vinodes.nodes where id = $1", id).Scan(&n.Name, &kind, &n.Owner, &p)
	if err != nil {
		return nil, err
	}
	if p.Valid {
		n.Parent = node.ID(p.String)
	}

	if kind == "folder" {
		n.Kind = node.KindFolder
	}

	if kind == "report" {
		n.Kind = node.KindReport
	}

	return n, nil
}

func (repo *repository) Put(n *node.Node) (node.ID, error) {
	id := uuid.New()
	n.ID = node.ID(id)
	if n.Parent == node.EmptyID {
		return repo.putRoot(n)
	}
	return repo.putChild(n)
}

func (repo *repository) putRoot(n *node.Node) (node.ID, error) {
	_, err := repo.db.Exec("insert into vinodes.nodes(id, parent, name, kind, owner, protected) values ($1, $2, $3, $4, $5, $6)",
		n.ID, nil, n.Name, n.Kind, n.Owner, true)
	if err != nil {
		return "", err
	}
	return n.ID, nil
}

func (repo *repository) putChild(n *node.Node) (node.ID, error) {
	var lft, rgt int
	repo.db.QueryRow("select lft, rgt from vinodes.nodes where id = $1", n.Parent).Scan(&lft, &rgt)
	if lft-(rgt-1) == 0 {
		return repo.putFirstChild(n, lft)
	}
	return repo.putSecondChild(n, rgt)
}

func (repo *repository) putFirstChild(n *node.Node, lft int) (node.ID, error) {
	_, err := repo.db.Exec("update vinodes.nodes set rgt = rgt + 2 where rgt > $1", lft)
	if err != nil {
		return "", err
	}
	_, err = repo.db.Exec("update vinodes.nodes set lft = lft + 2 where lft > $1", lft)
	if err != nil {
		return "", err
	}

	_, err = repo.db.Exec("insert into vinodes.nodes(id, parent, lft, rgt, name, kind, owner, protected) values ($1, $2, $3, $4, $5, $6, $7, $8)",
		n.ID, n.Parent, lft+1, lft+2, n.Name, n.Kind, n.Owner, true)
	if err != nil {
		return "", err
	}

	return n.ID, nil
}

func (repo *repository) putSecondChild(n *node.Node, lft int) (node.ID, error) {
	_, err := repo.db.Exec("update vinodes.nodes set rgt = rgt + 2 where rgt >= $1", lft)
	if err != nil {
		return "", err
	}
	_, err = repo.db.Exec("update vinodes.nodes set lft = lft + 2 where lft >= $1", lft)
	if err != nil {
		return "", err
	}

	_, err = repo.db.Exec("insert into vinodes.nodes(id, parent, lft, rgt, name, kind, owner, protected) values ($1, $2, $3, $4, $5, $6, $7, $8)",
		n.ID, n.Parent, lft, lft+1, n.Name, n.Kind, n.Owner, true)
	if err != nil {
		return "", err
	}
	return n.ID, nil
}
