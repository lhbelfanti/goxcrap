package chromedriver

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"regexp"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

const (
	chromeDriverPath        string = "./internal/chromedriver/driver"
	chromeDriverServicePort int    = 4444
)

// InitWebDriverService creates a new Chrome Driver Service
func InitWebDriverService() (*selenium.Service, error) {
	return selenium.NewChromeDriverService(chromeDriverPath, chromeDriverServicePort)
}

func StopWebDriverService(service *selenium.Service) {
	err := service.Stop()
	if err != nil {
		panic(err)
	}
}

// InitWebDriver creates a new Chrome WebDriver
func InitWebDriver() (selenium.WebDriver, error) {
	caps := selenium.Capabilities{}
	args := make([]string, 0, 3)
	args = append(args, "--headless-new") // comment out this line for testing
	customUserAgent := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36"
	args = append(args, "--user-agent="+customUserAgent)
	// TODO: Review why the page doesn't load if this arg is added
	/* proxyServerURL, err := getRandomProxyServerURL(customUserAgent)
	if err == nil && proxyServerURL != "" {
		args = append(args, "--proxy-server="+proxyServerURL)
	}*/
	caps.AddChrome(chrome.Capabilities{Args: args})

	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		return nil, err
	}

	// maximize the current window to avoid responsive rendering
	err = driver.MaximizeWindow("")
	if err != nil {
		return nil, err
	}

	return driver, nil
}

func getRandomProxyServerURL(userAgent string) (string, error) {
	// Create a new HTTP client
	client := http.Client{}

	// Create a new GET request
	req, err := http.NewRequest("GET", "https://free-proxy-list.net/#", nil)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error creating request: %v", err))
	}

	// Set any request headers if needed
	req.Header.Set("User-Agent", userAgent)

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error sending request: %v", err))
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error reading response: %v", err))
	}

	response := string(body)

	numBlock := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	port := "(:[0-9]{1,4})"
	regexPattern := numBlock + "\\." + numBlock + "\\." + numBlock + "\\." + numBlock + port
	regEx := regexp.MustCompile(regexPattern)
	ips := regEx.FindAllString(response, -1)

	source := rand.NewSource(time.Now().Unix())
	rng := rand.New(source)
	randomIndex := rng.Intn(len(ips))

	fmt.Println(ips[randomIndex])

	return ips[randomIndex], nil
}
