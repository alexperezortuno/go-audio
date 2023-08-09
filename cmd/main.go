package main

import (
	"github.com/alexperezortuno/go-audio/internal/encode"
	"github.com/alexperezortuno/go-audio/internal/record"
	"github.com/alexperezortuno/go-audio/internal/tts"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "debug",
				Usage: "debug mode",
			},
			&cli.BoolFlag{
				Name:    "play",
				Aliases: []string{"p"},
				Usage:   "play audio file after recording",
			},
			&cli.BoolFlag{
				Name:    "encode",
				Aliases: []string{"e"},
				Usage:   "encode audio file with custom codec, bitrate and format",
			},
			&cli.BoolFlag{
				Name:    "tts",
				Aliases: []string{"t"},
				Usage:   "Option to generate text to speech",
			},
			&cli.BoolFlag{
				Name:  "record",
				Usage: "Option to generate text to speech",
			},
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Value:   "audio",
				Usage:   "Output file directory",
			},
			&cli.IntFlag{
				Name:    "duration",
				Aliases: []string{"d"},
				Value:   5,
				Usage:   "Duration of recording in seconds",
			},
			&cli.IntFlag{
				Name:    "sample-rate",
				Aliases: []string{"r"},
				Value:   44100,
				Usage:   "Sample rate of recording",
			},
			&cli.IntFlag{
				Name:    "channels",
				Aliases: []string{"c"},
				Value:   2,
				Usage:   "Number of channels",
			},
			&cli.StringFlag{
				Name:  "device",
				Value: "microphone",
				Usage: "Device to record from, available devices: microphone, board",
			},
			&cli.IntFlag{
				Name:    "frames-per-buffer",
				Aliases: []string{"f"},
				Value:   64,
				Usage:   "Frames per buffer",
			},
			&cli.StringFlag{
				Name:    "sentence",
				Aliases: []string{"s"},
				Value:   "Hello!",
				Usage:   "input text to encode to speech",
			},
			&cli.StringFlag{
				Name:    "language",
				Aliases: []string{"l"},
				Value:   "en",
				Usage:   "language to use for text to speech, available languages: en, es, pt",
			},
			&cli.StringFlag{
				Name:    "input-file",
				Aliases: []string{"i"},
				Value:   "",
				Usage:   "input file to encode",
			},
			&cli.StringFlag{
				Name:  "log_format",
				Value: "txt",
				Usage: "log format, available formats: txt, json",
			},
			&cli.StringFlag{
				Name:  "codec",
				Value: "pcm_s16le",
				Usage: "codec to use for audio file, example codecs: pcm_s16le, pcm_s24le, pcm_s32le, pcm_f32le, pcm_f64le, libmp3lame",
			},
			&cli.StringFlag{
				Name:  "bitrate",
				Value: "128k",
				Usage: "bitrate to use for audio file, example bitrates: 128k, 256k, 512k",
			},
			&cli.StringFlag{
				Name:  "format",
				Value: "wav",
				Usage: "format to use for audio file, example formats: wav, mp3, ogg",
			},
			&cli.StringFlag{
				Name:  "file-name",
				Value: "recording",
				Usage: "file name to use for audio file",
			},
		},
		Action: func(c *cli.Context) error {
			if c.Bool("encode") {
				encode.Start(c)
				os.Exit(0)
			}

			if c.Bool("record") {
				record.Start(c)
				os.Exit(0)
			}

			if c.Bool("tts") {
				tts.Start(c)
				os.Exit(0)
			}
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
