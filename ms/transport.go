package microservice


import (
	"context"
	"encoding/json"
	"net/http"

)


func DecodeGetIdEndpointRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getIdRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
