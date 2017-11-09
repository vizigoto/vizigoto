package node_test

import (
	"os"
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/vizigoto/vizigoto/node"
	"github.com/vizigoto/vizigoto/node/storage/mem"
	"github.com/vizigoto/vizigoto/pkg/testutil"
)

func TestLoggingRepository(t *testing.T) {
	w := log.NewSyncWriter(os.Stdout)
	logger := log.NewLogfmtLogger(w)

	repo := mem.NewRepository()
	repo = node.NewLoggingRepository(logger, repo)
	folder := node.New(node.Folder, "")
	folderID, err := repo.Put(folder)
	testutil.FatalOnError(t, err)

	n, err := repo.Get(folderID)
	testutil.FatalOnError(t, err)

	if n.Kind != node.Folder {
		t.Fatal("kind error")
	}
}
