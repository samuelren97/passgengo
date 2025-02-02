package generator

import (
	"encoding/json"
	"fmt"
	"passgengo/internal/hashing"
	"passgengo/internal/logging"
	"passgengo/internal/models"
	"passgengo/internal/utils"
)

const (
	MinPassLen = 6
	MaxPassLen = 128

	lowerCharset   = "abcdefghijklmnopqrstuvwxyz"
	upperCharset   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	specialCharset = "!@#$%^&*()-_=+"
)

type Generator interface {
	Generate() (string, error)
	Json() (string, error)
}

type generator struct {
	length         int
	hashingMethod  hashing.Method
	noUpperChars   bool
	noSpecialChars bool
}

func (g *generator) Generate() (string, error) {
	var password []byte

	availableCharsets := g.getAvailableCharsets()

	for i := 0; i < g.length; i++ {
		randCharsetIndex, err := utils.RandomIntWithMax(len(availableCharsets))
		if err != nil {
			return "", err
		}

		num, err := utils.RandomIntWithMax(len(availableCharsets[randCharsetIndex.Int64()]))
		if err != nil {
			return "", err
		}
		password = append(password, availableCharsets[randCharsetIndex.Int64()][num.Int64()])
	}

	utils.LogDebug(string(password))

	switch {
	case g.hashingMethod == hashing.SHA256:
		logging.Info(fmt.Sprintf("Using a hashing method, cleartext password: %s", password))
		return hashing.HashSHA256(password)
	}

	return string(password), nil
}

func (g *generator) Json() (string, error) {
	jsonBytes, err := json.MarshalIndent(g.toModel(), "", "  ")
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}

func (g *generator) toModel() *models.GeneratorModel {
	return &models.GeneratorModel{
		Length:         g.length,
		HashingMethod:  g.hashingMethod,
		NoUpperChars:   g.noUpperChars,
		NoSpecialChars: g.noSpecialChars,
	}
}

func (g *generator) getAvailableCharsets() []string {
	availableCharsets := []string{lowerCharset}
	if !g.noUpperChars {
		availableCharsets = append(availableCharsets, upperCharset)
	}

	if !g.noSpecialChars {
		availableCharsets = append(availableCharsets, specialCharset)
	}

	return availableCharsets
}
