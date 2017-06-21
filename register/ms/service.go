package ms

import (
	"fmt"
)

type Node interface {
	Register(string) (string, error)
}

type NodeImpl struct {
}

func (NodeImpl) Register(s string) (string, error) {
	return fmt.Sprintf("Registering org:%s",s), nil

}
