package mrizkisaputra

type User struct {
	Username string
	Password string
}

func (user User) GetUsername() string {
	return user.Username
}

func (user User) Login(username, password string) string {
	if username == user.Username && password == user.Password {
		return "Login Success"
	}
	return "Login Fail"
}
