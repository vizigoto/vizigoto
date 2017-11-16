package main

import (
	"log"
	"net/http"

	"github.com/vizigoto/vizigoto/api"
	"github.com/vizigoto/vizigoto/content"
	"github.com/vizigoto/vizigoto/item"
	"github.com/vizigoto/vizigoto/mem"
	"github.com/vizigoto/vizigoto/user"
)

func main() {
	nodes := mem.NewNodeRepository()
	nodeRepo := mem.NewItemRepository(nodes)
	items := item.NewService(nodeRepo)

	userRepo := mem.NewUserRepository()
	users := user.NewService(userRepo)

	rootName, rootParent := "Home", ""
	rootID, err := items.AddFolder(rootName, rootParent)
	if err != nil {
		panic(err)
	}
	log.Println(rootID)

	reportName, reportContent := "report", "<h1>content"
	reportID, err := items.AddReport(reportName, rootID, reportContent)
	if err != nil {
		panic(err)
	}
	log.Println(reportID)

	content := content.NewService(rootID, items, users)
	v1Handler, err := api.NewV1(content)
	if err != nil {
		panic(err)
	}

	http.Handle("/v1/graphql", v1Handler)

	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	log.Fatal(http.ListenAndServe(":7171", nil))
}
