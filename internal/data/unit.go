package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidUnitFormat = errors.New("invalid unit format")

type Unit int64

func (u Unit) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d baht", u)

	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil
}

func (u *Unit) UnmarshalJSON(jsonValue []byte) error {
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidUnitFormat
	}

	parts := strings.Split(unquotedJSONValue, " ")

	if len(parts) != 2 || parts[1] != "baht" {
		return ErrInvalidUnitFormat
	}

	i, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return ErrInvalidUnitFormat
	}

	*u = Unit(i)

	return nil
}
