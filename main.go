package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	surf "gopkg.in/headzoo/surf.v1"
	"gopkg.in/headzoo/surf.v1/agent"
	"net/http/cookiejar"
	"net/url"
)

func main() {
	bow := surf.NewBrowser()
	jar, _ := cookiejar.New(nil)

	bow.AddRequestHeader("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	bow.AddRequestHeader("Accept-Charset", "utf8")

	//set user agent
	bow.SetUserAgent(agent.Firefox())

	if err := bow.Open("https://m.facebook.com"); err != nil {
		panic(err)
	}

	//set cookies since facebook need cookies
	u, _ := url.Parse("https://m.facebook.com")
	jar.SetCookies(u, bow.SiteCookies())
	bow.SetCookieJar(jar)

	fmt.Println(bow.Title())

	// create login form
	loginForm, err := bow.Form("form#login_form")
	if err != nil {
		panic(err)
	}
	/*
	fmt.Println("value")
	fmt.Println(loginForm.Value("lgnjs"))
	fmt.Println(loginForm.Value("lgndim"))
	fmt.Println(loginForm.Value("lgnrnd"))
	fmt.Println(loginForm.Value("prefill_contact_point"))
	fmt.Println(loginForm.Value("lsd"))
	fmt.Println(loginForm.Value("login_source"))
	fmt.Println(loginForm.Value("skstamp"))
	fmt.Println(loginForm.Value("ab_test_data"))
	fmt.Println(loginForm.Value("locale"))

	fmt.Println(loginForm.Value("lsd"))
	fmt.Println(loginForm.Value("m_ts"))
	fmt.Println(loginForm.Value("li"))
	*/

	if err := loginForm.Input("email", "xxx"); err != nil {
		fmt.Println(err)
	}

	if err := loginForm.Input("pass", "xxx"); err != nil {
		fmt.Println(err)
	}

	if err := loginForm.Submit(); err != nil {
		panic(err)
	}

	// set cookies again
	jar.SetCookies(u, bow.SiteCookies())
	bow.SetCookieJar(jar)

	fmt.Println(bow.Title())
	fmt.Println(bow.StatusCode())
	fmt.Println(bow.Url())

	if err := bow.Open("https://facebook.com"); err != nil {
		panic(err)
	}

	fmt.Println(bow.Title())
	fmt.Println(bow.StatusCode())
	fmt.Println(bow.Url())

	//set cookies 
	jar.SetCookies(u, bow.SiteCookies())
	bow.SetCookieJar(jar)

	// print blue
	bow.Find("div#pagelet_bluebar").Each(func(_ int, s *goquery.Selection) {
		fmt.Println(s.Text())
	})
}
