package server

import (
	"context"
	"reflect"
	"testing"

	"github.com/StainlessSteelSnake/gophkeeper/internal/auth"
	srs "github.com/StainlessSteelSnake/gophkeeper/internal/services"
	"github.com/StainlessSteelSnake/gophkeeper/internal/storage"
)

func TestNewServer(t *testing.T) {
	type args struct {
		storageController storage.Storager
		authenticator     auth.Authenticator
		network           string
		address           string
	}
	tests := []struct {
		name    string
		args    args
		want    *Server
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewServer(tt.args.storageController, tt.args.authenticator, tt.args.network, tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewServer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServer() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_AddBankCard(t *testing.T) {
	type fields struct {
		UnimplementedGophKeeperServer srs.UnimplementedGophKeeperServer
		storageController             storage.Storager
		authenticator                 auth.Authenticator
	}
	type args struct {
		ctx context.Context
		in  *srs.AddBankCardRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *srs.AddBankCardResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				UnimplementedGophKeeperServer: tt.fields.UnimplementedGophKeeperServer,
				storageController:             tt.fields.storageController,
				authenticator:                 tt.fields.authenticator,
			}
			got, err := s.AddBankCard(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddBankCard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddBankCard() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_AddBytes(t *testing.T) {
	type fields struct {
		UnimplementedGophKeeperServer srs.UnimplementedGophKeeperServer
		storageController             storage.Storager
		authenticator                 auth.Authenticator
	}
	type args struct {
		ctx context.Context
		in  *srs.AddBytesRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *srs.AddBytesResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				UnimplementedGophKeeperServer: tt.fields.UnimplementedGophKeeperServer,
				storageController:             tt.fields.storageController,
				authenticator:                 tt.fields.authenticator,
			}
			got, err := s.AddBytes(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddBytes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_AddLoginPassword(t *testing.T) {
	type fields struct {
		UnimplementedGophKeeperServer srs.UnimplementedGophKeeperServer
		storageController             storage.Storager
		authenticator                 auth.Authenticator
	}
	type args struct {
		ctx context.Context
		in  *srs.AddLoginPasswordRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *srs.AddLoginPasswordResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				UnimplementedGophKeeperServer: tt.fields.UnimplementedGophKeeperServer,
				storageController:             tt.fields.storageController,
				authenticator:                 tt.fields.authenticator,
			}
			got, err := s.AddLoginPassword(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddLoginPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddLoginPassword() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_AddText(t *testing.T) {
	type fields struct {
		UnimplementedGophKeeperServer srs.UnimplementedGophKeeperServer
		storageController             storage.Storager
		authenticator                 auth.Authenticator
	}
	type args struct {
		ctx context.Context
		in  *srs.AddTextRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *srs.AddTextResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				UnimplementedGophKeeperServer: tt.fields.UnimplementedGophKeeperServer,
				storageController:             tt.fields.storageController,
				authenticator:                 tt.fields.authenticator,
			}
			got, err := s.AddText(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddText() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_ChangeBankCard(t *testing.T) {
	type fields struct {
		UnimplementedGophKeeperServer srs.UnimplementedGophKeeperServer
		storageController             storage.Storager
		authenticator                 auth.Authenticator
	}
	type args struct {
		ctx context.Context
		in  *srs.ChangeBankCardRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *srs.ChangeBankCardResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				UnimplementedGophKeeperServer: tt.fields.UnimplementedGophKeeperServer,
				storageController:             tt.fields.storageController,
				authenticator:                 tt.fields.authenticator,
			}
			got, err := s.ChangeBankCard(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChangeBankCard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ChangeBankCard() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_ChangeBytes(t *testing.T) {
	type fields struct {
		UnimplementedGophKeeperServer srs.UnimplementedGophKeeperServer
		storageController             storage.Storager
		authenticator                 auth.Authenticator
	}
	type args struct {
		ctx context.Context
		in  *srs.ChangeBytesRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *srs.ChangeBytesResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				UnimplementedGophKeeperServer: tt.fields.UnimplementedGophKeeperServer,
				storageController:             tt.fields.storageController,
				authenticator:                 tt.fields.authenticator,
			}
			got, err := s.ChangeBytes(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChangeBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ChangeBytes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_ChangeLoginPassword(t *testing.T) {
	type fields struct {
		UnimplementedGophKeeperServer srs.UnimplementedGophKeeperServer
		storageController             storage.Storager
		authenticator                 auth.Authenticator
	}
	type args struct {
		ctx context.Context
		in  *srs.ChangeLoginPasswordRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *srs.ChangeLoginPasswordResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				UnimplementedGophKeeperServer: tt.fields.UnimplementedGophKeeperServer,
				storageController:             tt.fields.storageController,
				authenticator:                 tt.fields.authenticator,
			}
			got, err := s.ChangeLoginPassword(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChangeLoginPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ChangeLoginPassword() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_ChangeText(t *testing.T) {
	type fields struct {
		UnimplementedGophKeeperServer srs.UnimplementedGophKeeperServer
		storageController             storage.Storager
		authenticator                 auth.Authenticator
	}
	type args struct {
		ctx context.Context
		in  *srs.ChangeTextRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *srs.ChangeTextResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				UnimplementedGophKeeperServer: tt.fields.UnimplementedGophKeeperServer,
				storageController:             tt.fields.storageController,
				authenticator:                 tt.fields.authenticator,
			}
			got, err := s.ChangeText(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChangeText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ChangeText() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_ChangeUserRecord(t *testing.T) {
	type fields struct {
		UnimplementedGophKeeperServer srs.UnimplementedGophKeeperServer
		storageController             storage.Storager
		authenticator                 auth.Authenticator
	}
	type args struct {
		ctx context.Context
		in  *srs.ChangeUserRecordRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *srs.ChangeUserRecordResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				UnimplementedGophKeeperServer: tt.fields.UnimplementedGophKeeperServer,
				storageController:             tt.fields.storageController,
				authenticator:                 tt.fields.authenticator,
			}
			got, err := s.ChangeUserRecord(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChangeUserRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ChangeUserRecord() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_DeleteUserRecord(t *testing.T) {
	type fields struct {
		UnimplementedGophKeeperServer srs.UnimplementedGophKeeperServer
		storageController             storage.Storager
		authenticator                 auth.Authenticator
	}
	type args struct {
		ctx context.Context
		in  *srs.DeleteUserRecordRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *srs.DeleteUserRecordResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				UnimplementedGophKeeperServer: tt.fields.UnimplementedGophKeeperServer,
				storageController:             tt.fields.storageController,
				authenticator:                 tt.fields.authenticator,
			}
			got, err := s.DeleteUserRecord(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteUserRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteUserRecord() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_GetBankCard(t *testing.T) {
	type fields struct {
		UnimplementedGophKeeperServer srs.UnimplementedGophKeeperServer
		storageController             storage.Storager
		authenticator                 auth.Authenticator
	}
	type args struct {
		ctx context.Context
		in  *srs.GetBankCardRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *srs.GetBankCardResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				UnimplementedGophKeeperServer: tt.fields.UnimplementedGophKeeperServer,
				storageController:             tt.fields.storageController,
				authenticator:                 tt.fields.authenticator,
			}
			got, err := s.GetBankCard(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBankCard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBankCard() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_GetBytes(t *testing.T) {
	type fields struct {
		UnimplementedGophKeeperServer srs.UnimplementedGophKeeperServer
		storageController             storage.Storager
		authenticator                 auth.Authenticator
	}
	type args struct {
		ctx context.Context
		in  *srs.GetBytesRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *srs.GetBytesResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				UnimplementedGophKeeperServer: tt.fields.UnimplementedGophKeeperServer,
				storageController:             tt.fields.storageController,
				authenticator:                 tt.fields.authenticator,
			}
			got, err := s.GetBytes(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBytes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_GetLoginPassword(t *testing.T) {
	type fields struct {
		UnimplementedGophKeeperServer srs.UnimplementedGophKeeperServer
		storageController             storage.Storager
		authenticator                 auth.Authenticator
	}
	type args struct {
		ctx context.Context
		in  *srs.GetLoginPasswordRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *srs.GetLoginPasswordResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				UnimplementedGophKeeperServer: tt.fields.UnimplementedGophKeeperServer,
				storageController:             tt.fields.storageController,
				authenticator:                 tt.fields.authenticator,
			}
			got, err := s.GetLoginPassword(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLoginPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLoginPassword() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_GetText(t *testing.T) {
	type fields struct {
		UnimplementedGophKeeperServer srs.UnimplementedGophKeeperServer
		storageController             storage.Storager
		authenticator                 auth.Authenticator
	}
	type args struct {
		ctx context.Context
		in  *srs.GetTextRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *srs.GetTextResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				UnimplementedGophKeeperServer: tt.fields.UnimplementedGophKeeperServer,
				storageController:             tt.fields.storageController,
				authenticator:                 tt.fields.authenticator,
			}
			got, err := s.GetText(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetText() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_GetUserRecord(t *testing.T) {
	type fields struct {
		UnimplementedGophKeeperServer srs.UnimplementedGophKeeperServer
		storageController             storage.Storager
		authenticator                 auth.Authenticator
	}
	type args struct {
		ctx context.Context
		in  *srs.GetUserRecordRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *srs.GetUserRecordResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				UnimplementedGophKeeperServer: tt.fields.UnimplementedGophKeeperServer,
				storageController:             tt.fields.storageController,
				authenticator:                 tt.fields.authenticator,
			}
			got, err := s.GetUserRecord(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserRecord() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_GetUserRecords(t *testing.T) {
	type fields struct {
		UnimplementedGophKeeperServer srs.UnimplementedGophKeeperServer
		storageController             storage.Storager
		authenticator                 auth.Authenticator
	}
	type args struct {
		ctx context.Context
		in  *srs.GetUserRecordsRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *srs.GetUserRecordsResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				UnimplementedGophKeeperServer: tt.fields.UnimplementedGophKeeperServer,
				storageController:             tt.fields.storageController,
				authenticator:                 tt.fields.authenticator,
			}
			got, err := s.GetUserRecords(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserRecords() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserRecords() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_Login(t *testing.T) {
	type fields struct {
		UnimplementedGophKeeperServer srs.UnimplementedGophKeeperServer
		storageController             storage.Storager
		authenticator                 auth.Authenticator
	}
	type args struct {
		ctx context.Context
		in  *srs.LoginRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *srs.LoginResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				UnimplementedGophKeeperServer: tt.fields.UnimplementedGophKeeperServer,
				storageController:             tt.fields.storageController,
				authenticator:                 tt.fields.authenticator,
			}
			got, err := s.Login(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Login() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_Logout(t *testing.T) {
	type fields struct {
		UnimplementedGophKeeperServer srs.UnimplementedGophKeeperServer
		storageController             storage.Storager
		authenticator                 auth.Authenticator
	}
	type args struct {
		ctx context.Context
		in  *srs.LogoutRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *srs.LogoutResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				UnimplementedGophKeeperServer: tt.fields.UnimplementedGophKeeperServer,
				storageController:             tt.fields.storageController,
				authenticator:                 tt.fields.authenticator,
			}
			got, err := s.Logout(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Logout() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Logout() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_Register(t *testing.T) {
	type fields struct {
		UnimplementedGophKeeperServer srs.UnimplementedGophKeeperServer
		storageController             storage.Storager
		authenticator                 auth.Authenticator
	}
	type args struct {
		ctx context.Context
		in  *srs.RegisterRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *srs.RegisterResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				UnimplementedGophKeeperServer: tt.fields.UnimplementedGophKeeperServer,
				storageController:             tt.fields.storageController,
				authenticator:                 tt.fields.authenticator,
			}
			got, err := s.Register(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Register() got = %v, want %v", got, tt.want)
			}
		})
	}
}
