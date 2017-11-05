package node

import (
	"github.com/vizigoto/vizigoto/pkg/user"
	"github.com/vizigoto/vizigoto/pkg/uuid"
)

type ID string
type Kind string

const (
	KindFolder Kind = "folder"
	KindReport Kind = "report"
)

type Node struct {
	ID       ID
	Name     string
	Kind     Kind
	Parent   ID
	Owner    user.ID
	Children []ID
}

var EmptyID = ID("")

func NewID() ID {
	id := uuid.New()
	return ID(id)
}

func NewNode(name string, kind Kind, parent ID, owner user.ID) *Node {
	return &Node{
		Name:     name,
		Kind:     kind,
		Parent:   parent,
		Owner:    owner,
		Children: []ID{},
	}
}

type Repository interface {
	Get(id ID) (*Node, error)
	Put(*Node) (ID, error)
}

type Service interface {
	Get(id ID) (*Node, error)
	Add(name string, kind Kind, parent ID, owner user.ID) (ID, error)
}

type service struct {
	repo   Repository
	rootID ID
}

func NewService(id ID, repo Repository) Service {
	return &service{repo, id}
}

func (s *service) Get(id ID) (*Node, error) {
	return s.repo.Get(id)
}

func (s *service) Add(name string, kind Kind, parent ID, owner user.ID) (ID, error) {
	i := NewNode(name, kind, parent, owner)

	id, err := s.repo.Put(i)
	if err != nil {
		return "", err
	}
	return id, nil
}
