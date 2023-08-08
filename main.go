package main

import (
	"github.com/alexperezortuno/go-audio/core/commons"
	"github.com/alexperezortuno/go-audio/core/commons/structs"
	"github.com/alexperezortuno/go-audio/core/tts"
	"github.com/pteich/configstruct"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var logger = commons.GetLogger()

	conf := structs.Flags{
		Tts:   false,
		Debug: false,
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		os.Exit(0)
	}()

	cmd := configstruct.NewCommand(
		"Go Audio",
		"CLI tool to record and play audio",
		&conf,
		func(c *configstruct.Command, cfg interface{}) error {
			logger.Info(`Starting...`)
			//if conf.Debug {
			//}

			if conf.Tts {
				tts.Start()
			}
			return nil
		},
	)

	err := cmd.ParseAndRun(os.Args)
	if err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
