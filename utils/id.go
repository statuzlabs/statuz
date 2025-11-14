package utils

import (
	"fmt"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

const (
	nanoidAlphabet  = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_"
	defaultIDLength = 32
)

func GenerateID() (string, error) {
	id, err := gonanoid.Generate(nanoidAlphabet, defaultIDLength)
	if err != nil {
		return "", fmt.Errorf("failed to generate ID: %w", err)
	}
	return id, err
}

func MustGenerateID() string {
	id, err := gonanoid.Generate(nanoidAlphabet, defaultIDLength)
	if err != nil {
		panic(fmt.Sprintf("failed to generate ID: %v", err))
	}
	return id
}

func NewID(size int) (string, error) {
	if size < 1 || size > 100 {
		size = defaultIDLength
	}
	id, err := gonanoid.Generate(nanoidAlphabet, defaultIDLength)
	if err != nil {
		return "", fmt.Errorf("failed to generate ID: %w", err)
	}
	return id, err
}
