package main

import (
	"github.com/alexperezortuno/go-audio/internal/encode"
	"github.com/alexperezortuno/go-audio/internal/record"
	"github.com/alexperezortuno/go-audio/internal/tts"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "recorder"
	app.Usage = "Record audio"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "output, o",
			Value: "audio",
			Usage: "Output file directory",
		},
		cli.IntFlag{
			Name:  "duration, d",
			Value: 0,
			Usage: "Duration of recording in seconds",
		},
		cli.IntFlag{
			Name:  "sample-rate, r",
			Value: 44100,
			Usage: "Sample rate of recording",
		},
		cli.IntFlag{
			Name:  "channels, c",
			Value: 2,
			Usage: "Number of channels",
		},
		cli.StringFlag{
			Name:  "device",
			Value: "",
			Usage: "Device to record from, available devices: microphone, board",
		},
		cli.IntFlag{
			Name:  "frames-per-buffer, f",
			Value: 64,
			Usage: "Frames per buffer",
		},
		cli.BoolFlag{
			Name:  "tts, t",
			Usage: "Option to generate text to speech",
		},
		cli.StringFlag{
			Name:  "sentence, s",
			Value: "Hello!",
			Usage: "input text to encode to speech",
		},
		cli.StringFlag{
			Name:  "language, l",
			Value: "en",
			Usage: "language to use for text to speech, available languages: en, es, pt",
		},
		cli.BoolFlag{
			Name:  "debug",
			Usage: "debug mode",
		},
		cli.StringFlag{
			Name:  "log_format",
			Value: "txt",
			Usage: "log format, available formats: txt, json",
		},
		cli.BoolFlag{
			Name:  "play, p",
			Usage: "play audio file after recording",
		},
		cli.StringFlag{
			Name:  "codec",
			Value: "pcm_s16le",
			Usage: "codec to use for audio file, example codecs: pcm_s16le, pcm_s24le, pcm_s32le, pcm_f32le, pcm_f64le, libmp3lame",
		},
		cli.StringFlag{
			Name:  "bitrate",
			Value: "128k",
			Usage: "bitrate to use for audio file, example bitrates: 128k, 256k, 512k",
		},
		cli.StringFlag{
			Name:  "format",
			Value: "wav",
			Usage: "format to use for audio file, example formats: wav, mp3, ogg",
		},
		cli.StringFlag{
			Name:  "file-name",
			Value: "recording",
			Usage: "file name to use for audio file",
		},
		cli.BoolFlag{
			Name:  "encode, e",
			Usage: "encode audio file with custom codec, bitrate and format",
		},
		cli.StringFlag{
			Name:  "input-file, i",
			Value: "",
			Usage: "input file to encode",
		},
	}

	app.Action = func(c *cli.Context) {
		if c.Bool("encode") {
			encode.Start(c)
			os.Exit(0)
		}

		if c.String("device") == "microphone" || c.String("device") == "board" {
			record.Start(c)
			os.Exit(0)
		}

		if c.Bool("tts") {
			tts.Start(c)
			os.Exit(0)
		}
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
