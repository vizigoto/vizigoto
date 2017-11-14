package main

import (
	"log"
	"net/http"

	"github.com/vizigoto/vizigoto/api"
	"github.com/vizigoto/vizigoto/item"
	"github.com/vizigoto/vizigoto/item/storage/mem"
	nodeRepo "github.com/vizigoto/vizigoto/node/storage/mem"
)

func main() {
	nodes := nodeRepo.NewRepository()
	repo := mem.NewRepository(nodes)
	service := item.NewService(repo)

	rootName, rootParent := "Home", ""
	rootID, err := service.AddFolder(rootName, rootParent)
	if err != nil {
		panic(err)
	}
	log.Println(rootID)

	reportName, reportContent := "report", "<h1>content"
	reportID, err := service.AddReport(reportName, rootID, reportContent)
	if err != nil {
		panic(err)
	}
	log.Println(reportID)

	v1Handler, err := api.NewV1(service)
	if err != nil {
		panic(err)
	}

	http.Handle("/v1/graphql", v1Handler)

	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	log.Fatal(http.ListenAndServe(":7171", nil))
}
