package utils

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

func OpenFileWithDefaultMediaPlayer(file string) error {
	var err error

	switch runtime.GOOS {
	case "darwin":
		err = exec.Command("open", file).Start()
	case "linux":
		err = exec.Command("xdg-open", file).Start()
	case "windows":
		err = exec.Command("cmd", "/c", "start", file).Start()
	default:
		err = fmt.Errorf("unsupported platform %q", runtime.GOOS)
	}

	return err
}

func ConvertMediaFile(fileName string, context *cli.Context, remove bool) (string, error) {
	var err error
	var f string
	f = GenerateRandomString(40)
	f = context.String("language") + "_" + f + "." + context.String("format")
	filePath := context.String("output") + "/" + f

	cmd := exec.Command("ffmpeg",
		"-i",
		fileName,
		"-codec:a",
		context.String("codec"),
		"-b:a",
		context.String("bitrate"),
		"-ar",
		strconv.Itoa(context.Int("sample_rate")),
		"-ac",
		context.String("channels"),
		filePath,
		"-y",
	)

	e := cmd.Run()

	if e != nil {
		return "", e
	}

	if remove {
		err = os.Remove(fileName)
		if err != nil {
			fmt.Println("An error occurred when remove file", err)
			return filePath, err
		}
	}

	return filePath, err
}
