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

const (
	chromeDriverPath        string = "./internal/driver/chrome"
	chromeDriverServicePort int    = 9515
)

// InitWebDriverService creates a new Chrome Driver Service
func InitWebDriverService(localMode bool) (*selenium.Service, error) {
	driverPath := os.Getenv("DRIVER_PATH")
	if driverPath == "" || localMode {
		driverPath = chromeDriverPath
	}

	slog.Info(fmt.Sprintf(color.BlueString("Initializing Chrome Driver Service using driver from:\n%s"), color.GreenString(driverPath)))
	return selenium.NewChromeDriverService(driverPath, chromeDriverServicePort)
}

// StopWebDriverService stops web driver service to avoid leaks if the app is terminated
func StopWebDriverService(service *selenium.Service) {
	err := service.Stop()
	if err != nil {
		panic(err)
	}
}

// QuitWebDriver quits web driver to avoid leaks if the app is terminated
func QuitWebDriver(webDriver selenium.WebDriver) {
	err := webDriver.Quit()
	if err != nil {
		panic(err)
	}
}

// InitWebDriver creates a new Chrome WebDriver
func InitWebDriver(localMode bool) (selenium.WebDriver, error) {
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
	}

	chromeCaps := chrome.Capabilities{}
	if !localMode {
		browserPath := os.Getenv("BROWSER_PATH")
		args = append(args, "--headless")

		if browserPath != "" {
			chromeCaps.Path = browserPath
		}
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
	driver, err := selenium.NewRemote(caps, remotePath)
	if err != nil {
		return nil, err
	}

	if localMode {
		// maximize the current window to avoid responsive rendering
		err = driver.MaximizeWindow("")
		if err != nil {
			return nil, err
		}
	}

	return driver, nil
}
