package mem_test

import (
	"testing"

	"github.com/vizigoto/vizigoto/pkg/item"
	"github.com/vizigoto/vizigoto/pkg/report"
	"github.com/vizigoto/vizigoto/pkg/report/storage/mem"
)

func TestPutGet(t *testing.T) {
	repo := mem.NewRepository()
	report := report.NewReport("report", "Report", "root", "content")
	repo.Put(report)

	i, err := repo.Get(report.ID)
	if err != nil {
		t.Fatal(err)
	}

	if report.ID != i.ID {
		t.Fatal("name error")
	}

	_, err = repo.Get("unknow")
	if err != item.ErrItemNotFound {
		t.Fatal("error expected")
	}
}
