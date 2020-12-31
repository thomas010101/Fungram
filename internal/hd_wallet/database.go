package hd_wallet

import (
	"fmt"
	"io/ioutil"
	"os"
)

type writeFile struct {
	path string
}

func (w writeFile) Write(data []byte) (int, error) {
	file, err := os.OpenFile(w.path, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return 0, fmt.Errorf("openFile: %w", err)
	}
	defer file.Close()
	n, err := file.Write(data)
	if err != nil {
		return 0, fmt.Errorf("write: %w", err)
	}
	return n, nil
}

func (w writeFile) ReadAll() ([]byte, error) {
	file, err := os.Open(w.path)
	if err != nil {
		return nil, fmt.Errorf("openFile: %w", err)
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("readAll: %w", err)
	}
	return data, nil
}

type emptyWrite struct{}

func (e emptyWrite) Write(_ []byte) (int, error) {
	return 0, nil
}

func (e emptyWrite) ReadAll() ([]byte, error) {
	return nil, fmt.Errorf("empty data")
}
