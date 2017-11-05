package postgres

import (
	"database/sql"
	"errors"

	"github.com/vizigoto/vizigoto/pkg/item"
	"github.com/vizigoto/vizigoto/pkg/node"
)

type repository struct {
	db    *sql.DB
	nodes node.Repository
}

func NewRepository(db *sql.DB, nodes node.Repository) item.Repository {
	return &repository{db, nodes}
}

func (repo *repository) Get(id item.ID) (interface{}, error) {
	n, err := repo.nodes.Get(node.ID(id))
	if err != nil {
		return nil, err
	}

	if n.Kind == node.KindFolder {
		return repo.getFolder(n)
	}

	if n.Kind == node.KindReport {
		return repo.getReport(n)
	}

	return nil, errors.New("item not found")
}

func (repo *repository) getFolder(n *node.Node) (interface{}, error) {
	folder := &item.Folder{}
	folder.ID = item.ID(n.ID)
	folder.Name = n.Name
	folder.Owner = n.Owner
	var id string
	err := repo.db.QueryRow("select id from viitems.folders where id = $1", n.ID).Scan(&id)
	if err != nil {
		return nil, err
	}
	if item.ID(id) != folder.ID {
		return nil, errors.New("internal id error")
	}
	return folder, nil
}

func (repo *repository) getReport(n *node.Node) (interface{}, error) {
	report := &item.Report{}
	report.ID = item.ID(n.ID)
	report.Name = n.Name
	report.Owner = n.Owner
	err := repo.db.QueryRow("select content from viitems.reports where id = $1", n.ID).Scan(&report.Content)
	if err != nil {
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

	panic("type not identified")
}

func (repo *repository) putFolder(folder *item.Folder) (item.ID, error) {
	n := node.NewNode(folder.Name, node.KindFolder, node.ID(folder.Parent), folder.Owner)
	nodeID, err := repo.nodes.Put(n)
	if err != nil {
		return "", err
	}
	_, err = repo.db.Exec("insert into viitems.folders values($1)", nodeID)
	if err != nil {
		return "", err
	}

	return item.ID(nodeID), nil
}

func (repo *repository) putReport(report *item.Report) (item.ID, error) {
	n := node.NewNode(report.Name, node.KindReport, node.ID(report.Parent), report.Owner)
	nodeID, err := repo.nodes.Put(n)
	if err != nil {
		return "", err
	}
	_, err = repo.db.Exec("insert into viitems.reports(id, content) values($1, $2)", nodeID, report.Content)
	if err != nil {
		return "", err
	}

	return item.ID(nodeID), nil
}
