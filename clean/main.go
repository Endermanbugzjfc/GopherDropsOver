package main

import (
	"fmt"
	"github.com/Bios-Marcel/wastebasket"
	"os"
	"path/filepath"
	"time"
)

func main() {
	ec := 0
	for i, f := range os.Args[1:] {
		if i != 0 {
			fmt.Println()
		}
		fmt.Println(f)
		dir := filepath.Dir(f)
		lock := filepath.Join(
			dir,
			"clean.gopher-drops-over",
		)
		t := time.Now()
		if _, err := os.Stat(f); os.IsNotExist(err) {
			fmt.Println("Already cleaned")
		} else if _, err := os.Stat(lock); os.IsNotExist(err) {
			fmt.Println("File is protected")
		} else if err := wastebasket.Trash(dir); err != nil {
			ec++
			fmt.Println(err)
		} else {
			fmt.Printf(
				"Moved to system recycle bin (%s)\n",
				time.Since(t),
			)
		}
	}
	fmt.Println("Error count: ", ec)
}
