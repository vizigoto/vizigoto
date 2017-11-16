// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package api

import (
	"github.com/graphql-go/graphql"
	"github.com/vizigoto/vizigoto/user"
)

func getUserType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name:        "User",
		Description: "User represents a user in the content tree",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if u, ok := p.Source.(user.User); ok {
						return u.ID, nil
					}
					return nil, nil
				},
			},
			"name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if u, ok := p.Source.(user.User); ok {
						return u.Name, nil
					}
					return nil, nil
				},
			},
		},
		IsTypeOf: func(p graphql.IsTypeOfParams) bool {
			_, ok := p.Value.(user.User)
			return ok
		},
	})
}
