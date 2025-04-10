package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	//добавил ошибки
	ErrDirectoryIsNotIndicated = errors.New("the directory is not indicated")
	ErrDirectoryIsNotExist     = errors.New("directory \\tmp is not exist")

	bytesLen                      int64
	wrritenBytes, sumWrritenBytes int64
)

func Copy(fromPath, toPath string, offset, limit int64) error {

	//проверяем указаны ли директории
	if fromPath == "" || toPath == "" {
		return ErrDirectoryIsNotIndicated
	}

	//получаем fileInfo без отрытия файла
	fileInfo, err := os.Stat(fromPath)

	if err != nil {
		fmt.Println(err)
		return ErrUnsupportedFile
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
		log.Fatal(err, from)
	}

	err = os.Mkdir("tmp", 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(ErrDirectoryIsNotExist)
	}

	outFile, err := os.Create(toPath)
	if err != nil {
		log.Fatal(err, from)
	}

	//копируем по 5% байт от limit
	bytesLen = limit / 20

	for sumWrritenBytes < limit {

		readFile.Seek(offset, 0)

		if wrritenBytes, err = io.CopyN(outFile, readFile, bytesLen); err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		sumWrritenBytes += wrritenBytes
		offset += wrritenBytes
		showProgress(limit, sumWrritenBytes)

		time.Sleep(time.Millisecond * 100)
	}

	return nil
}

func showProgress(limit, sumWrritenBytes int64) {

	fmt.Print("\033[2K\r")
	fmt.Printf("Прогресс копирования %.0f %%", (float32(sumWrritenBytes)/float32(limit))*100)

}
