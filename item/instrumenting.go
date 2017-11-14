// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

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

func (repo *instrumentingRepository) Get(id string) (interface{}, error) {
	defer func(begin time.Time) {
		repo.counter.With(prometheus.Labels{"method": "get"}).Inc()
	}(time.Now())

	return repo.Repository.Get(id)
}

func (repo *instrumentingRepository) Put(i interface{}) (string, error) {
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

func (s *instrumentingService) Get(id string) (interface{}, error) {
	defer func(begin time.Time) {
		s.counter.With(prometheus.Labels{"method": "get"}).Inc()
	}(time.Now())

	return s.Service.Get(id)
}

func (s *instrumentingService) AddFolder(name, parent string) (string, error) {
	defer func(begin time.Time) {
		s.counter.With(prometheus.Labels{"method": "addfolder"}).Inc()
	}(time.Now())

	return s.Service.AddFolder(name, parent)
}

func (s *instrumentingService) AddReport(name, parent, content string) (string, error) {
	defer func(begin time.Time) {
		s.counter.With(prometheus.Labels{"method": "addreport"}).Inc()
	}(time.Now())

	return s.Service.AddReport(name, parent, content)
}
