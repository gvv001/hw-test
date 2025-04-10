package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
// fromPath string = "testdata/input.txt"
)

func TestCopy(t *testing.T) {
	// Test: проверка наличия исходного тестового файла
	fromPath := "testdata/input.txt"
	fileInfo, err := os.Stat(fromPath)
	require.NoError(t, err)

	// Test: offset больше, чем размер файла - невалидная ситуация
	fileSize := fileInfo.Size()
	err = Copy("/"+fromPath, "./tmp/out_offset0_limit0.txt", fileSize+1, 0)
	require.Error(t, err)
}
