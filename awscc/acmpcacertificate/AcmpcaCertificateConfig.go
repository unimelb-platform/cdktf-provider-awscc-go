package acmpcacertificate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AcmpcaCertificateConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The Amazon Resource Name (ARN) for the private CA to issue the certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#certificate_authority_arn AcmpcaCertificate#certificate_authority_arn}
	CertificateAuthorityArn *string `field:"required" json:"certificateAuthorityArn" yaml:"certificateAuthorityArn"`
	// The certificate signing request (CSR) for the Certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#certificate_signing_request AcmpcaCertificate#certificate_signing_request}
	CertificateSigningRequest *string `field:"required" json:"certificateSigningRequest" yaml:"certificateSigningRequest"`
	// The name of the algorithm that will be used to sign the Certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#signing_algorithm AcmpcaCertificate#signing_algorithm}
	SigningAlgorithm *string `field:"required" json:"signingAlgorithm" yaml:"signingAlgorithm"`
	// The time before which the Certificate will be valid.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#validity AcmpcaCertificate#validity}
	Validity *AcmpcaCertificateValidity `field:"required" json:"validity" yaml:"validity"`
	// These are fields to be overridden in a certificate at the time of issuance.
	//
	// These requires an API_Passthrough template be used or they will be ignored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#api_passthrough AcmpcaCertificate#api_passthrough}
	ApiPassthrough *AcmpcaCertificateApiPassthrough `field:"optional" json:"apiPassthrough" yaml:"apiPassthrough"`
	// Specifies a custom configuration template to use when issuing a certificate.
	//
	// If this parameter is not provided, ACM Private CA defaults to the EndEntityCertificate/V1 template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#template_arn AcmpcaCertificate#template_arn}
	TemplateArn *string `field:"optional" json:"templateArn" yaml:"templateArn"`
	// The time after which the Certificate will be valid.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#validity_not_before AcmpcaCertificate#validity_not_before}
	ValidityNotBefore *AcmpcaCertificateValidityNotBefore `field:"optional" json:"validityNotBefore" yaml:"validityNotBefore"`
}

