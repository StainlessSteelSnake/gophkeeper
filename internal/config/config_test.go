package config

import (
	"reflect"
	"testing"
)

func TestNewConfiguration(t *testing.T) {
	tests := []struct {
		name string
		want *Configuration
	}{
		{
			name: "Значения настроек сервера по-умолчанию",
			want: &Configuration{
				ServerAddress: defaultServerAddress,
				DatabaseURI:   defaultDatabaseURI,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConfiguration(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfiguration() = %v, want %v", got, tt.want)
			}
		})
	}
}
