package greeting

import (
	"fmt"
	"strings"
)

type (
	greeting struct {
	}

	Greeting interface {
		Greet(name interface{}) string
	}
)

func NewGreeting() Greeting {
	return &greeting{}
}

func (this *greeting) Greet(name interface{}) string {
	statement := "Hello"
	greetingName := ""
	if _, isArray := name.([]string); isArray {
		if len(name.([]string)) == 2 {
			greetingName = fmt.Sprintf("%s and %s", name.([]string)[0], name.([]string)[1])
		}
	} else if name == "" {
		greetingName = "my friend"
	} else if name == strings.ToUpper(name.(string)) {
		statement = strings.ToUpper(statement)
	}
	if greetingName == "" {
		greetingName = name.(string)
	}
	return fmt.Sprintf("%s, %s.", statement, greetingName)
}
