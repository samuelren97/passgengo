package models

import "passgengo/internal/hashing"

type GeneratorModel struct {
	Length         int            `json:"length"`
	HashingMethod  hashing.Method `json:"hashingMethod"`
	NoUpperChars   bool           `json:"noUpperChars"`
	NoSpecialChars bool           `json:"noSpecialChars"`
}
