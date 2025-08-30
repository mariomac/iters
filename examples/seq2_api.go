package main

import (
	"cmp"
	"fmt"
	"maps"
	"os"

	"github.com/mariomac/gostream/item"
	"github.com/mariomac/gostream/order"
	"github.com/mariomac/iters"
)

func main_seq2_api() {
	// get a slice of os.DirEntry elements
	osFiles, err := os.ReadDir(".")
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}

	// Put osFiles slice into a stream and filter it
	justFiles := iters.OfSlice(osFiles).
		Filter(func(entry os.DirEntry) bool {
			return !entry.IsDir()
		})

	// 1. Convert each os.DirEntry to a name/size item.Pair (iters.Map function)
	// 2. Sorting the stream by file size in descending order
	//    (using Sorted method cmp.Compare and order.Inverse)
	// 3. Limit the stream size to the top 3 files by size (using Limit method)
	sizeTop3 := iters.Map(justFiles,
		func(entry os.DirEntry) item.Pair[string, int64] {
			info, _ := entry.Info()
			return item.Pair[string, int64]{
				Key: info.Name(),
				Val: info.Size(),
			}
		}).
		Sorted(order.Inverse(func(a, b item.Pair[string, int64]) int {
			return cmp.Compare(a.Val, b.Val)
		})).
		Limit(3)

	// iters.Seq2 function allows iterating the stream within a for..range
	fmt.Println("Top 3 files:")
	for k, v := range iters.Seq2(sizeTop3) {
		fmt.Printf("%v (%v)\n", k, v)
	}

	// iters.Seq2 function allows also connecting the stream to other Go
	// standard library functions that expect an iter.Seq2 input
	// for example, maps.Collect
	asGoMap := maps.Collect(iters.Seq2(sizeTop3))

	if _, ok := asGoMap["README.md"]; ok {
		fmt.Println("README.md is in the top 3 files")
	}
	if _, ok := asGoMap["foo.go"]; !ok {
		fmt.Println("foo.go is NOT in the top 3 files")
	}
}
