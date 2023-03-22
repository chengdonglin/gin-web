package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(r *gin.Engine, srvName string, addr string, stop func()) {
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	go func() {
		log.Printf("%s web server running is %s \n", srvName, srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln(err)
		}
	}()
	quit := make(chan os.Signal)
	// SIGIN 用户发送 INTR 字符（ctrl + c）触发
	// SIGTERM 结束程序，可以被捕获，阻塞或忽略
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Printf("shutting Down project %s ....\n", srvName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if stop != nil {
		stop()
	}
	err := srv.Shutdown(ctx)
	if err != nil {
		log.Printf("%s server shutdown, cause by : %v \n", srvName, err)
	}
	select {
	case <-ctx.Done():
		log.Println("关闭超时")
	}
	log.Printf("%s web server stop success...\n", srvName)
}
