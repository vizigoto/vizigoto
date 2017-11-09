package item_test

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/vizigoto/vizigoto/item"
	"github.com/vizigoto/vizigoto/item/storage/mem"
	node "github.com/vizigoto/vizigoto/node/storage/mem"
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
	prometheus.MustRegister(counter)

	nodes := node.NewRepository()
	repo := mem.NewRepository(nodes)
	repo = item.NewInstrumentingRepository(counter, repo)

	name, parent := "Home", item.ID("")
	root := item.NewFolder(name, parent)
	id, err := repo.Put(root)
	testutil.FatalOnError(t, err)

	n, err := repo.Get(id)
	testutil.FatalOnError(t, err)

	if folder, ok := n.(*item.Folder); ok {
		if root.Name != folder.Name {
			t.Fatal("error")
		}
	}
}

func TestInstrumentingService(t *testing.T) {
	counter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name:      "name",
			Namespace: "namespace",
			Subsystem: "subsystem",
			Help:      "help.",
		},
		[]string{"method"},
	)
	nodes := node.NewRepository()
	repo := mem.NewRepository(nodes)
	service := item.NewService(repo)
	service = item.NewInstrumentingService(counter, service)

	rootID, err := service.AddFolder("Home", item.ID(""))
	testutil.FatalOnError(t, err)

	reportID, err := service.AddReport("report", rootID, "report content")
	testutil.FatalOnError(t, err)

	r, err := service.Get(reportID)
	testutil.FatalOnError(t, err)

	rep, ok := r.(*item.Report)
	if !ok {
		t.Fatal("type error")
	}
	if rep.Name != "report" {
		t.Fatal("report error")
	}
}
