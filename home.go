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

type HomeData struct {
    Location          string
    Title             string
    Date              string
    School            string
    About             string
    AboutBlurb        string
    Sponsorship       string
    SponsorshipBlurb  string
    Registration      string
    RegistrationBlurb string
}

// CommitteeMember is info for each committee member
type CommitteeMember struct {
    Name        string
    Class       string
    Role        string
    PicturePath string
    Info        string
}

type LocationData struct {
    Blurb          string
    LocationSource string
}

func loadHomeData(lang string) HomeData {
    dataFilePath := path.Join(filepath.Dir(os.Args[0]), "data", lang, "home.yml")
    dataFile, err := ioutil.ReadFile(dataFilePath)
    check(err, true)
    var data HomeData
    err = yaml.Unmarshal(dataFile, &data)
    check(err, true)
    return data
}

func home(w http.ResponseWriter, r *http.Request, data HomeData) {
    log.Debug("Loading home page for " + r.RemoteAddr)
    t := template.Must(template.ParseFiles(
        absPath("templates/base.html"),
        absPath("templates/home.html")))
    err := t.ExecuteTemplate(w, "base", data)
    check(err, true)
}

func loadHomeCommitteeData(lang string) []CommitteeMember {
    dataFilePath := path.Join(filepath.Dir(os.Args[0]), "data", lang, "homeCommittee.yml")
    dataFile, err := ioutil.ReadFile(dataFilePath)
    check(err, true)
    var data []CommitteeMember
    err = yaml.Unmarshal(dataFile, &data)
    check(err, true)
    return data
}

func homeCommittee(w http.ResponseWriter, r *http.Request, data []CommitteeMember) {
    log.Debug("Loading committee page for " + r.RemoteAddr)
    t := template.Must(template.ParseFiles(
        absPath("templates/base.html"),
        absPath("templates/committee.html")))
    err := t.ExecuteTemplate(w, "base", data)
    check(err, true)
}

func loadHomeUofCData(lang string) LocationData {
    dataFilePath := path.Join(filepath.Dir(os.Args[0]), "data", lang, "homeUofC.yml")
    dataFile, err := ioutil.ReadFile(dataFilePath)
    check(err, true)
    var data LocationData
    err = yaml.Unmarshal(dataFile, &data)
    check(err, true)
    return data
}

func homeUofC(w http.ResponseWriter, r *http.Request, data LocationData) {
    log.Debug("Loading uofc page for " + r.RemoteAddr)
    t := template.Must(template.ParseFiles(
        absPath("templates/base.html"),
        absPath("templates/uofc.html"),
        absPath("templates/widgets/didyouknow.html"),
        absPath("templates/widgets/location.html")))
    err := t.ExecuteTemplate(w, "base", data)
    check(err, true)
}

func loadHomeCalgaryData(lang string) LocationData {
    dataFilePath := path.Join(filepath.Dir(os.Args[0]), "data", lang, "homeCalgary.yml")
    dataFile, err := ioutil.ReadFile(dataFilePath)
    check(err, true)
    var data LocationData
    err = yaml.Unmarshal(dataFile, &data)
    check(err, true)
    return data
}

func homeCalgary(w http.ResponseWriter, r *http.Request, data LocationData) {
    log.Debug("Loading calgary page for " + r.RemoteAddr)
    t := template.Must(template.ParseFiles(
        absPath("templates/base.html"),
        absPath("templates/calgary.html"),
        absPath("templates/widgets/didyouknow.html"),
        absPath("templates/widgets/location.html")))
    err := t.ExecuteTemplate(w, "base", data)
    check(err, true)
}

func loadHomeHotelData(lang string) LocationData {
    dataFilePath := path.Join(filepath.Dir(os.Args[0]), "data", lang, "homeHotel.yml")
    dataFile, err := ioutil.ReadFile(dataFilePath)
    check(err, true)
    var data LocationData
    err = yaml.Unmarshal(dataFile, &data)
    check(err, true)
    return data
}

func homeHotel(w http.ResponseWriter, r *http.Request, data LocationData) {
    log.Debug("Loading hotel page for " + r.RemoteAddr)
    t := template.Must(template.ParseFiles(
        absPath("templates/base.html"),
        absPath("templates/hotel.html"),
        absPath("templates/widgets/location.html")))
    err := t.ExecuteTemplate(w, "base", data)
    check(err, true)
}
