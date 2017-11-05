package report

import "github.com/vizigoto/vizigoto/pkg/item"

type Report struct {
	ID      string
	Name    string
	Parent  string
	Content string
}

func NewReport(id, name, parent, content string) *Report {
	return &Report{ID: id, Name: name, Parent: parent, Content: content}
}

type Repository interface {
	Get(id string) (*Report, error)
	Put(*Report) error
}

type Service interface {
	GetReport(id string) (*Report, error)
	AddReport(name, parent, content string) (string, error)
}

type service struct {
	repo        Repository
	itemService item.Service
}

func NewService(repo Repository, itemService item.Service) Service {
	return &service{repo, itemService}
}

func (s service) GetReport(id string) (*Report, error) {
	i, err := s.itemService.GetItem(id)
	if err != nil {
		return nil, err
	}
	report, err := s.repo.Get(i.ID)
	if err != nil {
		return nil, err
	}
	return report, nil
}

func (s service) AddReport(name, parent, content string) (string, error) {
	id, err := s.itemService.AddItem(parent)
	if err != nil {
		return "", err
	}
	report := NewReport(id, name, parent, content)
	err = s.repo.Put(report)
	if err != nil {
		return "", err
	}
	return id, nil
}
