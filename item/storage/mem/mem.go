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

func NewRepository(nodes node.Repository) item.Repository {
	return &repository{items: make(map[item.ID]interface{}), nodes: nodes}
}

func (repo *repository) Get(id item.ID) (interface{}, error) {
	repo.mtx.RLock()
	defer repo.mtx.RUnlock()
	if i, ok := repo.items[id]; ok {
		if folder, ok := i.(*item.Folder); ok {
			return folder, nil
		}
	}
	return nil, errors.New("item not found")
}

func (repo *repository) Put(i interface{}) (item.ID, error) {
	repo.mtx.Lock()
	defer repo.mtx.Unlock()

	folder, ok := i.(*item.Folder)

	if ok {
		n := node.New(folder.Name, node.Folder, node.ID(folder.Parent), folder.Owner)
		id, err := repo.nodes.Put(n)
		if err != nil {
			return "", err
		}
		folder.ID = item.ID(id)
	}

	repo.items[folder.ID] = folder
	return folder.ID, nil
}
