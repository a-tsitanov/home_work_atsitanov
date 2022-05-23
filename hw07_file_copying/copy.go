package main

import (
	"errors"
	"github.com/cheggaaa/pb/v3"
	"io"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	fileInfo, err := os.Stat(fromPath)

	if err != nil {
		return err
	}

	if !fileInfo.Mode().IsRegular() {
		return ErrUnsupportedFile
	}

	if offset > fileInfo.Size() {
		return ErrOffsetExceedsFileSize
	}

	if limit == 0 {
		limit = fileInfo.Size()
	}

	destination, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer destination.Close()

	bar := pb.StartNew(int(fileInfo.Size()))
	f, err := os.OpenFile(fromPath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Seek(offset, 0)
	if err != nil {
		return err
	}
	buf := make([]byte, 1)
	for i := 0; i < int(limit); i++ {
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		if _, err := destination.Write(buf[:n]); err != nil {
			return err
		}
		bar.Increment()
	}
	bar.Finish()
	return nil
}
