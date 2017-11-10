// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

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
