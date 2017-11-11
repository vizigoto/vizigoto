// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package auth

type Repository interface {
	Get(id string) (i *Group, err error)
	Put(g *Group) (id string, err error)
}
