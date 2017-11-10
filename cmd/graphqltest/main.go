package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
	"github.com/vizigoto/vizigoto/item"
	"github.com/vizigoto/vizigoto/item/storage/mem"
	node "github.com/vizigoto/vizigoto/node/storage/mem"
)

func main() {

	nodes := node.NewRepository()
	repo := mem.NewRepository(nodes)
	service := item.NewService(repo)

	rootName, rootParent := "Home", item.ID("")
	rootID, err := service.AddFolder(rootName, rootParent)
	if err != nil {
		panic(err)
	}
	reportName, reportContent := "report", "<h1>content"
	reportID, err := service.AddReport(reportName, rootID, reportContent)
	if err != nil {
		panic(err)
	}

	var folderType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Folder",
		Description: "Folder represents a folder in the content tree",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if f, ok := p.Source.(*item.Folder); ok {
						return f.ID, nil
					}
					return nil, nil
				},
			},
			"name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if f, ok := p.Source.(*item.Folder); ok {
						return f.Name, nil
					}
					return nil, nil
				},
			},
			"parent": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if f, ok := p.Source.(*item.Folder); ok {
						return f.Parent, nil
					}
					return nil, nil
				},
			},
		},
		IsTypeOf: func(p graphql.IsTypeOfParams) bool {
			_, ok := p.Value.(*item.Folder)
			return ok
		},
	})

	var reportType = graphql.NewObject(graphql.ObjectConfig{
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

	var itemType = graphql.NewUnion(graphql.UnionConfig{
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

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"folder": &graphql.Field{
				Type: folderType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := item.ID(p.Args["id"].(string))
					i, err := service.Get(id)
					if err != nil {
						return nil, err
					}
					if v, ok := i.(*item.Folder); ok {
						return v, nil
					}
					return nil, nil
				},
			},
			"report": &graphql.Field{
				Type: reportType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := item.ID(p.Args["id"].(string))
					i, err := service.Get(id)
					if err != nil {
						return nil, err
					}
					if v, ok := i.(*item.Report); ok {
						return v, nil
					}
					return nil, nil
				},
			},
			"item": &graphql.Field{
				Type: itemType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := item.ID(p.Args["id"].(string))
					i, err := service.Get(id)
					if err != nil {
						return nil, err
					}
					switch v := i.(type) {
					case *item.Folder:
						return v, nil
					case *item.Report:
						return v, nil
					}
					return nil, nil
				},
			},
		},
	})

	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
		Types: []graphql.Type{itemType},
	})

	query := `query f($id: String!) {folder(id: $id) {id, name}}`
	params := graphql.Params{Schema: schema, RequestString: query, VariableValues: map[string]interface{}{"id": rootID}}
	rf := graphql.Do(params)
	if len(rf.Errors) > 0 {
		log.Printf("failed to execute graphql operation, errors: %+v", rf.Errors)
	} else {
		rJSON, _ := json.Marshal(rf)
		fmt.Printf("%s \n", rJSON)
	}

	log.Println("-")
	log.Println("----------------------------------------------")

	query = `query r($id: String!) {report(id: $id) {id, name}}`
	params = graphql.Params{Schema: schema, RequestString: query, VariableValues: map[string]interface{}{"id": reportID}}
	rr := graphql.Do(params)
	if len(rr.Errors) > 0 {
		log.Printf("failed to execute graphql operation, errors: %+v", rr.Errors)
	} else {
		rJSON, _ := json.Marshal(rr)
		fmt.Printf("%s \n", rJSON)
	}

	log.Println("-")
	log.Println("----------------------------------------------")

	query = `query r($id: String!) {item(id: $id) { ... on Folder{id, name}}}`
	params = graphql.Params{Schema: schema, RequestString: query, VariableValues: map[string]interface{}{"id": rootID}}
	rr = graphql.Do(params)
	if len(rr.Errors) > 0 {
		log.Printf("failed to execute graphql operation, errors: %+v", rr.Errors)
	} else {
		rJSON, _ := json.Marshal(rr)
		fmt.Printf("%s \n", rJSON)
	}

	log.Println("-")
	log.Println("----------------------------------------------")

	query = `query r($id: String!) {item(id: $id) { ... on Report{id, name, content} ... on Folder{id, name}}}`
	params = graphql.Params{Schema: schema, RequestString: query, VariableValues: map[string]interface{}{"id": rootID}}
	rr = graphql.Do(params)
	if len(rr.Errors) > 0 {
		log.Printf("failed to execute graphql operation, errors: %+v", rr.Errors)
	} else {
		rJSON, _ := json.Marshal(rr)
		fmt.Printf("%s \n", rJSON)
	}
}
