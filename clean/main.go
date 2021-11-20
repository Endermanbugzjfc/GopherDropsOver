package main

import (
	"fmt"
	"github.com/Bios-Marcel/wastebasket"
	"gopher-drops-over/utils"
	"os"
	"path/filepath"
)

func main() {
	for _, f := range os.Args[1:] {
		lock := filepath.Join(
			filepath.Dir(f),
			string(utils.DropOver+"-clean.lock"),
		)
		if _, err := os.Stat(lock); os.IsNotExist(err) {
			fmt.Println()
		} else if err := wastebasket.Trash(f); err != nil {
			e(f, err)
		}
	}
}

func e(f string, err error) {
	fmt.Println("Cannot clean \"", f, "\": ", err)
}
