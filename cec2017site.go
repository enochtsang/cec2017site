package main

import (
    "github.com/op/go-logging"
    "gopkg.in/natefinch/lumberjack.v2"
    "net/http"
    "os"
    "path"
    "path/filepath"
)

var log = logging.MustGetLogger("gpu-test-controller")

func check(err error, exit bool) {
    defer func() {
        if r := recover(); r != nil {
            log.Critical(r)
        }
    }()

    if err != nil {
        panic(err)
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
    // Intialize Logging
    logFormat := logging.MustStringFormatter(
        `%{time:15:04:05} [%{level}] ▶ %{message}`,
    )
    lumberLog := lumberjack.Logger{
        Filename:   "/var/log/cec2017site/cec2017site.log",
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
    http.HandleFunc("/home", home)
    http.HandleFunc("/home/committee", home_committee)
    http.HandleFunc("/home/uofc", home_uofc)
    http.HandleFunc("/home/calgary", home_calgary)
    http.HandleFunc("/home/hotel", home_hotel)
    http.HandleFunc("/competitions", competitions)
    http.HandleFunc("/sponsors", sponsors)
    http.HandleFunc("/sponsors/package", sponsors_package)
    http.HandleFunc("/attending", attending)
    http.HandleFunc("/contact", contact)
    http.HandleFunc("/contact/faq", contact_faq)
    http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir(absPath("resources")))))
    http.HandleFunc("/favicon.ico", faviconHandler)

    // Run Site
    log.Notice("CEC 2017 Site Running on 8443")
    err := http.ListenAndServeTLS(":8443", absPath("temp.crt"), absPath("temp.key"), nil)
    check(err, true)
}
