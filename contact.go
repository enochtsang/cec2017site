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
        log.Debug("Received contact POST Request from " + r.RemoteAddr)
        contactRequest := make(map[string]string)
        r.ParseForm()
        for k, v := range r.Form {
            contactRequest[strings.TrimSpace(k)] = strings.TrimSpace(strings.Join(v, " "))
        }

        fmt.Println("contact request: ", contactRequest)

        if !checkCaptcha(contactRequest["captcha"]).Success {
            sendMail(contactRequest["name"],
                contactRequest["organization"],
                contactRequest["returnEmail"],
                contactRequest["message"])
        } else {
            log.Warning("contact POST request with failed captcha")
        }

    } else {
        log.Debug("Loading contact page for " + r.RemoteAddr)
        t := template.Must(template.ParseFiles(
            absPath("templates/base.html"),
            absPath("templates/contact.html")))
        err := t.ExecuteTemplate(w, "base", nil)
        check(err, true)
    }
}

func contact_faq(w http.ResponseWriter, r *http.Request) {
    log.Debug("Loading FAQ page for " + r.RemoteAddr)
    t := template.Must(template.ParseFiles(
        absPath("templates/base.html"),
        absPath("templates/faq.html")))
    err := t.ExecuteTemplate(w, "base", nil)
    check(err, true)
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
        RecipientEmail: "etsang1@hotmail.com",
        Name:           name,
        Organization:   organization,
        ReturnEmail:    email,
        Message:        body,
    }

    if emailDataIsMalicious(*context) {
        log.Warning("Found malicious emailData: ", *context)
        return
    }

    t, err := template.ParseFiles(path.Join(filepath.Dir(os.Args[0]), "templates/other/email.tmpl"))
    check(err, true)

    var doc bytes.Buffer
    err = t.Execute(&doc, context)
    check(err, false)

    auth := smtp.PlainAuth("",
        context.SenderEmail,
        "fMNNAKH9YofCVH%d0^",
        "smtp.gmail.com",
    )

    log.Info("Sending email: ", doc.String())
    err = smtp.SendMail("smtp.gmail.com:587",
        auth,
        context.SenderEmail,
        []string{context.RecipientEmail},
        doc.Bytes())
    check(err, false)
}

func emailDataIsMalicious(data EmailData) bool {
    if strings.ContainsAny(data.SenderEmail, "\r\n") ||
        strings.ContainsAny(data.RecipientEmail, "\r\n") ||
        strings.ContainsAny(data.Name, "\r\n") ||
        strings.ContainsAny(data.Organization, "\r\n") {
        return true
    }
    return false
}
