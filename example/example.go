package main

import (
	"github.com/jander/golog/logger"
	"log"
)

func main() {
	rotatingHandler := logger.NewRotatingHandler("./", "test.log", 4, 4*1024*1024)

	logger.SetHandlers([]logger.Handler{logger.Console, rotatingHandler})

	defer logger.Close()

	logger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	logger.SetLevel(logger.INFO)

	m := make(chan int)
	n := make(chan int)

	go func() {
		for {
			select {
			case v := <-m:
				logger.Warn("m", v)
			case v := <-n:
				logger.Info("n", v)
			}
		}
	}()

	for i := 0; i < 1000000; i++ {
		if i%2 == 0 {
			m <- i
		} else {
			n <- i
		}
	}

}
