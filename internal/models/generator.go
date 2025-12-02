package models

import "passgengo/internal/hashing"

type GeneratorModel struct {
	Length         int            `json:"length"`
	HexString      bool           `json:"hexString"`
	Base64String   bool           `json:"base64String"`
	HashingMethod  hashing.Method `json:"hashingMethod"`
	NoUpperChars   bool           `json:"noUpperChars"`
	NoSpecialChars bool           `json:"noSpecialChars"`
}
