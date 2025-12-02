package generator

import (
	"fmt"
	"passgengo/internal/hashing"
)

type GeneratorBuilder struct {
	instance *generator
}

func NewGeneratorBuilder() *GeneratorBuilder {
	return &GeneratorBuilder{instance: &generator{}}
}

func (gb *GeneratorBuilder) HexString() {
	gb.instance.hexString = true
}

func (gb *GeneratorBuilder) Length(length int) error {
	if length < MinPassLen || length > MaxPassLen {
		return fmt.Errorf("incorrect password length")
	}

	gb.instance.length = length
	return nil
}

func (gb *GeneratorBuilder) HashingMethod(m int) error {
	method, err := hashing.IntToMethod(int(m))
	if err != nil {
		return err
	}

	gb.instance.hashingMethod = method
	return nil
}

func (gb *GeneratorBuilder) NoUpperChars() {
	gb.instance.noUpperChars = true
}

func (gb *GeneratorBuilder) NoSpecialChars() {
	gb.instance.noSpecialChars = true
}

func (gb *GeneratorBuilder) Build() Generator {
	return gb.instance
}
