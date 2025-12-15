package lickey_test

import (
	"testing"

	"github.com/TheManticoreProject/ExtractAS400LickeysFromDisk/licenceinfo/lickey"
)

func TestLICKEY_Unmarshal(t *testing.T) {
	type unmarshallTest struct {
		name      string
		input     []byte
		expectErr bool
		expectVal string
	}
	tests := []unmarshallTest{
		{
			name:      "LICKEY valid",
			input:     []byte("123456ABCDEF7890AB"),
			expectErr: false,
			expectVal: "123456ABCDEF7890AB",
		},
		{
			name:      "LICKEY invalid length",
			input:     []byte("1234567890ABCDEF"),
			expectErr: true,
		},
		{
			name:      "LICKEY non-hex",
			input:     []byte("123456ABCDEF7890AZ"),
			expectErr: true,
		},
		{
			name:      "LICKEY lowercase",
			input:     []byte("123456abcdef7890ab"),
			expectErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var l lickey.LICKEY
			err := l.Unmarshal(tc.input)
			if tc.expectErr {
				if err == nil {
					t.Errorf("Expected error for input %q, got nil", string(tc.input))
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				}
				if l.Value != tc.expectVal {
					t.Errorf("Expected Value %s, got %s", tc.expectVal, l.Value)
				}
			}
		})
	}
}

func TestLICKEY_Marshal(t *testing.T) {
	tests := []struct {
		name string
		val  string
	}{
		{
			name: "LICKEY marshall",
			val:  "123456ABCDEF7890AB",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			l := lickey.LICKEY{Value: tc.val}
			if l.Marshal() != tc.val {
				t.Errorf("Expected Marshall to return %s, got %s", tc.val, l.Marshal())
			}
		})
	}

	t.Run("LICKEY_GetCLValue", func(t *testing.T) {
		val := "123456ABCDEF7890AB"
		l := lickey.LICKEY{Value: val}
		expected := "123456 ABCDEF 7890AB"
		got := l.GetCLValue()
		if got != expected {
			t.Errorf("Expected GetCLValue to return %s, got %s", expected, got)
		}
	})
}

func TestLICKEY_MarshallUnmarshalInvolution(t *testing.T) {
	tests := []struct {
		name string
		val  string
	}{
		{
			name: "LICKEY involution",
			val:  "123456ABCDEF7890AB",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			l := lickey.LICKEY{}
			err := l.Unmarshal([]byte(tc.val))
			if err != nil {
				t.Fatalf("Unmarshal failed: %v", err)
			}
			marshalled := l.Marshal()
			if marshalled != tc.val {
				t.Errorf("Expected Marshall after Unmarshal to return %s, got %s", tc.val, marshalled)
			}
		})
	}
}

func TestLICKEY_GetCLValue(t *testing.T) {
	tests := []struct {
		name     string
		val      string
		expected string
		wantErr  bool
	}{
		{
			name:     "Valid LICKEY",
			val:      "123456ABCDEF7890AB",
			expected: "123456 ABCDEF 7890AB",
			wantErr:  false,
		},
		{
			name:     "All F",
			val:      "FFFFFFFFFFFFFFFFFF",
			expected: "FFFFFF FFFFFF FFFFFF",
			wantErr:  false,
		},
		{
			name:     "Lowercase input (should error on Unmarshal)",
			val:      "abcdefabcdefabcdef",
			expected: "",
			wantErr:  true,
		},
		{
			name:     "Too short",
			val:      "1234567890ABCDEF",
			expected: "",
			wantErr:  true,
		},
		{
			name:     "Too long",
			val:      "1234567890ABCDEF1234",
			expected: "",
			wantErr:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var l lickey.LICKEY
			err := l.Unmarshal([]byte(tc.val))
			if (err != nil) != tc.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tc.wantErr)
			}
			if !tc.wantErr {
				got := l.GetCLValue()
				if got != tc.expected {
					t.Errorf("GetCLValue() = %q, want %q", got, tc.expected)
				}
			}
		})
	}
}
