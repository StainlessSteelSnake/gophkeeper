package coder

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

type coder struct {
	keyHex   string
	key      [32]byte
	aesblock cipher.Block
	aescgm   cipher.AEAD
	nonce    []byte
}

type Coder interface {
	SetKeyPhrase(string) (string, error)
	SetKeyHex(string) error
	Encode([]byte) ([]byte, error)
	Decode([]byte) ([]byte, error)
}

func NewCoder() Coder {
	return &coder{}
}

func (c *coder) SetKeyPhrase(keyphrase string) (string, error) {
	c.key = sha256.Sum256([]byte(keyphrase))
	c.keyHex = hex.EncodeToString(c.key[:])

	err := c.init()
	if err != nil {
		return "", err
	}

	return c.keyHex, nil
}

func (c *coder) SetKeyHex(keyHex string) error {
	c.keyHex = keyHex

	key, err := hex.DecodeString(keyHex)
	if err != nil {
		return err
	}

	copy(c.key[:], key)

	return c.init()
}

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

func (c *coder) Encode(source []byte) ([]byte, error) {
	if c.keyHex == "" || c.aescgm == nil {
		return nil, errors.New("не задан ключ шифрования")
	}

	return c.aescgm.Seal(nil, c.nonce, source, nil), nil
}

func (c *coder) Decode(source []byte) ([]byte, error) {
	if c.keyHex == "" || c.aescgm == nil {
		return nil, errors.New("не задан ключ шифрования")
	}

	return c.aescgm.Open(nil, c.nonce, source, nil)
}
