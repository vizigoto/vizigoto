// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package node

import (
	"github.com/vizigoto/vizigoto/pkg/user"
)

type ID string
type Kind string

const (
	KindFolder Kind = "folder"
	KindReport Kind = "report"
)

type Node struct {
	ID       ID
	Name     string
	Kind     Kind
	Parent   ID
	Owner    user.ID
	Children []ID
}

var EmptyID = ID("")

func NewNode(name string, kind Kind, parent ID, owner user.ID) *Node {
	return &Node{
		Name:     name,
		Kind:     kind,
		Parent:   parent,
		Owner:    owner,
		Children: []ID{},
	}
}

type Repository interface {
	Get(id ID) (*Node, error)
	Put(*Node) (ID, error)
}
