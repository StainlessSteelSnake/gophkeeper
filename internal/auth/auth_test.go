package auth

import (
	"context"
	"reflect"
	"testing"
)

func TestNewAuthentication(t *testing.T) {
	type args struct {
		userController UserAdderGetter
	}
	tests := []struct {
		name    string
		args    args
		want    Authenticator
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAuthentication(tt.args.userController)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAuthentication() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthentication() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authentication_Authenticate(t *testing.T) {
	type fields struct {
		users          map[string]*user
		userController UserAdderGetter
	}
	type args struct {
		ctx context.Context
		t   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		want1   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &authentication{
				users:          tt.fields.users,
				userController: tt.fields.userController,
			}
			got, got1, err := a.Authenticate(tt.args.ctx, tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("Authenticate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Authenticate() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Authenticate() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_authentication_Login(t *testing.T) {
	type fields struct {
		users          map[string]*user
		userController UserAdderGetter
	}
	type args struct {
		ctx      context.Context
		login    string
		password string
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
			a := &authentication{
				users:          tt.fields.users,
				userController: tt.fields.userController,
			}
			got, err := a.Login(tt.args.ctx, tt.args.login, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Login() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authentication_Logout(t *testing.T) {
	type fields struct {
		users          map[string]*user
		userController UserAdderGetter
	}
	type args struct {
		ctx context.Context
		t   string
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
			a := &authentication{
				users:          tt.fields.users,
				userController: tt.fields.userController,
			}
			if err := a.Logout(tt.args.ctx, tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("Logout() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_authentication_Register(t *testing.T) {
	type fields struct {
		users          map[string]*user
		userController UserAdderGetter
	}
	type args struct {
		ctx      context.Context
		login    string
		password string
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
			a := &authentication{
				users:          tt.fields.users,
				userController: tt.fields.userController,
			}
			got, err := a.Register(tt.args.ctx, tt.args.login, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Register() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authentication_checkPassword(t *testing.T) {
	type fields struct {
		users          map[string]*user
		userController UserAdderGetter
	}
	type args struct {
		ctx      context.Context
		login    string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		want1   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &authentication{
				users:          tt.fields.users,
				userController: tt.fields.userController,
			}
			got, got1, err := a.checkPassword(tt.args.ctx, tt.args.login, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("checkPassword() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("checkPassword() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_authentication_createToken(t *testing.T) {
	type fields struct {
		users          map[string]*user
		userController UserAdderGetter
	}
	type args struct {
		loginHash string
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
			a := &authentication{
				users:          tt.fields.users,
				userController: tt.fields.userController,
			}
			got, err := a.createToken(tt.args.loginHash)
			if (err != nil) != tt.wantErr {
				t.Errorf("createToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("createToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getHash(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getHash(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("getHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getHash() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getRandom(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getRandom(tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("getRandom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getRandom() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSign(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getSign(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("getSign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getSign() got = %v, want %v", got, tt.want)
			}
		})
	}
}
