package internal

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func Run(connectionString, proxyServer string) {
	wg := sync.WaitGroup{}

	// 主 goroutine 结束时, 所有子 goroutine 也结束
	ancientContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 监听系统中断信号
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signals
		cancel()
	}()

	// 使用gin创建一个http服务器
	server := &http.Server{
		Addr:    connectionString,
		Handler: NewRouter(proxyServer),
	}

	// 启动服务监听和优雅退出，退出是直接shutdown服务器，不等待业务正常终止
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(errors.New("Failed to start server: " + err.Error()))
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ancientContext.Done()
		ctxExpired, timeoutCancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer timeoutCancel()
		if err := server.Shutdown(ctxExpired); err != nil {
			panic(errors.New("Failed to gracefully shutdown server: " + err.Error()))
		}
	}()
	fmt.Printf("Server listening on %s\n", connectionString)
	wg.Wait()
}
