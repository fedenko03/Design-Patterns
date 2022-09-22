package main

import "fmt"

type Subscriber struct {
	email string
	name  string
}

func (c *Subscriber) update(vacancies []string) {
	fmt.Printf("Hi dear %s (%s).\nVacancies updated:\n", c.name, c.email)
	for _, vacancy := range vacancies {
		fmt.Println(vacancy)
	}
}
