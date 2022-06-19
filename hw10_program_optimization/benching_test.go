package hw10programoptimization

import (
	"archive/zip"
	"testing"
)

func BenchmarkProcessingFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r, err := zip.OpenReader("testdata/users.dat.zip")
		if err != nil {
			b.Fatal("Error read file")
		}
		defer r.Close()
		data, err := r.File[0].Open()
		if err != nil {
			b.Fatal("Error read file")
		}
		_, err = GetDomainStat(data, "biz")
		if err != nil {
			b.Fatal("Error processing DomainStat")
		}
	}
}
