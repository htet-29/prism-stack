package data

import (
	"fmt"
	"strconv"
)

type Unit int64

func (u Unit) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d baht", u)
	
	quotedJSONValue := strconv.Quote(jsonValue)
	
	return []byte(quotedJSONValue), nil
}