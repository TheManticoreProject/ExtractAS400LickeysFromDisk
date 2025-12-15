package prcgrp_test

import (
	"reflect"
	"testing"

	"github.com/TheManticoreProject/ExtractAS400LickeysFromDisk/licenceinfo/prcgrp"
)

func TestPRCGRP_MarshalUnmarshal_RoundTrip(t *testing.T) {
	original := prcgrp.PRCGRP{Value: "P01"}
	marshalled := original.Marshal()

	var unmarshalled prcgrp.PRCGRP
	err := unmarshalled.Unmarshal(marshalled)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if !reflect.DeepEqual(original, unmarshalled) {
		t.Errorf("Round-trip failed: got %+v, want %+v", unmarshalled, original)
	}
}

func TestPRCGRP_Marshal(t *testing.T) {
	tests := []struct {
		name  string
		input prcgrp.PRCGRP
		want  []byte
	}{
		{
			name:  "P01 group",
			input: prcgrp.PRCGRP{Value: "P01"},
			want:  []byte("P01"),
		},
		{
			name:  "*ANY group",
			input: prcgrp.PRCGRP{Value: "*ANY"},
			want:  []byte("999"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.Marshal()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Marshal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPRCGRP_Unmarshal(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr bool
		want    prcgrp.PRCGRP
	}{
		{
			name:    "valid P group",
			input:   []byte{'P', '0', '1'},
			wantErr: false,
			want:    prcgrp.PRCGRP{Value: "P01"},
		},
		{
			name:    "valid 999 group",
			input:   []byte{'9', '9', '9'},
			wantErr: false,
			want:    prcgrp.PRCGRP{Value: "*ANY"},
		},
		{
			name:    "invalid too short",
			input:   []byte{'P', '0'},
			wantErr: true,
		},
		{
			name:    "invalid prefix",
			input:   []byte{'X', '0', '1'},
			wantErr: true,
		},
		{
			name:    "invalid P group non-digit",
			input:   []byte{'P', 'A', '1'},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var p prcgrp.PRCGRP
			err := p.Unmarshal(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil && !tt.wantErr && !reflect.DeepEqual(p, tt.want) {
				t.Errorf("Unmarshal() got = %+v, want %+v", p, tt.want)
			}
		})
	}
}
