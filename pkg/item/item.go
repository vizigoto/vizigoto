package item

import (
	"github.com/vizigoto/vizigoto/pkg/user"
)

type ID string

var EmptyID = ID("")

type Folder struct {
	ID       ID
	Name     string
	Parent   ID
	Owner    user.ID
	Readme   string
	Children []ID
}

func NewFolder(name string, parent ID, owner user.ID) *Folder {
	return &Folder{Name: name, Parent: parent, Owner: owner, Children: []ID{}}
}

type Report struct {
	ID      ID
	Name    string
	Parent  ID
	Owner   user.ID
	Content string
}

func NewReport(name string, parent ID, owner user.ID, content string) *Report {
	return &Report{Name: name, Parent: parent, Owner: owner, Content: content}
}

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
