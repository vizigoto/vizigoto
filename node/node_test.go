// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package node_test

import (
	"testing"

	"github.com/vizigoto/vizigoto/mem"
	"github.com/vizigoto/vizigoto/node"
	"github.com/vizigoto/vizigoto/pkg/testutil"
)

func TestNewNode(t *testing.T) {
	repo := mem.NewNodeRepository()

	parent := ""
	n := node.New(parent)

	id, err := repo.Put(n)
	testutil.FatalOnError(t, err)

	n, err = repo.Get(id)
	testutil.FatalOnError(t, err)

	if n.PathID() != id {
		t.Fatal("pathID error")
	}
}
