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

func absPath(relativePath string) string {
    dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    check(err)
    result := path.Join(dir, relativePath)
    log.Println("Created absPath " + result)
    return result
}

func redirectToHome(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "/home", 301)
}

func home(w http.ResponseWriter, r *http.Request) {
    t := template.Must(template.ParseFiles(
        absPath("templates/base.html"),
        absPath("templates/home.html")))
    err := t.ExecuteTemplate(w, "base", nil)
    check(err)
}

func home_committee(w http.ResponseWriter, r *http.Request) {
    t := template.Must(template.ParseFiles(
        absPath("templates/base.html"),
        absPath("templates/committee.html")))
    err := t.ExecuteTemplate(w, "base", nil)
    check(err)
}

func home_uofc(w http.ResponseWriter, r *http.Request) {
    t := template.Must(template.ParseFiles(
        absPath("templates/base.html"),
        absPath("templates/uofc.html"),
        absPath("templates/widgets/didyouknow.html"),
        absPath("templates/widgets/location.html")))
    err := t.ExecuteTemplate(w, "base", nil)
    check(err)
}

func home_calgary(w http.ResponseWriter, r *http.Request) {
    t := template.Must(template.ParseFiles(
        absPath("templates/base.html"),
        absPath("templates/calgary.html"),
        absPath("templates/widgets/didyouknow.html"),
        absPath("templates/widgets/location.html")))
    err := t.ExecuteTemplate(w, "base", nil)
    check(err)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, absPath("resources/images/favicon.ico"))
}

func main() {
    http.HandleFunc("/", redirectToHome)
    http.HandleFunc("/home", home)
    http.HandleFunc("/home/committee", home_committee)
    http.HandleFunc("/home/uofc", home_uofc)
    http.HandleFunc("/home/calgary", home_calgary)
    http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir(absPath("resources")))))
    http.HandleFunc("/favicon.ico", faviconHandler)

    log.Println("CEC 2017 Site Running on 8443")
    err := http.ListenAndServeTLS(":8443", absPath("temp.crt"), absPath("temp.key"), nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
