package main

import (
    "html/template"
    "log"
    "net/http"
    "os"
    "path"
    "path/filepath"
)

func check(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func home(w http.ResponseWriter, r *http.Request) {
    t := template.Must(template.ParseFiles(path.Join(filepath.Dir(os.Args[0]), "templates/index.html")))
    err := t.Execute(w, nil)
    check(err)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, path.Join(filepath.Dir(os.Args[0]), "resources/images/favicon.ico"))
}

func main() {
    http.HandleFunc("/", home)
    http.Handle("/resources/", http.StripPrefix("/resources/",
        http.FileServer(http.Dir(path.Join(filepath.Dir(os.Args[0]), "resources")))))
    // http.HandleFunc("/favicon.ico", faviconHandler)

    err := http.ListenAndServeTLS(":8443", "temp.crt", "temp.key", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
