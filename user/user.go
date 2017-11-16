// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package user

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
	Get(id string) (user User, err error)
	Put(user User) (id string, err error)
}

//Service is the interface that provides the basic User methods.
type Service interface {
	Get(id string) (user User, err error)
	AddUser(name string) (id string, err error)
}

type service struct {
	repo Repository
}

//NewService returns a new instance of the default user Service.
func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) Get(id string) (User, error) {
	return s.repo.Get(id)
}

func (s *service) AddUser(name string) (string, error) {
	u := New(name)
	return s.repo.Put(u)
}
