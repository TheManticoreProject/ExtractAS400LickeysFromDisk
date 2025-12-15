package licenceinfo

import (
	"fmt"

	"github.com/TheManticoreProject/ExtractAS400LickeysFromDisk/encoding"
	"github.com/TheManticoreProject/ExtractAS400LickeysFromDisk/licenceinfo/lickey"
	"github.com/TheManticoreProject/ExtractAS400LickeysFromDisk/licenceinfo/lictrm"
	"github.com/TheManticoreProject/ExtractAS400LickeysFromDisk/licenceinfo/prcgrp"
	"github.com/TheManticoreProject/ExtractAS400LickeysFromDisk/licenceinfo/prdid"
	"github.com/TheManticoreProject/ExtractAS400LickeysFromDisk/licenceinfo/serial"
)

const LICENCEINFO1_STRUCT_SIZE = 221 // 20+7+6+4+3+8+10+108+10+18+8+9+7+3

type LICENCEINFO1 struct {
	Prdid   prdid.PRDID   // 7 bytes
	Lictrm  lictrm.LICTRM // 6 bytes
	Feature []byte        // 4 bytes
	Lickey  lickey.LICKEY // 18 bytes
	Serial  serial.SERIAL // 7 bytes
	Prcgrp  prcgrp.PRCGRP // 3 bytes
	Vnddta  []byte        // 108 bytes

	// Internal
	RawStructure []byte
}

// Unmarshal the licenceinfo structure from the block of bytes
func (l *LICENCEINFO1) Unmarshal(block []byte) (int, error) {
	index := 20

	// Convert to ASCII
	block = encoding.EbcdicToAscii(block)

	// Parse the Prdid
	err := l.Prdid.Unmarshal(block[index : index+7])
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal PRDID: %v", err)
	}
	index += 7

	// Parse the Lictrm
	err = l.Lictrm.Unmarshal(block[index : index+6])
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal LICTRM: %v", err)
	}
	index += 6

	// Parse the Feature
	l.Feature = block[index : index+4]
	index += 4

	index += 3  // nulls
	index += 8  // FFs
	index += 10 // nulls

	l.Vnddta = block[index : index+108]
	index += 108

	index += 10 // nulls

	err = l.Lickey.Unmarshal(block[index : index+18])
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal LICKEY: %v", err)
	}
	index += 18
	// Check that l.Lickey is 18 hex chars in uppercase

	index += 8 // FFs

	// TODO: This is a space with 9 x 0x40 bytes of placeholders
	// These probably are placeholders for the USGLMT, EXPDATE
	index += 9

	// Parse the Serial
	err = l.Serial.Unmarshal(block[index : index+7])
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal SERIAL: %v", err)
	}
	index += 7

	// Parse the Prcgrp
	err = l.Prcgrp.Unmarshal(block[index : index+3])
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal PRCGRP: %v", err)
	}
	index += 3

	l.RawStructure = block[:LICENCEINFO1_STRUCT_SIZE]

	return index, nil
}

// Convert the licenceinfo structure to a CL ADDLICKEY command
func (l *LICENCEINFO1) ToCLAddLickey() string {
	marshalled_prdid := l.GetPrdidString()
	marshalled_lictrm := l.GetLictrmString()
	marshalled_feature := l.GetFeatureString()
	marshalled_lickey := l.Lickey.GetCLValue()
	marshalled_serial := l.GetSerialString()

	marshalled_vnddta := "TODO"
	marshalled_usglmt := "TODO"
	marshalled_expdate := "TODO"

	cl_command := "ADDLICKEY"

	cl_command += fmt.Sprintf(" PRDID(%s)", marshalled_prdid)

	cl_command += fmt.Sprintf(" LICTRM(%s)", marshalled_lictrm)

	cl_command += fmt.Sprintf(" FEATURE(%s)", marshalled_feature)

	cl_command += fmt.Sprintf(" SERIAL(%s)", marshalled_serial)

	cl_command += fmt.Sprintf(" PRCGRP(%s)", l.Prcgrp.Value)

	cl_command += fmt.Sprintf(" LICKEY(%s)", marshalled_lickey)

	cl_command += fmt.Sprintf(" USGLMT(%s)", marshalled_usglmt)

	cl_command += fmt.Sprintf(" EXPDATE(%s)", marshalled_expdate)

	cl_command += fmt.Sprintf(" VNDDTA(%s)", marshalled_vnddta)

	return cl_command
}

// Get the PRDID string
func (l *LICENCEINFO1) GetPrdidString() string {
	return l.Prdid.Marshal()
}

// Get the LICTRM string
func (l *LICENCEINFO1) GetLictrmString() string {
	return l.Lictrm.Marshal()
}

// Get the FEATURE string
func (l *LICENCEINFO1) GetFeatureString() string {
	return string(l.Feature)
}

// Get the LICKEY string
func (l *LICENCEINFO1) GetLickeyString() string {
	return l.Lickey.Value
}

// Get the SERIAL string
func (l *LICENCEINFO1) GetSerialString() string {
	return l.Serial.Marshal()
}

// Get the PRCGRP string
func (l *LICENCEINFO1) GetPrcgrpString() string {
	return string(l.Prcgrp.Marshal())
}

// Print the licenceinfo structure
func (l *LICENCEINFO1) Print() string {
	marshalled_prdid := l.GetPrdidString()
	marshalled_lictrm := l.GetLictrmString()
	marshalled_feature := l.GetFeatureString()
	marshalled_lickey := l.GetLickeyString()
	marshalled_serial := l.GetSerialString()
	marshalled_prcgrp := l.GetPrcgrpString()

	// marshalled_vnddta := string(l.Vnddta)

	return fmt.Sprintf(
		"PRDID=[%s] LICTRM=[%s] FEATURE=[%s] LICKEY=[%s] SERIAL=[%s] PRCGRP=[%s]",
		marshalled_prdid,
		marshalled_lictrm,
		marshalled_feature,
		marshalled_lickey,
		marshalled_serial,
		marshalled_prcgrp,
	)
}

func (l *LICENCEINFO1) GetRawStructure() []byte {
	return l.RawStructure
}
