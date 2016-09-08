package main

import (
    "html/template"
    "net/http"
)

// CommitteeMember is info for each committee member
type CommitteeMember struct {
    Name        string
    Class       string
    Role        string
    PicturePath string
    Info        string
}

type Location struct {
    LocationSource string
}

func home(w http.ResponseWriter, r *http.Request) {
    log.Debug("Loading home page for " + r.RemoteAddr)
    t := template.Must(template.ParseFiles(
        absPath("templates/base.html"),
        absPath("templates/home.html")))
    err := t.ExecuteTemplate(w, "base", nil)
    check(err, true)
}

func home_committee(w http.ResponseWriter, r *http.Request) {
    log.Debug("Loading committee page for " + r.RemoteAddr)
    t := template.Must(template.ParseFiles(
        absPath("templates/base.html"),
        absPath("templates/committee.html")))
    committee := []CommitteeMember{
        {
            Name:        "Christine Cao",
            Class:       "christine-cao",
            Role:        "Co-Chair",
            PicturePath: "/resources/images/placeholders/profile-placeholder.png",
            Info:        ` Christine is a 4th Year Geomatics Engineering Student with a Biomedical Engineering Minor on her internship term with a software engineering company, called INSERT NAME HERE. In previous years she has been actively involved in the UofC Engineering Students’ Society and the Geomatics Engineering Students’ Society.`,
        },
        {
            Name:        "Kaylyn Schnell",
            Class:       "kaylyn-schnell",
            Role:        "Co-Chair",
            PicturePath: "/resources/images/placeholders/profile-placeholder.png",
            Info:        `Kaylyn is a 4th Year Chemical Engineering Student currently on her internship with Canadian Natural Resources. Previously Kaylyn has held positions on the University of Calgary Students’ Union and the UofC Engineering Students’ Society. She has also been involved in coordinating the FIRST Robotics Western Canada Regional Competition for the past several years.`,
        },
        {
            Name:        "Arsalan Fardi",
            Class:       "arsalan-fardi",
            Role:        "VP Hospitality",
            PicturePath: "/resources/images/placeholders/profile-placeholder.png",
            Info:        `Arsalan is a 4th Year Chemical Engineering Student with a Petroleum Minor currently on his internship term with TransCanada. He is an active member of the Chemical Engineering Students’ Society and the Petroleum and Energy Society on campus.`,
        },
        {
            Name:        "Leah Loeppky",
            Class:       "leah-loeppky",
            Role:        "VP Events",
            PicturePath: "/resources/images/placeholders/profile-placeholder.png",
            Info:        `Leah is a 3rd Year Mechanical Engineering Student and the current President of the Mechanical Engineering Students’ Society. She recently completed a summer work term in Lloydminster, Saskatchewan with Husky Energy.`,
        },
        {
            Name:        "Cindy Cai",
            Class:       "cindy-cai",
            Role:        "Co-VP Sponsorship",
            PicturePath: "/resources/images/placeholders/profile-placeholder.png",
            Info:        `Cindy is a 3rd Year Chemical Engineering Student with a Petroleum Minor and spends her spare time volunteering as an Energy Engineering Instructor for Minds in Motion, helping to engage youth in STEM subjects.`,
        },
        {
            Name:        "Manpreet Deol",
            Class:       "manpreet-deol",
            Role:        "Co-VP Sponsorship",
            PicturePath: "/resources/images/placeholders/profile-placeholder.png",
            Info:        `Manpreet is a 2nd Year Mechanical Engineering Student and also the current VP External of the Mechanical Engineering Students’ Society. At the University of Calgary she has been involved in the Sikh Students’ Association and the Engineering Students’ Society. Manpreet has won a variety of awards and scholarships including a 2015 Schulich Leader of Canada Scholarship and she has been named one of Calgary’s Top 5 Youth to Watch.`,
        },
        {
            Name:        "Jennifer Busser",
            Class:       "jenniger-busser",
            Role:        "VP Translations",
            PicturePath: "/resources/images/placeholders/profile-placeholder.png",
            Info:        `Jenn is a 4th Year Geomatics Engineering Student currently on her internship with NovAtel Inc. She has previously been involved in the Geomatics Engineering Students’ Society and the 2015 Alberta Student Energy Conference.`,
        },
        {
            Name:        "Hugo Olaciregui",
            Class:       "justin-olaciregui",
            Role:        "VP Finance",
            PicturePath: "/resources/images/placeholders/profile-placeholder.png",
            Info:        `Hugo is a 4th Year Mechanical Engineering Student on his internship term with Crescent Point Energy. He has previously been involved in exchange programs at the Université de Montréal and University of Stuttgart in Germany. He is also bilingual, speaking both English and Spanish.`,
        },
        {
            Name:        "Matthew Buchko",
            Class:       "matthew-buchko",
            Role:        "VP Logistics",
            PicturePath: "/resources/images/placeholders/profile-placeholder.png",
            Info:        `Matt is a 4th Year Chemical Engineering Student with a Petroleum Minor on his internship with Crescent Point Energy. He has previously been involved in the UofC Engineering Students’ Society as the VP External and the Chemical Engineering Students’ Society. Originally from Canmore, Matt enjoys escaping to the Rocky Mountains to enjoy hiking and photography. `,
        },
        {
            Name:        "Jasmine Zhao",
            Class:       "jasmine-zhao",
            Role:        "VP Competitions",
            PicturePath: "/resources/images/placeholders/profile-placeholder.png",
            Info:        `Jasmine is a 4th Year Chemical Engineering Student on her internship with TransCanada. She has previously been involved with the Petroleum and Energy Society, the Chemical Engineering Students’ Society and the UofC Engineering Students’ Society.`,
        },
    }
    err := t.ExecuteTemplate(w, "base", committee)
    check(err, true)
}

func home_uofc(w http.ResponseWriter, r *http.Request) {
    log.Debug("Loading uofc page for " + r.RemoteAddr)
    t := template.Must(template.ParseFiles(
        absPath("templates/base.html"),
        absPath("templates/uofc.html"),
        absPath("templates/widgets/didyouknow.html"),
        absPath("templates/widgets/location.html")))
    location := Location{"https://www.google.com/maps/embed/v1/place?q=univeristy%20of%20calgary&key=AIzaSyDtvcBmchfCPR_BxsOhP8UWOvamaNEQbQA&zoom=12"}
    err := t.ExecuteTemplate(w, "base", location)
    check(err, true)
}

func home_calgary(w http.ResponseWriter, r *http.Request) {
    log.Debug("Loading calgary page for " + r.RemoteAddr)
    t := template.Must(template.ParseFiles(
        absPath("templates/base.html"),
        absPath("templates/calgary.html"),
        absPath("templates/widgets/didyouknow.html"),
        absPath("templates/widgets/location.html")))
    location := Location{"https://www.google.com/maps/embed/v1/place?q=calgary%20ab&key=AIzaSyDtvcBmchfCPR_BxsOhP8UWOvamaNEQbQA&zoom=4"}
    err := t.ExecuteTemplate(w, "base", location)
    check(err, true)
}

func home_hotel(w http.ResponseWriter, r *http.Request) {
    log.Debug("Loading hotel page for " + r.RemoteAddr)
    t := template.Must(template.ParseFiles(
        absPath("templates/base.html"),
        absPath("templates/hotel.html"),
        absPath("templates/widgets/location.html")))
    location := Location{"https://www.google.com/maps/embed/v1/place?q=radisson%20hotel%20calgary&key=AIzaSyDtvcBmchfCPR_BxsOhP8UWOvamaNEQbQA&zoom=12"}
    err := t.ExecuteTemplate(w, "base", location)
    check(err, true)
}
