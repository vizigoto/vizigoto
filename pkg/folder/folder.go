package folder

import "github.com/vizigoto/vizigoto/pkg/item"

type Folder struct {
	ID       string
	Name     string
	Parent   string
	Children []string
}

func NewFolder(id, name, parent string) *Folder {
	return &Folder{ID: id, Name: name, Parent: parent, Children: []string{}}
}

type Repository interface {
	Get(id string) (*Folder, error)
	Put(*Folder) error
}

type Service interface {
	GetRoot() (*Folder, error)
	GetFolder(id string) (*Folder, error)
	AddFolder(name, parent string) (string, error)
}

type service struct {
	repo        Repository
	itemService item.Service
}

func NewService(repo Repository, itemService item.Service) Service {
	return &service{repo, itemService}
}

func (s service) GetRoot() (*Folder, error) {
	i, err := s.itemService.GetRoot()
	if err != nil {
		return nil, err
	}
	folder, err := s.repo.Get(i.ID)
	if err != nil {
		return nil, err
	}
	return folder, nil
}

func (s service) GetFolder(id string) (*Folder, error) {
	i, err := s.itemService.GetItem(id)
	if err != nil {
		return nil, err
	}
	folder, err := s.repo.Get(i.ID)
	if err != nil {
		return nil, err
	}
	return folder, nil
}

func (s service) AddFolder(name, parent string) (string, error) {
	id, err := s.itemService.AddItem(parent)
	if err != nil {
		return "", err
	}
	folder := NewFolder(id, name, parent)
	err = s.repo.Put(folder)
	if err != nil {
		return "", err
	}
	return id, nil
}
