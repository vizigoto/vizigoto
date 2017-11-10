package node_test

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/vizigoto/vizigoto/node"
	"github.com/vizigoto/vizigoto/node/storage/mem"
	"github.com/vizigoto/vizigoto/pkg/testutil"
)

func TestInstrumentingRepository(t *testing.T) {
	counter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "repository",
			Name:      "request_count",
			Help:      "Number of requests received.",
		},
		[]string{"method"},
	)

	repo := mem.NewRepository()
	repo = node.NewInstrumentingRepository(counter, repo)

	parent := node.ID("")
	folder := node.New(parent)

	folderID, err := repo.Put(folder)
	testutil.FatalOnError(t, err)

	_, err = repo.Get(folderID)
	testutil.FatalOnError(t, err)
}
