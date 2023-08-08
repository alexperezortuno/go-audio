package structs

type Flags struct {
	Tts   bool `cli:"tts" cliAlt:"t" usage:"Generate text to speech"`
	Debug bool `env:"APP_DEBUG" cli:"debug" usage:"debug mode"`
}
