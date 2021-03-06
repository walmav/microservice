package main

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/walmav/microservice/register/ms"
	"log"
	"net/http"

)

func main() {
	svc := ms.NodeImpl{}

	GetIdHandler := httptransport.NewServer(
		ms.MakeGetIdEndpoint(svc),
		ms.DecodeGetIdEndpointRequest,
		ms.EncodeResponse,
	)
	http.Handle("/register", GetIdHandler)
	log.Print(http.ListenAndServe(":8080", nil))
}

