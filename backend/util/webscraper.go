package util

import (
	"log"
	"strings"

	"github.com/playwright-community/playwright-go"
)

func GetQuestionInfo(url string) (title string, difficulty string) {
    pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("Could not start Playwright: %v", err)
	}
	browser, err := pw.Chromium.Launch() 
	if err != nil {
		log.Fatalf("Could not launch browser: %v", err)
	}
	page, err := browser.NewPage() 
	if err != nil {
		log.Fatalf("Could not create page: %v", err)
	}
	if _, err = page.Goto(url); err != nil {
		log.Fatalf("Could not goto: %v", err)
	}

	title, err = page.Locator("div.text-title-large").TextContent()
	if err != nil {
        log.Fatalf("Could not get title: %v", err)
    }
	title = ExtractTitle(title)

	difficulty, err = page.Locator("div.text-difficulty-easy, div.text-difficulty-medium, div.text-difficulty-hard").TextContent()
	if err != nil {
        log.Fatalf("Could not get difficulty: %v", err)
    }

	return title, difficulty
}

func ExtractTitle(title string) string {
	space := strings.Index(title, " ") + 1
	return title[space:]
}