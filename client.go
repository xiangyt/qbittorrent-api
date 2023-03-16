package qbittorrent_api

type Client struct {
	Authorization
	ApplicationApi
	TorrentsApi
}

type Option func(client *Client)

func NewClient(host string, opts ...Option) *Client {

	c := &Client{
		Authorization: Authorization{
			isLoggedIn: false,
			Request: Request{
				host: host,
			},
		},
	}
	c.Request.initialize()
	c.TorrentsApi.client = c
	c.ApplicationApi.client = c

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func WithAuth(username, password string) Option {
	return func(client *Client) {
		client.Authorization.username = username
		client.Authorization.password = password
	}
}
