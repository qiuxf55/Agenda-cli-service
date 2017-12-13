package entity

type User struct {
	Name string `xorm:"pk varchar(255) notnull "`
	Password string `xorm:"varchar(255) notnull"`
	Email    string `xorm:"varchar(255) notnull"`
	Phone    string `xorm:"varchar(255) notnull"`
}
func GetName(a User)string {
	return a.Name
}

func GetPhone(a User) string{
	return a.Phone
}
func GetEmail(a User) string{
	return a.Email
}
func GetPassword(a User) string{
	return a.Password
}