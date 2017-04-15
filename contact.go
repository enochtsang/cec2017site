package main

import (
    "bytes"
    "encoding/json"
    "gopkg.in/yaml.v2"
    "html/template"
    "io/ioutil"
    "net/http"
    "net/smtp"
    "net/url"
    "os"
    "path"
    "path/filepath"
    "strings"
    textTemplate "text/template"
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

type ContactData struct {
    Blurb        string
    NameLabel    string
    SchoolLabel  string
    EmailLabel   string
    MessageLabel string
    SubmitLabel  string
    ThankYou     string
}

type FAQData struct {
    Title string
    FAQs  []FAQ
}

type FAQ struct {
    Question string
    Answer   string
}

func loadContactData(lang string) ContactData {
    dataFilePath := path.Join(filepath.Dir(os.Args[0]), "data", lang, "contact.yml")
    dataFile, err := ioutil.ReadFile(dataFilePath)
    check(err, true)
    var data ContactData
    err = yaml.Unmarshal(dataFile, &data)
    check(err, true)
    return data
}

func contact(w http.ResponseWriter, r *http.Request, data ContactData) {
    if r.Method == "POST" {
        log.Debug("Received contact POST Request from " + r.RemoteAddr)
        contactRequest := make(map[string]string)
        r.ParseForm()
        for k, v := range r.Form {
            contactRequest[strings.TrimSpace(k)] = strings.TrimSpace(strings.Join(v, " "))
        }

        log.Info("contact request: ", contactRequest)

        if checkCaptcha(contactRequest["captcha"]).Success {
            sendMail(contactRequest["name"],
                contactRequest["organization"],
                contactRequest["returnEmail"],
                contactRequest["message"])
        } else {
            log.Warning("contact POST request with failed captcha")
        }

    } else {
        log.Debug("Loading contact page for " + r.RemoteAddr)
        t := textTemplate.Must(textTemplate.ParseFiles(
            absPath("templates/base.html"),
            absPath("templates/contact.html")))
        err := t.ExecuteTemplate(w, "base", data)
        check(err, true)
    }
}

func loadContactFaqData(lang string) FAQData {
    dataFilePath := path.Join(filepath.Dir(os.Args[0]), "data", lang, "faq.yml")
    dataFile, err := ioutil.ReadFile(dataFilePath)
    check(err, true)
    var data FAQData
    err = yaml.Unmarshal(dataFile, &data)
    check(err, true)
    return data
}

func contactFaq(w http.ResponseWriter, r *http.Request, data FAQData) {
    log.Debug("Loading FAQ page for " + r.RemoteAddr)
    t := template.Must(template.ParseFiles(
        absPath("templates/base.html"),
        absPath("templates/faq.html")))
    err := t.ExecuteTemplate(w, "base", data)
    check(err, true)
}

func checkCaptcha(response string) (r recaptchaResponse) {
    resp, err := http.PostForm("https://www.google.com/recaptcha/api/siteverify",
        url.Values{"secret": {"6Lce7SYTAAAAAPRGE4pVdTrtAFb6sDSmWnLJzirM"}, "response": {response}})
    if err != nil {
        log.Warning("Post error: %s\n", err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Warning("Read error: could not read body: %s", err)
    }
    err = json.Unmarshal(body, &r)
    if err != nil {
        log.Warning("Read error: got invalid JSON: %s", err)
    }
    return
}

func sendMail(name, organization, email, body string) {
    context := &EmailData{
        SenderEmail:    "cec2017site@gmail.com",
        RecipientEmail: "cec2017@cfes.ca",
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
        "fakepassword",
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
