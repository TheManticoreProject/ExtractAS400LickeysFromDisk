package encoding

import (
	"bytes"
	"testing"
)

func TestEbcdicToAscii(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected []byte
	}{
		{
			name:     "empty input",
			input:    []byte{},
			expected: []byte{},
		},
		{
			name:     "single character - space",
			input:    []byte{0x40}, // EBCDIC space
			expected: []byte{32},   // ASCII space
		},
		{
			name:     "single character - letter A",
			input:    []byte{0xC1}, // EBCDIC 'A'
			expected: []byte{65},   // ASCII 'A'
		},
		{
			name:     "single character - letter a",
			input:    []byte{0x81}, // EBCDIC 'a'
			expected: []byte{97},   // ASCII 'a'
		},
		{
			name:     "single character - digit 0",
			input:    []byte{0xF0}, // EBCDIC '0'
			expected: []byte{48},   // ASCII '0'
		},
		{
			name:     "single character - digit 9",
			input:    []byte{0xF9}, // EBCDIC '9'
			expected: []byte{57},   // ASCII '9'
		},
		{
			name:     "hello world",
			input:    []byte{0x88, 0x85, 0x93, 0x93, 0x96, 0x40, 0xA6, 0x96, 0x99, 0x93, 0x84}, // EBCDIC "hello world"
			expected: []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100}, // ASCII "hello world"
		},
		{
			name:     "mixed case",
			input:    []byte{0xC1, 0x81, 0xC2, 0x82}, // EBCDIC "AaBb"
			expected: []byte{65, 97, 66, 98},         // ASCII "AaBb"
		},
		{
			name:     "numbers and symbols",
			input:    []byte{0xF0, 0xF1, 0xF2, 0x4B, 0x4C, 0x4D}, // EBCDIC "012.<("
			expected: []byte{48, 49, 50, 46, 60, 40},             // ASCII "012.<("
		},
		{
			name:     "special characters",
			input:    []byte{0x7B, 0x7D, 0x50, 0x61, 0x7C}, // EBCDIC "#'&/@"
			expected: []byte{35, 39, 38, 47, 64},           // ASCII "#'&/@"
		},
		{
			name:     "null byte",
			input:    []byte{0x00},
			expected: []byte{0},
		},
		{
			name:     "high byte values",
			input:    []byte{0xFF, 0xFE, 0xFD, 0xFC},
			expected: []byte{255, 254, 253, 252},
		},
		{
			name:     "control characters",
			input:    []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F},
			expected: []byte{0, 1, 2, 3, 156, 9, 134, 127, 151, 141, 142, 11, 12, 13, 14, 15},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EbcdicToAscii(tt.input)
			if !bytes.Equal(result, tt.expected) {
				t.Errorf("EbcdicToAscii() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestEbcdicIsPrintable(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected bool
	}{
		{
			name:     "empty input",
			input:    []byte{},
			expected: true,
		},
		{
			name:     "single printable character",
			input:    []byte{0x40}, // EBCDIC space -> ASCII space (32)
			expected: true,
		},
		{
			name:     "single letter A",
			input:    []byte{0xC1}, // EBCDIC 'A' -> ASCII 'A' (65)
			expected: true,
		},
		{
			name:     "single letter a",
			input:    []byte{0x81}, // EBCDIC 'a' -> ASCII 'a' (97)
			expected: true,
		},
		{
			name:     "single digit 0",
			input:    []byte{0xF0}, // EBCDIC '0' -> ASCII '0' (48)
			expected: true,
		},
		{
			name:     "single digit 9",
			input:    []byte{0xF9}, // EBCDIC '9' -> ASCII '9' (57)
			expected: true,
		},
		{
			name:     "printable string",
			input:    []byte{0x88, 0x85, 0x93, 0x93, 0x96, 0x40, 0xA6, 0x96, 0x99, 0x93, 0x84}, // EBCDIC "HELLO WORLD"
			expected: true,
		},
		{
			name:     "mixed printable characters",
			input:    []byte{0xC1, 0x81, 0xC2, 0x82, 0xF0, 0xF1, 0x4B, 0x4C}, // EBCDIC "AaBb01+," -> ASCII "AaBb01+,"
			expected: true,
		},
		{
			name:     "contains null byte",
			input:    []byte{0x00}, // EBCDIC null -> ASCII null (0)
			expected: false,
		},
		{
			name:     "contains control character",
			input:    []byte{0x01}, // EBCDIC SOH -> ASCII SOH (1)
			expected: false,
		},
		{
			name:     "contains tab character",
			input:    []byte{0x05}, // EBCDIC HT -> ASCII HT (9)
			expected: false,
		},
		{
			name:     "contains newline character",
			input:    []byte{0x15}, // EBCDIC NL -> ASCII NL (10)
			expected: false,
		},
		{
			name:     "contains carriage return",
			input:    []byte{0x0D}, // EBCDIC CR -> ASCII CR (13)
			expected: false,
		},
		{
			name:     "contains DEL character",
			input:    []byte{0x07}, // EBCDIC DEL -> ASCII DEL (127)
			expected: false,
		},
		{
			name:     "contains high byte values",
			input:    []byte{0xFF}, // EBCDIC 0xFF -> ASCII 0xFF (255)
			expected: false,
		},
		{
			name:     "mixed printable and non-printable",
			input:    []byte{0xC1, 0x00, 0xC2}, // EBCDIC "A\0B" -> ASCII "A\0B"
			expected: false,
		},
		{
			name:     "all printable range",
			input:    []byte{0x40, 0x4B, 0x4C, 0x4D, 0x4E, 0x4F, 0x50, 0x51, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57, 0x58, 0x59, 0x5A, 0x5B, 0x5C, 0x5D, 0x5E, 0x5F, 0x60, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69, 0x6A, 0x6B, 0x6C, 0x6D, 0x6E, 0x6F, 0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77, 0x78, 0x79, 0x7A, 0x7B, 0x7C, 0x7D, 0x7E, 0x7F, 0x80, 0x81, 0x82, 0x83, 0x84, 0x85, 0x86, 0x87, 0x88, 0x89, 0x8A, 0x8B, 0x8C, 0x8D, 0x8E, 0x8F, 0x90, 0x91, 0x92, 0x93, 0x94, 0x95, 0x96, 0x97, 0x98, 0x99, 0x9A, 0x9B, 0x9C, 0x9D, 0x9E, 0x9F, 0xA0, 0xA1, 0xA2, 0xA3, 0xA4, 0xA5, 0xA6, 0xA7, 0xA8, 0xA9, 0xAA, 0xAB, 0xAC, 0xAD, 0xAE, 0xAF, 0xB0, 0xB1, 0xB2, 0xB3, 0xB4, 0xB5, 0xB6, 0xB7, 0xB8, 0xB9, 0xBA, 0xBB, 0xBC, 0xBD, 0xBE, 0xBF, 0xC0, 0xC1, 0xC2, 0xC3, 0xC4, 0xC5, 0xC6, 0xC7, 0xC8, 0xC9, 0xCA, 0xCB, 0xCC, 0xCD, 0xCE, 0xCF, 0xD0, 0xD1, 0xD2, 0xD3, 0xD4, 0xD5, 0xD6, 0xD7, 0xD8, 0xD9, 0xDA, 0xDB, 0xDC, 0xDD, 0xDE, 0xDF, 0xE0, 0xE1, 0xE2, 0xE3, 0xE4, 0xE5, 0xE6, 0xE7, 0xE8, 0xE9, 0xEA, 0xEB, 0xEC, 0xED, 0xEE, 0xEF, 0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9, 0xFA, 0xFB, 0xFC, 0xFD, 0xFE, 0xFF},
			expected: false, // Contains non-printable characters
		},
		{
			name:     "boundary test - ASCII 31 (below printable)",
			input:    []byte{0x1F}, // EBCDIC 0x1F -> ASCII 31
			expected: false,
		},
		{
			name:     "boundary test - ASCII 32 (space, printable)",
			input:    []byte{0x40}, // EBCDIC space -> ASCII 32
			expected: true,
		},
		{
			name:     "boundary test - ASCII 126 (tilde, printable)",
			input:    []byte{0x7E}, // EBCDIC 0x7E -> ASCII 126
			expected: true,
		},
		{
			name:     "boundary test - ASCII 127 (DEL, non-printable)",
			input:    []byte{0x07}, // EBCDIC DEL -> ASCII 127
			expected: false,
		},
		{
			name:     "long printable string",
			input:    []byte{0xC1, 0xC2, 0xC3, 0xC4, 0xC5, 0xC6, 0xC7, 0xC8, 0xC9, 0xCA, 0xCB, 0xCC, 0xCD, 0xCE, 0xCF, 0xD0, 0xD1, 0xD2, 0xD3, 0xD4, 0xD5, 0xD6, 0xD7, 0xD8, 0xD9, 0xDA, 0xDB, 0xDC, 0xDD, 0xDE, 0xDF, 0xE0, 0xE1, 0xE2, 0xE3, 0xE4, 0xE5, 0xE6, 0xE7, 0xE8, 0xE9, 0xEA, 0xEB, 0xEC, 0xED, 0xEE, 0xEF, 0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9, 0xFA, 0xFB, 0xFC, 0xFD, 0xFE, 0xFF},
			expected: false, // Contains non-printable characters
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EbcdicIsPrintable(tt.input)
			if result != tt.expected {
				t.Errorf("EbcdicIsPrintable() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestEbcdicToAsciiEdgeCases tests edge cases and boundary conditions
func TestEbcdicToAsciiEdgeCases(t *testing.T) {
	t.Run("nil input", func(t *testing.T) {
		result := EbcdicToAscii(nil)
		if string(result) != "" {
			t.Errorf("EbcdicToAscii(nil) = %q, want %q", result, "")
		}
	})

	t.Run("single byte 0x00", func(t *testing.T) {
		result := EbcdicToAscii([]byte{0x00})
		expected := "\x00"
		if string(result) != expected {
			t.Errorf("EbcdicToAscii([]byte{0x00}) = %q, want %q", result, expected)
		}
	})

	t.Run("single byte 0xFF", func(t *testing.T) {
		result := EbcdicToAscii([]byte{0xFF})
		expected := "\xFF"
		if string(result) != expected {
			t.Errorf("EbcdicToAscii([]byte{0xFF}) = %q, want %q", result, expected)
		}
	})

	t.Run("all possible byte values", func(t *testing.T) {
		input := make([]byte, 256)
		for i := 0; i < 256; i++ {
			input[i] = byte(i)
		}
		result := EbcdicToAscii(input)
		expected := string(ebcdicToAscii[:])
		if string(result) != expected {
			t.Errorf("EbcdicToAscii(all bytes) = %q, want %q", result, expected)
		}
	})
}

// TestEbcdicIsPrintableEdgeCases tests edge cases for the printable check
func TestEbcdicIsPrintableEdgeCases(t *testing.T) {
	t.Run("nil input", func(t *testing.T) {
		result := EbcdicIsPrintable(nil)
		if !result {
			t.Errorf("EbcdicIsPrintable(nil) = %v, want %v", result, true)
		}
	})

	t.Run("single byte 0x00", func(t *testing.T) {
		result := EbcdicIsPrintable([]byte{0x00})
		if result {
			t.Errorf("EbcdicIsPrintable([]byte{0x00}) = %v, want %v", result, false)
		}
	})

	t.Run("single byte 0xFF", func(t *testing.T) {
		result := EbcdicIsPrintable([]byte{0xFF})
		if result {
			t.Errorf("EbcdicIsPrintable([]byte{0xFF}) = %v, want %v", result, false)
		}
	})

	t.Run("all printable ASCII range", func(t *testing.T) {
		// Test all EBCDIC values that map to printable ASCII (32-126)
		printableEBCDIC := []byte{
			0x40, 0x4B, 0x4C, 0x4D, 0x4E, 0x4F, 0x50, 0x51, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57, 0x58, 0x59, 0x5A, 0x5B, 0x5C, 0x5D, 0x5E, 0x5F, 0x60, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69, 0x6A, 0x6B, 0x6C, 0x6D, 0x6E, 0x6F, 0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77, 0x78, 0x79, 0x7A, 0x7B, 0x7C, 0x7D, 0x7E, 0x7F, 0x80, 0x81, 0x82, 0x83, 0x84, 0x85, 0x86, 0x87, 0x88, 0x89, 0x8A, 0x8B, 0x8C, 0x8D, 0x8E, 0x8F, 0x90, 0x91, 0x92, 0x93, 0x94, 0x95, 0x96, 0x97, 0x98, 0x99, 0x9A, 0x9B, 0x9C, 0x9D, 0x9E, 0x9F, 0xA0, 0xA1, 0xA2, 0xA3, 0xA4, 0xA5, 0xA6, 0xA7, 0xA8, 0xA9, 0xAA, 0xAB, 0xAC, 0xAD, 0xAE, 0xAF, 0xB0, 0xB1, 0xB2, 0xB3, 0xB4, 0xB5, 0xB6, 0xB7, 0xB8, 0xB9, 0xBA, 0xBB, 0xBC, 0xBD, 0xBE, 0xBF, 0xC0, 0xC1, 0xC2, 0xC3, 0xC4, 0xC5, 0xC6, 0xC7, 0xC8, 0xC9, 0xCA, 0xCB, 0xCC, 0xCD, 0xCE, 0xCF, 0xD0, 0xD1, 0xD2, 0xD3, 0xD4, 0xD5, 0xD6, 0xD7, 0xD8, 0xD9, 0xDA, 0xDB, 0xDC, 0xDD, 0xDE, 0xDF, 0xE0, 0xE1, 0xE2, 0xE3, 0xE4, 0xE5, 0xE6, 0xE7, 0xE8, 0xE9, 0xEA, 0xEB, 0xEC, 0xED, 0xEE, 0xEF, 0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9, 0xFA, 0xFB, 0xFC, 0xFD, 0xFE, 0xFF,
		}
		
		// Filter out only the truly printable ones (32-126)
		var printableBytes []byte
		for _, b := range printableEBCDIC {
			ascii := ebcdicToAscii[b]
			if ascii >= 32 && ascii <= 126 {
				printableBytes = append(printableBytes, b)
			}
		}
		
		result := EbcdicIsPrintable(printableBytes)
		if !result {
			t.Errorf("EbcdicIsPrintable(printable bytes) = %v, want %v", result, true)
		}
	})
}

// BenchmarkEbcdicToAscii benchmarks the conversion function
func BenchmarkEbcdicToAscii(b *testing.B) {
	testData := []byte{0x88, 0x85, 0x93, 0x93, 0x96, 0x40, 0xA6, 0x96, 0x99, 0x93, 0x84} // "HELLO WORLD"
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = EbcdicToAscii(testData)
	}
}

// BenchmarkEbcdicIsPrintable benchmarks the printable check function
func BenchmarkEbcdicIsPrintable(b *testing.B) {
	testData := []byte{0x88, 0x85, 0x93, 0x93, 0x96, 0x40, 0xA6, 0x96, 0x99, 0x93, 0x84} // "HELLO WORLD"
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = EbcdicIsPrintable(testData)
	}
}
