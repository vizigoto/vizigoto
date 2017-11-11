// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package auth

type AuthServiceManager interface {
	Get(id string) (g *Group, err error)
	AddGroup(name, parent string) (id string, err error)
}

type authServiceManager struct {
	repo Repository
}

func NewAuthServiceManager(repo Repository) AuthServiceManager {
	return &authServiceManager{repo}
}

func (s *authServiceManager) Get(id string) (g *Group, err error) {
	return s.repo.Get(id)
}

func (s *authServiceManager) AddGroup(name, parent string) (id string, err error) {
	i := NewGroup(name, parent)
	return s.repo.Put(i)
}
