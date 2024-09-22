package utils

import "encoding/json"

func ParseToken(obj any) (string, error) {
	req, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}

	token := struct {
		Val string `json:"token"`
	}{}

	err = json.Unmarshal(req, &token)
	if err != nil {
		return "", nil
	}

	return token.Val, nil
}
