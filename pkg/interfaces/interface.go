package interfaces

import (
	"fmt"
	"log"
)

type Speaker interface {
	Speak()
}

type Dog struct {
	Name string
}

func (d *Dog) Speak() {
	fmt.Println(d.Name)
}

func TestingDogInterface() {
	var s1 Speaker
	var d1 *Dog = &Dog{Name: "Doggy-Don"}
	s1 = d1
	s1.Speak()
	// s1 = d2
	// s1.Speak()
}

type Notifier interface {
	Notify() error
}

func SendNotification(notify Notifier) error {
	return notify.Notify()
}

type User struct {
	Name  string
	Email string
}

func (u *User) Notify() error {
	log.Printf("User: Sending User Email To %s<%s>\n",
		u.Name,
		u.Email)

	return nil
}

func TestingNotifierInterface() {
	user := User{
		Name:  "janet jones",
		Email: "janet@email.com",
	}
	// var notifier Notifier = user // This will give a compile time error, the rules for a type to implement the interface are the following-
	// 1.) The method set of the corresponding pointer type *T is the set of all methods with receiver
	// type *T ot T.
	// 2.) The method set of any other type T consists of all method with receiver type T.
	// We are not using a pointer type so rule 1 is not satisfied and receiver type is a pointer so rule 2 is also
	// no satisfied, that is why we are getting the compile time error.
	user.Notify() // same as &user.Notify
	SendNotification(&user)
}
