package acmpcacertificate


type AcmpcaCertificateApiPassthroughExtensionsKeyUsage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#crl_sign AcmpcaCertificate#crl_sign}.
	CrlSign interface{} `field:"optional" json:"crlSign" yaml:"crlSign"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#data_encipherment AcmpcaCertificate#data_encipherment}.
	DataEncipherment interface{} `field:"optional" json:"dataEncipherment" yaml:"dataEncipherment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#decipher_only AcmpcaCertificate#decipher_only}.
	DecipherOnly interface{} `field:"optional" json:"decipherOnly" yaml:"decipherOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#digital_signature AcmpcaCertificate#digital_signature}.
	DigitalSignature interface{} `field:"optional" json:"digitalSignature" yaml:"digitalSignature"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#encipher_only AcmpcaCertificate#encipher_only}.
	EncipherOnly interface{} `field:"optional" json:"encipherOnly" yaml:"encipherOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#key_agreement AcmpcaCertificate#key_agreement}.
	KeyAgreement interface{} `field:"optional" json:"keyAgreement" yaml:"keyAgreement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#key_cert_sign AcmpcaCertificate#key_cert_sign}.
	KeyCertSign interface{} `field:"optional" json:"keyCertSign" yaml:"keyCertSign"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#key_encipherment AcmpcaCertificate#key_encipherment}.
	KeyEncipherment interface{} `field:"optional" json:"keyEncipherment" yaml:"keyEncipherment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#non_repudiation AcmpcaCertificate#non_repudiation}.
	NonRepudiation interface{} `field:"optional" json:"nonRepudiation" yaml:"nonRepudiation"`
}

