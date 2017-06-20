package ms

import "strings"

type Node interface {
	GetId(string) (string, error)
}

type NodeImpl struct {
}

func (NodeImpl) GetId(s string) (string, error) {
	return strings.ToLower("spiffe://{domain}/{element}/{service_account_name}/{service_name}"), nil

}
