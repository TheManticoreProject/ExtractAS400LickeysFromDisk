package lictrm_test

import (
	"testing"

	"github.com/TheManticoreProject/ExtractAS400LickeysFromDisk/licenceinfo/lictrm"
)

func TestLICTRM_Unmarshal(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr bool
		want    lictrm.LICTRM
	}{
		{
			name:    "valid lictrm V1R2M3",
			input:   []byte{'V', '1', 'R', '2', 'M', '3'},
			wantErr: false,
			want:    lictrm.LICTRM{Version: 1, Release: 2, Modification: 3},
		},
		{
			name:    "invalid prefix",
			input:   []byte{'X', '1', 'R', '2', 'M', '3'},
			wantErr: true,
		},
		{
			name:    "invalid version",
			input:   []byte{'V', 'x', 'R', '2', 'M', '3'},
			wantErr: true,
		},
		{
			name:    "invalid release",
			input:   []byte{'V', '1', 'R', 'x', 'M', '3'},
			wantErr: true,
		},
		{
			name:    "invalid modification",
			input:   []byte{'V', '1', 'R', '2', 'M', 'x'},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var l lictrm.LICTRM
			err := l.Unmarshal(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && (l != tt.want) {
				t.Errorf("Unmarshal() got = %+v, want %+v", l, tt.want)
			}
		})
	}
}

func TestLICTRM_Marshal(t *testing.T) {
	l := lictrm.LICTRM{Version: 7, Release: 8, Modification: 9}
	got := l.Marshal()
	want := "V7R8M9"
	if got != want {
		t.Errorf("Marshal() = %q, want %q", got, want)
	}
}
