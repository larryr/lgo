package dorand

import (
	"crypto/rand"
)

func U64() (uint64, error) {
	var b [8]byte

	_, err := rand.Read(b[:])
	if err != nil {
		return 0, err
	}
	u64 :=
		uint64(b[0])<<56 |
			uint64(b[1])<<48 |
			uint64(b[2])<<40 |
			uint64(b[3])<<32 |
			uint64(b[4])<<24 |
			uint64(b[5])<<16 |
			uint64(b[6])<<8 |
			uint64(b[7])
	return u64, nil
}

func Million() ([]byte, error) {

	//make an array to hold a million 8-byte random values
	b := make([]byte, 8*1000000)
	_, err := rand.Read(b)
	if err != nil {
		return nil, nil
	}
	return b, nil
}
