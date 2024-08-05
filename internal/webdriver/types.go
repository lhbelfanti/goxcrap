package webdriver

import "github.com/tebeka/selenium"

// Manager interface is adopted by the different implementations of a web driver manager used in this application
type Manager interface {
	// InitWebDriverService initializes a new *selenium.Service
	InitWebDriverService() error

	// InitWebDriver initializes a new selenium.WebDriver
	InitWebDriver() error

	// Quit stops the selenium.WebDriver and its *selenium.Service to avoid leaks if the app is terminated
	Quit() error

	// WebDriver returns the initialized selenium.WebDriver
	WebDriver() selenium.WebDriver
}

const (
	chromeDriverPath        string = "./internal/webdriver/chromedriver"
	chromeDriverServicePort int    = 9515
)

var (
	capabilitiesPreferences = map[string]interface{}{
		"profile.default_content_setting_values.media_stream": 2, // Disable media stream
		"profile.managed_default_content_settings.images":     2, // Disable images
	}

	capabilitiesArgs = []string{
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
)
