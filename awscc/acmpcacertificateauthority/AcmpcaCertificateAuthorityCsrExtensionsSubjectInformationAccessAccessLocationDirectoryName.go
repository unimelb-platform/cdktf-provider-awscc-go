package acmpcacertificateauthority


type AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationDirectoryName struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate_authority#common_name AcmpcaCertificateAuthority#common_name}.
	CommonName *string `field:"optional" json:"commonName" yaml:"commonName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate_authority#country AcmpcaCertificateAuthority#country}.
	Country *string `field:"optional" json:"country" yaml:"country"`
	// Array of X.500 attribute type and value. CustomAttributes cannot be used along with pre-defined attributes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate_authority#custom_attributes AcmpcaCertificateAuthority#custom_attributes}
	CustomAttributes interface{} `field:"optional" json:"customAttributes" yaml:"customAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate_authority#distinguished_name_qualifier AcmpcaCertificateAuthority#distinguished_name_qualifier}.
	DistinguishedNameQualifier *string `field:"optional" json:"distinguishedNameQualifier" yaml:"distinguishedNameQualifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate_authority#generation_qualifier AcmpcaCertificateAuthority#generation_qualifier}.
	GenerationQualifier *string `field:"optional" json:"generationQualifier" yaml:"generationQualifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate_authority#given_name AcmpcaCertificateAuthority#given_name}.
	GivenName *string `field:"optional" json:"givenName" yaml:"givenName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate_authority#initials AcmpcaCertificateAuthority#initials}.
	Initials *string `field:"optional" json:"initials" yaml:"initials"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate_authority#locality AcmpcaCertificateAuthority#locality}.
	Locality *string `field:"optional" json:"locality" yaml:"locality"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate_authority#organization AcmpcaCertificateAuthority#organization}.
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate_authority#organizational_unit AcmpcaCertificateAuthority#organizational_unit}.
	OrganizationalUnit *string `field:"optional" json:"organizationalUnit" yaml:"organizationalUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate_authority#pseudonym AcmpcaCertificateAuthority#pseudonym}.
	Pseudonym *string `field:"optional" json:"pseudonym" yaml:"pseudonym"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate_authority#serial_number AcmpcaCertificateAuthority#serial_number}.
	SerialNumber *string `field:"optional" json:"serialNumber" yaml:"serialNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate_authority#state AcmpcaCertificateAuthority#state}.
	State *string `field:"optional" json:"state" yaml:"state"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate_authority#surname AcmpcaCertificateAuthority#surname}.
	Surname *string `field:"optional" json:"surname" yaml:"surname"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate_authority#title AcmpcaCertificateAuthority#title}.
	Title *string `field:"optional" json:"title" yaml:"title"`
}

