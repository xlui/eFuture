package main

import (
	"regexp"
	"testing"
)

func TestConfigUsername(t *testing.T) {
	c, _ := regexp.Compile("^[\\w-.]+@[a-zA-Z0-9-]{1,200}(\\.[a-zA-Z0-9-]+)+$")
	if !c.MatchString(conf.Username) {
		t.Error("Username is invalid!")
	}
}

func TestConfigPassword(t *testing.T) {
	if conf.Password == "" {
		t.Error("Password should not be empty!")
	}
}

func TestConfigSmtp(t *testing.T) {
	if conf.Smtp == "" {
		t.Error("You haven't configure the right smtp server!")
	}
}
