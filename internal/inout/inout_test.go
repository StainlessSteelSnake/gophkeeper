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
			name:  "Чтение пустого набора строк",
			input: []string{},
			want:  []string{},
		},
		{
			name:  "Чтение одной строки",
			input: []string{"Строка для проверки"},
			want:  []string{"Строка для проверки"},
		},
		{
			name:  "Чтение нескольких строк",
			input: []string{"Строка для проверки 1", "Строка для проверки 2", "Строка для проверки 3"},
			want:  []string{"Строка для проверки 1", "Строка для проверки 2", "Строка для проверки 3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var in string
			for i, s := range tt.input {
				in = in + s
				if i < len(tt.input)-1 {
					in = in + "\n"
				}
			}

			result := ReadStrings(strings.NewReader(in))
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestReadStringAsBytes(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []byte
	}{
		{
			name:  "Чтение пустого набора строк как последовательности байт",
			input: []string{},
			want:  []byte{},
		},
		{
			name:  "Чтение одной строки как последовательности байт",
			input: []string{"Строка для проверки"},
			want:  []byte("Строка для проверки"),
		},
		{
			name:  "Чтение нескольких строк как последовательности байт",
			input: []string{"Строка для проверки 1", "Строка для проверки 2", "Строка для проверки 3"},
			want:  []byte("Строка для проверки 1\nСтрока для проверки 2\nСтрока для проверки 3"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var in string
			for i, s := range tt.input {
				in = in + s
				if i < len(tt.input)-1 {
					in = in + "\n"
				}
			}

			result := ReadStringAsBytes(strings.NewReader(in))
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
			name:  "Считывание пустой последовательности байт",
			input: []byte{},
			want:  []byte{},
		},
		{
			name:  "Считывание непустой последовательности байт",
			input: []byte("Последовательность байт"),
			want:  []byte("Последовательность байт"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ReadBytes(bytes.NewReader(tt.input))
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
			name:  "Запись пустого набора строк",
			input: []string{},
			want:  "",
		},
		{
			name:  "Запись одной строки",
			input: []string{"Строка для проверки"},
			want:  "Строка для проверки",
		},
		{
			name:  "Запись нескольких строк",
			input: []string{"Строка для проверки 1", "Строка для проверки 2", "Строка для проверки 3"},
			want:  "Строка для проверки 1Строка для проверки 2Строка для проверки 3",
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
			name:  "Запись пустой последовательности байт",
			input: []byte{},
			want:  []byte{},
		},
		{
			name:  "Запись непустой последовательности байт",
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
