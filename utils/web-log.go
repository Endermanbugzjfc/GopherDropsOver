package utils

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/skratchdot/open-golang/open"
	"net/http"
	"time"
)

//go:embed web-log.html
var html []byte

func WebError(ec int, log string) {
	fmt.Println("Starting web log-monitor...")

	const port = ":8080"
	srv := &http.Server{Addr: port}

	t, stop := time.NewTimer(time.Second*60), make(chan struct{})
	go func() {
		fmt.Println(t.C)
		select {
		case <-t.C:
		case <-stop:
			t.Stop()
		}
		_ = srv.Shutdown(context.Background())
	}()

	srv.Handler = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.String() == "/ok" {
			close(stop)
			return
		}
		if _, err := fmt.Fprintf(
			writer,
			string(html),
			ec,
			"https://github.com/Endermanbugzjfc/gopher-drops-over/issues",
			log,
		); err != nil {
			fmt.Println(err)
		}
	})

	_ = open.Start("http://localhost" + port)

	if err := srv.ListenAndServe(); err != nil {
		fmt.Println("Stopping web log-monitor...")
	}
}
