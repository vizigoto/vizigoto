// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package node_test

import (
	"testing"

	"github.com/vizigoto/vizigoto/node"
)

func TestNewNode(t *testing.T) {
	parent := "xyz"
	n := node.New(parent)
	if n.Parent != parent {
		t.Fatal("folder error")
	}
}
