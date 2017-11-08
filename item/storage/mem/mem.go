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
	items map[item.ID]interface{}
	nodes node.Repository
}

// NewRepository returns an instance of a item repository.
func NewRepository(nodes node.Repository) item.Repository {
	return &repository{items: make(map[item.ID]interface{}), nodes: nodes}
}

func (repo *repository) Get(id item.ID) (interface{}, error) {
	repo.mtx.RLock()
	defer repo.mtx.RUnlock()

	n, err := repo.nodes.Get(node.ID(id))
	if err != nil {
		return nil, err
	}

	if i, ok := repo.items[id]; ok {
		if folder, ok := i.(*item.Folder); ok {
			for _, c := range n.Children {
				folder.Children = append(folder.Children, item.ID(c))
			}
			return folder, nil
		}
		if report, ok := i.(*item.Report); ok {
			return report, nil
		}
	}
	return nil, errors.New("item not found")
}

func (repo *repository) Put(i interface{}) (item.ID, error) {
	repo.mtx.Lock()
	defer repo.mtx.Unlock()

	folder, ok := i.(*item.Folder)
	if ok {
		n := node.New(node.Folder, node.ID(folder.Parent))
		id, err := repo.nodes.Put(n)
		if err != nil {
			return "", err
		}
		repo.items[item.ID(id)] = folder
		return item.ID(id), nil
	}

	report, ok := i.(*item.Report)
	if ok {
		n := node.New(node.Report, node.ID(report.Parent))
		id, err := repo.nodes.Put(n)
		if err != nil {
			return "", err
		}
		repo.items[item.ID(id)] = report
		return item.ID(id), nil
	}

	panic("type error")
}
