# Hello, World! 👋🏻

```go
package jack

import (
    "github.com/Jack-Gledhill/experiences"
    "github.com/Jack-Gledhill/life"
    "github.com/Jack-Gledhill/locations"
)

var (
    Me = life.Person{
        Name:       "Jack Gledhill",
        Profession: "Software Engineer",
        Skills: []string{
            "Full-stack Development",
            "Systems Administration",
            "DevOps",
            "GitOps",
        },
    }
    Languages = []string{
        "Golang",
        "Python",
        "Java",
        "Ruby",
        "JavaScript",
        "HTML/CSS",
    }
    Tools = []string{
        "Git",
        "JetBrains IDEs",
        "Kubernetes",
        "Docker",
        "Ansible",
    }
    Education = []life.Education{
        {
            Name:   "University of Sheffield",
            Course: "BEng Software Engineering",
            URL:    "https://sheffield.ac.uk",
            Email:  "jgledhill2@sheffield.ac.uk",
            Year:   0,
        },
    }
    Experience = []life.Experience{
        {
            Name:     "HackSheffield 9",
            Location: locations.Sheffield,
            Type:     experiences.Hackathon,
            Role:     "Technical Lead",
            URL:      "https://hacksheffield.uk",
        },
        {
            Name:     "Sheffield CompSoc",
            Location: locations.Sheffield,
            Type:     experiences.Society,
            Role:     "First Year Representative",
            URL:      "https://shefcompsoc.uk",
        },
        {
            Name:     "Digital Zest",
            Location: locations.Scarborough,
            Type:     experiences.Employment,
            Role:     "Software Engineer",
            URL:      "https://digitalzest.co.uk",
        },
        {
            Name:     "Streamcord",
            Location: locations.Remote,
            Type:     experiences.Employment,
            Role:     "Software Engineer",
            URL:      "https://streamcord.io",
        },
    }
    Projects = []life.Project{
        {
            Name:        "Constellation",
            Description: "My Kubernetes homelab that runs various Open Source software, some of my own projects, and a bunch of game servers.",
            URL:         "https://github.com/constellation-net",
        },
        {
            Name:        "RoboJack",
            Description: "A Discord bot I wrote to do a few laborious things for me, using Discord's interaction tools as a replacement for a UI.",
            URL:         "https://github.com/Jack-Gledhill/robojack",
        },
    }
)

func init() {
    // Recover from panics
    defer func() {
        if recover() != nil {
            Me.TellSelf("There there, everything will be OK :)")
        }
    }()

    Me.DoStuff()
}
```
