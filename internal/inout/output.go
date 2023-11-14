package inout

import (
	"bufio"
	"os"
)

// WriteStrings записывает текстовые данные в стандартный источник вывода (консоль или файл).
func WriteStrings(s []string) error {
	buf := bufio.NewWriter(os.Stdout)

	for _, line := range s {
		_, err := buf.WriteString(line)
		if err != nil {
			return err
		}
	}

	buf.Flush()
	return nil
}

// WriteStrings записывает бинарные данные в стандартный источник вывода (консоль или файл).
func WriteBytes(b []byte) {
	buf := bufio.NewWriter(os.Stdout)

	buf.Write(b)
}
