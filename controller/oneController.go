package controller

type Client struct {
}

func (c *Client) Set(token string, id string) (err error) {
	return nil
}

func (c *Client) Get(token string) (id string, err error) {
	return token, nil
}
