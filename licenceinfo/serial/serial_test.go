package serial_test

import (
	"reflect"
	"testing"

	"github.com/TheManticoreProject/ExtractAS400LickeysFromDisk/licenceinfo/serial"
)

func TestSERIAL_MarshalUnmarshal_RoundTrip(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "valid serial all digits",
			input:   "1234567",
			wantErr: false,
		},
		{
			name:    "valid serial all hex uppercase",
			input:   "1A2B3C4",
			wantErr: false,
		},
		{
			name:    "invalid serial lowercase",
			input:   "1a2b3c4",
			wantErr: true,
		},
		{
			name:    "invalid serial non-hex",
			input:   "1A2B3CZ",
			wantErr: true,
		},
		{
			name:    "invalid serial too short",
			input:   "123456",
			wantErr: true,
		},
		{
			name:    "invalid serial too long",
			input:   "12345678",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := serial.SERIAL{Value: tt.input}
			marshalled := original.Marshal()

			var unmarshalled serial.SERIAL
			err := unmarshalled.Unmarshal([]byte(marshalled))
			if (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && !reflect.DeepEqual(original, unmarshalled) {
				t.Errorf("Round-trip failed: got %+v, want %+v", unmarshalled, original)
			}
		})
	}
}

func TestSERIAL_Unmarshal(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr bool
		want    serial.SERIAL
	}{
		{
			name:    "valid serial digits",
			input:   []byte("1234567"),
			wantErr: false,
			want:    serial.SERIAL{Value: "1234567"},
		},
		{
			name:    "valid serial hex",
			input:   []byte("1A2B3C4"),
			wantErr: false,
			want:    serial.SERIAL{Value: "1A2B3C4"},
		},
		{
			name:    "invalid serial lowercase",
			input:   []byte("1a2b3c4"),
			wantErr: true,
		},
		{
			name:    "invalid serial non-hex",
			input:   []byte("1A2B3CZ"),
			wantErr: true,
		},
		{
			name:    "invalid serial too short",
			input:   []byte("123456"),
			wantErr: true,
		},
		{
			name:    "invalid serial too long",
			input:   []byte("12345678"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s serial.SERIAL
			err := s.Unmarshal(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && s != tt.want {
				t.Errorf("Unmarshal() got = %+v, want %+v", s, tt.want)
			}
		})
	}
}

func TestSERIAL_Marshal(t *testing.T) {
	tests := []struct {
		name  string
		input serial.SERIAL
		want  string
	}{
		{
			name:  "serial digits",
			input: serial.SERIAL{Value: "1234567"},
			want:  "1234567",
		},
		{
			name:  "serial hex",
			input: serial.SERIAL{Value: "1A2B3C4"},
			want:  "1A2B3C4",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.Marshal()
			if got != tt.want {
				t.Errorf("Marshal() = %q, want %q", got, tt.want)
			}
		})
	}
}
