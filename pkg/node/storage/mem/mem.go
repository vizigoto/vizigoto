package mem

import (
	"errors"
	"sync"

	"github.com/vizigoto/vizigoto/pkg/node"
)

type repository struct {
	mtx   sync.RWMutex
	nodes map[node.ID]*node.Node
}

func NewRepository() node.Repository {
	return &repository{nodes: make(map[node.ID]*node.Node)}
}

func (repo *repository) Get(id node.ID) (*node.Node, error) {
	repo.mtx.RLock()
	defer repo.mtx.RUnlock()
	if i, ok := repo.nodes[id]; ok {
		repo.assembleItem(i)
		return i, nil
	}
	return nil, errors.New("node not found")
}

func (repo *repository) Put(i *node.Node) (node.ID, error) {
	repo.mtx.Lock()
	defer repo.mtx.Unlock()
	id := node.NewID()
	i.ID = id
	repo.nodes[id] = i
	return id, nil
}

func (repo *repository) assembleItem(i *node.Node) {
	i.Children = []node.ID{}
	for _, v := range repo.nodes {
		if v.Parent == i.ID {
			i.Children = append(i.Children, v.ID)
		}
	}
}
