package model

type User struct {
	name, password string
	id int64
}

func (u *User) String() string {
	return u.name + u.password + string(u.id)
}

var nextId int64 = 0

func CreateUser(name, password string) User {
	nextId++
	return User{name:name, password:password, id:nextId}
}
