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
	GetHome() (item.Folder, error)
	GetItem(id string) (interface{}, error)
	AddFolder(name, parent string) (string, error)
	AddReport(name, parent, content string) (string, error)

	GetUser(id string) (user.User, error)
	AddUser(name string) (id string, err error)
}

type service struct {
	homeID string
	items  item.Service
	users  user.Service
}

// NewService returns an instance of an content Service.
func NewService(homeID string, items item.Service, users user.Service) Service {
	return &service{homeID, items, users}
}

func (s *service) GetHome() (item.Folder, error) {
	i, err := s.items.Get(s.homeID)
	if err != nil {
		return item.Folder{}, nil
	}
	if h, ok := i.(item.Folder); ok {
		return h, nil
	}
	panic("home is not a folder")
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
