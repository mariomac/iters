package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"

	"github.com/mariomac/iters"
)

func main_seq_api() {
	// get a slice of os.DirEntry elements
	osFiles, err := os.ReadDir(".")
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}
	// sort files by top size, using the standard Go slices package
	slices.SortFunc(osFiles, func(a, b os.DirEntry) int {
		ainf, _ := a.Info()
		binf, _ := b.Info()
		return -cmp.Compare(ainf.Size(), binf.Size())
	})

	// Put osFiles slice into a sequence and filter it
	// This operation doesn't involve any slice copy
	justFiles := iters.Filter(
		slices.Values(osFiles),
		func(entry os.DirEntry) bool {
			return !entry.IsDir()
		})

	// Limit the sequence to the top 3 files
	sizeTop3 := iters.Limit(3, justFiles)
	fmt.Println("Top 3 files:")
	for v := range sizeTop3 {
		fmt.Println(v.Name())
	}
}
