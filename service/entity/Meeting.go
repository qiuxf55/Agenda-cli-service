package entity

type Met struct {
	Title     string     `xorm:"pk varchar(255) notnull "`
	Sponsor   string     `xorm:"varchar(255) notnull"`
	Start string 	 `xorm:"varchar(255) notnull"`
	End   string 	 `xorm:"varchar(255) notnull"`
	Participators string `xorm:"varchar(255) notnull"`
}


type Meeting struct {
	Title     string
	Sponsor   string
	Start Date
	End   Date
	Participator []string 
}
func GetSponsor(a Meeting) string {
	return a.Sponsor
}
func GetParticipator(a Meeting) []string{
	return a.Participator
}

func GetStart(a Meeting) Date{
	return a.Start
}
func GetEnd(a Meeting) Date{
	return a.End
}

func GetTitle(a Meeting)string {
	return a.Title
}