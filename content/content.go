// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package content

import (
	"github.com/vizigoto/vizigoto/item"
)

type contentService struct {
	item.Service
}

// NewContentService returns an instance of an content Service.
func NewContentService(s item.Service) item.Service {
	return &contentService{s}
}

func (s *contentService) Get(id string) (interface{}, error) {
	return s.Service.Get(id)
}

func (s *contentService) AddFolder(name, parent string) (string, error) {
	return s.Service.AddFolder(name, parent)
}

func (s *contentService) AddReport(name, parent, content string) (string, error) {
	return s.Service.AddReport(name, parent, content)
}
