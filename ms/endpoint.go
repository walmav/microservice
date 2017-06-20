package microservice

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	_ "github.com/go-kit/kit/log"
)

func MakeGetIdEndpoint(svc Node) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(getIdRequest)
		v, err := svc.GetId(req.S)
		if err != nil {
			return getIdResponse{v}, err

		}
		return getIdResponse{v}, nil
	}
}

type getIdRequest struct {
	S string `json:"s"`
}

type getIdResponse struct {
	S   string `json:"s"`
}
