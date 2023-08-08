# go-audio

### Requirements

```shell
sudo apt install portaudio19-dev
```


```shell
go mod init github.com/your-username/your-repo
go mod tidy
```

### Run

```shell
go run cmd/main.go
```

### Build

```shell
make build
```

### Examples

```shell
./build/go_audio -t -p -l "en" -s "How are you?"
```
    
```shell
./build/go_audio -t -p -l "es" -s "¿En que puedo ayudarte?"
```

```shell
./build/go_audio -t -p -l "pt" -s "Muito longe, no entanto, que a estrada é muito longa."
```

```shell
./build/go_audio -t -p --format "mp3" --codec "libmp3lame" -i "audio/file_name_to_mp3.wav" -e
```

