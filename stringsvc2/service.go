package main

import (
	"errors"
	"fmt"
	"strings"
)

func Here() {
	fmt.Print("here")
}
type StringService interface {
	Uppercase(string) (string,error)
	Count(string) int
}


type stringService struct {

}

func (stringService) Uppercase (s string) (string, error) {
	fmt.Println("called")
	if s == "" {
		return "",ErrEmpty
	}

	return strings.ToUpper(s), nil

}

func (stringService) Count (s string) int {
	return len(s)
}

var ErrEmpty = errors.New("string is empty")
