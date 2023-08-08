package constant

import "strings"

const (
	Es = "es"
	En = "en"
	Pt = "pt"
)

var Languages = []string{Es, En, Pt}

func GetLanguage(language string) string {
	language = strings.ToLower(language)
	for _, lang := range Languages {
		if lang == language {
			return lang
		}
	}
	return En
}
