package hashing

import "fmt"

type Method int

const (
	None = iota
	SHA256
)

var HashMethods = [...]string{
	"None",
	"SHA256",
}

func IntToMethod(m int) (Method, error) {
	if m < 0 || m >= len(HashMethods) {
		return 0, fmt.Errorf("Could not find hash method with int: %d", m)
	}

	return Method(m), nil
}

func (m Method) String() string {
	return HashMethods[m]
}
