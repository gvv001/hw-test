package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	// добавил ошибки.
	ErrFromPathNotSpecified = errors.New("'fromPath' not specified")
	ErrToPathNotSpecified   = errors.New("'toPath' not specified")

	bytesLen, wrritenBytes, sumWrritenBytes int64
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	// проверяем указаны ли директории
	if fromPath == "" {
		return ErrFromPathNotSpecified
	}
	if toPath == "" {
		return ErrToPathNotSpecified
	}

	// получаем fileInfo без отрытия файла
	fileInfo, err := os.Stat(fromPath)
	if err != nil {
		return err
	}

	fileSize := fileInfo.Size()
	if offset > fileSize {
		return ErrOffsetExceedsFileSize
	}

	if limit == 0 || limit >= fileSize {
		limit = fileSize
	}

	readFile, err := os.Open(from)
	if err != nil {
		return err
	}

	outFile, err := os.Create(toPath)
	if err != nil {
		return err
	}

	if limit > fileSize-offset {
		limit = fileSize - offset
	}

	// копируем по 10% байт от limit
	bytesLen = limit / 10

	for sumWrritenBytes < limit {
		readFile.Seek(offset, 0)

		if wrritenBytes, err = io.CopyN(outFile, readFile, bytesLen); err != nil {
			if !errors.Is(err, io.EOF) {
				return err
			}
		}

		sumWrritenBytes += wrritenBytes
		offset += wrritenBytes

		// обновляем прогресс бар
		showProgress(limit, sumWrritenBytes)
	}

	return nil
}

func showProgress(limit, sumWrritenBytes int64) {
	fmt.Print("\033[2K\r")
	fmt.Printf("Прогресс копирования %.0f %%", (float32(sumWrritenBytes)/float32(limit))*100)
}
