package conformance

import (
	"crypto/rand"
	"io"
	"testing"
)

func BenchmarkToStringCopy(b *testing.B) {
	buf, _ := io.ReadAll(io.LimitReader(rand.Reader, 4096))

	sum := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c := []byte(string(buf))
		sum += len(c)
	}
}

func BenchmarkExplicitCopy(b *testing.B) {
	buf, _ := io.ReadAll(io.LimitReader(rand.Reader, 4096))

	sum := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c := make([]byte, len(buf))
		copy(c, buf)
		sum += len(c)
	}
}

func BenchmarkAppendCopy(b *testing.B) {
	buf, _ := io.ReadAll(io.LimitReader(rand.Reader, 4096))

	sum := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c := append([]byte{}, buf...)
		sum += len(c)
	}
}
