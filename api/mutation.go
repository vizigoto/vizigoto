// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package api

import (
	"github.com/graphql-go/graphql"
	"github.com/vizigoto/vizigoto/item"
)

func (v1 *v1) getMutationType(folderType, reportType *graphql.Object, itemType *graphql.Union) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"addFolder": &graphql.Field{
				Type: folderType,
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"parent": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					name, parent := p.Args["name"].(string), p.Args["parent"].(string)
					id, err := v1.service.AddFolder(name, item.ID(parent))
					if err != nil {
						return nil, err
					}

					f, err := v1.service.Get(id)
					if err != nil {
						return nil, err
					}
					return f, nil
				},
			},
		},
	})
}
