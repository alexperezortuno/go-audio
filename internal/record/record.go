package record

import (
	"encoding/binary"
	"fmt"
	"github.com/MarkKremer/microphone"
	"github.com/alexperezortuno/go-audio/kit/logger"
	"github.com/alexperezortuno/go-audio/kit/utils"
	"github.com/faiface/beep"
	"github.com/faiface/beep/wav"
	"github.com/gordonklaus/portaudio"
	"github.com/urfave/cli/v2"
	"os"
	"os/signal"
	"time"
)

func recordFromMicrophone(c *cli.Context) {
	var log = logger.GetLogger(c)
	var output = c.String("output")
	var duration int = c.Int("duration")
	var sampleRate int = c.Int("sample-rate")
	var channels int = c.Int("channels")
	err := microphone.Init()
	chk(err, c)
	defer func() {
		err := microphone.Terminate()
		chk(err, c)
	}()

	stream, format, err := microphone.OpenDefaultStream(
		beep.SampleRate(sampleRate),
		channels)
	chk(err, c)

	defer func(stream *microphone.Streamer) {
		err := stream.Close()
		chk(err, c)
	}(stream)

	filename := output + "/mic-" + utils.GenerateRandomString(40) + "." + c.String("format")

	f, err := os.Create(filename)
	chk(err, c)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)
	go func() {
		<-sig
		err := stream.Stop()
		chk(err, c)
		err = stream.Close()
		chk(err, c)
		return
	}()

	err = stream.Start()
	if err != nil {
		return
	}

	if duration > 0 {
		log.Debugln("Recording for", duration, "seconds")

		go func() {
			for {
				select {
				case <-time.After(time.Duration(duration) * time.Second):
					err := stream.Stop()
					chk(err, c)
					err = stream.Close()
					chk(err, c)
					return
				}
			}
		}()
	}

	err = wav.Encode(f, stream, format)
	chk(err, c)

	if c.Bool("encode") {
		_, err = utils.ConvertMediaFile(filename, c, false)
		chk(err, c)
	}

	if c.Bool("play") {
		err := utils.OpenFileWithDefaultMediaPlayer(filename)
		chk(err, c)
	}
}

func recordFromDevice(c *cli.Context) {
	var log = logger.GetLogger(c)
	var output = c.String("output")
	var sampleRate = c.Int("sample-rate")
	var framesPerBuffer = c.Int("frames-per-buffer")
	log.Println("Recording.  Press Ctrl-C to stop.")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)

	fileName := output + "/device-" + utils.GenerateRandomString(40) + "." + c.String("format")
	f, err := os.Create(fileName)
	chk(err, c)

	// form chunk
	_, err = f.WriteString("FORM")
	chk(err, c)
	chk(binary.Write(f, binary.BigEndian, int32(0)), c) //total bytes
	_, err = f.WriteString("WAV")
	chk(err, c)

	// common chunk
	_, err = f.WriteString("COMM")
	chk(err, c)
	chk(binary.Write(f, binary.BigEndian, int32(18)), c)               //size
	chk(binary.Write(f, binary.BigEndian, int16(1)), c)                //channels
	chk(binary.Write(f, binary.BigEndian, int32(0)), c)                //number of samples
	chk(binary.Write(f, binary.BigEndian, int16(32)), c)               //bits per sample
	_, err = f.Write([]byte{0x40, 0x0e, 0xac, 0x44, 0, 0, 0, 0, 0, 0}) //80-bit sample rate 44100
	chk(err, c)

	// sound chunk
	_, err = f.WriteString("SSND")
	chk(err, c)
	chk(binary.Write(f, binary.BigEndian, int32(0)), c) //size
	chk(binary.Write(f, binary.BigEndian, int32(0)), c) //offset
	chk(binary.Write(f, binary.BigEndian, int32(0)), c) //block
	nSamples := 0
	defer func() {
		// fill in missing sizes
		totalBytes := 4 + 8 + 18 + 8 + 8 + 4*nSamples
		_, err = f.Seek(4, 0)
		chk(err, c)
		chk(binary.Write(f, binary.BigEndian, int32(totalBytes)), c)
		_, err = f.Seek(22, 0)
		chk(err, c)
		chk(binary.Write(f, binary.BigEndian, int32(nSamples)), c)
		_, err = f.Seek(42, 0)
		chk(err, c)
		chk(binary.Write(f, binary.BigEndian, int32(4*nSamples+8)), c)
		chk(f.Close(), c)
	}()

	err = portaudio.Initialize()
	if err != nil {
		log.Errorln(err)
		return
	}
	defer func() {
		err := portaudio.Terminate()
		if err != nil {
			log.Errorln(err)
		}
	}()
	in := make([]int32, framesPerBuffer)
	stream, err := portaudio.OpenDefaultStream(1, 0, float64(sampleRate), len(in), in)
	chk(err, c)
	defer func(stream *portaudio.Stream) {
		err := stream.Close()
		chk(err, c)
	}(stream)

	chk(stream.Start(), c)
	for {
		chk(stream.Read(), c)
		chk(binary.Write(f, binary.BigEndian, in), c)
		nSamples += len(in)
		select {
		case <-sig:
			return
		default:
		}
	}
	chk(stream.Stop(), c)
}

func chk(err error, c *cli.Context) {
	log := logger.GetLogger(c)
	if err != nil {
		log.Errorln(err)
	}
}

func Start(c *cli.Context) {
	fmt.Println("Recording. Press Ctrl-C to stop.")

	if c.String("device") == "microphone" {
		recordFromMicrophone(c)
	}

	if c.String("device") == "board" {
		recordFromDevice(c)
	}
}
