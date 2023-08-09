package handler

import (
	"fmt"
	"github.com/alexperezortuno/go-audio/kit/logger"
	"github.com/alexperezortuno/go-audio/kit/utils"
	"github.com/urfave/cli/v2"
	"os"
	"time"
)

type AdditionalProcess struct {
	Context cli.Context
}

func (a AdditionalProcess) Play(fileName string) error {
	var log = logger.GetLogger(&a.Context)
	var err error
	if a.Context.Bool("play") {
		log.Debugln("Playing TTS...")
		time.Sleep(2 * time.Second)
		fileName, err = utils.ConvertMediaFile(fileName, &a.Context, true)
		if err != nil {
			fmt.Println("An error occurred while converting the file:", err)
			os.Exit(1)
		}
		err = utils.OpenFileWithDefaultMediaPlayer(fileName)
		if err != nil {
			fmt.Println("An error occurred when open the file: ", err)
			os.Exit(1)
		}
	}
	return nil
}
