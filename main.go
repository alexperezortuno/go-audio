package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gordonklaus/portaudio"
	"github.com/urfave/cli"
)

var (
	recording bool
	file      *os.File
)

func startRecording(c *cli.Context) {
	portaudio.Initialize()
	defer portaudio.Terminate()

	// Open output file
	f, err := os.Create(c.String("output"))
	if err != nil {
		fmt.Println(err)
		return
	}
	file = f
	defer f.Close()

	// Open default input stream
	in := make([]byte, 64)
	inputStream, err := portaudio.OpenDefaultStream(1, 0, 44100, len(in), in)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer inputStream.Close()

	// Start input stream
	if err = inputStream.Start(); err != nil {
		fmt.Println(err)
		return
	}
	recording = true
	timeOut := time.NewTimer(time.Duration(c.Int("duration")) * time.Second)
	defer timeOut.Stop()
	for {
		select {
		case <-timeOut.C:
			stopRecording()
			return
		default:
			if !recording {
				return
			}
			// Read audio from input
			if err = inputStream.Read(); err != nil {
				fmt.Println(err)
				return
			}
			// Write audio to file
			n, err := file.Write(in)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Printf("Wrote %d bytes to file\n", n)
		}
	}
}

func stopRecording() {
	recording = false
	file.Close()
}

func main() {
	app := cli.NewApp()
	app.Name = "recorder"
	app.Usage = "Record audio"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "output, o",
			Value: "recording.wav",
			Usage: "Output file name",
		},
		cli.IntFlag{
			Name:  "duration, d",
			Value: 60,
			Usage: "Duration of recording in seconds",
		},
	}
	app.Action = startRecording

	app.Run(os.Args)
}
