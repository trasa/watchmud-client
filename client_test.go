package main

func NewTestClient() *Client {
	return &Client{
		source: make(chan interface{}, 2),
	}
}
