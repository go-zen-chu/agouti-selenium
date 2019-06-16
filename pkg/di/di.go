package di

import (
	"fmt"
	"github.com/go-zen-chu/agouti-selenium/pkg/domain/sample"
	"github.com/go-zen-chu/agouti-selenium/pkg/domain/selenium"
	is "github.com/go-zen-chu/agouti-selenium/pkg/infrastructure/selenium"
)

type DI map[string]interface{}

// NewDI creates di container
func NewDI() DI {
	return map[string]interface{}{}
}

func (di DI) Config() *sample.Config {
	key := "Config"
	return di.retrieveCache(key, func() (interface{}, error) {
		return sample.LoadConfigFromEnv()
	}).(*sample.Config)
}

func (di DI) Sample() sample.Sample {
	key := "Sample"
	return di.retrieveCache(key, func() (interface{}, error) {
		val := sample.NewSample(di.Config(), di.BrowserTester())
		return val, nil
	}).(sample.Sample)
}

func (di DI) BrowserTester() selenium.BrowserTester {
	key := "BrowserTester"
	return di.retrieveCache(key, func() (interface{}, error) {
		cnf := di.Config()
		val := is.NewBrowserTester(cnf.Driver)
		return val, nil
	}).(selenium.BrowserTester)
}

func (di DI) retrieveCache(key string, initFunc func() (interface{}, error)) interface{} {
	if di[key] == nil {
		val, err := initFunc()
		if err != nil {
			panic(fmt.Errorf("error while initializing in DI: %s", err.Error()))
		}
		di[key] = val
	}
	return di[key]
}