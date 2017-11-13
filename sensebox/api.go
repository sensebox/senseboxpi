package sensebox

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
)

type APIConnection struct {
	baseUrl string
}

type

func NewAPIConnection(baseUrl string) *APIConnection {
	return &APIConnection{baseUrl: baseUrl}
}

func (a *APIConnection) FetchBox() {
	_, body, errs := gorequest.New().Get("https://api.osem.vo1d.space/boxes/57fb712811347b0011c10e80").End()
	if errs != nil {
		fmt.Println(errs)
	}
	fmt.Println(body)
}
