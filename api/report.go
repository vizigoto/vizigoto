package api

import (
	"github.com/graphql-go/graphql"
	"github.com/vizigoto/vizigoto/item"
)

func getReportType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name:        "Report",
		Description: "Report represents a report in the content tree",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if f, ok := p.Source.(*item.Report); ok {
						return f.ID, nil
					}
					return nil, nil
				},
			},
			"name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if f, ok := p.Source.(*item.Report); ok {
						return f.Name, nil
					}
					return nil, nil
				},
			},
			"parent": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if f, ok := p.Source.(*item.Report); ok {
						return f.Parent, nil
					}
					return nil, nil
				},
			},
			"content": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if f, ok := p.Source.(*item.Report); ok {
						return f.Content, nil
					}
					return nil, nil
				},
			},
		},
		IsTypeOf: func(p graphql.IsTypeOfParams) bool {
			_, ok := p.Value.(*item.Report)
			return ok
		},
	})
}
