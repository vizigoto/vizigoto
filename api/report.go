// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package api

import (
	"github.com/graphql-go/graphql"
	"github.com/vizigoto/vizigoto/authz"
	"github.com/vizigoto/vizigoto/item"
)

func getReportType(pathType, permissionType *graphql.Object) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name:        "Report",
		Description: "Report represents a report in the content tree",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if f, ok := p.Source.(item.Report); ok {
						return f.ID, nil
					}
					return nil, nil
				},
			},
			"name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if f, ok := p.Source.(item.Report); ok {
						return f.Name, nil
					}
					return nil, nil
				},
			},
			"parent": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if f, ok := p.Source.(item.Report); ok {
						return f.Parent, nil
					}
					return nil, nil
				},
			},
			"content": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if f, ok := p.Source.(item.Report); ok {
						return f.Content, nil
					}
					return nil, nil
				},
			},
			"path": &graphql.Field{
				Type: graphql.NewList(pathType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if r, ok := p.Source.(item.Report); ok {
						return r.Path, nil
					}
					return nil, nil
				},
			},
			"permission": &graphql.Field{
				Type: permissionType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return authz.Permission{Write: true}, nil
				},
			},
		},
		IsTypeOf: func(p graphql.IsTypeOfParams) bool {
			_, ok := p.Value.(item.Report)
			return ok
		},
	})
}
