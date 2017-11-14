// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package node_test

import (
	"fmt"
	"log"

	"github.com/vizigoto/vizigoto/node"
	"github.com/vizigoto/vizigoto/node/storage/mem"
)

func ExampleNew() {
	repo := mem.NewRepository()

	parent := ""
	no := node.New(parent)

	id, err := repo.Put(no)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("node ID: %v", id)
}
