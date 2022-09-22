package main

func main() {

	publisher := newPublisher("hh.ru")

	observerFirst := &Subscriber{"211524@astanait.edu.kz", "alexx"}
	publisher.addVacancy("Java Developer")
	publisher.subscribe(observerFirst)
	publisher.notifyAll()

	publisher.addVacancy("Full Stack Developer")
}
