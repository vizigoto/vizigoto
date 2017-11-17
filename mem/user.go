// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package mem

import (
	"sync"

	"github.com/pkg/errors"

	"github.com/vizigoto/vizigoto/pkg/uuid"
	"github.com/vizigoto/vizigoto/user"
)

type repository struct {
	sync.RWMutex
	users map[string]user.User
}

// NewUserRepository returns an instance of a user repository using an in-memory storage engine.
// All data will be lost after instance release. Useful for testing purposes.
func NewUserRepository() user.Repository {
	return &repository{users: make(map[string]user.User)}
}

func (repo *repository) Get(id string) (user.User, error) {
	repo.RLock()
	defer repo.RUnlock()
	if u, ok := repo.users[id]; ok {
		return u, nil
	}
	return user.User{}, errors.New("user not found")
}

func (repo *repository) Put(u user.User) (string, error) {
	repo.Lock()
	defer repo.Unlock()
	u.ID = uuid.New()
	repo.users[u.ID] = u
	return u.ID, nil
}
