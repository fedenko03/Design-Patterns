package main

type Publisher struct {
	subscribers []Observer
	vacancies   []string
	siteTitle   string
}

func newPublisher(title string) *Publisher {
	return &Publisher{
		siteTitle: title,
	}
}

func (i *Publisher) addVacancy(vacancy string) {
	i.vacancies = append(i.vacancies, vacancy)
	i.notifyAll()
}

func (i *Publisher) subscribe(o Observer) {
	i.subscribers = append(i.subscribers, o)
}

func (i *Publisher) notifyAll() {
	for _, observer := range i.subscribers {
		observer.update(i.vacancies)
	}
}
