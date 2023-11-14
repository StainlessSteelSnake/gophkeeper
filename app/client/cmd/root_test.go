package cmd

import (
	"testing"

	conf "github.com/StainlessSteelSnake/gophkeeper/app/client/config"
	"github.com/StainlessSteelSnake/gophkeeper/internal/services"
)

func TestExecute(t *testing.T) {
	type args struct {
		initFunc func(cfg conf.Configurator) (services.GophKeeperClient, func() error)
		cfg      conf.Configurator
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Execute(tt.args.initFunc, tt.args.cfg)
		})
	}
}

func Test_checkCvc(t *testing.T) {
	type args struct {
		c string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := checkCvc(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkCvc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("checkCvc() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkMonth(t *testing.T) {
	type args struct {
		m string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := checkMonth(tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkMonth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("checkMonth() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkYear(t *testing.T) {
	type args struct {
		y string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := checkYear(tt.args.y)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkYear() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("checkYear() got = %v, want %v", got, tt.want)
			}
		})
	}
}
