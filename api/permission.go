package api

import (
	"github.com/graphql-go/graphql"
	"github.com/vizigoto/vizigoto/authz"
)

func getPermissionType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name:        "Permission",
		Description: "Permission represents a user permission in the content tree",
		Fields: graphql.Fields{
			"write": &graphql.Field{
				Type: graphql.Boolean,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return true, nil
				},
			},
		},
		IsTypeOf: func(p graphql.IsTypeOfParams) bool {
			_, ok := p.Value.(authz.Permission)
			return ok
		},
	})
}
