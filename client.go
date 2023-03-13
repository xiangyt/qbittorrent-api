package qbittorrent_api

type Client struct {
	authorization
	applicationApi
	torrentsApi
}

type Option func(client *Client)

func NewClient(host string, opts ...Option) *Client {

	c := &Client{
		authorization: authorization{
			isLoggedIn: false,
			request: request{
				host: host,
			},
		},
	}
	c.request.initialize()
	c.torrentsApi.client = c
	c.applicationApi.client = c

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func WithAuth(username, password string) Option {
	return func(client *Client) {
		client.authorization.username = username
		client.authorization.password = password
	}
}
