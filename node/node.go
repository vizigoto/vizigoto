// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

// Package node provides a low level API to handle nodes in a content tree.
package node

// ID uniquely identifies a particular node.
type ID string

// Node represents one single node in the content tree.
type Node struct {
	ID       ID
	Parent   ID
	Children []ID
}

// New allocates a node and returns a pointer to it.
func New(parent ID) *Node {
	return &Node{Parent: parent, Children: []ID{}}
}

// Repository provides a limited interface to a storage layer.
type Repository interface {
	Get(ID) (*Node, error)
	Put(*Node) (ID, error)
}
