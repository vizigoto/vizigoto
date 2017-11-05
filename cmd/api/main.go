package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
)

type folder struct {
	Name string `json:"name"`
}

type report struct {
	Name string `json:"name"`
}

var odie = &folder{"Folder A"}

var folderType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Folder",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
		},
	},
	IsTypeOf: func(p graphql.IsTypeOfParams) bool {
		_, ok := p.Value.(*folder)
		return ok
	},
})

var reportType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Report",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
		},
	},
	IsTypeOf: func(p graphql.IsTypeOfParams) bool {
		_, ok := p.Value.(*report)
		return ok
	},
})
var itemType = graphql.NewUnion(graphql.UnionConfig{
	Name: "Item",
	Types: []*graphql.Object{
		folderType, reportType,
	},
	ResolveType: func(p graphql.ResolveTypeParams) *graphql.Object {
		if _, ok := p.Value.(*folder); ok {
			return folderType
		}
		if _, ok := p.Value.(*report); ok {
			return reportType
		}
		return nil
	},
})

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"item": &graphql.Field{
			Type: itemType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return odie, nil
			},
		},
	},
})

var unionInterfaceTestSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: queryType,
	Types: []graphql.Type{itemType},
})

func main() {

	q := `
		{
			item {
				... on Folder {
					name
				}
			}
		}
	`

	params := graphql.Params{Schema: unionInterfaceTestSchema, RequestString: q}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed. errors: %+v", r.Errors)
	}

	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON)
}
