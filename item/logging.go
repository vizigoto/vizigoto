package item

import (
	"time"

	"github.com/vizigoto/vizigoto/pkg/log"
)

type loggingRepository struct {
	logger log.Logger
	Repository
}

// NewLoggingRepository returns a new instance of a logging Repository.
func NewLoggingRepository(logger log.Logger, r Repository) Repository {
	return &loggingRepository{logger, r}
}

func (repo *loggingRepository) Get(id ID) (i interface{}, err error) {
	defer func(begin time.Time) {
		repo.logger.Log("method", "get", "id", id)
	}(time.Now())

	i, err = repo.Repository.Get(id)
	return
}

func (repo *loggingRepository) Put(i interface{}) (id ID, err error) {
	defer func(begin time.Time) {
		repo.logger.Log("method", "put", "id", id)
	}(time.Now())

	id, err = repo.Repository.Put(i)
	return
}

type loggingService struct {
	logger log.Logger
	Service
}

// NewLoggingService returns a new instance of a logging Service.
func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (s *loggingService) Get(id ID) (interface{}, error) {
	defer func(begin time.Time) {
		s.logger.Log("method", "get", "id", id)
	}(time.Now())

	return s.Service.Get(id)
}

func (s *loggingService) AddFolder(name string, parent ID) (ID, error) {
	defer func(begin time.Time) {
		s.logger.Log("addfolder", "get", "name", name, "parent", parent)
	}(time.Now())

	return s.Service.AddFolder(name, parent)
}

func (s *loggingService) AddReport(name string, parent ID, content string) (ID, error) {
	defer func(begin time.Time) {
		s.logger.Log("addreport", "get", "name", name, "parent", parent)
	}(time.Now())

	return s.Service.AddReport(name, parent, content)
}
