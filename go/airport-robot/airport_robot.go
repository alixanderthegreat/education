package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet() string
}

func SayHello(name string, lang Greeter) string {
	return fmt.Sprintf("I can speak %s: %s %s!", lang.LanguageName(), lang.Greet(), name)
}

type Italian struct{}

func (i Italian) LanguageName() string { return "Italian" }
func (i Italian) Greet() string        { return "Ciao" }

type Portuguese struct{}

func (i Portuguese) LanguageName() string { return "Portuguese" }
func (i Portuguese) Greet() string        { return "Olá" }
