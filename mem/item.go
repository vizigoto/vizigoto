// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package mem

import (
	"context"
	"sync"

	"github.com/pkg/errors"

	"github.com/vizigoto/vizigoto/item"
	"github.com/vizigoto/vizigoto/node"
)

type itemRepository struct {
	sync.RWMutex
	items map[string]interface{}
	nodes node.Repository
}

// NewItemRepository returns an instance of a item repository.
func NewItemRepository(nodes node.Repository) item.Repository {
	return &itemRepository{items: make(map[string]interface{}), nodes: nodes}
}

func (repo *itemRepository) Get(ctx context.Context, id string) (interface{}, error) {
	repo.RLock()
	defer repo.RUnlock()

	n, err := repo.nodes.Get(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "could not find the node")
	}

	path := repo.assemblePath(n)

	if i, ok := repo.items[id]; ok {
		if folder, ok := i.(item.Folder); ok {
			for _, c := range n.Children {
				folder.Children = append(folder.Children, c)
			}
			folder.Path = path
			return folder, nil
		}
		if report, ok := i.(item.Report); ok {
			report.Path = path
			return report, nil
		}
	}
	return nil, errors.New("item not found")
}

func (repo *itemRepository) Put(ctx context.Context, i interface{}) (string, error) {
	repo.Lock()
	defer repo.Unlock()

	folder, ok := i.(item.Folder)
	if ok {
		n := node.New(folder.Parent)
		id, err := repo.nodes.Put(ctx, n)
		if err != nil {
			return "", errors.Wrap(err, "could not put the node")
		}
		folder.ID = id
		repo.items[id] = folder
		return id, nil
	}

	report, ok := i.(item.Report)
	if ok {
		n := node.New(report.Parent)
		id, err := repo.nodes.Put(ctx, n)
		if err != nil {
			return "", errors.Wrap(err, "could not put the node")
		}
		report.ID = id
		repo.items[id] = report
		return id, nil
	}

	panic("type error")
}

func (repo *itemRepository) assemblePath(n node.Node) []item.Path {
	paths := []item.Path{}
	for _, v := range n.Path {
		if i, ok := repo.items[v.ID]; ok {
			switch el := i.(type) {
			case item.Folder:
				paths = append(paths, el)
			case item.Report:
				paths = append(paths, el)
			}
		}
	}
	return paths
}
