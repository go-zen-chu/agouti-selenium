//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE

package sample

import (
	"fmt"
	"github.com/go-zen-chu/agouti-selenium/pkg/domain/selenium"
)

type Sample interface {
	Start() error
	LoginGoogle() error
	Stop() error
}

type sample struct {
	config        *Config
	browserTester selenium.BrowserTester
}

const loginURL = "https://accounts.google.com/signin/v2/identifier"

func NewSample(config *Config, browserTester selenium.BrowserTester) Sample {
	s := &sample{
		config:        config,
		browserTester: browserTester,
	}
	return s
}

func (s *sample) Start() error {
	return s.browserTester.Start()
}

func (s *sample) Stop() error {
	return s.browserTester.Stop()
}

func (s *sample) LoginGoogle() error {
	err := s.browserTester.Navigate(loginURL)
	if err != nil {
		return fmt.Errorf("error to navigate: %s", err.Error())
	}
	err = s.browserTester.FillByID("identifierId", s.config.User)
	if err != nil {
		return  fmt.Errorf("error filling user: %s", err.Error())
	}
	// in actual case google randomly generate submit button
	// so code below doesn't work :(
	err = s.browserTester.ClickByID("submit")
	if err != nil {
		return  fmt.Errorf("error submitting : %s", err.Error())
	}
	err = s.browserTester.FillByID("password", s.config.Password)
	if err != nil {
		return  fmt.Errorf("error filling password : %s", err.Error())
	}
	err = s.browserTester.ClickByID("submit")
	if err != nil {
		return  fmt.Errorf("error submitting : %s", err.Error())
	}
	return nil
}
