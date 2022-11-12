package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"task_list/config"
	"task_list/route"
)

func main() {
	defer cmdRecover()
	r := route.Init()
	port := config.GlobalConfig.ServerConfig.ServerPort
	errChan := make(chan error, 1)

	// Start http server
	go func() {
		errChan <- r.Run(":" + port)
	}()
	// Graceful shutdown
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	<-errChan

	fmt.Println("http server shutdown...")
}

func cmdRecover() {
	if r := recover(); r != nil {
		var msg string
		for i := 2; ; i++ {
			_, file, line, ok := runtime.Caller(i)
			if !ok {
				break
			}
			msg = msg + fmt.Sprintf("%s:%d\n", file, line)
		}
		fmt.Printf("%s\n↧↧↧↧↧↧ PANIC ↧↧↧↧↧↧\n%s↥↥↥↥↥↥ PANIC ↥↥↥↥↥↥", r, msg)
	}
}
