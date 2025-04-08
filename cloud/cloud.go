package cloud

// CloudProvider representa o tipo de provedor de nuvem
type CloudProvider string

const (
	AWS   CloudProvider = "AWS"
	GCP   CloudProvider = "GCP"
	AZURE CloudProvider = "AZURE"
	OCI   CloudProvider = "OCI"
)
