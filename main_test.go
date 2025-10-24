package main

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)


func TestCheckToken (t *testing.T) {
	app := InitApp()

	// request without token
	req := httptest.NewRequest("GET" , "/check-token" , nil)
	req.Header.Set("Authorization" , "")
	res , _ := app.Test(req)
    if res.StatusCode != fiber.StatusUnauthorized {
        t.Errorf("Expected status %d, got %d" , fiber.StatusUnauthorized , res.StatusCode)
	}

	// request with token
	req = httptest.NewRequest("GET" , "/check-token" , nil)
	req.Header.Set("Authorization" , "Bearer mock-token")
	res , _ = app.Test(req)
    if res.StatusCode != fiber.StatusOK {
        t.Errorf("Expected status %d, got %d" , fiber.StatusUnauthorized , res.StatusCode)
	}
}