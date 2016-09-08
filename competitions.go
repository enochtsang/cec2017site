package main

import (
    "html/template"
    "net/http"
)

// var log defined in cec2017site.go

type Competition struct {
    ID    string
    Title string
    Blurb string
}

func competitions(w http.ResponseWriter, r *http.Request) {
    log.Debug("Loading competitions page for " + r.RemoteAddr)
    t := template.Must(template.ParseFiles(
        absPath("templates/base.html"),
        absPath("templates/competitions.html")))

    competitions := []Competition{
        {
            ID:    "competitions-general",
            Title: "General",
            Blurb: "blurb,",
        },
        {
            ID:    "senior-design",
            Title: "Senior Design",
            Blurb: "Senior design challenges a team of four students to design and construct a working prototype that can complete in a given engineering objective. These tasks are designed to test the competitor’s knowledge in design theory, technical skills, time management, and teamwork. The team must present their prototype to a panel of judges at the end of their design period.",
        },
        {
            ID:    "junior-design",
            Title: "Junior Design",
            Blurb: "Junior design challenges a team of four competitors, whose current standing is equivalent to 1st or 2nd year of an undergraduate engineering program, to design and build a prototype to address a technical problem within a day. This is a counterpart to the Senior Design competition, but emphasizes on prototype functionality of the design. The competitors must present their final prototype to a panel of judges at the end of their design period.",
        },
        {
            ID:    "innovative-design",
            Title: "Innovative Design",
            Blurb: "Innovative design competitors must prepare entirely prior to the commencement of the Canadian Engineering Competition. They must choose their own topic, prepare research, and develop a novel and groundbreaking design that addresses a void in society. The winner will be selected for their technical breakthrough and their excellence in implementing the engineering design procedure.",
        },
        {
            ID:    "re-engineering",
            Title: "Re Engineering",
            Blurb: "blurb",
        },
        {
            ID:    "engineering-communication",
            Title: "Engineering Communication",
            Blurb: "Engineering communications challenges a team of one or two competitors to describe a complicated technical process or issue in terms that the general public can understand. Evaluation is based on the competitor’s public speaking skills and the ability to persuade their audience.",
        },
        {
            ID:    "consulting",
            Title: "Consulting",
            Blurb: "Consulting engineering challenges a team of four competitors to design a detailed solution to a large-scale engineering problem. The proposal must be made in a way that promotes the solution to the client and demonstrates resourcefulness while acting in good faith with the spirit of the competition, and keeping in consideration the economic feasibility and environmental impacts of their design. The winners are selected for their excellence in the implementation of market research, feasibility studies, and design prototyping.",
        },
        {
            ID:    "extemporaneous-debate",
            Title: "Extemporaneous Debate",
            Blurb: "Commonly known as impromptu debate, here, a team of two competitors must present a series of convincing and well-constructed arguments given the minimal preparation time. The topic and viewpoints are determined at the beginning of each debate. The goal is to assess the competitors’ ability to convey ideas and develop arguments. These debates will be conducted in a shortened Canadian National Style debate format.",
        },
    }
    err := t.ExecuteTemplate(w, "base", competitions)
    check(err, true)
}
