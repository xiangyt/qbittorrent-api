package qbittorrent_api

import "testing"

func TestLogin(t *testing.T) {
	a := authorization{
		isLoggedIn: false,
		request: request{
			Jar:  nil,
			host: host,
		},
	}
	a.request.initialize()
	a.Login(name, pwd)
}
