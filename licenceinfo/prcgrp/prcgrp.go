package prcgrp

import (
	"fmt"
	"strings"
)

type PRCGRP struct {
	Value string
}

func (p *PRCGRP) Unmarshal(data []byte) error {
	if len(data) < 3 {
		return fmt.Errorf("PRCGRP must be at least 3 characters long")
	}
	if data[0] != '9' && data[0] != 'P' {
		return fmt.Errorf("PRCGRP must start with 9 or P, got '%s'", string(data))
	}

	// Accept "999" or "P" + 2 digits
	if strings.EqualFold(string(data), "999") {
		p.Value = "*ANY"
		return nil
	} else if data[0] == 'P' && data[1] >= '0' && data[1] <= '9' && data[2] >= '0' && data[2] <= '9' {
		p.Value = string(data)
		return nil
	} else {
		return fmt.Errorf("invalid PRCGRP value: %q", string(data))
	}
}

func (p *PRCGRP) Marshal() []byte {
	if p.Value == "*ANY" {
		return []byte("999")
	} else {
		return []byte(p.Value)
	}
}
