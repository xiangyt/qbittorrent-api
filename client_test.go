package qbittorrent_api

import "testing"

func TestNewClient(t *testing.T) {
	c := NewClient("http://nas.zerotier.xyt:8999", "admin", "xyt303229577")
	c.Login()
	if ts, err := c.Torrents(nil); err != nil {
		t.Error(err)
	} else {
		for _, torrent := range ts {
			t.Log(torrent)
		}
	}
}
