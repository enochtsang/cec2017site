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

type AttendingData struct {
    Title         string
    Blurb         string
    DownloadLabel string
    DelegatePdf   string
    DelegateLabel string
    RulebookPdf   string
    RulebookLabel string
}

func loadAttendingData(lang string) AttendingData {
    dataFilePath := path.Join(filepath.Dir(os.Args[0]), "data", lang, "attending.yml")
    dataFile, err := ioutil.ReadFile(dataFilePath)
    check(err, true)
    var data AttendingData
    err = yaml.Unmarshal(dataFile, &data)
    check(err, true)
    return data
}

func attending(w http.ResponseWriter, r *http.Request, data AttendingData) {
    log.Debug("Loading attending page for " + r.RemoteAddr)
    t := template.Must(template.ParseFiles(
        absPath("templates/base.html"),
        absPath("templates/attending.html")))
    err := t.ExecuteTemplate(w, "base", data)
    check(err, true)
}
