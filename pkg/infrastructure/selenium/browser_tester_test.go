package selenium_test

import (
	"github.com/go-zen-chu/agouti-selenium/pkg/infrastructure/selenium"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

)

var _ = Describe("Browser", func() {

	Describe("NewBrowserTester", func() {
		It("should be created", func() {
			b := selenium.NewBrowserTester("phantomjs")
			Expect(b).NotTo(BeNil())
		})
	})
})
