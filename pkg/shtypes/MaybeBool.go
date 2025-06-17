<<<<<<<< HEAD:pkg/types/MaybeBool.go
package types //nolint:var-naming
========
package shtypes
>>>>>>>> 13f5db68e6a2a700ee2cee509efb0b481bb4fff4:pkg/shtypes/MaybeBool.go

import (
	"strconv"
	"strings"
)

// MaybeBool is a simple wrapper unmarshalling json responses that could be either an int or an int in a string.
type MaybeBool bool

// UnmarshalJSON is a helper interface for dealing with things that may or may not be a string representing a bool, or  number representing a bool.
func (fi *MaybeBool) UnmarshalJSON(b []byte) error {
	maybeBool, err := strconv.ParseBool(strings.Trim(string(b), "\""))
	if err != nil {
		return err
	}

	*fi = MaybeBool(maybeBool)
	return nil
}
