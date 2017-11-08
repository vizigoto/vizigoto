package node_test

import (
	"testing"

	"github.com/vizigoto/vizigoto/node"
)

func TestNewFolder(t *testing.T) {
	kind, parent := node.Folder, node.ID("xyz")
	n := node.New(kind, parent)
	if n.Kind != kind || n.Parent != parent {
		t.Fatal("folder error")
	}
}

func TestNewReport(t *testing.T) {
	kind, parent := node.Report, node.ID("abc")
	n := node.New(kind, parent)
	if n.Kind != kind || n.Parent != parent {
		t.Fatal("report error")
	}
}
