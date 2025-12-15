package licenceinfo

// LicenceInfo defines the common behavior for licence information records.
// It is implemented by the different record layouts (e.g., LicenceInfo1, LicenceInfo2).
type LicenceInfoInterface interface {
	// Unmarshal parses the licence record from the provided bytes and returns the next index.
	Unmarshal(block []byte) (int, error)

	// ToCLAddLickey returns the IBM i CL command to add this licence key.
	ToCLAddLickey() string

	// Field accessors as strings
	GetPrdidString() string
	GetLictrmString() string
	GetFeatureString() string
	GetLickeyString() string
	GetSerialString() string
	GetPrcgrpString() string

	// Get the raw structure as a byte slice
	GetRawStructure() []byte

	// Human-readable single-line representation.
	Print() string
}
