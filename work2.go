package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"

	"golang.org/x/sync/errgroup"
)

func StartHttpServer(srv *http.Server) error {
	http.HandleFunc("/hello", HelloServer)
	fmt.Println("start")
	err := srv.ListenAndServe()
	return err
}

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world\n")
}

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)

	group, errCtx := errgroup.WithContext(ctx)

	srv := &http.Server{Addr: ":9090"}

	group.Go(func() error {
		return StartHttpServer(srv)
	})

	group.Go(func() error {
		<-errCtx.Done()
		fmt.Println("stop")
		return srv.Shutdown(errCtx)
	})

	channel := make(chan os.Signal, 1)
	signal.Notify(channel)

	group.Go(func() error {
		for {
			select {
			case <-errCtx.Done():
				return errCtx.Err()
			case <-channel:
				cancel()
			}
		}
		return nil
	})

	if err := group.Wait(); err != nil {
		fmt.Println("group error:", err)
	}

	fmt.Println("all group done")
}
