package language

import (
	"io/ioutil"
	"log"
	"encoding/json"
)

type Lang struct {
	Response map[string]string
}

func New(languageType interface{}) *Lang {
	var selectLang string

	switch v := languageType.(type) {
		case string:
			selectLang = v
		default:
			selectLang = "zhTw"
	}

	readFile, err := ioutil.ReadFile("language/" + selectLang + ".json")
	
	if err != nil {
		log.Panicln(err)
	}

	var parseLang Lang

	err = json.Unmarshal(readFile, &parseLang)

	if err != nil {
		log.Panicln(err)
	}

	return &parseLang
}

