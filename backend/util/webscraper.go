package util

import (
	"fmt"
	"log"
	"strings"

	"github.com/playwright-community/playwright-go"
)

func InitBrowser() (browser playwright.Browser){
	var err error
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("Could not start Playwright: %v", err)
	}
	browser, err = pw.Firefox.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
	})
	if err != nil {
		log.Fatalf("Could not launch browser: %v", err)
	}

	fmt.Println("Successfully started webscraper.")

	return browser
}


func GetQuestionInfo(url string, browser playwright.Browser) (title string, difficulty string) {
	page, err := browser.NewPage(playwright.BrowserNewPageOptions{
		BypassCSP: playwright.Bool(true),
	}) 
	if err != nil {
		log.Fatalf("Could not create page: %v", err)
	}
	defer page.Close()

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