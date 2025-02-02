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
		return 0, fmt.Errorf("could not find hash method with int: %d", m)
	}

	return Method(m), nil
}

func (m Method) String() string {
	return HashMethods[m]
}

func GetHashingMethodsString() string {
	hashingMethods := "\n"
	for i, method := range HashMethods {
		hashingMethods += fmt.Sprintf("\t%d -> %s", i, method)
		if i != len(HashMethods)-1 {
			hashingMethods += "\n"
		}
	}
	return hashingMethods
}
