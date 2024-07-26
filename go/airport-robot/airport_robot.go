package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

func SayHello(name string, g Greeter) string {
	language := g.LanguageName()
	greet := g.Greet(name)

	return fmt.Sprintf("I can speak %s: %s", language, greet)
}

type German struct {}

func (g German) LanguageName() string {
	return "German"
}

func (g German) Greet(name string) string {
	return fmt.Sprintf("Hallo %s!", name)
}

type Italian struct {}

func (i Italian) LanguageName() string {
	return "Italian"
}

func (i Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s!", name)
}

type Portuguese struct {}

func (p Portuguese) LanguageName() string {
	return "Portuguese"
}

func (p Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %s!", name)
}
