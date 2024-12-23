package endpoints_test

import (
	"regexp"
	"sync"
	"testing"

	"github.com/amfonelic/gomatcher/internal/endpoints"
	"github.com/amfonelic/gomatcher/pkg/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockEndpoint struct {
	mock.Mock
	Data chan string
}

func (m *MockEndpoint) HandleRequest(req any) error {
	args := m.Called(req)
	return args.Error(0)
}

func (m *MockEndpoint) SetupServer() {
	m.Called()
}

func (m *MockEndpoint) RunServer() {
	m.Called()
}

func (m *MockEndpoint) String() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockEndpoint) GetData() chan string {
	return m.Data
}

func TestRunServers(t *testing.T) {
	mockEndpoint1 := new(MockEndpoint)
	mockEndpoint2 := new(MockEndpoint)
	mockEndpoint1.On("SetupServer").Return()
	mockEndpoint2.On("SetupServer").Return()

	endpoints.RunServers([]endpoints.IEndpoint{mockEndpoint1, mockEndpoint2})

	mockEndpoint1.AssertCalled(t, "SetupServer")
	mockEndpoint2.AssertCalled(t, "SetupServer")
}

func TestComposeMatchPrintData(t *testing.T) {
	mockEndpoint := new(MockEndpoint)
	mockEndpoint.Data = make(chan string, 1)
	mockEndpoint.On("String").Return("MockEndpoint")
	mockEndpoint.Data <- "testData"

	pattern := regexp.MustCompile(`testData`)

	manager := &endpoints.DataManager{Data: make(map[string]string)}
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		endpoints.ComposeData(mockEndpoint, manager, &wg)
	}()

	wg.Wait()

	values := helpers.MapToSlice(manager.Data)
	patterns, err := helpers.FindPatterns(pattern, values)
	assert.NoError(t, err, "Expected no error in finding patterns")

	if len(patterns) > 1 {
		isMatched, err := helpers.AllStringsAreEqual(patterns)
		assert.NoError(t, err, "Expected no error in checking equality of strings")
		assert.True(t, isMatched, "Expected all strings to be equal")
	} else {
		t.Log("Skipping AllStringsAreEqual as the pattern slice has only one element")
	}
}

func TestComposeData(t *testing.T) {
	mockEndpoint := new(MockEndpoint)
	mockEndpoint.Data = make(chan string, 1)
	mockEndpoint.On("String").Return("MockEndpoint")
	mockEndpoint.Data <- "testData"

	manager := &endpoints.DataManager{Data: make(map[string]string)}
	var wg sync.WaitGroup
	wg.Add(1)

	endpoints.ComposeData(mockEndpoint, manager, &wg)

	wg.Wait()

	assert.Equal(t, "testData", manager.Data["MockEndpoint"], "Expected ComposeData to correctly populate manager data")
}
