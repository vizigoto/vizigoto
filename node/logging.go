// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package node

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

func (repo *loggingRepository) Get(id string) (n *Node, err error) {
	defer func(begin time.Time) {
		t := toFields(n)
		repo.logger.Log(
			"method", "get",
			"id", id,
			"parent", t.parent,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	n, err = repo.Repository.Get(id)
	return
}

func (repo *loggingRepository) Put(n *Node) (id string, err error) {
	defer func(begin time.Time) {
		repo.logger.Log(
			"method", "put",
			"id", id,
			"parent", n.Parent,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	id, err = repo.Repository.Put(n)
	return
}

func toFields(n *Node) no {
	r := no{}
	if n != nil {
		r.parent = string(n.Parent)
	}
	return r
}

type no struct {
	id     string
	parent string
}
