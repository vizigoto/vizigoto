// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

// Package node provides a low level API to handle nodes in a content tree.
package node

// Node represents one single node in the content tree.
type Node struct {
	ID       string
	Parent   string
	Children []string
	Path     []string
}

// New allocates a node and returns a pointer to it.
func New(parent string) *Node {
	return &Node{Parent: parent, Children: []string{}, Path: []string{}}
}

// Repository provides a limited interface to a storage layer.
type Repository interface {
	Get(id string) (n *Node, err error)
	Put(n *Node) (id string, err error)
}
