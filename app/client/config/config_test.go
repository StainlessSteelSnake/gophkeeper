package config

import (
	"reflect"
	"testing"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func TestReadConfig(t *testing.T) {
	tests := []struct {
		name    string
		want    Configurator
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadConfig()
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_config_BindPFlag(t *testing.T) {
	type fields struct {
		v *viper.Viper
	}
	type args struct {
		key  string
		flag *pflag.Flag
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
			c := &config{
				v: tt.fields.v,
			}
			if err := c.BindPFlag(tt.args.key, tt.args.flag); (err != nil) != tt.wantErr {
				t.Errorf("BindPFlag() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_config_GetKeyPhrase(t *testing.T) {
	type fields struct {
		v *viper.Viper
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &config{
				v: tt.fields.v,
			}
			if got := c.GetKeyPhrase(); got != tt.want {
				t.Errorf("GetKeyPhrase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_config_GetServerAddress(t *testing.T) {
	type fields struct {
		v *viper.Viper
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &config{
				v: tt.fields.v,
			}
			if got := c.GetServerAddress(); got != tt.want {
				t.Errorf("GetServerAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_config_GetToken(t *testing.T) {
	type fields struct {
		v *viper.Viper
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &config{
				v: tt.fields.v,
			}
			if got := c.GetToken(); got != tt.want {
				t.Errorf("GetToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_config_GetVersion(t *testing.T) {
	type fields struct {
		v *viper.Viper
	}
	tests := []struct {
		name          string
		fields        fields
		wantVersion   string
		wantBuildTime string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &config{
				v: tt.fields.v,
			}
			gotVersion, gotBuildTime := c.GetVersion()
			if gotVersion != tt.wantVersion {
				t.Errorf("GetVersion() gotVersion = %v, want %v", gotVersion, tt.wantVersion)
			}
			if gotBuildTime != tt.wantBuildTime {
				t.Errorf("GetVersion() gotBuildTime = %v, want %v", gotBuildTime, tt.wantBuildTime)
			}
		})
	}
}

func Test_config_SetKeyPhrase(t *testing.T) {
	type fields struct {
		v *viper.Viper
	}
	type args struct {
		keyphrase string
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
			c := &config{
				v: tt.fields.v,
			}
			if err := c.SetKeyPhrase(tt.args.keyphrase); (err != nil) != tt.wantErr {
				t.Errorf("SetKeyPhrase() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_config_SetServerAddress(t *testing.T) {
	type fields struct {
		v *viper.Viper
	}
	type args struct {
		a string
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
			c := &config{
				v: tt.fields.v,
			}
			if err := c.SetServerAddress(tt.args.a); (err != nil) != tt.wantErr {
				t.Errorf("SetServerAddress() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_config_SetToken(t *testing.T) {
	type fields struct {
		v *viper.Viper
	}
	type args struct {
		token string
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
			c := &config{
				v: tt.fields.v,
			}
			if err := c.SetToken(tt.args.token); (err != nil) != tt.wantErr {
				t.Errorf("SetToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_config_SetVersion(t *testing.T) {
	type fields struct {
		v *viper.Viper
	}
	type args struct {
		version   string
		buildTime string
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
			c := &config{
				v: tt.fields.v,
			}
			if err := c.SetVersion(tt.args.version, tt.args.buildTime); (err != nil) != tt.wantErr {
				t.Errorf("SetVersion() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
