// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package auth

type Group struct {
	ID     string
	Name   string
	Parent string
}

func NewGroup(name, parent string) *Group {
	return &Group{Name: name, Parent: parent}
}
