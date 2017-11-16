// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

// Package node provides a low level API to handle nodes in a content tree.
package node

// Path to the node.
type Path []PathNode

// PathNode is the single node in the path to the node.
type PathNode interface {
	PathID() string
}

// Node represents one single node in the content tree.
type Node struct {
	ID       string
	Parent   string
	Children []string
	Path     Path
}

// New allocates a node and returns a pointer to it.
func New(parent string) *Node {
	return &Node{Parent: parent, Children: []string{}, Path: Path{}}
}

// PathID returns the id.
func (n *Node) PathID() string {
	return n.ID
}

// Repository provides a limited interface to a storage layer.
type Repository interface {
	Get(id string) (n *Node, err error)      // Get a node from content tree
	Put(n *Node) (id string, err error)      // Put a node in the content tree
	Move(n *Node, parent string) (err error) // Move a node to another parent node in the content tree
}
