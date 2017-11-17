// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package content

import (
	"context"

	"github.com/vizigoto/vizigoto/item"
	"github.com/vizigoto/vizigoto/user"
)

//Service is the interface that provides the content methods.
type Service interface {
	GetHome(ctx context.Context) (i item.Folder, err error)
	GetItem(ctx context.Context, id string) (i interface{}, err error)
	AddFolder(ctx context.Context, name, parent string) (id string, err error)
	AddReport(ctx context.Context, name, parent, content string) (id string, err error)

	GetUser(ctx context.Context, id string) (u user.User, err error)
	AddUser(ctx context.Context, name string) (id string, err error)
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

func (s *service) GetHome(ctx context.Context) (item.Folder, error) {
	i, err := s.items.Get(ctx, s.homeID)
	if err != nil {
		return item.Folder{}, nil
	}
	if h, ok := i.(item.Folder); ok {
		return h, nil
	}
	panic("home is not a folder")
}

func (s *service) GetItem(ctx context.Context, id string) (interface{}, error) {
	return s.items.Get(ctx, id)
}

func (s *service) AddFolder(ctx context.Context, name, parent string) (string, error) {
	return s.items.AddFolder(ctx, name, parent)
}

func (s *service) AddReport(ctx context.Context, name, parent, content string) (string, error) {
	return s.items.AddReport(ctx, name, parent, content)
}

func (s *service) GetUser(ctx context.Context, id string) (user.User, error) {
	return s.users.Get(id)
}

func (s *service) AddUser(ctx context.Context, name string) (string, error) {
	return s.users.AddUser(name)
}
