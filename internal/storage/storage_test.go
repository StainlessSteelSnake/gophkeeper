package storage

import (
	"context"
	"reflect"
	"testing"

	"github.com/jackc/pgx/v5"
)

func TestNewStorage(t *testing.T) {
	type args struct {
		ctx         context.Context
		databaseUri string
	}
	tests := []struct {
		name string
		args args
		want Storager
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStorage(tt.args.ctx, tt.args.databaseUri); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStorage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_AddBankCard(t *testing.T) {
	type fields struct {
		conn *pgx.Conn
		user string
	}
	type args struct {
		ctx       context.Context
		userLogin string
		name      string
		bankCard  *BankCard
		metadata  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				conn: tt.fields.conn,
				user: tt.fields.user,
			}
			got, err := s.AddBankCard(tt.args.ctx, tt.args.userLogin, tt.args.name, tt.args.bankCard, tt.args.metadata)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddBankCard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AddBankCard() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_AddBinary(t *testing.T) {
	type fields struct {
		conn *pgx.Conn
		user string
	}
	type args struct {
		ctx       context.Context
		userLogin string
		name      string
		binary    []byte
		metadata  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				conn: tt.fields.conn,
				user: tt.fields.user,
			}
			got, err := s.AddBinary(tt.args.ctx, tt.args.userLogin, tt.args.name, tt.args.binary, tt.args.metadata)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddBinary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AddBinary() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_AddLoginPassword(t *testing.T) {
	type fields struct {
		conn *pgx.Conn
		user string
	}
	type args struct {
		ctx       context.Context
		userLogin string
		name      string
		login     []byte
		password  []byte
		metadata  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				conn: tt.fields.conn,
				user: tt.fields.user,
			}
			got, err := s.AddLoginPassword(tt.args.ctx, tt.args.userLogin, tt.args.name, tt.args.login, tt.args.password, tt.args.metadata)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddLoginPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AddLoginPassword() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_AddText(t *testing.T) {
	type fields struct {
		conn *pgx.Conn
		user string
	}
	type args struct {
		ctx       context.Context
		userLogin string
		name      string
		text      []byte
		metadata  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				conn: tt.fields.conn,
				user: tt.fields.user,
			}
			got, err := s.AddText(tt.args.ctx, tt.args.userLogin, tt.args.name, tt.args.text, tt.args.metadata)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AddText() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_AddUser(t *testing.T) {
	type fields struct {
		conn *pgx.Conn
		user string
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
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				conn: tt.fields.conn,
				user: tt.fields.user,
			}
			if err := s.AddUser(tt.args.ctx, tt.args.login, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("AddUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_ChangeBankCard(t *testing.T) {
	type fields struct {
		conn *pgx.Conn
		user string
	}
	type args struct {
		ctx       context.Context
		userLogin string
		id        int
		bankCard  *BankCard
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
			s := &Storage{
				conn: tt.fields.conn,
				user: tt.fields.user,
			}
			if err := s.ChangeBankCard(tt.args.ctx, tt.args.userLogin, tt.args.id, tt.args.bankCard); (err != nil) != tt.wantErr {
				t.Errorf("ChangeBankCard() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_ChangeBinary(t *testing.T) {
	type fields struct {
		conn *pgx.Conn
		user string
	}
	type args struct {
		ctx       context.Context
		userLogin string
		id        int
		binary    []byte
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
			s := &Storage{
				conn: tt.fields.conn,
				user: tt.fields.user,
			}
			if err := s.ChangeBinary(tt.args.ctx, tt.args.userLogin, tt.args.id, tt.args.binary); (err != nil) != tt.wantErr {
				t.Errorf("ChangeBinary() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_ChangeLoginPassword(t *testing.T) {
	type fields struct {
		conn *pgx.Conn
		user string
	}
	type args struct {
		ctx       context.Context
		userLogin string
		id        int
		login     []byte
		password  []byte
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
			s := &Storage{
				conn: tt.fields.conn,
				user: tt.fields.user,
			}
			if err := s.ChangeLoginPassword(tt.args.ctx, tt.args.userLogin, tt.args.id, tt.args.login, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("ChangeLoginPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_ChangeRecord(t *testing.T) {
	type fields struct {
		conn *pgx.Conn
		user string
	}
	type args struct {
		ctx context.Context
		r   *Record
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
			s := &Storage{
				conn: tt.fields.conn,
				user: tt.fields.user,
			}
			if err := s.ChangeRecord(tt.args.ctx, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("ChangeRecord() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_ChangeText(t *testing.T) {
	type fields struct {
		conn *pgx.Conn
		user string
	}
	type args struct {
		ctx       context.Context
		userLogin string
		id        int
		text      []byte
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
			s := &Storage{
				conn: tt.fields.conn,
				user: tt.fields.user,
			}
			if err := s.ChangeText(tt.args.ctx, tt.args.userLogin, tt.args.id, tt.args.text); (err != nil) != tt.wantErr {
				t.Errorf("ChangeText() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_Close(t *testing.T) {
	type fields struct {
		conn *pgx.Conn
		user string
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				conn: tt.fields.conn,
				user: tt.fields.user,
			}
			s.Close(tt.args.ctx)
		})
	}
}

func TestStorage_DeleteRecord(t *testing.T) {
	type fields struct {
		conn *pgx.Conn
		user string
	}
	type args struct {
		ctx       context.Context
		userLogin string
		id        int
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
			s := &Storage{
				conn: tt.fields.conn,
				user: tt.fields.user,
			}
			if err := s.DeleteRecord(tt.args.ctx, tt.args.userLogin, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteRecord() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_GetBankCard(t *testing.T) {
	type fields struct {
		conn *pgx.Conn
		user string
	}
	type args struct {
		ctx       context.Context
		userLogin string
		id        int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *BankCard
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				conn: tt.fields.conn,
				user: tt.fields.user,
			}
			got, err := s.GetBankCard(tt.args.ctx, tt.args.userLogin, tt.args.id)
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

func TestStorage_GetBinary(t *testing.T) {
	type fields struct {
		conn *pgx.Conn
		user string
	}
	type args struct {
		ctx       context.Context
		userLogin string
		id        int
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
			s := &Storage{
				conn: tt.fields.conn,
				user: tt.fields.user,
			}
			got, err := s.GetBinary(tt.args.ctx, tt.args.userLogin, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBinary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBinary() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_GetLoginPassword(t *testing.T) {
	type fields struct {
		conn *pgx.Conn
		user string
	}
	type args struct {
		ctx       context.Context
		userLogin string
		id        int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		want1   []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				conn: tt.fields.conn,
				user: tt.fields.user,
			}
			got, got1, err := s.GetLoginPassword(tt.args.ctx, tt.args.userLogin, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLoginPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLoginPassword() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetLoginPassword() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestStorage_GetRecord(t *testing.T) {
	type fields struct {
		conn *pgx.Conn
		user string
	}
	type args struct {
		ctx       context.Context
		userLogin string
		id        int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Record
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				conn: tt.fields.conn,
				user: tt.fields.user,
			}
			got, err := s.GetRecord(tt.args.ctx, tt.args.userLogin, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRecord() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_GetRecords(t *testing.T) {
	type fields struct {
		conn *pgx.Conn
		user string
	}
	type args struct {
		ctx       context.Context
		userLogin string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Record
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				conn: tt.fields.conn,
				user: tt.fields.user,
			}
			got, err := s.GetRecords(tt.args.ctx, tt.args.userLogin)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRecords() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRecords() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_GetText(t *testing.T) {
	type fields struct {
		conn *pgx.Conn
		user string
	}
	type args struct {
		ctx       context.Context
		userLogin string
		id        int
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
			s := &Storage{
				conn: tt.fields.conn,
				user: tt.fields.user,
			}
			got, err := s.GetText(tt.args.ctx, tt.args.userLogin, tt.args.id)
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

func TestStorage_GetUser(t *testing.T) {
	type fields struct {
		conn *pgx.Conn
		user string
	}
	type args struct {
		ctx   context.Context
		login string
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		wantPasswordHash string
		wantRecordCount  int
		wantErr          bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				conn: tt.fields.conn,
				user: tt.fields.user,
			}
			gotPasswordHash, gotRecordCount, err := s.GetUser(tt.args.ctx, tt.args.login)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotPasswordHash != tt.wantPasswordHash {
				t.Errorf("GetUser() gotPasswordHash = %v, want %v", gotPasswordHash, tt.wantPasswordHash)
			}
			if gotRecordCount != tt.wantRecordCount {
				t.Errorf("GetUser() gotRecordCount = %v, want %v", gotRecordCount, tt.wantRecordCount)
			}
		})
	}
}

func TestStorage_addRecord(t *testing.T) {
	type fields struct {
		conn *pgx.Conn
		user string
	}
	type args struct {
		ctx context.Context
		r   *Record
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				conn: tt.fields.conn,
				user: tt.fields.user,
			}
			got, err := s.addRecord(tt.args.ctx, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("addRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("addRecord() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_addTextOrBinary(t *testing.T) {
	type fields struct {
		conn *pgx.Conn
		user string
	}
	type args struct {
		ctx        context.Context
		userLogin  string
		name       string
		binary     []byte
		metadata   string
		recordType string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				conn: tt.fields.conn,
				user: tt.fields.user,
			}
			got, err := s.addTextOrBinary(tt.args.ctx, tt.args.userLogin, tt.args.name, tt.args.binary, tt.args.metadata, tt.args.recordType)
			if (err != nil) != tt.wantErr {
				t.Errorf("addTextOrBinary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("addTextOrBinary() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_changeTextOrBinary(t *testing.T) {
	type fields struct {
		conn *pgx.Conn
		user string
	}
	type args struct {
		ctx       context.Context
		userLogin string
		id        int
		binary    []byte
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
			s := &Storage{
				conn: tt.fields.conn,
				user: tt.fields.user,
			}
			if err := s.changeTextOrBinary(tt.args.ctx, tt.args.userLogin, tt.args.id, tt.args.binary); (err != nil) != tt.wantErr {
				t.Errorf("changeTextOrBinary() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_getTextOrBinary(t *testing.T) {
	type fields struct {
		conn *pgx.Conn
		user string
	}
	type args struct {
		ctx       context.Context
		userLogin string
		id        int
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
			s := &Storage{
				conn: tt.fields.conn,
				user: tt.fields.user,
			}
			got, err := s.getTextOrBinary(tt.args.ctx, tt.args.userLogin, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("getTextOrBinary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTextOrBinary() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_init(t *testing.T) {
	type fields struct {
		conn *pgx.Conn
		user string
	}
	type args struct {
		ctx context.Context
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
			s := &Storage{
				conn: tt.fields.conn,
				user: tt.fields.user,
			}
			if err := s.init(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getRecordType(t *testing.T) {
	type args struct {
		rt string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRecordType(tt.args.rt); got != tt.want {
				t.Errorf("getRecordType() = %v, want %v", got, tt.want)
			}
		})
	}
}
