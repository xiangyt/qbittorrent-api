package qbittorrent_api

import "testing"

func TestLogin(t *testing.T) {
	a := Authorization{
		isLoggedIn: false,
		Request: Request{
			Jar:  nil,
			host: host,
		},
	}
	a.Request.initialize()
	a.Login(name, pwd)
}
