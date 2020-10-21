package lib

import (
	"crypto/sha256"
	"hash"

	"github.com/zeebo/blake3"
	"golang.org/x/crypto/blake2b"
)

type Hasher func([]byte) []byte

func checksums(h Hasher, chunks [][]byte) [][]byte {
	checksums := [][]byte{}
	for _, c := range chunks {
		checksums = append(checksums, h(c))
	}
	return checksums
}

func checksumsOptimized(hasher hash.Hash, chunks [][]byte) [][]byte {
	checksums := make([][]byte, len(chunks))
	for i, c := range chunks {
		hasher.Write(c)
		checksums[i] = hasher.Sum(nil)
		hasher.Reset()
	}
	return checksums
}

func hashSHA2(b []byte) []byte {
	hash := sha256.Sum256(b)
	return hash[:]
}

func hashBlake2(b []byte) []byte {
	hash := blake2b.Sum256(b)
	return hash[:]
}

func hashBlake3(b []byte) []byte {
	hash := blake3.Sum256(b)
	return hash[:]
}
