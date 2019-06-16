package selenium

import (
	"fmt"
	ds "github.com/go-zen-chu/agouti-selenium/pkg/domain/selenium"
	"github.com/sclevine/agouti"
)

type browserTester struct {
	driver *agouti.WebDriver
	page   *agouti.Page
}

func NewBrowserTester(driver string) ds.BrowserTester {
	b := &browserTester{}
	switch driver {
	case "chrome":
		b.driver = agouti.ChromeDriver(agouti.Debug)
	case "phantomjs":
		b.driver = agouti.PhantomJS(agouti.Debug)
	}
	return b
}

func (b *browserTester) Start() error {
	err := b.driver.Start()
	if err != nil {
		return fmt.Errorf("error starting %s", err.Error())
	}
	b.page, err = b.driver.NewPage()
	if err != nil {
		return fmt.Errorf("error while creating new page %s", err.Error())
	}
	return nil
}

func (b *browserTester) Stop() error {
	if b.driver != nil {
		err := b.driver.Stop()
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *browserTester) Navigate(url string) error {
	return b.page.Navigate(url)
}

func (b *browserTester) FillByID(id, text string) error {
	elem := b.page.FindByID(id)
	if elem == nil {
		return fmt.Errorf("not found")
	}
	err := elem.Fill(text)
	return err
}

func (b *browserTester) ClickByID(id string) error {
	elem := b.page.FindByID(id)
	if elem == nil {
		return fmt.Errorf("not found")
	}
	err := elem.Click()
	return err
}

func (b *browserTester) ScreenShot(path string) error {
	if b.page != nil {
		b.page.Screenshot(path)
		return nil
	} else {
		return fmt.Errorf("no page for screenshot")
	}
}