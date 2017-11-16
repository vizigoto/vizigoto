package main

import (
	"fmt"
	"time"

	"github.com/vizigoto/vizigoto/item"
	"github.com/vizigoto/vizigoto/mem"
)

func main() {
	nodes := mem.NewNodeRepository()
	repo := mem.NewItemRepository(nodes)
	service := item.NewService(repo)

	rootName, rootParent := "Home", ""
	rootID, err := service.AddFolder(rootName, rootParent)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 1000; i++ {
		childID, err := service.AddFolder(fmt.Sprintf("no %d", i), rootID)
		if err != nil {
			fmt.Println(err.Error())
		}
		c, err := service.Get(childID)
		if err != nil {
			fmt.Println(err.Error())
		}
		if child, ok := c.(*item.Folder); ok {
			fmt.Println(child.Name)
		}
		time.Sleep(time.Second)
	}
}
