package acmpcacertificate


type AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNames struct {
	// Structure that contains X.500 distinguished name information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#directory_name AcmpcaCertificate#directory_name}
	DirectoryName *AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryName `field:"optional" json:"directoryName" yaml:"directoryName"`
	// String that contains X.509 DnsName information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#dns_name AcmpcaCertificate#dns_name}
	DnsName *string `field:"optional" json:"dnsName" yaml:"dnsName"`
	// Structure that contains X.509 EdiPartyName information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#edi_party_name AcmpcaCertificate#edi_party_name}
	EdiPartyName *AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesEdiPartyName `field:"optional" json:"ediPartyName" yaml:"ediPartyName"`
	// String that contains X.509 IpAddress information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#ip_address AcmpcaCertificate#ip_address}
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// Structure that contains X.509 OtherName information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#other_name AcmpcaCertificate#other_name}
	OtherName *AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOtherName `field:"optional" json:"otherName" yaml:"otherName"`
	// String that contains X.509 ObjectIdentifier information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#registered_id AcmpcaCertificate#registered_id}
	RegisteredId *string `field:"optional" json:"registeredId" yaml:"registeredId"`
	// String that contains X.509 Rfc822Name information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#rfc_822_name AcmpcaCertificate#rfc_822_name}
	Rfc822Name *string `field:"optional" json:"rfc822Name" yaml:"rfc822Name"`
	// String that contains X.509 UniformResourceIdentifier information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#uniform_resource_identifier AcmpcaCertificate#uniform_resource_identifier}
	UniformResourceIdentifier *string `field:"optional" json:"uniformResourceIdentifier" yaml:"uniformResourceIdentifier"`
}

