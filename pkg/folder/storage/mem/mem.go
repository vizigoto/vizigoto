package mem

import (
	"sync"

	"github.com/vizigoto/vizigoto/pkg/folder"
	"github.com/vizigoto/vizigoto/pkg/item"
)

// Repository ...
type repository struct {
	mtx     sync.RWMutex
	folders map[string]*folder.Folder
}

// NewRepository ...
func NewRepository() folder.Repository {
	return &repository{folders: make(map[string]*folder.Folder)}
}

func (repo *repository) Get(id string) (*folder.Folder, error) {
	repo.mtx.RLock()
	defer repo.mtx.RUnlock()
	if i, ok := repo.folders[id]; ok {
		repo.assembleFolder(i)
		return i, nil
	}
	return nil, item.ErrItemNotFound
}

func (repo *repository) Put(f *folder.Folder) error {
	repo.mtx.Lock()
	defer repo.mtx.Unlock()
	repo.folders[f.ID] = f
	return nil
}

func (repo *repository) assembleFolder(f *folder.Folder) {
	f.Children = []string{}
	for _, v := range repo.folders {
		if v.Parent == f.ID {
			f.Children = append(f.Children, v.ID)
		}
	}
}
