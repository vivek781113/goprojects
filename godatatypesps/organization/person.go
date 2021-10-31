package organization

import (
	"errors"
	"fmt"
	"strings"
)

type TwitterHandler = string
type Identifiable interface {
	ID() string
}

type Person struct {
	firstName      string
	lastName       string
	twitterHandler string
}

func NewPerson(fn, ln string) Person {
	return Person{
		firstName: fn,
		lastName:  ln,
	}
}

func (p *Person) FullName() string {
	return fmt.Sprintf("%s %s", p.firstName, p.lastName)
}

func (p *Person) ID() string {
	return "12345"
}

func (p *Person) SetTwitterhandler(handler TwitterHandler) error {
	if len(handler) == 0 {
		p.twitterHandler = handler
	} else if !strings.HasPrefix(handler, "@") {
		return errors.New("twitter handler must start with @ symbol")
	}
	p.twitterHandler = handler
	return nil
}

func (p *Person) TwitterHandler() string {
	return p.twitterHandler
}
