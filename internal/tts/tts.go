package tts

import (
	"github.com/alexperezortuno/go-audio/kit/logger"
	"github.com/alexperezortuno/go-audio/platform/constant"
	"github.com/alexperezortuno/go-audio/platform/handler"
	"github.com/hegedustibor/htgo-tts"
	"github.com/hegedustibor/htgo-tts/voices"
	"github.com/urfave/cli/v2"
)

func Start(c *cli.Context) {
	var log = logger.GetLogger(c)
	log.Debugln("Generating TTS...")
	var l = c.String("language")
	l = constant.GetLanguage(l)
	if l == "en" {
		speak(c, voices.English)
	}
	if l == "es" {
		speak(c, voices.Spanish)
	}
	if l == "pt" {
		speak(c, voices.Portuguese)
	}

	log.Debugln("TTS generated")
}

func speak(c *cli.Context, language string) {
	var log = logger.GetLogger(c)
	speech := htgotts.Speech{Folder: c.String("output"), Language: language, Handler: &handler.AdditionalProcess{Context: *c}}
	err := speech.Speak(c.String("sentence"))
	if err != nil {
		log.Errorln(err)
		return
	}
}
