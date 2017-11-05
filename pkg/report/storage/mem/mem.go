package mem

import (
	"sync"

	"github.com/vizigoto/vizigoto/pkg/item"
	"github.com/vizigoto/vizigoto/pkg/report"
)

// Repository ...
type repository struct {
	mtx     sync.RWMutex
	reports map[string]*report.Report
}

// NewRepository ...
func NewRepository() report.Repository {
	return &repository{reports: make(map[string]*report.Report)}
}

func (repo *repository) Get(id string) (*report.Report, error) {
	repo.mtx.RLock()
	defer repo.mtx.RUnlock()
	if r, ok := repo.reports[id]; ok {
		return r, nil
	}
	return nil, item.ErrItemNotFound
}

func (repo *repository) Put(r *report.Report) error {
	repo.mtx.Lock()
	defer repo.mtx.Unlock()
	repo.reports[r.ID] = r
	return nil
}
