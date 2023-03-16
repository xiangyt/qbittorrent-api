package qbittorrent_api

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
)

type Authorization struct {
	isLoggedIn bool
	username   string
	password   string
	Request
}

func (a *Authorization) IsLoggedIn() bool {
	return a.isLoggedIn
}

func (a *Authorization) Login(username, password string) error {
	a.username = username
	a.password = password
	resp, err := a.post(APINameAuthorization, "login", map[string]string{
		"username": a.username,
		"password": a.password,
	})
	if err != nil {
		logrus.Error("Login failed")
		return err
	}
	defer resp.Body.Close()

	logrus.Debug("Login successful")
	a.isLoggedIn = true

	if cookies := resp.Cookies(); len(cookies) > 0 {
		cookieURL, _ := url.Parse(a.host)
		a.Request.Jar.SetCookies(cookieURL, cookies)
	}

	// create a new client with the cookie jar and replace the old one
	// so that all our later requests are authenticated
	a.http = &http.Client{
		Jar: a.Jar,
	}
	return nil
}

func (a *Authorization) Logout() {
	resp, err := a.post(APINameAuthorization, "logout", nil)
	if err == nil {
		resp.Body.Close()
	}
	a.Request.initialize()
}
