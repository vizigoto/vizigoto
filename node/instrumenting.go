package node

import (
	"time"

	"github.com/vizigoto/vizigoto/pkg/metrics"
)

type instrumentingRepository struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	Repository
}

// NewInstrumentingRepository returns an instance of an instrumenting Repository.
func NewInstrumentingRepository(counter metrics.Counter, latency metrics.Histogram, r Repository) Repository {
	return &instrumentingRepository{counter, latency, r}
}

func (repo *instrumentingRepository) Get(id ID) (*Node, error) {
	defer func(begin time.Time) {
		repo.requestCount.With("method", "get").Add(1)
		repo.requestLatency.With("method", "get").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return repo.Repository.Get(id)
}

func (repo *instrumentingRepository) Put(i *Node) (ID, error) {
	defer func(begin time.Time) {
		repo.requestCount.With("method", "put").Add(1)
		repo.requestLatency.With("method", "put").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return repo.Put(i)
}
