package nanoid

import (
	goNanoId "github.com/matoous/go-nanoid/v2"
)

const (
	defaultSize = 12

	defaultCharacters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

type NanoID string

func GetID() (NanoID, error) {
	id, err := goNanoId.Generate(defaultCharacters, defaultSize)
	if err != nil {
		return "", err
	}

	return NanoID(id), nil
}

func (nanoID NanoID) String() string {
	return string(nanoID)
}
