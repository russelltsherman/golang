package main

import (
	"log"
	"os"
	"time"
)

type User struct {
	Email string
	Name  string
}

var DataBase = []User{
	{Email: "alexander.davis@example.com", Name: "Alexander Davis"},
	{Email: "alexander.jackson@example.com", Name: "Alexander Jackson"},
	{Email: "avery.williams@example.com", Name: "Avery Williams"},
	{Email: "charolette.smith@example.com", Name: "Charlotte Smith"},
	{Email: "daniel.miller@example.com", Name: "Daniel Miller"},
	{Email: "ella.smith@example.com", Name: "Ella Smith"},
	{Email: "jacob.white@example.com", Name: "Jacob White"},
	{Email: "james.martin@example.com", Name: "James Martin"},
	{Email: "james.miller@example.com", Name: "James Miller"},
	{Email: "jayden.jackson@example.com", Name: "Jayden Jackson"},
	{Email: "liam.robinson@example.com", Name: "Liam Robinson"},
	{Email: "mason.martin@example.com", Name: "Mason Martin"},
	{Email: "mathew.jackson@example.com", Name: "Mathew Jackson"},
	{Email: "mia.smith@example.com", Name: "Mia Smith"},
	{Email: "micheal.white@example.com", Name: "Micheal White"},
	{Email: "natalie.martin@example.com", Name: "Natalie Martin"},
	{Email: "sofia.garcia@example.com", Name: "Sofia Garcia"},
	{Email: "william.brown@example.com", Name: "William Brown"},
}

type Worker struct {
	users []User
	ch    chan *User
}

func NewWorker(users []User, ch chan *User) *Worker {
	return &Worker{users: users, ch: ch}
}

func (w *Worker) Find(email string) {
	for i := range w.users {
		user := &w.users[i]
		if user.Email == email {
			w.ch <- user
		}
	}
}

func main() {
	email := os.Args[1]

	ch := make(chan *User)

	// multiple worers search slices of database
	go NewWorker(DataBase[:6], ch).Find(email)
	go NewWorker(DataBase[6:12], ch).Find(email)
	go NewWorker(DataBase[12:], ch).Find(email)

	log.Printf("searching for %s", email)
	select {
	case user := <-ch:
		log.Printf("email %s is owned by %s", email, user.Name)
	case <-time.After(100 * time.Millisecond):
		log.Printf("email %s not found", email)
	}
}
