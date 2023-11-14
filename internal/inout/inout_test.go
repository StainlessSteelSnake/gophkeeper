package inout

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadStrings(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{
			name:  "Успешное чтение строк",
			input: []string{"Строка для проверки"},
			want:  []string{"Строка для проверки"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			for _, s := range tt.input {
				r := strings.NewReader(s)
				result := ReadStrings(r)
				assert.Equal(t, tt.want, result)
			}

		})
	}
}

func TestReadStringAsBytes(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []byte
	}{
		{
			name:  "Успешное чтение строк как последовательности байт",
			input: "Строка для проверки",
			want:  []byte("Строка для проверки"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(tt.input)
			result := ReadStringAsBytes(r)
			assert.Equal(t, tt.want, result)

		})
	}
}

func TestReadBytes(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		want  []byte
	}{
		{
			name:  "Успешное чтение строк как последовательности байт",
			input: []byte("Последовательность байт"),
			want:  []byte("Последовательность байт"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := bytes.NewReader(tt.input)
			result, err := ReadBytes(r)
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestWriteStrings(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  string
	}{
		{
			name:  "Успешная запись строк",
			input: []string{"Строка для проверки 1", "Строка для проверки 2"},
			want:  "Строка для проверки 1Строка для проверки 2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			b := new(bytes.Buffer)

			err := WriteStrings(tt.input, b)
			if err != nil {
				t.Error(err)
			}

			result, err := b.ReadString('\n')
			if err != nil && err != io.EOF {
				t.Error(err)
			}

			assert.Equal(t, tt.want, result)
		})
	}
}

func TestWriteBytes(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		want  []byte
	}{
		{
			name:  "Успешная запись последовательности байт",
			input: []byte("Последовательность байт"),
			want:  []byte("Последовательность байт"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			buf := new(bytes.Buffer)
			WriteBytes(tt.input, buf)

			result := make([]byte, 0)

			for {
				b, err := buf.ReadByte()
				if err != nil && err == io.EOF {
					break
				}

				if err != nil {
					t.Error(err)
				}

				result = append(result, b)
			}

			assert.Equal(t, tt.want, result)
		})
	}
}
