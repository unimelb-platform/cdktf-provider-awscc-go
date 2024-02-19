package acmpcacertificate


type AcmpcaCertificateApiPassthroughSubject struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#common_name AcmpcaCertificate#common_name}.
	CommonName *string `field:"optional" json:"commonName" yaml:"commonName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#country AcmpcaCertificate#country}.
	Country *string `field:"optional" json:"country" yaml:"country"`
	// Array of X.500 attribute type and value. CustomAttributes cannot be used along with pre-defined attributes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#custom_attributes AcmpcaCertificate#custom_attributes}
	CustomAttributes interface{} `field:"optional" json:"customAttributes" yaml:"customAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#distinguished_name_qualifier AcmpcaCertificate#distinguished_name_qualifier}.
	DistinguishedNameQualifier *string `field:"optional" json:"distinguishedNameQualifier" yaml:"distinguishedNameQualifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#generation_qualifier AcmpcaCertificate#generation_qualifier}.
	GenerationQualifier *string `field:"optional" json:"generationQualifier" yaml:"generationQualifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#given_name AcmpcaCertificate#given_name}.
	GivenName *string `field:"optional" json:"givenName" yaml:"givenName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#initials AcmpcaCertificate#initials}.
	Initials *string `field:"optional" json:"initials" yaml:"initials"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#locality AcmpcaCertificate#locality}.
	Locality *string `field:"optional" json:"locality" yaml:"locality"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#organization AcmpcaCertificate#organization}.
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#organizational_unit AcmpcaCertificate#organizational_unit}.
	OrganizationalUnit *string `field:"optional" json:"organizationalUnit" yaml:"organizationalUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#pseudonym AcmpcaCertificate#pseudonym}.
	Pseudonym *string `field:"optional" json:"pseudonym" yaml:"pseudonym"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#serial_number AcmpcaCertificate#serial_number}.
	SerialNumber *string `field:"optional" json:"serialNumber" yaml:"serialNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#state AcmpcaCertificate#state}.
	State *string `field:"optional" json:"state" yaml:"state"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#surname AcmpcaCertificate#surname}.
	Surname *string `field:"optional" json:"surname" yaml:"surname"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#title AcmpcaCertificate#title}.
	Title *string `field:"optional" json:"title" yaml:"title"`
}

