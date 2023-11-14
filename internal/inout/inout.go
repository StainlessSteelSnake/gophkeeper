// Пакет inout предоставляет функции для считывания и записи текстовых и бинарных данных
// из консоли, либо из внешнего источника (файла).
package inout

import (
	"bufio"
	"io"
)

// ReadStrings считывает текстовые данные из стандартного источника ввода (консоль или файл).
func ReadStrings(in io.Reader) []string {
	var result = make([]string, 0)

	buf := bufio.NewScanner(in)
	for buf.Scan() {
		t := buf.Text()
		result = append(result, t)
	}

	return result
}

// ReadStringAsBytes считывает текстовые данные в виде последовательности байт
// из стандартного источника ввода (консоль или файл).
func ReadStringAsBytes(in io.Reader) []byte {
	var result = make([]byte, 0)
	strings := ReadStrings(in)

	for i, s := range strings {
		if i != 0 {
			result = append(result, '\n')
		}

		result = append(result, []byte(s)...)
	}

	return result
}

// ReadBytes считывает бинарные данные в виде последовательности байт
// из стандартного источника ввода (консоль или файл).
func ReadBytes(in io.Reader) ([]byte, error) {
	buf := bufio.NewReader(in)
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
