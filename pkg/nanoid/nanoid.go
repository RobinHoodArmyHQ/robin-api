package nanoid

import (
	goNanoId "github.com/matoous/go-nanoid/v2"
)

const (
	defaultSize = 12

	defaultCharacters = "_-0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func GetID() (string, error) {
	id, err := goNanoId.Generate(defaultCharacters, defaultSize)
	if err != nil {
		return "", err
	}

	return id, nil
}
