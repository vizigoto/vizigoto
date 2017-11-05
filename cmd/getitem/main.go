package main

import (
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"strconv"

	_ "net/http/pprof"

	"github.com/vizigoto/vizigoto/pkg/item"
	itemRepo "github.com/vizigoto/vizigoto/pkg/item/storage/mem"
)

func main() {

	itemRepo := itemRepo.NewRepository()
	itemRepo.Put(item.NewItem("root", "Root", "", "", "folder", "user"))
	itemService := item.NewService("root", itemRepo)
	id, _ := itemService.AddItem("Financeiro", "root", "folder", "user")
	id2, _ := itemService.AddItem("Controladoria", id, "folder", "user")
	log.Println(id)
	log.Println(id2)

	handler := item.MakeHandler(itemService)
	_ = handler
	//http.Handle("/", handler)
	//http.ListenAndServe(":7070", nil)

	for kx := 0; kx < 1000; kx++ {
		itemService.AddItem(strconv.Itoa(kx), strconv.Itoa(kx), "folder", "user")
	}

	f, err := os.Create("mem.prof")
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}
	runtime.GC() // get up-to-date statistics
	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Fatal("could not write memory profile: ", err)
	}
	f.Close()
}
