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
			Name:      "name",
			Namespace: "namespace",
			Subsystem: "subsystem",
			Help:      "help.",
		},
		[]string{"method"},
	)

	repo := mem.NewRepository()
	repo = node.NewInstrumentingRepository(counter, repo)
	folder := node.New("")
	folderID, err := repo.Put(folder)
	testutil.FatalOnError(t, err)

	_, err = repo.Get(folderID)
	testutil.FatalOnError(t, err)
}
