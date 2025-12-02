package utils

import (
	"crypto/rand"
	"math/big"
)

func RandomIntWithMax(max int) (*big.Int, error) {
	return rand.Int(rand.Reader, big.NewInt(int64(max)))
}

func ShuffleBytes(b *[]byte) error {
	nBytes := *b
	for j := range nBytes {
		randIndex, err := RandomIntWithMax(len(nBytes))
		if err != nil {
			return err
		}
		tmpByte := nBytes[j]
		nBytes[j] = nBytes[randIndex.Int64()]
		nBytes[randIndex.Int64()] = tmpByte
	}
	*b = nBytes
	return nil
}
