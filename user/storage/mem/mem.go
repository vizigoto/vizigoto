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
	sync.RWMutex
	users map[string]*user.User
}

// NewRepository returns an instance of a user repository using an in-memory storage engine.
// All data will be lost after instance release. Useful for testing purposes.
func NewRepository() user.Repository {
	return &repository{users: make(map[string]*user.User)}
}

func (repo *repository) Get(id string) (*user.User, error) {
	repo.RLock()
	defer repo.RUnlock()
	if i, ok := repo.users[id]; ok {
		return i, nil
	}
	return nil, errors.New("user not found")
}

func (repo *repository) Put(i *user.User) (string, error) {
	repo.Lock()
	defer repo.Unlock()
	i.ID = uuid.New()
	repo.users[i.ID] = i
	return i.ID, nil
}
