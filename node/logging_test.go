// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package node_test

import (
	"os"
	"testing"

	"github.com/vizigoto/vizigoto/node"
	"github.com/vizigoto/vizigoto/node/storage/mem"
	"github.com/vizigoto/vizigoto/pkg/log"
	"github.com/vizigoto/vizigoto/pkg/testutil"
)

func TestLoggingRepository(t *testing.T) {
	w := log.NewSyncWriter(os.Stdout)
	logger := log.NewLogfmtLogger(w)

	repo := mem.NewRepository()
	repo = node.NewLoggingRepository(logger, repo)

	parent := ""
	folder := node.New(parent)

	folderID, err := repo.Put(folder)
	testutil.FatalOnError(t, err)

	_, err = repo.Get(folderID)
	testutil.FatalOnError(t, err)
}
