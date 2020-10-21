# go-blake2-benchmark


Benchmark blake3, blake2b and sha2 hash functions:
+ "crypto/sha256"
+ "golang.org/x/crypto/blake2b"
+ "github.com/zeebo/blake3"


Running:
```
go test -test.benchtime 4s -test.benchmem -bench=.
```

### Results:

Test data:
200 chunks, 10MB each

```
goos: linux
goarch: amd64
BenchmarkBlake3-4        918159396 ns/op	   18672 B/op	     209 allocs/op
BenchmarkBlake3Opt-4    1009439487 ns/op	   19840 B/op	     401 allocs/op
BenchmarkBlake2-4       2259191472 ns/op	   18672 B/op	     209 allocs/op
BenchmarkBlake2Opt-4    2178748968 ns/op	   11456 B/op	     201 allocs/op
BenchmarkSHA2-4         6219570322 ns/op	   18672 B/op	     209 allocs/op
BenchmarkSHA2Opt-4      6257553206 ns/op	   11392 B/op	     202 allocs/op
```
