package encode

import (
	"github.com/alexperezortuno/go-audio/kit/logger"
	"github.com/alexperezortuno/go-audio/kit/utils"
	"github.com/urfave/cli/v2"
	"os"
)

func Start(c *cli.Context) {
	var log = logger.GetLogger(c)
	var err error
	var fileName string
	f := c.String("input-file")

	if f == "" {
		log.Errorln("You must provide a file to convert.")
		os.Exit(1)
	}

	fileName, err = utils.ConvertMediaFile(f, c, false)
	if err != nil {
		log.Errorln("An error occurred while converting the file:", err)
		os.Exit(1)
	}
	err = utils.OpenFileWithDefaultMediaPlayer(fileName)
	if err != nil {
		log.Errorln("An error occurred when open the file: ", err)
		os.Exit(1)
	}
}
