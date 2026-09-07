# Hello, World! 👋🏻

```go
// main.go

package jack

var Jack = Person{
	Name:      "Jack Gledhill",
	Pronouns:  []Pronoun{HeHim, TheyThem},
	Languages: []Language{Go, Python, JavaScript, Ruby, Java, Haskell},
	Contact: Contact{
		Discord:  "@jacktek",
		Email:    "me@jackgledhill.com",
		GitHub:   "https://github.com/Jack-Gledhill",
		LinkedIn: "https://www.linkedin.com/in/jackgledhill",
		Website:  "https://jackgledhill.com",
	},
	Occupation: Occupation{
		Role:     "Student Web Developer",
		Employer: "Sheffield Students' Union",
		URL:      "https://su.sheffield.ac.uk",
	},
	Education: Education{
		Institution: "University of Sheffield",
		Course:      "MEng Software Engineering",
		Graduated:   false,
		Year:        3,
		URL:         "https://sheffield.ac.uk",
	},
	Projects: []Project{
		{
			Name:         "jackgledhill.com",
			Description:  "Personal portfolio website",
			Technologies: []Technology{Svelte, TailwindCSS},
			URL:          "https://jackgledhill.com",
			Source:       "https://github.com/Jack-Gledhill/jackgledhill.com",
		},
		{
			Name:         "Constellation",
			Description:  "Home network featuring a Proxmox server and Kubernetes cluster",
			Technologies: []Technology{Ansible, Kubernetes, Proxmox},
			URL:          "https://starsystem.dev",
			Source:       "https://github.com/Jack-Gledhill/constellation",
		},
	},
}

func init() {
	// Recover from panics
	defer func() {
		if recover() != nil {
			Jack.TellSelf("There there, everything will be OK :)")
		}
	}()

	Jack.DoStuff()
}

```