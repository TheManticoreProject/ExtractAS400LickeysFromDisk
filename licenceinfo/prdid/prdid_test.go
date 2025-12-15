package prdid_test

import (
	"testing"

	"github.com/TheManticoreProject/ExtractAS400LickeysFromDisk/licenceinfo/prdid"
)

func TestPRDID_Unmarshal(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr bool
		want    prdid.PRDID
	}{
		{
			name:    "valid PRDID",
			input:   []byte{'1', '2', '3', '4', 'A', 'B', 'C'},
			wantErr: false,
			want:    prdid.PRDID{Product: "1234", Option: "ABC"},
		},
		{
			name:    "IBM OS400 5770SS1",
			input:   []byte{'5', '7', '7', '0', 'S', 'S', '1'},
			wantErr: false,
			want:    prdid.PRDID{Product: "5770", Option: "SS1"},
		},
		{
			name:    "IBM DB2 5770DB2",
			input:   []byte{'5', '7', '7', '0', 'D', 'B', '2'},
			wantErr: false,
			want:    prdid.PRDID{Product: "5770", Option: "DB2"},
		},
		{
			name:    "IBM BRMS 5770BR1",
			input:   []byte{'5', '7', '7', '0', 'B', 'R', '1'},
			wantErr: false,
			want:    prdid.PRDID{Product: "5770", Option: "BR1"},
		},
		{
			name:    "IBM QMGTOOLS 5733QMG",
			input:   []byte{'5', '7', '3', '3', 'Q', 'M', 'G'},
			wantErr: false,
			want:    prdid.PRDID{Product: "5733", Option: "QMG"},
		},
		{
			name:    "IBM HTTP Server 5770DG1",
			input:   []byte{'5', '7', '7', '0', 'D', 'G', '1'},
			wantErr: false,
			want:    prdid.PRDID{Product: "5770", Option: "DG1"},
		},
		{
			name:    "IBM Java 5770JV1",
			input:   []byte{'5', '7', '7', '0', 'J', 'V', '1'},
			wantErr: false,
			want:    prdid.PRDID{Product: "5770", Option: "JV1"},
		},
		{
			name:    "IBM PASE 5770SS1",
			input:   []byte{'5', '7', '7', '0', 'S', 'S', '1'},
			wantErr: false,
			want:    prdid.PRDID{Product: "5770", Option: "SS1"},
		},
		{
			name:    "IBM Rational Developer 5733RDI",
			input:   []byte{'5', '7', '3', '3', 'R', 'D', 'I'},
			wantErr: false,
			want:    prdid.PRDID{Product: "5733", Option: "RDI"},
		},
		{
			name:    "IBM i Access 5770XE1",
			input:   []byte{'5', '7', '7', '0', 'X', 'E', '1'},
			wantErr: false,
			want:    prdid.PRDID{Product: "5770", Option: "XE1"},
		},
		{
			name:    "invalid - too short",
			input:   []byte{'1', '2', '3', '4', 'A', 'B'},
			wantErr: true,
		},
		{
			name:    "invalid - too long",
			input:   []byte{'1', '2', '3', '4', 'A', 'B', 'C', 'D'},
			wantErr: true,
		},
		{
			name:    "invalid - first 4 not digits",
			input:   []byte{'A', '2', '3', '4', 'A', 'B', 'C'},
			wantErr: true,
		},
		{
			name:    "invalid - last 3 not uppercase or digit",
			input:   []byte{'1', '2', '3', '4', 'a', 'b', 'c'},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var p prdid.PRDID
			err := p.Unmarshal(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && (p.Product != tt.want.Product || p.Option != tt.want.Option) {
				t.Errorf("Unmarshal() got = %+v, want %+v", p, tt.want)
			}
		})
	}
}

func TestPRDID_GetCLValue(t *testing.T) {
	tests := []struct {
		name   string
		prdid  prdid.PRDID
		wantCL string
	}{
		{
			name:   "Standard product",
			prdid:  prdid.PRDID{Product: "5770", Option: "SS1"},
			wantCL: "PRDID(5770SS1)",
		},
		{
			name:   "Alphanumeric option",
			prdid:  prdid.PRDID{Product: "5733", Option: "RDI"},
			wantCL: "PRDID(5733RDI)",
		},
		{
			name:   "Option with digit",
			prdid:  prdid.PRDID{Product: "5770", Option: "XE1"},
			wantCL: "PRDID(5770XE1)",
		},
		{
			name:   "All digits",
			prdid:  prdid.PRDID{Product: "1234", Option: "567"},
			wantCL: "PRDID(1234567)",
		},
		{
			name:   "Empty fields",
			prdid:  prdid.PRDID{Product: "", Option: ""},
			wantCL: "PRDID()",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.prdid.GetCLValue()
			if got != tt.wantCL {
				t.Errorf("GetCLValue() = %q, want %q", got, tt.wantCL)
			}
		})
	}
}
