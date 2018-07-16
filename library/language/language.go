package language

import (
    "io/ioutil"
    "log"
    "encoding/json"
)

type Lang struct {
    Response map[string]string
}

func New(languageType string) *Lang {

    // Load language json file
    readFile, err := ioutil.ReadFile("language/" + languageType + ".json")
    
    if err != nil {
        log.Panicln(err)
    }

    // Parse json file
    var parseLang Lang

    err = json.Unmarshal(readFile, &parseLang)

    if err != nil {
        log.Panicln(err)
    }

    return &parseLang
}

