package coder

import (
	"crypto/cipher"
	"encoding/hex"
	"reflect"
	"testing"
)

func TestNewCoder(t *testing.T) {
	tests := []struct {
		name string
		want Coder
	}{
		{
			name: "Шифровальщик по-умолчанию",
			want: &coder{},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCoder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_coder_Decode(t *testing.T) {
	type fields struct {
		keyHex   string
		key      [32]byte
		aesblock cipher.Block
		aescgm   cipher.AEAD
		nonce    []byte
	}
	type args struct {
		source []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &coder{
				keyHex:   tt.fields.keyHex,
				key:      tt.fields.key,
				aesblock: tt.fields.aesblock,
				aescgm:   tt.fields.aescgm,
				nonce:    tt.fields.nonce,
			}
			got, err := c.Decode(tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_coder_Encode(t *testing.T) {
	type fields struct {
		keyHex   string
		key      [32]byte
		aesblock cipher.Block
		aescgm   cipher.AEAD
		nonce    []byte
	}
	type args struct {
		source []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Расшифровка фразы",
			fields: fields{
				keyHex: "4c07b5decb9ae4f2f899c2df24ac8cf041702f5a65184fad2b49be7bd473ea12", // "Ключ шифрования"
				key: [32]byte{0x4c, 0x07, 0xb5, 0xde, 0xcb, 0x9a, 0xe4, 0xf2,
					0xf8, 0x99, 0xc2, 0xdf, 0x24, 0xac, 0x8c, 0xf0,
					0x41, 0x70, 0x2f, 0x5a, 0x65, 0x18, 0x4f, 0xad,
					0x2b, 0x49, 0xbe, 0x7b, 0xd4, 0x73, 0xea, 0x12},
			},
			args: args{
				source: []byte("Зашифрованная фраза"),
			},
			want:    "1234",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &coder{
				keyHex:   tt.fields.keyHex,
				key:      tt.fields.key,
				aesblock: tt.fields.aesblock,
				aescgm:   tt.fields.aescgm,
				nonce:    tt.fields.nonce,
			}

			err := c.init()

			if (err != nil) != tt.wantErr {
				t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got, err := c.Encode(tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			wantByte, err := hex.DecodeString(tt.want)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, wantByte) {
				t.Errorf("Encode() got = %v, want %v", got, wantByte)
			}
		})
	}
}

func Test_coder_SetKeyHex(t *testing.T) {
	type fields struct {
		keyHex   string
		key      [32]byte
		aesblock cipher.Block
		aescgm   cipher.AEAD
		nonce    []byte
	}
	type args struct {
		keyHex string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &coder{
				keyHex:   tt.fields.keyHex,
				key:      tt.fields.key,
				aesblock: tt.fields.aesblock,
				aescgm:   tt.fields.aescgm,
				nonce:    tt.fields.nonce,
			}
			if err := c.SetKeyHex(tt.args.keyHex); (err != nil) != tt.wantErr {
				t.Errorf("SetKeyHex() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_coder_SetKeyPhrase(t *testing.T) {
	type fields struct {
		keyHex   string
		key      [32]byte
		aesblock cipher.Block
		aescgm   cipher.AEAD
		nonce    []byte
	}
	type args struct {
		keyphrase string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &coder{
				keyHex:   tt.fields.keyHex,
				key:      tt.fields.key,
				aesblock: tt.fields.aesblock,
				aescgm:   tt.fields.aescgm,
				nonce:    tt.fields.nonce,
			}
			got, err := c.SetKeyPhrase(tt.args.keyphrase)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetKeyPhrase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SetKeyPhrase() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_coder_init(t *testing.T) {
	type fields struct {
		keyHex   string
		key      [32]byte
		aesblock cipher.Block
		aescgm   cipher.AEAD
		nonce    []byte
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &coder{
				keyHex:   tt.fields.keyHex,
				key:      tt.fields.key,
				aesblock: tt.fields.aesblock,
				aescgm:   tt.fields.aescgm,
				nonce:    tt.fields.nonce,
			}
			if err := c.init(); (err != nil) != tt.wantErr {
				t.Errorf("init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
