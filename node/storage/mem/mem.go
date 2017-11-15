// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package mem

import (
	"errors"
	"sync"

	"github.com/vizigoto/vizigoto/node"
	"github.com/vizigoto/vizigoto/pkg/uuid"
)

type repository struct {
	mtx   sync.RWMutex
	nodes map[string]*node.Node
}

// NewRepository returns an instance of a node repository using an in-memory storage engine.
// All data will be lost after instance release. Useful for testing purposes.
func NewRepository() node.Repository {
	return &repository{nodes: make(map[string]*node.Node)}
}

func (repo *repository) Get(id string) (*node.Node, error) {
	repo.mtx.RLock()
	defer repo.mtx.RUnlock()
	if i, ok := repo.nodes[id]; ok {
		repo.assembleChildren(i)
		i.Path = repo.path(i.ID)
		return i, nil
	}
	return nil, errors.New("node not found")
}

func (repo *repository) Put(n *node.Node) (string, error) {
	repo.mtx.Lock()
	defer repo.mtx.Unlock()
	n.ID = uuid.New()
	repo.nodes[n.ID] = n
	return n.ID, nil
}

func (repo *repository) assembleChildren(n *node.Node) {
	n.Children = []string{}
	for _, v := range repo.nodes {
		if v.Parent == n.ID {
			n.Children = append(n.Children, v.ID)
		}
	}
}

func (repo *repository) path(id string) []string {
	n := repo.nodes[id]
	if n.Parent == "" {
		return []string{id}
	}
	paths := repo.path(n.Parent)
	paths = append(paths, id)
	return paths
}
