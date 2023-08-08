package tts

import (
	"github.com/alexperezortuno/go-audio/kit/logger"
	"github.com/alexperezortuno/go-audio/platform/constant"
	"github.com/alexperezortuno/go-audio/platform/handler"
	"github.com/hegedustibor/htgo-tts"
	"github.com/hegedustibor/htgo-tts/voices"
	"github.com/urfave/cli"
)

func Start(c *cli.Context) {
	var log = logger.GetLogger(c)
	log.Debugln("Generating TTS...")
	var l = c.String("language")
	l = constant.GetLanguage(l)
	if l == "en" {
		speak(c, voices.English, "audio")
	}
	if l == "es" {
		speak(c, voices.Spanish, "audio")
	}
	if l == "pt" {
		speak(c, voices.Portuguese, "audio")
	}

	log.Debugln("TTS generated")
}

func speak(c *cli.Context, language string, folder string) {
	var log = logger.GetLogger(c)
	speech := htgotts.Speech{Folder: c.String("output"), Language: language, Handler: &handler.AdditionalProcess{Context: *c}}
	err := speech.Speak(c.String("sentence"))
	if err != nil {
		log.Errorln(err)
		return
	}
}
