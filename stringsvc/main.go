package main1

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"log"
	"net/http"
	"strings"
)

// interface
type StringService interface {
	Uppercase(string) (string,error)
	Count(string) int
}

// implementation

type stringService struct {

}

var ErrEmpty = errors.New("string is empty")


func (stringService) Uppercase (s string) (string, error) {
	if s == "" {
		return "",ErrEmpty
	}

	return strings.ToUpper(s), nil

}

func (stringService) Count (s string) int {
	return len(s)
}


// defining request and response

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

// defining endpoints

func makeUppercaseEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(uppercaseRequest)
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
func makeCountEndpoint(svc stringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(countRequest)
		count := svc.Count(req.S)
		return countResponse{V:count},nil
	}
}
//defining transport



func main() {

	svc := stringService{}
	uppercaseHandler := httptransport.NewServer(
			makeUppercaseEndpoint(svc),
			decodeUppercaseRequest,
			encodeResponse,
		)
	countHandler := httptransport.NewServer(
		makeCountEndpoint(svc),
		decodeCountRequest,
		encodeResponse,
	)

	http.Handle("/uppercase",uppercaseHandler)
	http.Handle("/count",countHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
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