// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package mem

import (
	"errors"
	"sync"

	"github.com/vizigoto/vizigoto/item"
	"github.com/vizigoto/vizigoto/node"
)

type repository struct {
	mtx   sync.RWMutex
	items map[string]interface{}
	nodes node.Repository
}

// NewRepository returns an instance of a item repository.
func NewRepository(nodes node.Repository) item.Repository {
	return &repository{items: make(map[string]interface{}), nodes: nodes}
}

func (repo *repository) Get(id string) (interface{}, error) {
	repo.mtx.RLock()
	defer repo.mtx.RUnlock()

	n, err := repo.nodes.Get(id)
	if err != nil {
		return nil, err
	}

	if i, ok := repo.items[id]; ok {
		if folder, ok := i.(*item.Folder); ok {
			for _, c := range n.Children {
				folder.Children = append(folder.Children, c)
			}
			return folder, nil
		}
		if report, ok := i.(*item.Report); ok {
			return report, nil
		}
	}
	return nil, errors.New("item not found")
}

func (repo *repository) Put(i interface{}) (string, error) {
	repo.mtx.Lock()
	defer repo.mtx.Unlock()

	folder, ok := i.(*item.Folder)
	if ok {
		n := node.New(folder.Parent)
		id, err := repo.nodes.Put(n)
		if err != nil {
			return "", err
		}
		folder.ID = id
		repo.items[id] = folder
		return id, nil
	}

	report, ok := i.(*item.Report)
	if ok {
		n := node.New(report.Parent)
		id, err := repo.nodes.Put(n)
		if err != nil {
			return "", err
		}
		report.ID = id
		repo.items[id] = report
		return id, nil
	}

	panic("type error")
}
