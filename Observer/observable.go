package main

type Observable interface {
	subscribe(observer Observer)
	addVacancy(vacancy string)
	notifyAll()
}
