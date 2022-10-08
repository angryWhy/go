package main

import "sync"

type user struct {
	name int
	age  int
}

func main() {

}
func newUser() *user {
	return &user{
		name: 1,
		age:  2,
	}
}

var (
	u        *user
	useronce sync.Once
)

func getInstance() *user {
	useronce.Do(func() {
		if u == nil {
			u = &user{}
		}
	})
	return u
}
