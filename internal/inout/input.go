package inout

import (
	"bufio"
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

func ReadBytes() []byte {
	var result = make([]byte, 0)
	buf := bufio.NewScanner(os.Stdin)
	for buf.Scan() {
		s := buf.Bytes()
		result = append(result, s...)
	}

	return result
}
