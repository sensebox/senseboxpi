package sensebox

import (
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/parnurzeal/gorequest"
)

const (
	baseURL = "https://api.osem.vo1d.space/boxes"
)

type APIConnection struct {
	boxID string
}

func validateID(id string) (err error) {
	hex, err := hex.DecodeString(id)
	if err != nil {
		return
	}

	if len(hex) != 12 {
		return errors.New("id must be exactly 24 characters long")
	}

	return
}

func NewAPIConnection(boxID string) (*APIConnection, error) {
	err := validateID(boxID)
	if err != nil {
		return nil, err
	}

	return &APIConnection{boxID: boxID}, nil
}

func (a *APIConnection) FetchBox() {
	var box SenseBox
	_, _, errs := gorequest.New().Get(baseURL + "/" + a.boxID).EndStruct(&box)
	if errs != nil {
		fmt.Println(errs)
	}

	fmt.Println(box)
}
