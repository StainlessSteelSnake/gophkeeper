package coder

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

// coder содержит настройки для шифрования и расшифровки данных.
type coder struct {
	keyHex   string       // Ключ шифрования в виде шестнадцатиричной строки.
	key      [32]byte     // Ключ шифрования в виде последовательности байт.
	aesblock cipher.Block // Шифрователь блока данных.
	aescgm   cipher.AEAD  // Шифрователь данных произвольной длины.
	nonce    []byte       // Дополнительный код шифрования.
}

// Coder предоставляет методы для шифрования и расшифровки данных, а также установки ключа шифрования.
type Coder interface {
	SetKeyPhrase(string) (string, error) // Установка ключа шифрования в виде строки.
	SetKeyHex(string) error              // Установка ключа шифрования в виде шестнадцатиричной строки.
	Encode([]byte) ([]byte, error)       // Шифрование данных.
	Decode([]byte) ([]byte, error)       // Расшифровка данных.
}

// NewCoder создаёт экземпляр шифрователя данных.
func NewCoder() Coder {
	return &coder{}
}

// SetKeyPhrase задаёт ключ шифрования в виде строки.
// Ключ шифрования представляет хэш от исходной строки, вычисленный по алгоритму SHA256.
func (c *coder) SetKeyPhrase(keyphrase string) (string, error) {
	c.key = sha256.Sum256([]byte(keyphrase))
	c.keyHex = hex.EncodeToString(c.key[:])

	err := c.init()
	if err != nil {
		return "", err
	}

	return c.keyHex, nil
}

// SetKeyHex задаёт ключ шифрования в виде шестнадцатиричной строки.
func (c *coder) SetKeyHex(keyHex string) error {
	c.keyHex = keyHex

	key, err := hex.DecodeString(keyHex)
	if err != nil {
		return err
	}

	copy(c.key[:], key)

	return c.init()
}

// init создаёт шифрователь данных после задания или смены ключа шифрования.
func (c *coder) init() error {
	if c.keyHex == "" {
		return errors.New("не задан ключ шифрования")
	}

	var err error

	c.aesblock, err = aes.NewCipher(c.key[:])
	if err != nil {
		return err
	}

	c.aescgm, err = cipher.NewGCM(c.aesblock)
	if err != nil {
		return err
	}

	c.nonce = c.key[len(c.key)-c.aescgm.NonceSize():]

	return nil
}

// Encode шифрует переданную последовательность байт.
func (c *coder) Encode(source []byte) ([]byte, error) {
	if c.keyHex == "" || c.aescgm == nil {
		return nil, errors.New("не задан ключ шифрования")
	}

	return c.aescgm.Seal(nil, c.nonce, source, nil), nil
}

// Decode расшифровывает переданную последовательность байт.
func (c *coder) Decode(source []byte) ([]byte, error) {
	if c.keyHex == "" || c.aescgm == nil {
		return nil, errors.New("не задан ключ шифрования")
	}

	return c.aescgm.Open(nil, c.nonce, source, nil)
}
