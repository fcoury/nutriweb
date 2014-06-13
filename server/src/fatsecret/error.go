package fatsecret

import (
	"strings"
)

func CheckError(result []byte) (*Error, error) {
	if strings.Contains(string(result), "<error") {
		error, err := ParseError(result)
		if err != nil {
			return nil, err
		}
		return &error, nil
	}
	return nil, nil
}
