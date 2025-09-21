package main

import "fmt"

type User struct {
	username string
	age      int
	mobile   string
}

type IUser interface {
	CreateUser(user *User) error
	ReadUserName(username string) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(username string) error
}

func NewUser(username string, age int, mobile string) (*User, error) {
	return &User{
		username: username,
		age:      age,
		mobile:   mobile,
	}, nil

}

func (s *User) CreateUser(user *User) error {

}

func (s *User) ReadUserName(username string) (*User, error) {

}

func (s *User) UpdateUser(user *User) error {

}

func (s *User) DeleteUser(username string) error {

}

func main() {
	fmt.Println("vim-go")

	person := CreateUser("Arkadiy", 20, "")

	delPerson := DeleteUser(person)
	fmt.Println(delPerson)

}
