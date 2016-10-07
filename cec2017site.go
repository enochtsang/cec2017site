package main

import (
    "github.com/op/go-logging"
    "gopkg.in/natefinch/lumberjack.v2"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "net/http"
    "os"
    "path"
    "path/filepath"
    "strconv"
    "time"
)

var log = logging.MustGetLogger("gpu-test-controller")

type ConfigInfo struct {
    Port     int    `yaml:"Port"`
    CertFile string `yaml:"CertFile"`
    KeyFile  string `yaml:"KeyFile"`
    LogFile  string `yaml:"LogFile"`
}

func check(err error, exit bool) {
    if exit {
        defer func() {
            if r := recover(); r != nil {
                log.Critical(r)
                os.Exit(1)
            }
        }()
        if err != nil {
            panic(err)
        }
    } else {
        log.Error(err)
    }
}

func absPath(relativePath string) string {
    dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    check(err, true)
    result := path.Join(dir, relativePath)
    return result
}

func redirectToHome(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "/home", 301)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, absPath("resources/images/favicon.ico"))
}

func main() {
    configFile, err := ioutil.ReadFile(path.Join(filepath.Dir(os.Args[0]), "config.yml"))
    check(err, true)
    var config ConfigInfo
    err = yaml.Unmarshal(configFile, &config)
    check(err, true)

    // Intialize Logging
    logFormat := logging.MustStringFormatter(
        `%{time:15:04:05} [%{level}] ▶ %{message}`,
    )
    lumberLog := lumberjack.Logger{
        Filename:   config.LogFile,
        MaxSize:    999999999, // megabytes
        MaxBackups: 0,
        MaxAge:     0, //days
    }
    go func() {
        for {
            <-time.After(time.Hour * 24)
            lumberLog.Rotate() // Rotate once per day
        }
    }()
    logBackend := logging.NewLogBackend(&lumberLog, "", 0)
    backendFormatter := logging.NewBackendFormatter(logBackend, logFormat)
    logging.SetBackend(backendFormatter)

    //Initialize loggers
    http.HandleFunc("/", redirectToHome)
    http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
        home(w, r, loadHomeData("en"))
    })
    http.HandleFunc("/home/committee", func(w http.ResponseWriter, r *http.Request) {
        homeCommittee(w, r, loadHomeCommitteeData("en"))
    })
    http.HandleFunc("/home/uofc", func(w http.ResponseWriter, r *http.Request) {
        homeUofC(w, r, loadHomeUofCData("en"))
    })
    http.HandleFunc("/home/calgary", func(w http.ResponseWriter, r *http.Request) {
        homeCalgary(w, r, loadHomeCalgaryData("en"))
    })
    http.HandleFunc("/home/hotel", func(w http.ResponseWriter, r *http.Request) {
        homeHotel(w, r, loadHomeHotelData("en"))
    })
    http.HandleFunc("/competitions", competitions)
    // http.HandleFunc("/competitions", func(w http.ResponseWriter, r *http.Request) {
    //     competitions(w, r, loadCompetitionsData("en"))
    // })
    http.HandleFunc("/sponsors", sponsors)
    // http.HandleFunc("/sponsors", func(w http.ResponseWriter, r *http.Request) {
    //     sponsors(w, r, loadSponsorsData("en"))
    // })
    http.HandleFunc("/sponsors/package", sponsorsPackage)
    // http.HandleFunc("/sponsors/package", func(w http.ResponseWriter, r *http.Request) {
    //     sponsorsPackage(w, r, loadSponsorsPackageData("en"))
    // })
    http.HandleFunc("/attending", attending)
    // http.HandleFunc("/attending", func(w http.ResponseWriter, r *http.Request) {
    //     attending(w, r, loadAttendingData("en"))
    // })
    http.HandleFunc("/contact", contact)
    // http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
    //     contact(w, r, loadContactData("en"))
    // })
    http.HandleFunc("/contact/faq", contactFaq)
    // http.HandleFunc("/contact/faq", func(w http.ResponseWriter, r *http.Request) {
    //     contactFaq(w, r, loadContactFaqData("en"))
    // })
    http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir(absPath("resources")))))
    http.HandleFunc("/favicon.ico", faviconHandler)

    // Run Site
    log.Notice("CEC 2017 Site running on 8443")
    err = http.ListenAndServeTLS(":"+strconv.Itoa(config.Port), absPath(config.CertFile), absPath(config.KeyFile), nil)
    check(err, true)
    log.Notice("CEC 2017 Site closing")
}
