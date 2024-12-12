package models

type Language struct {
	Label string
	Value string
}

var (
	EN_LANG = Language{"English", "en"}
	FR_LANG = Language{"French", "fr"}
)

func LanguageFromValue(value string) (Language, bool) {
	for _, lang := range GetLanguages() {
		if lang.Value == value {
			return lang, true
		}
	}
	return Language{}, false
}

func GetLanguages() []Language {
	return []Language{EN_LANG, FR_LANG}
}
