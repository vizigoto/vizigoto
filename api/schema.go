package api

import "github.com/graphql-go/graphql"

func getSchema(query *graphql.Object, types []graphql.Type) (*graphql.Schema, error) {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: query,
		Types: types,
	})

	if err != nil {
		return nil, err
	}

	return &schema, nil
}
