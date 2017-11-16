// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package api

import (
	"github.com/graphql-go/graphql"
	"github.com/vizigoto/vizigoto/item"
	"github.com/vizigoto/vizigoto/user"
)

func (v1 *v1) getQueryType(folderType, reportType, userType *graphql.Object, itemType *graphql.Union) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{

			// me
			"me": &graphql.Field{
				Type: userType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					u := user.New("maria")
					u.ID = "xsadfuoaudsfosuf"
					return u, nil
				},
			},

			// folder
			"folder": &graphql.Field{
				Type: folderType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"].(string)
					i, err := v1.content.GetItem(id)
					if err != nil {
						return nil, err
					}
					if v, ok := i.(item.Folder); ok {
						return v, nil
					}
					return nil, nil
				},
			},

			// report
			"report": &graphql.Field{
				Type: reportType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"].(string)
					i, err := v1.content.GetItem(id)
					if err != nil {
						return nil, err
					}
					if v, ok := i.(item.Report); ok {
						return v, nil
					}
					return nil, nil
				},
			},

			// item
			"item": &graphql.Field{
				Type: itemType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"].(string)
					i, err := v1.content.GetItem(id)
					if err != nil {
						return nil, err
					}
					switch v := i.(type) {
					case item.Folder:
						return v, nil
					case item.Report:
						return v, nil
					}
					return nil, nil
				},
			},

			// home
			"home": &graphql.Field{
				Type: folderType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					h, err := v1.content.GetHome()
					if err != nil {
						return nil, err
					}
					return h, nil
				},
			},
		},
	})
}
