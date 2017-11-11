// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package mem

import (
	"errors"
	"sync"

	"github.com/vizigoto/vizigoto/authz"
	"github.com/vizigoto/vizigoto/pkg/uuid"
)

type repository struct {
	mtx   sync.RWMutex
	nodes map[string]*auth.Group
}

func NewRepository() auth.Repository {
	return &repository{nodes: make(map[string]*auth.Group)}
}

func (repo *repository) Get(id string) (*auth.Group, error) {
	repo.mtx.RLock()
	defer repo.mtx.RUnlock()

	if i, ok := repo.nodes[id]; ok {
		return i, nil
	}
	return nil, errors.New("node not found")
}

func (repo *repository) Put(g *auth.Group) (string, error) {
	repo.mtx.Lock()
	defer repo.mtx.Unlock()

	g.ID = uuid.New()
	repo.nodes[g.ID] = g
	return g.ID, nil
}
