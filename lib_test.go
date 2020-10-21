package lib

import (
	"crypto/sha256"
	"log"
	"math/rand"
	"testing"

	"github.com/zeebo/blake3"
	"golang.org/x/crypto/blake2b"
)

var chunks [][]byte

func init() {
	nChunks := 200
	chunkSize := int(10e6) // 10MB
	c, err := genBytes(chunkSize)
	if err != nil {
		log.Fatal(err)
	}
	c2, err := genBytes(chunkSize)
	if err != nil {
		log.Fatal(err)
	}
	chunks = make([][]byte, nChunks)
	for i := 0; i < nChunks; i++ {
		if i%2 == 0 {
			chunks[i] = c
		} else {
			chunks[i] = c2
		}
	}
}

func genBytes(size int) (blk []byte, err error) {
	blk = make([]byte, size)
	_, err = rand.Read(blk)
	return
}

func BenchmarkBlake3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		checksums(hashBlake3, chunks)
	}
}

func BenchmarkBlake3Opt(b *testing.B) {
	hasher := blake3.New()
	for n := 0; n < b.N; n++ {
		checksumsOptimized(hasher, chunks)
	}
}

func BenchmarkBlake2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		checksums(hashBlake2, chunks)
	}
}

func BenchmarkBlake2Opt(b *testing.B) {
	hasher, err := blake2b.New256(nil)
	if err != nil {
		log.Fatal(err)
	}
	for n := 0; n < b.N; n++ {
		checksumsOptimized(hasher, chunks)
	}
}

func BenchmarkSHA2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		checksums(hashSHA2, chunks)
	}
}

func BenchmarkSHA2Opt(b *testing.B) {
	hasher := sha256.New()
	for n := 0; n < b.N; n++ {
		checksumsOptimized(hasher, chunks)
	}
}
