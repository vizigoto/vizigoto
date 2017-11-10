// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package user

// ID uniquely identifies a particular user.
type ID string

// User represents one single user.
type User struct {
	ID   ID
	Name string
}

// New allocates an user and returns a pointer to it.
func New(name string) *User {
	return &User{Name: name}
}

// Repository provides a limited interface to a storage layer.
type Repository interface {
	Get(ID) (*User, error)
	Put(*User) (ID, error)
}
