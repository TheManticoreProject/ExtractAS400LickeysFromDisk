package lictrm

import (
	"fmt"
	"strconv"
)

type LICTRM struct {
	Version      int
	Release      int
	Modification int
}

func (l *LICTRM) Unmarshal(data []byte) error {
	if len(data) < 6 {
		return fmt.Errorf("lictrm must be at least 6 characters long")
	}

	if data[0] != 'V' || data[2] != 'R' || data[4] != 'M' {
		return fmt.Errorf("lictrm does not have the form V_R_M_")
	}

	ver, err := strconv.Atoi(string(data[1]))
	if err != nil {
		return fmt.Errorf("invalid version: %v", err)
	}
	rel, err := strconv.Atoi(string(data[3]))
	if err != nil {
		return fmt.Errorf("invalid release: %v", err)
	}
	mod, err := strconv.Atoi(string(data[5]))
	if err != nil {
		return fmt.Errorf("invalid modification: %v", err)
	}
	l.Version = ver
	l.Release = rel
	l.Modification = mod

	return nil
}

func (l *LICTRM) Marshal() string {
	return fmt.Sprintf("V%dR%dM%d", l.Version, l.Release, l.Modification)
}
