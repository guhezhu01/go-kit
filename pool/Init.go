package pool

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	go func() {
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGINT)
		for {
			select {
			case <-quit:
			default:
				log.Print(" go-kit init")
				time.Sleep(100)

			}

		}

	}()

}
