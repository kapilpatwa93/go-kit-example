package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"net/http"
)

type uppercaseRequest struct {
	S string `json:"s"`
}

type uppercaseResponse struct {
	V string `json:"v"`
	Err string `json:"err,omitempty"`
}

type countRequest struct {
	S string `json:"s"`
}

type countResponse struct {
	V int `json:"v"`
}


func makeUppercaseEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(uppercaseRequest)
		fmt.Println("inside uppercase endpoint function")
		v , err := svc.Uppercase(req.S)
		if err != nil {
			return uppercaseResponse{
				v,
				err.Error(),
			}, nil
		}
		return uppercaseResponse{

			v,
			"nil",
		}, nil
	}
}
func makeCountEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(countRequest)
		count := svc.Count(req.S)
		return countResponse{V:count},nil
	}
}

func decodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var uppercaseRequest  uppercaseRequest
	err := json.NewDecoder(r.Body).Decode(&uppercaseRequest)
	if err != nil {
		return nil, err
	}
	return uppercaseRequest, nil
}

func decodeCountRequest (_ context.Context, r *http.Request) (interface{}, error) {
	var request countRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context,w  http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}