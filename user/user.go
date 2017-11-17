// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package user

import "context"

// User represents one single user.
type User struct {
	ID   string
	Name string
}

// New allocates an user and returns a pointer to it.
func New(name string) User {
	return User{Name: name}
}

// Repository provides a limited interface to a user storage layer.
type Repository interface {
	Get(ctx context.Context, id string) (u User, err error)
	Put(ctx context.Context, user User) (id string, err error)
}

//Service is the interface that provides the basic User methods.
type Service interface {
	Get(ctx context.Context, id string) (u User, err error)
	AddUser(ctx context.Context, name string) (id string, err error)
}

type service struct {
	repo Repository
}

//NewService returns a new instance of the default user Service.
func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) Get(ctx context.Context, id string) (User, error) {
	return s.repo.Get(ctx, id)
}

func (s *service) AddUser(ctx context.Context, name string) (string, error) {
	u := New(name)
	return s.repo.Put(ctx, u)
}
