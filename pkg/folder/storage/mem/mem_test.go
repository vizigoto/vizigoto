package mem_test

import (
	"testing"

	"github.com/vizigoto/vizigoto/pkg/folder"
	"github.com/vizigoto/vizigoto/pkg/folder/storage/mem"
	"github.com/vizigoto/vizigoto/pkg/item"
)

func TestPutGet(t *testing.T) {
	repo := mem.NewRepository()
	root := folder.NewFolder("root", "", "Root")
	repo.Put(root)

	i, err := repo.Get(root.ID)
	if err != nil {
		t.Fatal(err)
	}

	if root.ID != i.ID {
		t.Fatal("id error")
	}

	_, err = repo.Get("unknow")
	if err != item.ErrItemNotFound {
		t.Fatal("error expected")
	}

	children := []string{"a", "b", "c", "d"}
	for _, c := range children {
		repo.Put(folder.NewFolder(c, root.ID, c))
	}

	i, err = repo.Get(root.ID)
	if err != nil {
		t.Fatal(err)
	}

	ok := true
	for _, j := range i.Children {
		for _, c := range children {
			if j == c {
				ok = false
			}
		}
	}
	if ok {
		t.Fatal("children fail")
	}
}
