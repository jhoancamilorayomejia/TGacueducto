package models

import (
	"encoding/json"
)

func Parser[T any](body []byte) (*T, error) {
	var user T

	err := json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
