package item_test

import (
	"testing"

	"github.com/vizigoto/vizigoto/item"
	"github.com/vizigoto/vizigoto/item/storage/mem"
	node "github.com/vizigoto/vizigoto/node/storage/mem"
)

func TestNewFolder(t *testing.T) {
	name, parent := "Home", item.ID("")
	i := item.NewFolder(name, parent)
	if i.Name != name ||
		i.Parent != parent {
		t.Fatal("folder error")
	}
}

func TestNewReport(t *testing.T) {
	name, parent, content := "Home", item.ID(""), "<h1>report"
	i := item.NewReport(name, parent, content)
	if i.Name != name ||
		i.Parent != parent {
		t.Fatal("report error")
	}
}

func TestService(t *testing.T) {
	nodes := node.NewRepository()
	repo := mem.NewRepository(nodes)
	service := item.NewService(repo)

	rootName, rootParent := "Home", item.ID("")
	rootID, err := service.AddFolder(rootName, rootParent)
	if err != nil {
		t.Fatal(err)
	}

	reportName, reportContent := "report", "<h1>content"
	reportID, err := service.AddReport(reportName, rootID, reportContent)
	if err != nil {
		t.Fatal(err)
	}

	f, err := service.Get(rootID)
	if err != nil {
		t.Fatal(err)
	}

	r, err := service.Get(reportID)
	if err != nil {
		t.Fatal(err)
	}

	folder, ok := f.(*item.Folder)
	if !ok {
		t.Fatal("type error")
	}

	if folder.Name != rootName ||
		folder.Parent != rootParent {
		t.Fatal("folder error")
	}

	report, ok := r.(*item.Report)
	if !ok {
		t.Fatal("type error")
	}

	if report.Name != reportName ||
		report.Content != reportContent ||
		report.Parent != rootID {
		t.Fatal("report error")
	}
}
