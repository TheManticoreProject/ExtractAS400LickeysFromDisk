package serial

import (
	"fmt"
	"strings"
)

type SERIAL struct {
	Value string
}

func (s *SERIAL) Unmarshal(data []byte) error {
	if len(data) != 7 {
		return fmt.Errorf("SERIAL must be exactly 7 characters long")
	}

	for _, c := range string(data) {
		if !((c >= '0' && c <= '9') || (c >= 'A' && c <= 'F')) {
			return fmt.Errorf("SERIAL contains non-hex character: %c", c)
		}
	}
	if string(data) != strings.ToUpper(string(data)) {
		return fmt.Errorf("SERIAL is not uppercase: %s", data)
	}

	s.Value = string(data)

	return nil
}

func (s *SERIAL) Marshal() string {
	return s.Value
}
