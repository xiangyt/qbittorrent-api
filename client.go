package qbittorrent_api

type Client struct {
	Authorization
	TorrentsApi
}

type Option func(client *Client)

func NewClient(host, username, password string, opts ...Option) *Client {

	c := &Client{
		Authorization: Authorization{
			isLoggedIn: false,
			username:   username,
			password:   password,
			Request: Request{
				host: host,
			},
		},
	}
	c.Request.initialize()
	c.TorrentsApi.client = c

	for _, opt := range opts {
		opt(c)
	}

	return c
}
