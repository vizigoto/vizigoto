package api

import (
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/vizigoto/vizigoto/item"
)

type v1 struct {
	service item.Service
}

// NewV1 returns a http.Handler which serves the v1 graphql api.
func NewV1(service item.Service) (http.Handler, error) {
	v1 := &v1{service}

	folderType := getFolderType()
	reportType := getReportType()
	itemType := getItemType(folderType, reportType)

	query := v1.getQueryType(folderType, reportType, itemType)
	types := []graphql.Type{itemType}
	schema, err := getSchema(query, types)
	if err != nil {
		return nil, err
	}

	h := handler.New(&handler.Config{Schema: schema, Pretty: true})
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ContextHandler(r.Context(), w, r)
	}), nil
}
