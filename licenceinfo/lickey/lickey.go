package lickey

import (
	"fmt"
	"strings"
)

type LICKEY struct {
	Value string
}

// Unmarshal the LICKEY structure from the block of bytes
func (l *LICKEY) Unmarshal(data []byte) error {
	if len(data) != 3*6 {
		return fmt.Errorf("LICKEY must be exactly 18 characters long")
	}

	for _, c := range string(data) {
		if !((c >= '0' && c <= '9') || (c >= 'A' && c <= 'F')) {
			return fmt.Errorf("LICKEY contains non-hex character: %c", c)
		}
	}
	if string(data) != strings.ToUpper(string(data)) {
		return fmt.Errorf("LICKEY is not uppercase: %s", data)
	}

	l.Value = string(data)

	return nil
}

// Marshall the LICKEY structure to a string
func (l *LICKEY) Marshal() string {
	return l.Value
}

// Get the CL value of the LICKEY structure
func (l *LICKEY) GetCLValue() string {
	return fmt.Sprintf("%s %s %s", l.Value[0:6], l.Value[6:12], l.Value[12:18])
}
