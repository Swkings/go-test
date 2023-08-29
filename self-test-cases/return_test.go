package test

import (
	"fmt"
	"testing"
)

type RPCClient struct {
	Name string
}

func newClient() RPCClient {
	return RPCClient{
		Name: "test",
	}
}

func NewClient[Client any](empty bool, newClient func() Client) (emptyClient *Client) {
	if empty {
		return emptyClient
	}

	return PtrConvert(newClient())
}
func TestReturn(t *testing.T) {
	fmt.Println(NewClient(true, newClient))
	fmt.Println(NewClient(false, newClient))
	fmt.Println(NewClient(false, newClient).Name)
}
