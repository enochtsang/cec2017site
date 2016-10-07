package main

import (
    "gopkg.in/yaml.v2"
    "html/template"
    "io/ioutil"
    "net/http"
    "os"
    "path"
    "path/filepath"
)

// var log defined in cec2017site.go

type Competition struct {
    ID    string
    Title string
    Blurb string
}

func loadCompetitionsData(lang string) []Competition {
    dataFilePath := path.Join(filepath.Dir(os.Args[0]), "data", lang, "competitions.yml")
    dataFile, err := ioutil.ReadFile(dataFilePath)
    check(err, true)
    var data []Competition
    err = yaml.Unmarshal(dataFile, &data)
    check(err, true)
    return data
}

func competitions(w http.ResponseWriter, r *http.Request, data []Competition) {
    log.Debug("Loading competitions page for " + r.RemoteAddr)
    t := template.Must(template.ParseFiles(
        absPath("templates/base.html"),
        absPath("templates/competitions.html")))
    err := t.ExecuteTemplate(w, "base", data)
    check(err, true)
}
