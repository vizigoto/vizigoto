// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package api

import (
	"github.com/graphql-go/graphql"
	"github.com/vizigoto/vizigoto/item"
)

func getFolderType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name:        "Folder",
		Description: "Folder represents a folder in the content tree",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if f, ok := p.Source.(*item.Folder); ok {
						return f.ID, nil
					}
					return nil, nil
				},
			},
			"name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if f, ok := p.Source.(*item.Folder); ok {
						return f.Name, nil
					}
					return nil, nil
				},
			},
			"parent": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if f, ok := p.Source.(*item.Folder); ok {
						return f.Parent, nil
					}
					return nil, nil
				},
			},
		},
		IsTypeOf: func(p graphql.IsTypeOfParams) bool {
			_, ok := p.Value.(*item.Folder)
			return ok
		},
	})
}
