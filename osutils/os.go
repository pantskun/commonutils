package osutils

import (
	"context"
	"fmt"
	"os"
	"os/signal"
)

func ListenSystemSignal(ctrlBreakChan chan os.Signal, ctx context.Context, cancel context.CancelFunc) {
	signal.Notify(ctrlBreakChan)

	for {
		select {
		case <-ctx.Done():
			return
		case s := <-ctrlBreakChan:
			if s == os.Interrupt {
				fmt.Println("got signal:", s)
				cancel()

				return
			}
		}
	}
}
