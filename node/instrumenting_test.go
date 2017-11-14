// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package node_test

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/vizigoto/vizigoto/node"
	"github.com/vizigoto/vizigoto/node/storage/mem"
	"github.com/vizigoto/vizigoto/pkg/testutil"
)

func TestInstrumentingRepository(t *testing.T) {
	counter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "repository",
			Name:      "request_count",
			Help:      "Number of requests received.",
		},
		[]string{"method"},
	)

	repo := mem.NewRepository()
	repo = node.NewInstrumentingRepository(counter, repo)

	parent := ""
	folder := node.New(parent)

	folderID, err := repo.Put(folder)
	testutil.FatalOnError(t, err)

	_, err = repo.Get(folderID)
	testutil.FatalOnError(t, err)
}
