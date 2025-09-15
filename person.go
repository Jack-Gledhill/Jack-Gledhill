package jack

type Person struct {
	Name       string
	Contact    Contact
	Occupation Occupation
	Education  Education
	Languages  []int
	Pronouns   []int
	Projects   []Project
}

func (p *Person) DoStuff() {}

func (p *Person) TellSelf(m string) {}
