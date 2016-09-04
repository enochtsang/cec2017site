package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "html/template"
    "io/ioutil"
    "net/http"
    "net/smtp"
    "net/url"
    "os"
    "path"
    "path/filepath"
    "strings"
    // "text/template"
    "time"
)

type contactResponse struct {
    success bool
    error   string
}

type recaptchaResponse struct {
    Success     bool      `json:"success"`
    ChallengeTS time.Time `json:"challenge_ts"`
    Hostname    string    `json:"hostname"`
    ErrorCodes  []int     `json:"error-codes"`
}

type EmailData struct {
    SenderEmail    string
    RecipientEmail string
    Name           string
    Organization   string
    ReturnEmail    string
    Message        string
}

func contact(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        contactRequest := make(map[string]string)
        r.ParseForm()
        for k, v := range r.Form {
            contactRequest[strings.TrimSpace(k)] = strings.TrimSpace(strings.Join(v, " "))
        }

        fmt.Println("contact request: ", contactRequest)

        if checkCaptcha(contactRequest["captcha"]).Success {
            fmt.Println("captcha success!")
            sendMail(contactRequest["name"],
                contactRequest["organization"],
                contactRequest["returnEmail"],
                contactRequest["message"])
        } else {
            fmt.Println("captcha fail :(")
        }

    } else {
        t := template.Must(template.ParseFiles(
            absPath("templates/base.html"),
            absPath("templates/contact.html")))
        err := t.ExecuteTemplate(w, "base", nil)
        check(err)
    }
}

func contact_faq(w http.ResponseWriter, r *http.Request) {
    t := template.Must(template.ParseFiles(
        absPath("templates/base.html"),
        absPath("templates/faq.html")))
    err := t.ExecuteTemplate(w, "base", nil)
    check(err)
}

func checkCaptcha(response string) (r recaptchaResponse) {
    resp, err := http.PostForm("https://www.google.com/recaptcha/api/siteverify",
        url.Values{"secret": {"6Lce7SYTAAAAAPRGE4pVdTrtAFb6sDSmWnLJzirM"}, "response": {response}})
    if err != nil {
        fmt.Printf("Post error: %s\n", err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Read error: could not read body: %s", err)
    }
    err = json.Unmarshal(body, &r)
    if err != nil {
        fmt.Println("Read error: got invalid JSON: %s", err)
    }
    return
}

func sendMail(name, organization, email, body string) {
    context := &EmailData{
        SenderEmail:    "cec2017site@gmail.com",
        RecipientEmail: "echtsang@gmail.com",
        Name:           name,
        Organization:   organization,
        ReturnEmail:    email,
        Message:        body,
    }

    t, err := template.ParseFiles(path.Join(filepath.Dir(os.Args[0]), "templates/other/email.tmpl"))
    check(err)

    var doc bytes.Buffer
    err = t.Execute(&doc, context)
    fmt.Println(doc.String())
    check(err)

    auth := smtp.PlainAuth("",
        context.SenderEmail,
        "fMNNAKH9YofCVH%d0^",
        "smtp.gmail.com",
    )

    err = smtp.SendMail("smtp.gmail.com:587",
        auth,
        context.SenderEmail,
        []string{context.RecipientEmail},
        doc.Bytes())
    check(err)
}
