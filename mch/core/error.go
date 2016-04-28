package core

import (
	"fmt"
)

var _ error = (*Error)(nil)

type Error struct {
	XMLName    struct{} `xml:"xml"                  json:"-"`
	ReturnCode string   `xml:"return_code"          json:"return_code"`
	ReturnMsg  string   `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
}

func (err *Error) Error() string {
	return fmt.Sprintf("return_code: %s, return_msg: %s", err.ReturnCode, err.ReturnMsg)
}
