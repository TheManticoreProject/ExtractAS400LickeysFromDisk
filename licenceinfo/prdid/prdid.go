package prdid

import (
	"fmt"
)

type PRDID struct {
	Product string
	Option  string
}

func (p *PRDID) Unmarshal(data []byte) error {

	if len(data) != 7 {
		return fmt.Errorf("PRDID must be 7 characters long")
	}
	for i := 0; i < 4; i++ {
		if data[i] < '0' || data[i] > '9' {
			return fmt.Errorf("PRDID: first 4 characters must be digits")
		}
	}
	for i := 4; i < 7; i++ {
		c := data[i]
		if !((c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')) {
			return fmt.Errorf("PRDID: last 3 characters must be uppercase letters or digits")
		}
	}

	p.Product = string(data[:4])
	p.Option = string(data[4:7])

	return nil
}

func (p *PRDID) Marshal() string {
	return p.Product + p.Option
}

func (p *PRDID) GetCLValue() string {
	return fmt.Sprintf("PRDID(%s)", p.Product+p.Option)
}
