package main

import (
	"fmt"
	"github.com/Bios-Marcel/wastebasket"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var (
		ec atomic.Value
		wg sync.WaitGroup
	)
	ec.Store(0)
	fmt.Println()
	for _, f := range os.Args[1:] {
		log := []string{f}
		dir := filepath.Dir(f)
		lock := filepath.Join(
			dir,
			"clean.gopher-drops-over",
		)
		wg.Add(1)
		f := f
		go func() {
			t := time.Now()
			if _, err := os.Stat(f); os.IsNotExist(err) {
				log = append(log, "Already cleaned")
			} else if _, err := os.Stat(lock); os.IsNotExist(err) {
				log = append(log, "File is protected")
			} else if err := wastebasket.Trash(dir); err != nil {
				ec.Store(ec.Load().(int) + 1)
				log = append(log, fmt.Sprint(err))
			} else {
				log = append(log, fmt.Sprintf(
					"Moved to system recycle bin (%s)",
					time.Since(t),
				))
			}
			fmt.Println(strings.Join(append(log, ""), "\n"))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Error count:", ec.Load())
}
