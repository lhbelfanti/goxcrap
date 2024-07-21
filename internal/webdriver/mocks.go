package webdriver

import (
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/log"
)

// MockNewManager mocks NewManager function
func MockNewManager(manager Manager) NewManager {
	return func() Manager {
		return manager
	}
}

// Mock is a mock implementation of selenium.WebDriver.
type Mock struct {
	mock.Mock
}

func (m *Mock) Status() (*selenium.Status, error) {
	args := m.Called()
	return args.Get(0).(*selenium.Status), args.Error(1)
}

func (m *Mock) NewSession() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *Mock) SessionId() string {
	args := m.Called()
	return args.String(0)
}

func (m *Mock) SessionID() string {
	args := m.Called()
	return args.String(0)
}

func (m *Mock) SwitchSession(sessionID string) error {
	args := m.Called(sessionID)
	return args.Error(0)
}

func (m *Mock) Capabilities() (selenium.Capabilities, error) {
	args := m.Called()
	return args.Get(0).(selenium.Capabilities), args.Error(1)
}

func (m *Mock) SetAsyncScriptTimeout(timeout time.Duration) error {
	args := m.Called(timeout)
	return args.Error(0)
}

func (m *Mock) SetImplicitWaitTimeout(timeout time.Duration) error {
	args := m.Called(timeout)
	return args.Error(0)
}

func (m *Mock) SetPageLoadTimeout(timeout time.Duration) error {
	args := m.Called(timeout)
	return args.Error(0)
}

func (m *Mock) Quit() error {
	args := m.Called()
	return args.Error(0)
}

func (m *Mock) CurrentWindowHandle() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *Mock) WindowHandles() ([]string, error) {
	args := m.Called()
	return args.Get(0).([]string), args.Error(1)
}

func (m *Mock) CurrentURL() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *Mock) Title() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *Mock) PageSource() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *Mock) Close() error {
	args := m.Called()
	return args.Error(0)
}

func (m *Mock) SwitchFrame(frame interface{}) error {
	args := m.Called(frame)
	return args.Error(0)
}

func (m *Mock) SwitchWindow(name string) error {
	args := m.Called(name)
	return args.Error(0)
}

func (m *Mock) CloseWindow(name string) error {
	args := m.Called(name)
	return args.Error(0)
}

func (m *Mock) MaximizeWindow(name string) error {
	args := m.Called(name)
	return args.Error(0)
}

func (m *Mock) ResizeWindow(name string, width, height int) error {
	args := m.Called(name, width, height)
	return args.Error(0)
}

func (m *Mock) Get(url string) error {
	args := m.Called(url)
	return args.Error(0)
}

func (m *Mock) Forward() error {
	args := m.Called()
	return args.Error(0)
}

func (m *Mock) Back() error {
	args := m.Called()
	return args.Error(0)
}

func (m *Mock) Refresh() error {
	args := m.Called()
	return args.Error(0)
}

func (m *Mock) FindElement(by, value string) (selenium.WebElement, error) {
	args := m.Called(by, value)
	return args.Get(0).(selenium.WebElement), args.Error(1)
}

func (m *Mock) FindElements(by, value string) ([]selenium.WebElement, error) {
	args := m.Called(by, value)
	return args.Get(0).([]selenium.WebElement), args.Error(1)
}

func (m *Mock) ActiveElement() (selenium.WebElement, error) {
	args := m.Called()
	return args.Get(0).(selenium.WebElement), args.Error(1)
}

func (m *Mock) DecodeElement(data []byte) (selenium.WebElement, error) {
	args := m.Called(data)
	return args.Get(0).(selenium.WebElement), args.Error(1)
}

func (m *Mock) DecodeElements(data []byte) ([]selenium.WebElement, error) {
	args := m.Called(data)
	return args.Get(0).([]selenium.WebElement), args.Error(1)
}

func (m *Mock) GetCookies() ([]selenium.Cookie, error) {
	args := m.Called()
	return args.Get(0).([]selenium.Cookie), args.Error(1)
}

func (m *Mock) GetCookie(name string) (selenium.Cookie, error) {
	args := m.Called(name)
	return args.Get(0).(selenium.Cookie), args.Error(1)
}

func (m *Mock) AddCookie(cookie *selenium.Cookie) error {
	args := m.Called(cookie)
	return args.Error(0)
}

func (m *Mock) DeleteAllCookies() error {
	args := m.Called()
	return args.Error(0)
}

func (m *Mock) DeleteCookie(name string) error {
	args := m.Called(name)
	return args.Error(0)
}

func (m *Mock) Click(button int) error {
	args := m.Called(button)
	return args.Error(0)
}

func (m *Mock) DoubleClick() error {
	args := m.Called()
	return args.Error(0)
}

func (m *Mock) ButtonDown() error {
	args := m.Called()
	return args.Error(0)
}

func (m *Mock) ButtonUp() error {
	args := m.Called()
	return args.Error(0)
}

func (m *Mock) SendModifier(modifier string, isDown bool) error {
	args := m.Called(modifier, isDown)
	return args.Error(0)
}

func (m *Mock) KeyDown(keys string) error {
	args := m.Called(keys)
	return args.Error(0)
}

func (m *Mock) KeyUp(keys string) error {
	args := m.Called(keys)
	return args.Error(0)
}

func (m *Mock) Screenshot() ([]byte, error) {
	args := m.Called()
	return args.Get(0).([]byte), args.Error(1)
}

func (m *Mock) Log(typ log.Type) ([]log.Message, error) {
	args := m.Called(typ)
	return args.Get(0).([]log.Message), args.Error(1)
}

func (m *Mock) DismissAlert() error {
	args := m.Called()
	return args.Error(0)
}

func (m *Mock) AcceptAlert() error {
	args := m.Called()
	return args.Error(0)
}

func (m *Mock) AlertText() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *Mock) SetAlertText(text string) error {
	args := m.Called(text)
	return args.Error(0)
}

func (m *Mock) ExecuteScript(script string, args []interface{}) (interface{}, error) {
	callArgs := m.Called(script, args)
	return callArgs.Get(0), callArgs.Error(1)
}

func (m *Mock) ExecuteScriptAsync(script string, args []interface{}) (interface{}, error) {
	callArgs := m.Called(script, args)
	return callArgs.Get(0), callArgs.Error(1)
}

func (m *Mock) ExecuteScriptRaw(script string, args []interface{}) ([]byte, error) {
	callArgs := m.Called(script, args)
	return callArgs.Get(0).([]byte), callArgs.Error(1)
}

func (m *Mock) ExecuteScriptAsyncRaw(script string, args []interface{}) ([]byte, error) {
	callArgs := m.Called(script, args)
	return callArgs.Get(0).([]byte), callArgs.Error(1)
}

func (m *Mock) WaitWithTimeoutAndInterval(condition selenium.Condition, timeout, interval time.Duration) error {
	args := m.Called(condition, timeout, interval)
	return args.Error(0)
}

func (m *Mock) WaitWithTimeout(condition selenium.Condition, timeout time.Duration) error {
	args := m.Called(condition, timeout)
	return args.Error(0)
}

func (m *Mock) Wait(condition selenium.Condition) error {
	args := m.Called(condition)
	return args.Error(0)
}

// MockManager is a mock implementation of Manager.
type MockManager struct {
	mock.Mock
}

func (m *MockManager) InitWebDriverService() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockManager) InitWebDriver() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockManager) Quit() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockManager) WebDriver() selenium.WebDriver {
	args := m.Called()
	return args.Get(0).(selenium.WebDriver)
}
