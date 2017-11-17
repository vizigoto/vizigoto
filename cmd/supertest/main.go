package main

import (
	"context"
	"fmt"
	"time"

	"github.com/vizigoto/vizigoto/item"
	"github.com/vizigoto/vizigoto/mem"
)

func main() {
	ctx := context.Background()
	nodes := mem.NewNodeRepository()
	repo := mem.NewItemRepository(nodes)
	service := item.NewService(repo)

	rootName, rootParent := "Home", ""
	rootID, err := service.AddFolder(ctx, rootName, rootParent)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 1000; i++ {
		childID, err := service.AddFolder(ctx, fmt.Sprintf("no %d", i), rootID)
		if err != nil {
			fmt.Println(err.Error())
		}
		c, err := service.Get(ctx, childID)
		if err != nil {
			fmt.Println(err.Error())
		}
		if child, ok := c.(*item.Folder); ok {
			fmt.Println(child.Name)
		}
		time.Sleep(time.Second)
	}
}
