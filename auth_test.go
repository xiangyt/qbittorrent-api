package qbittorrent_api

import "testing"

func TestLogin(t *testing.T) {
	a := Authorization{
		isLoggedIn: false,
		username:   "admin",
		password:   "xyt303229577",
		Request: Request{
			Jar:  nil,
			host: "http://nas.zerotier.xyt:8999",
		},
	}
	a.Request.initialize()
	a.Login()
}
