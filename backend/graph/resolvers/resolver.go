package resolvers

//go:generate go run github.com/99designs/gqlgen generate

import (
	"github.com/playwright-community/playwright-go"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	DB *gorm.DB
	Browser playwright.Browser
}
