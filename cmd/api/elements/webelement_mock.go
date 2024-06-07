package elements

import (
	"github.com/stretchr/testify/mock"
	"github.com/tebeka/selenium"
)

// MockWebElement is a mock implementation of WebElement for testing purposes
type MockWebElement struct {
	mock.Mock
}

func (m *MockWebElement) Click() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockWebElement) SendKeys(keys string) error {
	args := m.Called(keys)
	return args.Error(0)
}

func (m *MockWebElement) Submit() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockWebElement) Clear() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockWebElement) MoveTo(xOffset, yOffset int) error {
	args := m.Called(xOffset, yOffset)
	return args.Error(0)
}

func (m *MockWebElement) FindElement(by, value string) (selenium.WebElement, error) {
	args := m.Called(by, value)
	return args.Get(0).(selenium.WebElement), args.Error(1)
}

func (m *MockWebElement) FindElements(by, value string) ([]selenium.WebElement, error) {
	args := m.Called(by, value)
	return args.Get(0).([]selenium.WebElement), args.Error(1)
}

func (m *MockWebElement) TagName() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *MockWebElement) Text() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *MockWebElement) IsSelected() (bool, error) {
	args := m.Called()
	return args.Bool(0), args.Error(1)
}

func (m *MockWebElement) IsEnabled() (bool, error) {
	args := m.Called()
	return args.Bool(0), args.Error(1)
}

func (m *MockWebElement) IsDisplayed() (bool, error) {
	args := m.Called()
	return args.Bool(0), args.Error(1)
}

func (m *MockWebElement) GetAttribute(name string) (string, error) {
	args := m.Called(name)
	return args.String(0), args.Error(1)
}

func (m *MockWebElement) Location() (*selenium.Point, error) {
	args := m.Called()
	return args.Get(0).(*selenium.Point), args.Error(1)
}

func (m *MockWebElement) LocationInView() (*selenium.Point, error) {
	args := m.Called()
	return args.Get(0).(*selenium.Point), args.Error(1)
}

func (m *MockWebElement) Size() (*selenium.Size, error) {
	args := m.Called()
	return args.Get(0).(*selenium.Size), args.Error(1)
}

func (m *MockWebElement) CSSProperty(name string) (string, error) {
	args := m.Called(name)
	return args.String(0), args.Error(1)
}

func (m *MockWebElement) Screenshot(scroll bool) ([]byte, error) {
	args := m.Called(scroll)
	return args.Get(0).([]byte), args.Error(1)
}
