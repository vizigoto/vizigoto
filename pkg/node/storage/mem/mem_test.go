package mem_test

import (
	"testing"

	"github.com/vizigoto/vizigoto/pkg/node"
	"github.com/vizigoto/vizigoto/pkg/node/storage/mem"
)

func TestPutGet(t *testing.T) {
	var repo node.Repository = mem.NewRepository()
	root := node.NewNode("Home", node.KindFolder, node.ID(""), "x")
	rootID, err := repo.Put(root)
	if err != nil {
		t.Fatal(err)
	}

	i, err := repo.Get(rootID)
	if err != nil {
		t.Fatal(err)
	}

	if root.ID != i.ID {
		t.Fatal("id error")
	}

	_, err = repo.Get("unknow")
	if err == nil {
		t.Fatal("error expected")
	}

	children := []string{"a", "b", "c", "d"}
	childrenIDs := []node.ID{}
	for _, c := range children {
		id, err := repo.Put(node.NewNode(c, node.KindFolder, rootID, "x"))
		if err != nil {
			t.Fatal(err)
		}
		childrenIDs = append(childrenIDs, id)
	}

	i, err = repo.Get(rootID)
	if err != nil {
		t.Fatal(err)
	}
	ok := true
	for _, j := range i.Children {
		for _, c := range childrenIDs {
			if j == c {
				ok = false
			}
		}
	}
	if ok {
		t.Fatal("children fail")
	}
}
