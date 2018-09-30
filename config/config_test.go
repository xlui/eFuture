package config

import (
	"regexp"
	"testing"
)

// Check send mail username. Must be an email address
func TestConfigUsername(t *testing.T) {
	c, _ := regexp.Compile("^[\\w-.]+@[a-zA-Z0-9-]{1,200}(\\.[a-zA-Z0-9-]+)+$")
	if !c.MatchString(Configuration.Username) {
		t.Error("Username is invalid!")
	}
}

// Check password.
func TestConfigPassword(t *testing.T) {
	if Configuration.Password == "" {
		t.Error("Password should not be empty!")
	}
}

// Check smtp server
func TestConfigSmtp(t *testing.T) {
	if Configuration.Smtp == "" {
		t.Error("You haven't configure the right smtp server!")
	}
}

// Check smtp port
func TestSmtpPort(t *testing.T) {
	c, _ := regexp.Compile("^\\d+$")
	if !c.MatchString(Configuration.SmtpPort) {
		t.Error("Smtp port must be numbers!")
	}
}
