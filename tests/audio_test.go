package main

import (
	"github.com/alexperezortuno/go-audio/kit/utils"
	"github.com/urfave/cli/v2"
	"os"
	"strconv"
	"testing"
)

func TestConvertMediaFile(t *testing.T) {
	context := cli.NewContext(nil, nil, nil)
	var err error
	err = context.Set("language", "en")
	if err != nil {
		return
	}
	err = context.Set("format", "mp3")
	if err != nil {
		return
	}
	err = context.Set("output", "audio")
	if err != nil {
		return
	}
	err = context.Set("codec", "libmp3lame")
	if err != nil {
		return
	}
	err = context.Set("bitrate", "128k")
	if err != nil {
		return
	}
	err = context.Set("sample_rate", strconv.Itoa(44100))
	if err != nil {
		return
	}
	err = context.Set("channels", "2")
	if err != nil {
		return
	}

	fileName := "example.wav"

	filePath, err := utils.ConvertMediaFile(fileName, context, false)
	if err != nil {
		t.Errorf("Error al convertir el archivo de medios: %s", err.Error())
	}

	// Verify if file exist
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Errorf("El archivo de salida no existe: %s", filePath)
	}

	// Remove file after test
	err = os.Remove(filePath)
	if err != nil {
		t.Errorf("Error al eliminar el archivo de salida: %s", err.Error())
	}
}

func TestOpenFileWithDefaultMediaPlayer(t *testing.T) {
	context := cli.NewContext(nil, nil, nil)
	var err error
	err = context.Set("language", "en")
	if err != nil {
		return
	}
	err = context.Set("format", "mp3")
	if err != nil {
		return
	}
	err = context.Set("output", "audio")
	if err != nil {
		return
	}
	err = context.Set("codec", "libmp3lame")
	if err != nil {
		return
	}
	err = context.Set("bitrate", "128k")
	if err != nil {
		return
	}
	err = context.Set("sample_rate", strconv.Itoa(44100))
	if err != nil {
		return
	}
	err = context.Set("channels", "1")
	if err != nil {
		return
	}

	fileName := "example.wav"

	filePath, err := utils.ConvertMediaFile(fileName, context, false)
	if err != nil {
		t.Errorf("Error al convertir el archivo de medios: %s", err.Error())
	}

	err = utils.OpenFileWithDefaultMediaPlayer(filePath)
	if err != nil {
		return
	}

	// Verify if file exist
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Errorf("El archivo de salida no existe: %s", filePath)
	}

	// Remove file after test
	err = os.Remove(filePath)
	if err != nil {
		t.Errorf("Error al eliminar el archivo de salida: %s", err.Error())
	}
}
