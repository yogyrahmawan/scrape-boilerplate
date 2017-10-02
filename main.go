package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	surf "gopkg.in/headzoo/surf.v1"
	"gopkg.in/headzoo/surf.v1/agent"
)

func main() {
	bow := surf.NewBrowser()
	bow.AddRequestHeader("Accept", "text/html")
	bow.AddRequestHeader("Accept-Charset", "utf8")

	//set user agent
	bow.SetUserAgent(agent.Chrome())

	if err := bow.Open("https://goodfil.ms/users/sign_in"); err != nil {
		panic(err)
	}

	// create login form
	loginForm, err := bow.Form("form.new_user")
	if err != nil {
		panic(err)
	}

	loginForm.Input("user[email]", "your email")
	loginForm.Input("user[password]", "your password")

	if err := loginForm.Submit(); err != nil {
		panic(err)
	}

	fmt.Println(bow.Title())

	// click the user
	if err := bow.Click("a.session-user"); err != nil {
		panic(err)
	}

	fmt.Println(bow.Title())

	// print reviews
	bow.Find("div.column.reviews.padded").Each(func(_ int, s *goquery.Selection) {
		fmt.Println(s.Text())
	})
}
