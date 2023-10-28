package dynamodb

import (
	"bytes"
	"fmt"
	"testing"
)

func BenchmarkString(b *testing.B) {
	buffer := []byte{'y', 'o', 's', 'h', 'i', 'm', 'a', 'c', 'o', 'v', 'i', 'l'}
	for i := 0; i < b.N; i++ {
		str1 := string(buffer)

		_ = str1
	}
}

func BenchmarkStringV2(b *testing.B) {
	buffer := []byte{'y', 'o', 's', 'h', 'i', 'm', 'a', 'c', 'o', 'v', 'i', 'l'}
	for i := 0; i < b.N; i++ {
		str1 := bytes.NewBuffer(buffer).String()

		_ = str1
	}
}

func BenchmarkStringV3(b *testing.B) {
	buffer := []byte{'y', 'o', 's', 'h', 'i', 'm', 'a', 'c', 'o', 'v', 'i', 'l'}
	for i := 0; i < b.N; i++ {
		str1 := fmt.Sprintf("%s", buffer)

		_ = str1
	}
}
