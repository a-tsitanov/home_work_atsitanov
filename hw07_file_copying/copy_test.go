package main

import (
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestCopy(t *testing.T) {
	t.Run("simple case", func(t *testing.T) {
		tempDir, err := ioutil.TempDir("", "hw_temp_dir")
		if err != nil {
			panic(err)
		}
		defer os.Remove(tempDir)

		dstFile := filepath.Join(tempDir, "dst.txt")
		err = Copy("testdata/input.txt", dstFile, 0, 0)
		require.Nil(t, err)
		_, err = os.Stat(dstFile)
		require.Nil(t, err)

	})
}
