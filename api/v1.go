// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package api

import (
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/vizigoto/vizigoto/content"
)

type v1 struct {
	content content.Service
}

// NewV1 returns a http.Handler which serves the v1 graphql api.
func NewV1(content content.Service) (http.Handler, error) {
	v1 := &v1{content}

	pathType := getPathType()
	permissionType := getPermissionType()
	folderType := getFolderType(pathType, permissionType)
	reportType := getReportType(pathType, permissionType)
	userType := getUserType()
	itemType := getItemType(folderType, reportType)

	query := v1.getQueryType(folderType, reportType, userType, itemType)
	mutation := v1.getMutationType(folderType, reportType, itemType)

	types := []graphql.Type{itemType}
	schema, err := getSchema(query, mutation, types)
	if err != nil {
		return nil, err
	}

	h := handler.New(&handler.Config{Schema: schema, Pretty: true})
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ContextHandler(r.Context(), w, r)
	}), nil
}
