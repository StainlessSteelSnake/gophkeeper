package inout

import (
	"bufio"
	"io"
	"os"
)

func ReadStrings() []string {
	var result = make([]string, 0)

	buf := bufio.NewScanner(os.Stdin)
	for buf.Scan() {
		t := buf.Text()
		result = append(result, t)
	}

	return result
}

func ReadStringAsBytes() []byte {
	var result = make([]byte, 0)
	strings := ReadStrings()

	for i, s := range strings {
		if i != 0 {
			result = append(result, '\n')
		}

		result = append(result, []byte(s)...)
	}

	return result
}

func ReadBytes() ([]byte, error) {
	buf := bufio.NewReader(os.Stdin)
	var result []byte
	var byteBuffer = make([]byte, buf.Size())

	var count int
	var err error

	for err == nil {
		count, err = buf.Read(byteBuffer)

		if count > 0 {
			result = append(result, byteBuffer[:count]...)
		}
	}

	if err == io.EOF {
		return result, nil
	}

	if err != nil {
		return nil, err
	}

	return result, nil
}
