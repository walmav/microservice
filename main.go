package main

import(
	httptransport "github.com/go-kit/kit/transport/http"
	"log"
	"net/http"
	"github.com/walmav/microservice/microservice"

)
func main() {
svc := microservice.NodeImpl{}

GetIdHandler := httptransport.NewServer(
	microservice.MakeGetIdEndpoint(svc),
	microservice.DecodeGetIdEndpointRequest,
	microservice.EncodeResponse,
)
	http.Handle("/getId", GetIdHandler)
	log.Print(http.ListenAndServe(":8080", nil))

}