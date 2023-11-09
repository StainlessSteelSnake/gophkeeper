package inout

import (
	"bufio"
	"os"
)

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

func WriteBytes(b []byte) {
	buf := bufio.NewWriter(os.Stdout)

	buf.Write(b)
}
