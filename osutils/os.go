package osutils

import (
	"context"
	"log"
	"os"
	"os/signal"
)

func ListenSystemSignalsWithCtx(ctx context.Context, cancel context.CancelFunc, signalChan chan os.Signal, signals ...os.Signal) {
	signal.Notify(signalChan)

	for {
		select {
		case <-ctx.Done():
			return
		case acceptS := <-signalChan:
			for _, listenS := range signals {
				if acceptS == listenS {
					log.Println("get signal: ", acceptS)
					cancel()

					return
				}
			}
		}
	}
}
