// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package api

import (
	"github.com/graphql-go/graphql"
	"github.com/vizigoto/vizigoto/item"
)

func (v1 *v1) getQueryType(folderType, reportType *graphql.Object, itemType *graphql.Union) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"folder": &graphql.Field{
				Type: folderType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"].(string)
					i, err := v1.service.Get(id)
					if err != nil {
						return nil, err
					}
					if v, ok := i.(*item.Folder); ok {
						return v, nil
					}
					return nil, nil
				},
			},
			"report": &graphql.Field{
				Type: reportType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"].(string)
					i, err := v1.service.Get(id)
					if err != nil {
						return nil, err
					}
					if v, ok := i.(*item.Report); ok {
						return v, nil
					}
					return nil, nil
				},
			},
			"item": &graphql.Field{
				Type: itemType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"].(string)
					i, err := v1.service.Get(id)
					if err != nil {
						return nil, err
					}
					switch v := i.(type) {
					case *item.Folder:
						return v, nil
					case *item.Report:
						return v, nil
					}
					return nil, nil
				},
			},
		},
	})
}
