package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/neilotoole/errgroup"
	"log"
	"net/http"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Golang"))
}
func main() {
	http.HandleFunc("/", Index)
	srv := &http.Server{Addr: ":8080"}
	g, ctx1 := errgroup.WithContext(context.Background())
	g.Go(AppServer1)
	go func(ctx1 context.Context) {
		select {
		case <-ctx1.Done():
			if err := srv.Shutdown(ctx1); err != nil {
				log.Println("Server Shutdown Error")
				return
			}
			log.Println("Server Stop")
		}
	}(ctx1)
	if err := srv.ListenAndServe(); err != nil {
		log.Printf("error:%s", "Server Error")
	}
}
func AppServer1() error {
	for i := 0; i < 10; i++ {

		time.Sleep(time.Second * 1)
		fmt.Println(i)
	}
	return errors.New("Server Error")
}
