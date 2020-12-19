package osutils

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/pantskun/commonutils/pathutils"
)

func ListenSystemSignalsWithCtx(ctx context.Context, cancel context.CancelFunc, signalChan chan os.Signal, signals ...os.Signal) {
	go func() {
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
	}()
}

// WriteToFile
// 将content写入filePath文件中.
func WriteToFile(filePath string, content []byte) error {
	file, err := CreateFile(filePath)
	if err != nil {
		return err
	}

	_, err = file.Write(content)
	if err != nil {
		return err
	}

	file.Close()

	return nil
}

// CreateFile
// 创建fp文件.
func CreateFile(fp string) (*os.File, error) {
	p := pathutils.GetParentPath(fp)
	if p != fp {
		err := os.MkdirAll(p, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}

	return os.Create(fp)
}
