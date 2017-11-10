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
			if _, ok := p.Value.(*item.Folder); ok {
				return folderType
			}
			if _, ok := p.Value.(*item.Report); ok {
				return reportType
			}
			return nil
		},
	})
}
