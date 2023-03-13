package qbittorrent_api

import (
	"net/url"
	"testing"
)

func TestJoinUrl(t *testing.T) {
	t.Log(url.JoinPath("example.com", "auth", "login"))
}
