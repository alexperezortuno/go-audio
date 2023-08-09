# go-audio

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