// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package content

import (
	"github.com/vizigoto/vizigoto/item"
	"github.com/vizigoto/vizigoto/user"
)

//Service is the interface that provides the content methods.
type Service interface {
	GetItem(id string) (interface{}, error)
	AddFolder(name, parent string) (string, error)
	AddReport(name, parent, content string) (string, error)

	GetUser(id string) (user.User, error)
	AddUser(name string) (id string, err error)
}

type service struct {
	items item.Service
	users user.Service
}

// NewContentService returns an instance of an content Service.
func NewContentService(items item.Service, users user.Service) Service {
	return &service{items, users}
}

func (s *service) GetItem(id string) (interface{}, error) {
	return s.items.Get(id)
}

func (s *service) AddFolder(name, parent string) (string, error) {
	return s.items.AddFolder(name, parent)
}

func (s *service) AddReport(name, parent, content string) (string, error) {
	return s.items.AddReport(name, parent, content)
}

func (s *service) GetUser(id string) (user.User, error) {
	return s.users.Get(id)
}

func (s *service) AddUser(name string) (string, error) {
	return s.users.AddUser(name)
}
