package types

type ApiKubeconfig struct {
	CurrentContext           string
	CertificateAuthorityData string
	InsecureSkipTLSVerify    bool
	ContextUser              string
	ContextCluster           string
	Token                    string
	BaseURL                  string
	ClientCertificateData    string
	ClientKeyData            string
}
