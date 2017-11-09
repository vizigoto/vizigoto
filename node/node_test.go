package node_test

import (
	"testing"

	"github.com/vizigoto/vizigoto/node"
)

func TestNewFolder(t *testing.T) {
	parent := node.ID("xyz")
	n := node.New(parent)
	if n.Parent != parent {
		t.Fatal("folder error")
	}
}

func TestNewReport(t *testing.T) {
	parent := node.ID("abc")
	n := node.New(parent)
	if n.Parent != parent {
		t.Fatal("report error")
	}
}
