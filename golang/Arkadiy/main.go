package main

import (
	"errors"
	"fmt"
)

type User struct {
	Username string
	Age      int
	Mobile   string
}

type UserService struct {
	users map[string]*User
}

type IUser interface {
	CreateUser(user *User) error
	GetUser(username string) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(username string) error
	ListUsers() []*User
}

func NewUserService() IUser {
	return &UserService{
		users: make(map[string]*User),
	}
}

func NewUser(username string, age int, mobile string) (*User, error) {
	if username == "" {
		return nil, errors.New("username cannot be empty")
	}
	if age < 0 {
		return nil, errors.New("age cannot be negative")
	}

	return &User{
		Username: username,
		Age:      age,
		Mobile:   mobile,
	}, nil
}

func (us *UserService) CreateUser(user *User) error {
	if user == nil {
		return errors.New("user cannot be nil")
	}
	if user.Username == "" {
		return errors.New("username cannot be empty")
	}

	if _, exists := us.users[user.Username]; exists {
		return errors.New("user already exists")
	}

	us.users[user.Username] = user
	fmt.Printf("User %s created successfully\n", user.Username)
	return nil
}

func (us *UserService) GetUser(username string) (*User, error) {
	if username == "" {
		return nil, errors.New("username cannot be empty")
	}

	user, exists := us.users[username]
	if !exists {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (us *UserService) UpdateUser(user *User) error {
	if user == nil {
		return errors.New("user cannot be nil")
	}
	if user.Username == "" {
		return errors.New("username cannot be empty")
	}

	if _, exists := us.users[user.Username]; !exists {
		return errors.New("user not found")
	}

	us.users[user.Username] = user
	fmt.Printf("User %s updated successfully\n", user.Username)
	return nil
}

func (us *UserService) DeleteUser(username string) error {
	if username == "" {
		return errors.New("username cannot be empty")
	}

	if _, exists := us.users[username]; !exists {
		return errors.New("user not found")
	}

	delete(us.users, username)
	fmt.Printf("User %s deleted successfully\n", username)
	return nil
}

func (us *UserService) ListUsers() []*User {
	userList := make([]*User, 0, len(us.users))
	for _, user := range us.users {
		userList = append(userList, user)
	}
	return userList
}

func main() {
	fmt.Println("User CRUD Operations with Service Pattern")

	userService := NewUserService()

	// Create users
	user1, _ := NewUser("Arkadiy", 20, "+1234567890")
	user2, _ := NewUser("Maria", 25, "+0987654321")
	user3, _ := NewUser("Ivan", 30, "+1111111111")

	userService.CreateUser(user1)
	userService.CreateUser(user2)
	userService.CreateUser(user3)

	// List all users
	fmt.Println("\nAll users:")
	for _, user := range userService.ListUsers() {
		fmt.Printf("Username: %s, Age: %d, Mobile: %s\n", user.Username, user.Age, user.Mobile)
	}

	// Get specific user
	user, err := userService.GetUser("Arkadiy")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("\nFound user: %+v\n", user)
	}

	// Update user
	updatedUser, _ := NewUser("Arkadiy", 21, "+9999999999")
	userService.UpdateUser(updatedUser)

	// Delete user
	userService.DeleteUser("Ivan")

	// List users after operations
	fmt.Println("\nUsers after operations:")
	for _, user := range userService.ListUsers() {
		fmt.Printf("Username: %s, Age: %d, Mobile: %s\n", user.Username, user.Age, user.Mobile)
	}
}
