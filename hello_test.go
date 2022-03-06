package main

import (
	"regexp"
	"testing"
)

func TestGetFullNameValid(t *testing.T) {
	name := "Andre"
	want := regexp.MustCompile(`\b` + name + `\b`)
	fullName, err := GetFullName(name)
	if !want.MatchString(fullName) || err != nil {
		t.Fatalf(`Hello("%v") = %q, %v, want match for %#q, nil`, name, fullName, err, want)
	}
}

func TestGetFullNameEmpty(t *testing.T) {
	fullName, err := GetFullName("")
	if fullName != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, fullName, err)
	}
}
