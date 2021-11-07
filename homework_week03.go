package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"

	"golang.org/x/sync/errgroup"
)

/**
基于 errgroup 实现一个 http server 的启动和关闭 ，
以及 linux signal 信号的注册和处理，
要保证能够一个退出，全部注销退出。
*/
func main() {
	//用到了 errgroup 一个出错，其余取消的能力
	group, ctx := errgroup.WithContext(context.Background())
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(response, "week03, homework")
	})
	server := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}
	serverOut := make(chan int)
	//group1 启动Server
	group.Go(func() error {
		return server.ListenAndServe()
	})
	//group2 关闭Server
	group.Go(func() error {
		select {
		case <-serverOut:
			fmt.Println("server quit!")
		case <-ctx.Done():
			fmt.Println("error group exit")
		}
		return server.Shutdown(ctx)
	})
	//group3 signal注册
	group.Go(func() error {
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-ctx.Done():
			return ctx.Err()
		case signalRes := <-quit:
			return errors.Errorf("new os exit: %v", signalRes)
		}
	})
	if err := group.Wait(); err != nil {
		fmt.Println("http server exception occurred, exception info: ", err)
	}
}
