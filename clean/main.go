package main

import (
	"fmt"
	"github.com/Bios-Marcel/wastebasket"
	"os"
	"path/filepath"
	"time"
)

func main() {
	for i, f := range os.Args[1:] {
		if i != 0 {
			fmt.Println()
		}
		fmt.Println(f)
		lock := filepath.Join(
			filepath.Dir(f),
			"clean.gopher-drops-over",
		)
		t := time.Now()
		if _, err := os.Stat(lock); os.IsNotExist(err) {
			fmt.Println("File is protected")
		} else if err := wastebasket.Trash(f); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf(
				"Moved to system recycle bin (%s)\n",
				time.Since(t),
			)
		}
	}
}
