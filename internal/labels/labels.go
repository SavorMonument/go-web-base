package labels

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"generic.com/internal/models"
)

type LabelMapper struct {
	Labels   *map[string]map[string]string
	Language *models.Language
}

var WLOG = log.New(os.Stderr, "LABEL.WARNING\t", log.Ldate|log.Ltime)

func NewLabelMapper(path string, language *models.Language) (*LabelMapper, error) {

	allLabels := make(map[string]map[string]string)
	// root := "config/ui/"
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("Cannot load accessing a path %q: %v", path, err)
		} else if !info.IsDir() {
			// If it's a file open it
			f, err := os.Open(path)
			if err != nil {
				return fmt.Errorf("Error opening file: %q, %v", path, err)
			}
			defer f.Close()

			// Load it into a string map
			lang := strings.TrimSuffix(info.Name(), ".json")
			var labels map[string]string
			dec := json.NewDecoder(f)
			if err := dec.Decode(&labels); err != nil {
				return fmt.Errorf("Error decoding file: %q, %v", path, err)
			}
			// Put the string label map into the config
			allLabels[lang] = labels
		}
		return nil
	})
	if err != nil {
		fmt.Printf("error loading labels: %v", err)
	}

	return &LabelMapper{Labels: &allLabels, Language: language}, nil
}

func (l *LabelMapper) WithLanguage(language *models.Language) *LabelMapper {
	if _, ok := (*l.Labels)[language.Value]; !ok {
		WLOG.Printf("Tried to set unknown language %s", language)
		return l
	}
	return &LabelMapper{Labels: l.Labels, Language: language}
}

func (l *LabelMapper) GetLabel(label string) string {
	if l, ok := (*l.Labels)[l.Language.Value][label]; ok {
		return l
	} else {
		return label
	}
}
