package helper

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
)

type ArchieIO struct {
	Path string
}

func (io *ArchieIO) ReadStringStream() (string, error) {
	return dataStringify(doReadFileStream(io.Path))
}

func (io *ArchieIO) ReadByteStream() ([]byte, error) {
	return doReadFileStream(io.Path)
}

func (io *ArchieIO) ReadStringAll() (string, error) {
	return dataStringify(doReadFileAll(io.Path))
}

func (io *ArchieIO) ReadByteAll() ([]byte, error) {
	return doReadFileAll(io.Path)
}

func dataStringify(data []byte, err error) (string, error) {
	return string(data), err
}

func doReadFileAll(path string) ([]byte, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	data, err := ioutil.ReadAll(f)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func doReadFileStream(path string) ([]byte, error) {
	var data []byte

	err := ReadFileStream(path, func(line []byte) {
		if data == nil {
			data = line
			return
		}

		data = append(data, line...)
	})

	return data, err
}

// 用缓冲区按行读取文件
func ReadFileStream(path string, onData func(data []byte)) error {
	f, err := os.Open(path)

	if err != nil {
		return err
	}

	defer f.Close()

	buf := bufio.NewReader(f)

	for {
		line, _, err := buf.ReadLine()

		if err != nil {
			if err == io.EOF {
				return nil
			}

			return err
		}

		onData(append(line, '\n'))
	}
}
