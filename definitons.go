package qbittorrent_api

type APIName = string

const (
	APINameAuthorization APIName = "auth"
	APINameApplication           = "app"
	APINameLog                   = "log"
	APINameSync                  = "sync"
	APINameTransfer              = "transfer"
	APINameTorrents              = "torrents"
	APINameRSS                   = "rss"
	APINameSearch                = "search"
	APINameEMPTY                 = ""
)

const (
	GoQBitVersion  = "v1.0.0"
	WebQBitVersion = "v4.4.3.1"
)
