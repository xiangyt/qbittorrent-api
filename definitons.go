package qbittorrent_api

type apiName = string

const (
	apiNameAuthorization apiName = "auth"
	apiNameApplication           = "app"
	apiNameLog                   = "log"
	apiNameSync                  = "sync"
	apiNameTransfer              = "transfer"
	apiNameTorrents              = "torrents"
	apiNameRSS                   = "rss"
	apiNameSearch                = "search"
	apiNameEMPTY                 = ""
)

const (
	GoQBitVersion  = "v1.0.0"
	WebQBitVersion = "v4.4.3.1"
)
