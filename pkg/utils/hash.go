package utils

import "github.com/matthewhartstonge/argon2"

var argon argon2.Config

func InitHash() {
	argon = argon2.DefaultConfig()
}

func Hash(password *string) (string, error) {
	encoded, err := argon.HashEncoded([]byte(*password))
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}
