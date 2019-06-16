//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE

package selenium

// BrowserTester defines interface of selenium command
type BrowserTester interface {
	Start() error
	Stop() error
	Navigate(url string) error
	FillByID(id, text string) error
	ClickByID(id string) error
	ScreenShot(path string) error
}
