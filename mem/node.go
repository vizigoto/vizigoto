// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package mem

import (
	"context"
	"errors"
	"sync"

	"github.com/vizigoto/vizigoto/node"
	"github.com/vizigoto/vizigoto/pkg/uuid"
)

type nodeRepository struct {
	sync.RWMutex
	nodes map[string]node.Node
}

// NewNodeRepository returns an instance of a node repository using an in-memory storage engine.
// All data will be lost after instance release. Useful for testing purposes.
func NewNodeRepository() node.Repository {
	return &nodeRepository{nodes: make(map[string]node.Node)}
}

func (repo *nodeRepository) Get(ctx context.Context, id string) (node.Node, error) {
	repo.RLock()
	defer repo.RUnlock()
	if i, ok := repo.nodes[id]; ok {
		repo.assembleChildren(i)
		i.Path = repo.path(i.ID)
		return i, nil
	}
	return node.Node{}, errors.New("node not found")
}

func (repo *nodeRepository) Put(ctx context.Context, n node.Node) (string, error) {
	repo.Lock()
	defer repo.Unlock()
	n.ID = uuid.New()
	repo.nodes[n.ID] = n
	return n.ID, nil
}

func (repo *nodeRepository) Move(ctx context.Context, n node.Node, parent string) error {
	repo.Lock()
	defer repo.Unlock()
	n.Parent = parent
	repo.nodes[n.ID] = n
	return nil
}

func (repo *nodeRepository) assembleChildren(n node.Node) {
	n.Children = []string{}
	for _, v := range repo.nodes {
		if v.Parent == n.ID {
			n.Children = append(n.Children, v.ID)
		}
	}
}

func (repo *nodeRepository) path(id string) []node.Node {
	n := repo.nodes[id]
	if n.Parent == "" {
		return []node.Node{n}
	}
	paths := repo.path(n.Parent)
	paths = append(paths, n)
	return paths
}
