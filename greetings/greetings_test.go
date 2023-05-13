package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Joao"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Joao")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Joao") = %q, %v, want math for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
