package life

type Person struct {
	Name       string
	Profession string
	Skills     []string
}

func (p *Person) DoStuff() {}

type Experience struct {
	Name     string
	Location int
	Type     int
	Role     string
	URL      string
}

type Project struct {
	Name        string
	Description string
	URL         string
}

type Education struct {
	Name   string
	Course string
	URL    string
	Email  string
	Year   uint8
}
