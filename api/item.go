// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package api

import (
	"github.com/graphql-go/graphql"
	"github.com/vizigoto/vizigoto/item"
)

func getItemType(folderType, reportType *graphql.Object) *graphql.Union {

	return graphql.NewUnion(graphql.UnionConfig{
		Name: "Item",
		Types: []*graphql.Object{
			reportType, folderType,
		},
		ResolveType: func(p graphql.ResolveTypeParams) *graphql.Object {
			if _, ok := p.Value.(item.Folder); ok {
				return folderType
			}
			if _, ok := p.Value.(item.Report); ok {
				return reportType
			}
			return nil
		},
	})
}
