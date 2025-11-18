package redis

type Client struct {
	Addr string
}

func New(addr string) *Client {
	return &Client{Addr: addr}
}
