package utils

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"log"
	"sync"
)

type LocalizerService struct {
	Bundle *i18n.Bundle
}

var (
	localizerInstance *LocalizerService
	once              sync.Once
)

func GetLocalizerInstance() *LocalizerService {
	once.Do(func() {
		bundle := i18n.NewBundle(language.English)
		bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

		_, err := bundle.LoadMessageFile("internal/locales/translation.en.json")
		if err != nil {
			log.Println("Error loading translation.en.json")
		}

		_, err = bundle.LoadMessageFile("internal/locales/translation.tr.json")
		if err != nil {
			log.Println("Error loading translation.tr.json")
		}

		localizerInstance = &LocalizerService{Bundle: bundle}
	})
	return localizerInstance
}

func (s *LocalizerService) Localizer(lang string) *i18n.Localizer {
	return i18n.NewLocalizer(s.Bundle, lang)
}

func Translate(context *gin.Context, key string, values map[string]string) string {
	lang, exists := context.Get("lang")
	if !exists {
		lang = ""
	}

	localizer := GetLocalizerInstance().Localizer(lang.(string))

	if values != nil && len(values) > 0 {
		return localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID:    key,
			TemplateData: values,
		})
	}

	return localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: key,
	})
}
