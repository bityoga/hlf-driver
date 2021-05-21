package hlfsdk

type Wallet struct {
	Cert        string `json:"cert"`
	CertName    string `json:"certName"`
	PrivKeyName string `json:"privKeyName"`
}