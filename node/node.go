// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package node

import (
	"github.com/vizigoto/vizigoto/user"
)

// ID uniquely identifies a particular node.
type ID string

// Kind of node
type Kind string

// Kinds of nodes
const (
	Folder Kind = "folder"
	Report Kind = "report"
)

// Node represents one single node in the content tree.
type Node struct {
	ID       ID
	Name     string
	Kind     Kind
	Parent   ID
	Owner    user.ID
	Children []ID
}

// New allocates a node and returns a pointer to it.
func New(name string, kind Kind, parent ID, owner user.ID) *Node {
	return &Node{
		Name:     name,
		Kind:     kind,
		Parent:   parent,
		Owner:    owner,
		Children: []ID{},
	}
}

// Repository provides a limited interface to a storage layer.
type Repository interface {
	Get(ID) (*Node, error)
	Put(*Node) (ID, error)
}
