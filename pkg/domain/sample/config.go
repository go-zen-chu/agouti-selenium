package sample

import (
	"fmt"
	"os"
)

type Config struct {
	User     string
	Password string
	Driver   string
}

func LoadConfigFromEnv() (*Config, error) {
	c := &Config{}
	props := []struct {
		field *string
		name  string
	}{
		{field: &c.User, name: "USER"},
		{field: &c.Password, name: "PASSWORD"},
		{field: &c.Driver, name: "WEBDRIVER"},
	}
	emptyProp := ""
	for _, prop := range props {
		*prop.field = os.Getenv(prop.name)
		if *prop.field == "" {
			emptyProp += " " + prop.name
		}
	}
	if emptyProp != "" {
		return nil, fmt.Errorf("envvar not set %s", emptyProp)
	}
	return c, nil
}
