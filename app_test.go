package qbittorrent_api

import (
	"testing"
)

func TestVersion(t *testing.T) {
	c := NewClient(host)
	if version, err := c.Version(); err != nil {
		t.Error(err)
	} else {
		t.Log(version)
	}
}
