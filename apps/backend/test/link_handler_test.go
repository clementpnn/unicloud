package test

import (
	"backend/api/handler"
	"backend/domain/model"
	"backend/service"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockLinkService struct {
	mock.Mock
}

var _ service.LinkService = (*MockLinkService)(nil)

func (m *MockLinkService) CreateShortURL(longURL string) (*model.Link, error) {
	args := m.Called(longURL)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Link), args.Error(1)
}

func (m *MockLinkService) GetByShortURL(shortURL string) (*model.Link, error) {
	args := m.Called(shortURL)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Link), args.Error(1)
}

func (m *MockLinkService) GetOriginalURL(shortURL string) (string, error) {
	args := m.Called(shortURL)
	return args.String(0), args.Error(1)
}

func TestCreateShortURL(t *testing.T) {
	app := fiber.New()
	mockService := new(MockLinkService)
	handler := handler.NewLinkHandler(mockService)

	app.Post("/api/v1/shorten", handler.CreateShortURL)

	mockLink := &model.Link{ShortURL: "abc123", LongURL: "https://example.com"}
	mockService.On("CreateShortURL", "https://example.com").Return(mockLink, nil)

	reqBody := `{"url":"https://example.com"}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/shorten", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestRedirectToURL(t *testing.T) {
	app := fiber.New()
	mockService := new(MockLinkService)
	handler := handler.NewLinkHandler(mockService)
	app.Get("/:shortURL", handler.RedirectToURL)

	mockLink := &model.Link{ShortURL: "abc123", LongURL: "https://example.com"}
	mockService.On("GetByShortURL", "abc123").Return(mockLink, nil)

	req := httptest.NewRequest(http.MethodGet, "/abc123", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusMovedPermanently, resp.StatusCode)
}
