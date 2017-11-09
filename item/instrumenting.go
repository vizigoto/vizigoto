package item

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type instrumentingRepository struct {
	counter *prometheus.CounterVec
	Repository
}

// NewInstrumentingRepository returns an instance of an instrumenting Repository.
func NewInstrumentingRepository(c *prometheus.CounterVec, r Repository) Repository {
	return &instrumentingRepository{c, r}
}

func (repo *instrumentingRepository) Get(id ID) (interface{}, error) {
	defer func(begin time.Time) {
		repo.counter.With(prometheus.Labels{"method": "get"}).Inc()
	}(time.Now())

	return repo.Repository.Get(id)
}

func (repo *instrumentingRepository) Put(i interface{}) (ID, error) {
	defer func(begin time.Time) {
		repo.counter.With(prometheus.Labels{"method": "put"}).Inc()
	}(time.Now())

	return repo.Repository.Put(i)
}

type instrumentingService struct {
	counter *prometheus.CounterVec
	Service
}

// NewInstrumentingService returns an instance of an instrumenting Service.
func NewInstrumentingService(c *prometheus.CounterVec, s Service) Service {
	return &instrumentingService{c, s}
}

func (s *instrumentingService) Get(id ID) (interface{}, error) {
	defer func(begin time.Time) {
		s.counter.With(prometheus.Labels{"method": "get"}).Inc()
	}(time.Now())

	return s.Service.Get(id)
}

func (s *instrumentingService) AddFolder(name string, parent ID) (ID, error) {
	defer func(begin time.Time) {
		s.counter.With(prometheus.Labels{"method": "addfolder"}).Inc()
	}(time.Now())

	return s.Service.AddFolder(name, parent)
}

func (s *instrumentingService) AddReport(name string, parent ID, content string) (ID, error) {
	defer func(begin time.Time) {
		s.counter.With(prometheus.Labels{"method": "addreport"}).Inc()
	}(time.Now())

	return s.Service.AddReport(name, parent, content)
}
