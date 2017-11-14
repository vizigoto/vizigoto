// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package item

// Folder represents a folder in the content tree.
type Folder struct {
	ID       string
	Name     string
	Parent   string
	Children []string
}

// NewFolder allocates a folder and returns a pointer to it.
func NewFolder(name string, parent string) *Folder {
	return &Folder{Name: name, Parent: parent, Children: []string{}}
}

// Report represents a report in the content tree.
type Report struct {
	ID      string
	Name    string
	Parent  string
	Content string
}

// NewReport allocates a report and returns a pointer to it.
func NewReport(name, parent, content string) *Report {
	return &Report{Name: name, Parent: parent, Content: content}
}

// Repository provides a limited interface to a storage layer.
type Repository interface {
	Get(string) (interface{}, error)
	Put(interface{}) (string, error)
}

//Service is the interface that provides the basic Item methods.
type Service interface {
	Get(id string) (interface{}, error)
	AddFolder(name, parent string) (string, error)
	AddReport(name, parent, content string) (string, error)
}

type service struct {
	repo Repository
}

//NewService returns a new instance of the default item Service.
func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) Get(id string) (interface{}, error) {
	return s.repo.Get(id)
}

func (s *service) AddFolder(name, parent string) (string, error) {
	folder := NewFolder(name, parent)
	return s.repo.Put(folder)
}

func (s *service) AddReport(name, parent, content string) (string, error) {
	report := NewReport(name, parent, content)
	return s.repo.Put(report)
}
