package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/vizigoto/vizigoto/item"
	"github.com/vizigoto/vizigoto/item/storage/mem"
	"github.com/vizigoto/vizigoto/node"
	nodeRepo "github.com/vizigoto/vizigoto/node/storage/mem"
	"github.com/vizigoto/vizigoto/pkg/log"
)

func main() {
	nodeCounterRepo := prometheus.NewCounterVec(prometheus.CounterOpts{Namespace: "vizigoto", Subsystem: "node", Name: "repository", Help: "help"}, []string{"method"})
	prometheus.MustRegister(nodeCounterRepo)

	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(":8081", nil)

	w := log.NewSyncWriter(os.Stdout)
	logger := log.NewLogfmtLogger(w)

	nodes := nodeRepo.NewRepository()
	nodes = node.NewInstrumentingRepository(nodeCounterRepo, nodes)
	nodes = node.NewLoggingRepository(logger, nodes)

	itemCounterRepo := prometheus.NewCounterVec(prometheus.CounterOpts{Namespace: "vizigoto", Subsystem: "item", Name: "repository", Help: "help"}, []string{"method"})
	prometheus.MustRegister(itemCounterRepo)

	repo := mem.NewRepository(nodes)
	repo = item.NewInstrumentingRepository(itemCounterRepo, repo)
	repo = item.NewLoggingRepository(logger, repo)

	serviceCounterRepo := prometheus.NewCounterVec(prometheus.CounterOpts{Namespace: "vizigoto", Subsystem: "item", Name: "service", Help: "help"}, []string{"method"})
	prometheus.MustRegister(serviceCounterRepo)

	service := item.NewService(repo)
	service = item.NewInstrumentingService(serviceCounterRepo, service)
	service = item.NewLoggingService(logger, service)

	rootName, rootParent := "Home", item.ID("")
	rootID, err := service.AddFolder(rootName, rootParent)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 1000; i++ {
		childID, err := service.AddFolder(fmt.Sprintf("no %d", i), rootID)
		if err != nil {
			fmt.Println(err.Error())
		}
		c, err := service.Get(childID)
		if err != nil {
			fmt.Println(err.Error())
		}
		if child, ok := c.(*item.Folder); ok {
			fmt.Println(child.Name)
		}
		time.Sleep(time.Second)
	}
}
