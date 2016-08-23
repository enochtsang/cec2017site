package main

import (
    "log"
    "net/http"
    "os"
    "path"
    "path/filepath"
)

func check(err error) {
    if err != nil {
        panic(err)
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

func faviconHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, absPath("resources/images/favicon.ico"))
}

func main() {
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

    log.Println("CEC 2017 Site Running on 8443")
    err := http.ListenAndServeTLS(":8443", absPath("temp.crt"), absPath("temp.key"), nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
