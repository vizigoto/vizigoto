// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package mem

import (
	"errors"
	"sync"

	"github.com/vizigoto/vizigoto/pkg/uuid"
	"github.com/vizigoto/vizigoto/user"
)

type repository struct {
	mtx   sync.RWMutex
	users map[user.ID]*user.User
}

// NewRepository returns an instance of a user repository using an in-memory storage engine.
// All data will be lost after instance release. Useful for testing purposes.
func NewRepository() user.Repository {
	return &repository{users: make(map[user.ID]*user.User)}
}

func (repo *repository) Get(id user.ID) (*user.User, error) {
	repo.mtx.RLock()
	defer repo.mtx.RUnlock()
	if i, ok := repo.users[id]; ok {
		return i, nil
	}
	return nil, errors.New("user not found")
}

func (repo *repository) Put(i *user.User) (user.ID, error) {
	repo.mtx.Lock()
	defer repo.mtx.Unlock()
	id := uuid.New()
	i.ID = user.ID(id)
	repo.users[i.ID] = i
	return i.ID, nil
}
