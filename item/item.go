// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package item

import (
	"github.com/vizigoto/vizigoto/user"
)

// ID uniquely identifies a particular item.
type ID string

// Folder represents a folder in the content tree.
type Folder struct {
	ID       ID
	Name     string
	Parent   ID
	Owner    user.ID
	Readme   string
	Children []ID
}

// NewFolder allocates a folder and returns a pointer to it.
func NewFolder(name string, parent ID, owner user.ID) *Folder {
	return &Folder{Name: name, Parent: parent, Owner: owner, Children: []ID{}}
}

// Report represents a report in the content tree.
type Report struct {
	ID      ID
	Name    string
	Parent  ID
	Owner   user.ID
	Content string
}

// NewReport allocates a report and returns a pointer to it.
func NewReport(name string, parent ID, owner user.ID, content string) *Report {
	return &Report{Name: name, Parent: parent, Owner: owner, Content: content}
}

// Repository provides a limited interface to a storage layer.
type Repository interface {
	Get(ID) (interface{}, error)
	Put(interface{}) (ID, error)
}

type Service interface {
	Get(id ID) (interface{}, error)
	AddFolder(name string, parent ID, owner user.ID) (ID, error)
	AddReport(name string, parent ID, owner user.ID, content string) (ID, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) Get(id ID) (interface{}, error) {
	return s.repo.Get(id)
}

func (s *service) AddFolder(name string, parent ID, owner user.ID) (ID, error) {
	folder := NewFolder(name, parent, owner)
	return s.repo.Put(folder)
}

func (s *service) AddReport(name string, parent ID, owner user.ID, content string) (ID, error) {
	report := NewReport(name, parent, owner, content)
	return s.repo.Put(report)
}
