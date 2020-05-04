package model

type User struct{
	ID int
	Username string
	Fullname string
	Email string
	Password string
}

func (u *User)Hello()string{
	return "Hello " + u.Fullname
}

func(u *User) Create() (*User){
	u.ID = 2
	u.Username = "MrHide06"
	u.Fullname = "Wahid"
	u.Email = "wahid@gmail.com"
	u.Password = "wahid"

	return u
}

func CreateUser(id int, username, fullname, email, password string) *User{
	return &User{
		ID: id,
		Username: username,
		Fullname: fullname,
		Email : email,
		Password: password,
	}
}