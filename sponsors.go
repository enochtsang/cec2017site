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

type SponsorsData struct {
	Lead           string
	LeadSponsors   []Sponsor
	Gold           string
	GoldSponsors   []Sponsor
	Silver         string
	SilverSponsors []Sponsor
	Bronze         string
	BronzeSponsors []Sponsor
}

type Sponsor struct {
	Name    string
	ImgPath string
	Link    string
}

type SponsorsPackageData struct {
	Title         string
	Blurb         string
	Label         string
	PdfPath       string
	DownloadLabel string
}

func loadSponsorsData(lang string) SponsorsData {
	dataFilePath := path.Join(filepath.Dir(os.Args[0]), "data", lang, "sponsors.yml")
	dataFile, err := ioutil.ReadFile(dataFilePath)
	check(err, true)
	var data SponsorsData
	err = yaml.Unmarshal(dataFile, &data)
	check(err, true)
	return data
}

func sponsors(w http.ResponseWriter, r *http.Request, data SponsorsData) {
	log.Debug("Loading sponsors page for " + r.RemoteAddr)
	t := template.Must(template.ParseFiles(
		absPath("templates/base.html"),
		absPath("templates/sponsors.html")))
	err := t.ExecuteTemplate(w, "base", data)
	check(err, true)
}

func loadSponsorsPackageData(lang string) SponsorsPackageData {
	dataFilePath := path.Join(filepath.Dir(os.Args[0]), "data", lang, "sponsorsPackage.yml")
	dataFile, err := ioutil.ReadFile(dataFilePath)
	check(err, true)
	var data SponsorsPackageData
	err = yaml.Unmarshal(dataFile, &data)
	check(err, true)
	return data
}

func sponsorsPackage(w http.ResponseWriter, r *http.Request, data SponsorsPackageData) {
	log.Debug("Loading sponsorsPackage page for " + r.RemoteAddr)
	t := template.Must(template.ParseFiles(
		absPath("templates/base.html"),
		absPath("templates/sponsorsPackage.html")))
	err := t.ExecuteTemplate(w, "base", data)
	check(err, true)
}
