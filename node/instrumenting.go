// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package node

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

func (repo *instrumentingRepository) Get(id ID) (*Node, error) {
	defer func(begin time.Time) {
		repo.counter.With(prometheus.Labels{"method": "get"}).Inc()
	}(time.Now())

	return repo.Repository.Get(id)
}

func (repo *instrumentingRepository) Put(i *Node) (ID, error) {
	defer func(begin time.Time) {
		repo.counter.With(prometheus.Labels{"method": "put"}).Inc()
	}(time.Now())

	return repo.Repository.Put(i)
}
