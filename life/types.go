package life

type SoftwareEngineer struct {
	Name              string
	FavouriteLanguage string
	Socials           map[string]string
}

func (p *SoftwareEngineer) DoStuff() {}

func (p *SoftwareEngineer) TellSelf(m string) {}

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
	Year   uint8
}
