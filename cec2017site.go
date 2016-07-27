package main

import (
    "io"
    "log"
    "net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, "hello, world!\n")
}

func main() {
    http.HandleFunc("/", HelloServer)
    err := http.ListenAndServeTLS(":8443", "temp.crt", "temp.key", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
