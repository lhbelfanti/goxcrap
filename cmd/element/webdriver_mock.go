package element

import (
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/log"
)

// MockWebDriver is a mock implementation of WebDriver.
type MockWebDriver struct {
	mock.Mock
}

func (m *MockWebDriver) Status() (*selenium.Status, error) {
	args := m.Called()
	return args.Get(0).(*selenium.Status), args.Error(1)
}

func (m *MockWebDriver) NewSession() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *MockWebDriver) SessionId() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockWebDriver) SessionID() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockWebDriver) SwitchSession(sessionID string) error {
	args := m.Called(sessionID)
	return args.Error(0)
}

func (m *MockWebDriver) Capabilities() (selenium.Capabilities, error) {
	args := m.Called()
	return args.Get(0).(selenium.Capabilities), args.Error(1)
}

func (m *MockWebDriver) SetAsyncScriptTimeout(timeout time.Duration) error {
	args := m.Called(timeout)
	return args.Error(0)
}

func (m *MockWebDriver) SetImplicitWaitTimeout(timeout time.Duration) error {
	args := m.Called(timeout)
	return args.Error(0)
}

func (m *MockWebDriver) SetPageLoadTimeout(timeout time.Duration) error {
	args := m.Called(timeout)
	return args.Error(0)
}

func (m *MockWebDriver) Quit() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockWebDriver) CurrentWindowHandle() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *MockWebDriver) WindowHandles() ([]string, error) {
	args := m.Called()
	return args.Get(0).([]string), args.Error(1)
}

func (m *MockWebDriver) CurrentURL() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *MockWebDriver) Title() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *MockWebDriver) PageSource() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *MockWebDriver) Close() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockWebDriver) SwitchFrame(frame interface{}) error {
	args := m.Called(frame)
	return args.Error(0)
}

func (m *MockWebDriver) SwitchWindow(name string) error {
	args := m.Called(name)
	return args.Error(0)
}

func (m *MockWebDriver) CloseWindow(name string) error {
	args := m.Called(name)
	return args.Error(0)
}

func (m *MockWebDriver) MaximizeWindow(name string) error {
	args := m.Called(name)
	return args.Error(0)
}

func (m *MockWebDriver) ResizeWindow(name string, width, height int) error {
	args := m.Called(name, width, height)
	return args.Error(0)
}

func (m *MockWebDriver) Get(url string) error {
	args := m.Called(url)
	return args.Error(0)
}

func (m *MockWebDriver) Forward() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockWebDriver) Back() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockWebDriver) Refresh() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockWebDriver) FindElement(by, value string) (selenium.WebElement, error) {
	args := m.Called(by, value)
	return args.Get(0).(selenium.WebElement), args.Error(1)
}

func (m *MockWebDriver) FindElements(by, value string) ([]selenium.WebElement, error) {
	args := m.Called(by, value)
	return args.Get(0).([]selenium.WebElement), args.Error(1)
}

func (m *MockWebDriver) ActiveElement() (selenium.WebElement, error) {
	args := m.Called()
	return args.Get(0).(selenium.WebElement), args.Error(1)
}

func (m *MockWebDriver) DecodeElement(data []byte) (selenium.WebElement, error) {
	args := m.Called(data)
	return args.Get(0).(selenium.WebElement), args.Error(1)
}

func (m *MockWebDriver) DecodeElements(data []byte) ([]selenium.WebElement, error) {
	args := m.Called(data)
	return args.Get(0).([]selenium.WebElement), args.Error(1)
}

func (m *MockWebDriver) GetCookies() ([]selenium.Cookie, error) {
	args := m.Called()
	return args.Get(0).([]selenium.Cookie), args.Error(1)
}

func (m *MockWebDriver) GetCookie(name string) (selenium.Cookie, error) {
	args := m.Called(name)
	return args.Get(0).(selenium.Cookie), args.Error(1)
}

func (m *MockWebDriver) AddCookie(cookie *selenium.Cookie) error {
	args := m.Called(cookie)
	return args.Error(0)
}

func (m *MockWebDriver) DeleteAllCookies() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockWebDriver) DeleteCookie(name string) error {
	args := m.Called(name)
	return args.Error(0)
}

func (m *MockWebDriver) Click(button int) error {
	args := m.Called(button)
	return args.Error(0)
}

func (m *MockWebDriver) DoubleClick() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockWebDriver) ButtonDown() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockWebDriver) ButtonUp() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockWebDriver) SendModifier(modifier string, isDown bool) error {
	args := m.Called(modifier, isDown)
	return args.Error(0)
}

func (m *MockWebDriver) KeyDown(keys string) error {
	args := m.Called(keys)
	return args.Error(0)
}

func (m *MockWebDriver) KeyUp(keys string) error {
	args := m.Called(keys)
	return args.Error(0)
}

func (m *MockWebDriver) Screenshot() ([]byte, error) {
	args := m.Called()
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockWebDriver) Log(typ log.Type) ([]log.Message, error) {
	args := m.Called(typ)
	return args.Get(0).([]log.Message), args.Error(1)
}

func (m *MockWebDriver) DismissAlert() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockWebDriver) AcceptAlert() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockWebDriver) AlertText() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *MockWebDriver) SetAlertText(text string) error {
	args := m.Called(text)
	return args.Error(0)
}

func (m *MockWebDriver) ExecuteScript(script string, args []interface{}) (interface{}, error) {
	callArgs := m.Called(script, args)
	return callArgs.Get(0), callArgs.Error(1)
}

func (m *MockWebDriver) ExecuteScriptAsync(script string, args []interface{}) (interface{}, error) {
	callArgs := m.Called(script, args)
	return callArgs.Get(0), callArgs.Error(1)
}

func (m *MockWebDriver) ExecuteScriptRaw(script string, args []interface{}) ([]byte, error) {
	callArgs := m.Called(script, args)
	return callArgs.Get(0).([]byte), callArgs.Error(1)
}

func (m *MockWebDriver) ExecuteScriptAsyncRaw(script string, args []interface{}) ([]byte, error) {
	callArgs := m.Called(script, args)
	return callArgs.Get(0).([]byte), callArgs.Error(1)
}

func (m *MockWebDriver) WaitWithTimeoutAndInterval(condition selenium.Condition, timeout, interval time.Duration) error {
	args := m.Called(condition, timeout, interval)
	return args.Error(0)
}

func (m *MockWebDriver) WaitWithTimeout(condition selenium.Condition, timeout time.Duration) error {
	args := m.Called(condition, timeout)
	return args.Error(0)
}

func (m *MockWebDriver) Wait(condition selenium.Condition) error {
	args := m.Called(condition)
	return args.Error(0)
}
