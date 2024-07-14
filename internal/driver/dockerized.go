package driver

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

// DockerizedWebDriver represents a Web Driver for the dockerized version of this application
type DockerizedWebDriver struct{}

// NewDockerizedWebDriver creates a new DockerizedWebDriver
func NewDockerizedWebDriver() DockerizedWebDriver {
	return DockerizedWebDriver{}
}

// InitWebDriverService creates a new Chrome Driver Service
func (dwd DockerizedWebDriver) InitWebDriverService() (*selenium.Service, error) {
	driverPath := os.Getenv("DRIVER_PATH")
	if driverPath == "" {
		driverPath = chromeDriverPath
	}

	slog.Info(fmt.Sprintf(color.BlueString("Initializing Chrome Driver Service using driver from:\n%s"), color.GreenString(driverPath)))
	return selenium.NewChromeDriverService(driverPath, chromeDriverServicePort)
}

// StopWebDriverService stops web driver service to avoid leaks if the app is terminated
func (dwd DockerizedWebDriver) StopWebDriverService(service *selenium.Service) {
	err := service.Stop()
	if err != nil {
		panic(err)
	}
}

// QuitWebDriver quits web driver to avoid leaks if the app is terminated
func (dwd DockerizedWebDriver) QuitWebDriver(webDriver selenium.WebDriver) {
	err := webDriver.Quit()
	if err != nil {
		panic(err)
	}
}

// InitWebDriver creates a new Chrome WebDriver
func (dwd DockerizedWebDriver) InitWebDriver() (selenium.WebDriver, error) {
	browserPath := os.Getenv("BROWSER_PATH")

	args := []string{
		"--no-sandbox",
		"--disable-dev-shm-usage",
		"--disable-gpu",
		"--blink-settings=imagesEnabled=false",
		"--disable-extensions",
		"--disable-popup-blocking",
		"--disable-infobars",
		"--disable-logging",
		"--disable-notifications",
		"--disable-background-networking",
		"--disable-background-timer-throttling",
		"--disable-backgrounding-occluded-windows",
		"--disable-breakpad",
		"--disable-client-side-phishing-detection",
		"--disable-component-extensions-with-background-pages",
		"--disable-default-apps",
		"--disable-hang-monitor",
		"--disable-ipc-flooding-protection",
		"--disable-prompt-on-repost",
		"--disable-renderer-backgrounding",
		"--disable-sync",
		"--metrics-recording-only",
		"--mute-audio",
		"--no-first-run",
		"--safebrowsing-disable-auto-update",
		"--enable-automation",
		"--disable-blink-features=AutomationControlled",
		"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36",
		"--headless",
	}

	chromeCaps := chrome.Capabilities{}
	if browserPath != "" {
		chromeCaps.Path = browserPath
	}

	chromeCaps.Args = args

	slog.Info(fmt.Sprintf(color.BlueString("Setting up Chrome Capacities using the following Args:\n%s\n"), color.GreenString(strings.Join(chromeCaps.Args, "\n"))))
	if chromeCaps.Path != "" {
		slog.Info(fmt.Sprintf("and the following Path:\n%s", color.GreenString(chromeCaps.Path)))
	}

	caps := selenium.Capabilities{"browserName": "chrome"}
	caps.AddChrome(chromeCaps)

	remotePath := fmt.Sprintf("http://localhost:%d/wd/hub", chromeDriverServicePort)
	slog.Info(color.BlueString(fmt.Sprintf("Creating Remote Client at: \n%s", color.GreenString(remotePath))))
	wd, err := selenium.NewRemote(caps, remotePath)
	if err != nil {
		return nil, err
	}

	return wd, nil
}
