package main

import (
	"fmt"
	"github.com/Bios-Marcel/wastebasket"
	"gopher-drops-over/utils"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"time"
)

func main() {
	var (
		ec      atomic.Value
		rawLogs []string
		logs    = make(chan string)
	)
	ec.Store(0)
	fmt.Println("")
	items := os.Args[1:]
	for _, f := range items {
		log := []string{f}
		dir := filepath.Dir(f)
		lock := filepath.Join(
			dir,
			"clean.gopher-drops-over",
		)
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
				err = nil
				if err = os.MkdirAll(dir, 0750); err == nil {
					err = os.WriteFile(lock, []byte{}, 0644)
				}
				if err != nil {
					ec.Store(ec.Load().(int) + 1)
					log = append(log, fmt.Sprint(err))
				}
			}
			rawLog := strings.Join(append(log, ""), "\n")
			fmt.Println(rawLog)
			logs <- rawLog
		}()
	}
	for range items {
		rawLogs = append(rawLogs, <-logs)
	}
	vEc := ec.Load().(int)
	fmt.Println("Error count:", vEc)
	if vEc > 0 {
		utils.WebError(vEc, strings.Join(rawLogs, "\n\n"))
	}
}
