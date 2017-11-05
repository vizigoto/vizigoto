// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package mem_test

import (
	"testing"

	"github.com/vizigoto/vizigoto/pkg/item"
	"github.com/vizigoto/vizigoto/pkg/item/storage/mem"
	node "github.com/vizigoto/vizigoto/pkg/node/storage/mem"
)

func TestPutGet(t *testing.T) {
	nodes := node.NewRepository()
	var repo item.Repository = mem.NewRepository(nodes)
	root := item.NewFolder("Home", "", "")
	id, err := repo.Put(root)
	if err != nil {
		t.Fatal(err)
	}

	if id == "" {
		t.Fatal("id error")
	}

	f, err := repo.Get(id)
	if err != nil {
		t.Fatal(err)
	}

	folder, ok := f.(*item.Folder)
	if !ok {
		t.Fatal("type match fail")
	}

	if folder.Name != "Home" {
		t.Fatal("name error")
	}
}
