package main

import (
	"encoding/binary"
	"fmt"
	"github.com/MarkKremer/microphone"
	"github.com/faiface/beep"
	"github.com/faiface/beep/wav"
	"github.com/gordonklaus/portaudio"
	"github.com/urfave/cli"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"
)

func recordFromMicrophone(output string,
	duration int,
	sampleRate int,
	channels int) {
	err := microphone.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer microphone.Terminate()

	stream, format, err := microphone.OpenDefaultStream(
		beep.SampleRate(sampleRate),
		channels)
	if err != nil {
		log.Fatal(err)
	}

	defer stream.Close()

	filename := "mic-" + output
	if !strings.HasSuffix(filename, ".wav") {
		filename += ".wav"
	}

	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)
	go func() {
		<-sig
		stream.Stop()
		stream.Close()
	}()

	stream.Start()

	if duration > 0 {
		fmt.Println("Recording for", duration, "seconds")

		go func() {
			for {
				select {
				case <-time.After(time.Duration(duration) * time.Second):
					stream.Stop()
					stream.Close()
					return
				}
			}
		}()
	}

	err = wav.Encode(f, stream, format)
	if err != nil {
		log.Fatal(err)
	}
}

func recordFromDevice(output string,
	duration int,
	sampleRate int,
	channels int,
	framesPerBuffer int) {
	fmt.Println("Recording.  Press Ctrl-C to stop.")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)

	fileName := "device-" + output
	if !strings.HasSuffix(fileName, ".wav") {
		fileName += ".wav"
	}
	f, err := os.Create(fileName)
	chk(err)

	// form chunk
	_, err = f.WriteString("FORM")
	chk(err)
	chk(binary.Write(f, binary.BigEndian, int32(0))) //total bytes
	_, err = f.WriteString("WAV")
	chk(err)

	// common chunk
	_, err = f.WriteString("COMM")
	chk(err)
	chk(binary.Write(f, binary.BigEndian, int32(18)))                  //size
	chk(binary.Write(f, binary.BigEndian, int16(1)))                   //channels
	chk(binary.Write(f, binary.BigEndian, int32(0)))                   //number of samples
	chk(binary.Write(f, binary.BigEndian, int16(32)))                  //bits per sample
	_, err = f.Write([]byte{0x40, 0x0e, 0xac, 0x44, 0, 0, 0, 0, 0, 0}) //80-bit sample rate 44100
	chk(err)

	// sound chunk
	_, err = f.WriteString("SSND")
	chk(err)
	chk(binary.Write(f, binary.BigEndian, int32(0))) //size
	chk(binary.Write(f, binary.BigEndian, int32(0))) //offset
	chk(binary.Write(f, binary.BigEndian, int32(0))) //block
	nSamples := 0
	defer func() {
		// fill in missing sizes
		totalBytes := 4 + 8 + 18 + 8 + 8 + 4*nSamples
		_, err = f.Seek(4, 0)
		chk(err)
		chk(binary.Write(f, binary.BigEndian, int32(totalBytes)))
		_, err = f.Seek(22, 0)
		chk(err)
		chk(binary.Write(f, binary.BigEndian, int32(nSamples)))
		_, err = f.Seek(42, 0)
		chk(err)
		chk(binary.Write(f, binary.BigEndian, int32(4*nSamples+8)))
		chk(f.Close())
	}()

	portaudio.Initialize()
	defer portaudio.Terminate()
	in := make([]int32, framesPerBuffer)
	stream, err := portaudio.OpenDefaultStream(1, 0, 44100, len(in), in)
	chk(err)
	defer stream.Close()

	chk(stream.Start())
	for {
		chk(stream.Read())
		chk(binary.Write(f, binary.BigEndian, in))
		nSamples += len(in)
		select {
		case <-sig:
			return
		default:
		}
	}
	chk(stream.Stop())
}

func chk(err error) {
	if err != nil {
		log.Fatal(err)
		//panic(err)
	}
}

func start(c *cli.Context) {
	fmt.Println("Recording. Press Ctrl-C to stop.")

	if c.String("device") == "microphone" {
		recordFromMicrophone(c.String("output"),
			c.Int("duration"),
			c.Int("sample-rate"),
			c.Int("channels"))
	}

	if c.String("device") == "device" {
		recordFromDevice(c.String("output"),
			c.Int("duration"),
			c.Int("sample-rate"),
			c.Int("channels"),
			c.Int("frames-per-buffer"))
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "recorder"
	app.Usage = "Record audio"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "output, o",
			Value: "recording",
			Usage: "Output file name",
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
			Value: 1,
			Usage: "Number of channels",
		},
		cli.StringFlag{
			Name:  "device, i",
			Value: "microphone",
			Usage: "Device to record from",
		},
		cli.IntFlag{
			Name:  "frames-per-buffer, f",
			Value: 64,
			Usage: "Frames per buffer",
		},
	}

	app.Action = start

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
