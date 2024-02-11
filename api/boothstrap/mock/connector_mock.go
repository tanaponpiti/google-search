package boothstrap_test

import (
	"github.com/stretchr/testify/mock"
)

type MockHTMLRetrieverConnector struct {
	mock.Mock
}

func (m *MockHTMLRetrieverConnector) GetRenderedHTML(url string) (*string, error) {
	args := m.Called(url)
	return args.Get(0).(*string), args.Error(1)
}
