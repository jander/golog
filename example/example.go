package main

import (
	"bufio"
	"github.com/jander/golog/logger"
	"log"
	"os"
)

// This will run untill press enter key

func main() {
	rotatingHandler := logger.NewRotatingHandler(".", "test.log", 4, 4*1024*1024)

	logger.SetHandlers(logger.Console, rotatingHandler)

	defer logger.Close()

	logger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	//logger.DEBUG is default, nosense here.
	logger.SetLevel(logger.DEBUG)

	for i := 0; i < 10; i++ {
		go func(num int) {
			count := 0
			for {
				logger.Debug("log", num, "-", count)
				count++
			}
		}(i)
	}

	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}
