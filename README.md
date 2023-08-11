# go-audio

### CLI Options

| Flag                    | Default    | Environment Variable         | Description                                                                         |
|-------------------------|------------|------------------------------|-------------------------------------------------------------------------------------|
| --help, -h              |            |                              | show help                                                                           |
| --debug                 | false      | OPEN_AUDIO_DEBUG             | Debug mode                                                                          |
| -p, --play              | false      | OPEN_AUDIO_PLAY              | play audio file in default player                                                   |
| -e, --encode            | false      | OPEN_AUDIO_ENCODE            | encode audio file with custom codec, bitrate and format                             |
| -t, --tts               | false      | OPEN_AUDIO_TTS               | Option to generate text to speech                                                   |
| --record                | false      | OPEN_AUDIO_RECORD            | record audio from microphone or board                                               |
| -o, --output            | audio      | OPEN_AUDIO_OUTPUT            | output file directory                                                               |
| -d, --duration          | 5          | OPEN_AUDIO_DURATION          | duration of recording in seconds                                                    |
| -r, --sample-rate       | 44100      | OPEN_AUDIO_SAMPLE_RATE       | sample rate to use for recording audio                                              |
| -c, --channels          | 2          | OPEN_AUDIO_CHANNELS          | number of channels to use for recording audio                                       |
| --device                | microphone | OPEN_AUDIO_DEVICE            | device to record from, available devices: microphone, board                         |
| -f, --frames-per-buffer | 64         | OPEN_AUDIO_FRAMES_PER_BUFFER | frames per buffer for recording audio                                               |
| -s, --sentence          | Hello!     | OPEN_AUDIO_SENTENCE          | input text to encode to speech                                                      |
| -l, --language          | en         | OPEN_AUDIO_LANGUAGE          | language to use for text to speech, available languages: en, es, pt, fr, it, ru, de |
| -i, --input-file        |            | OPEN_AUDIO_INPUT_FILE        | input file to encode                                                                |
| --log_format            | txt        | OPEN_AUDIO_LOG_FORMAT        | log format, available formats: txt, json                                            |
| --codec                 | pcm_s16le  | OPEN_AUDIO_CODEC             | codec to use for audio file, example codecs: pcm_s16le, pcm_s24le                   |
| --bitrate               | 128k       | OPEN_AUDIO_BITRATE           | bitrate to use for audio file, example bitrates: 128k, 256k, 512k                   |
| --format                | wav        | OPEN_AUDIO_FORMAT            | format to use for audio file, example formats: wav, mp3, ogg                        |
| --file-name             | recording  | OPEN_AUDIO_FILE_NAME         | file name to use for audio file                                                     |

### Requirements

```bash
sudo apt install portaudio19-dev
```

```bash
brew install ffmpeg       # macOS
sudo apt-get install ffmpeg  # Linux
```

```bash
go mod init github.com/your-username/your-repo
go mod tidy
```

### Run

```bash
go run cmd/main.go
```

### Build

```bash
make build
```

### Examples

```bash
./build/go_audio -t -p -l "en" -s "How are you?"
```

```bash
./build/go_audio -t -p -l "es" -s "¿En que puedo ayudarte?"
```

```bash
./build/go_audio -t -p -l "pt" -s "Muito longe, no entanto, que a estrada é muito longa."
```

```bash
./build/go_audio -t -p --format "mp3" --codec "libmp3lame" -i "audio/file_name_to_mp3.wav" -e
```

### Test

```bash
make test
```