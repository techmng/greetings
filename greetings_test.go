package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Manoj"
	want := regexp.MustCompile("Hello " + name + " Welcome!")
	msg, err := Hello(name)
	if want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Manoj") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
