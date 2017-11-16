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
func New(name string) *User {
	return &User{Name: name}
}

// Repository provides a limited interface to a user storage layer.
type Repository interface {
	Get(id string) (user *User, err error)
	Put(user *User) (id string, err error)
}
