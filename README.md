# Hello, World! 👋🏻

```go
// main.go

package jack

import (
	"github.com/Jack-Gledhill/Jack-Gledhill/languages"
	"github.com/Jack-Gledhill/Jack-Gledhill/pronouns"
)

var Jack = Person{
	Name:      "Jack Gledhill",
	Pronouns:  []int{pronouns.HeHim, pronouns.TheyThem},
	Languages: []int{languages.Go, languages.Python, languages.JavaScript},
	Contact: Contact{
		Discord:  "@jacktek",
		Email:    "me@jackgledhill.com",
		GitHub:   "https://github.com/Jack-Gledhill",
		LinkedIn: "https://www.linkedin.com/in/jackgledhill",
		Website:  "https://jackgledhill.com",
	},
	Occupation: Occupation{
		Role:     "Web Developer Intern",
		Employer: "b:friend",
		URL:      "https://letsbfriend.org.uk",
	},
	Education: Education{
		Institution: "University of Sheffield",
		Course:      "MEng Software Engineering",
		Graduated:   false,
		Year:        2,
		URL:         "https://sheffield.ac.uk",
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