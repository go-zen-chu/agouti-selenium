package selenium_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSelenium(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Selenium Suite")
}
