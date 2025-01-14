package persist

import (
	"bufio"
	"os"
)

func Save(data []byte, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(file)
	for _, d := range data {
		err := writer.WriteByte(d)
		if err != nil {
			return err
		}
	}
	writer.Flush()

	return nil

}

func Load(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	buffered := bufio.NewScanner(file)
	buffered.Split(bufio.ScanBytes)

	data := make([]byte, 0)
	for buffered.Scan() {
		data = append(data, buffered.Bytes()...)
	}

	return data, nil

}
