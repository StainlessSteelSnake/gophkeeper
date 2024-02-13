package inout

import (
	"bufio"
	"io"
)

// WriteStrings записывает текстовые данные в стандартный источник вывода (консоль или файл).
func WriteStrings(s []string, out io.Writer) error {
	buf := bufio.NewWriter(out)

	for _, line := range s {
		_, err := buf.WriteString(line)
		if err != nil {
			return err
		}
	}

	buf.Flush()
	return nil
}

// WriteBytes записывает бинарные данные в стандартный источник вывода (консоль или файл).
func WriteBytes(b []byte, out io.Writer) {
	buf := bufio.NewWriter(out)

	buf.Write(b)
	buf.Flush()
}
