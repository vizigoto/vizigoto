// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package api

import (
	"github.com/graphql-go/graphql"
	"github.com/vizigoto/vizigoto/item"
)

func getPathType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name:        "Path",
		Description: "Path",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if p, ok := p.Source.(item.Path); ok {
						return p.PathID(), nil
					}
					return nil, nil
				},
			},
			"name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if p, ok := p.Source.(item.Path); ok {
						return p.PathName(), nil
					}
					return nil, nil
				},
			},
		},
		IsTypeOf: func(p graphql.IsTypeOfParams) bool {
			_, ok := p.Value.(item.Path)
			return ok
		},
	})
}
