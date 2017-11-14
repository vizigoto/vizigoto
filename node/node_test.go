// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package node_test

import (
	"testing"

	"github.com/vizigoto/vizigoto/node"
)

func TestNewFolder(t *testing.T) {
	parent := "xyz"
	n := node.New(parent)
	if n.Parent != parent {
		t.Fatal("folder error")
	}
}

func TestNewReport(t *testing.T) {
	parent := "abc"
	n := node.New(parent)
	if n.Parent != parent {
		t.Fatal("report error")
	}
}
