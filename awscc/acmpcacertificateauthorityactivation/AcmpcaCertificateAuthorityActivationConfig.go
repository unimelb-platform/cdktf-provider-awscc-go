package acmpcacertificateauthorityactivation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AcmpcaCertificateAuthorityActivationConfig struct {
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
	// Certificate Authority certificate that will be installed in the Certificate Authority.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate_authority_activation#certificate AcmpcaCertificateAuthorityActivation#certificate}
	Certificate *string `field:"required" json:"certificate" yaml:"certificate"`
	// Arn of the Certificate Authority.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate_authority_activation#certificate_authority_arn AcmpcaCertificateAuthorityActivation#certificate_authority_arn}
	CertificateAuthorityArn *string `field:"required" json:"certificateAuthorityArn" yaml:"certificateAuthorityArn"`
	// Certificate chain for the Certificate Authority certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate_authority_activation#certificate_chain AcmpcaCertificateAuthorityActivation#certificate_chain}
	CertificateChain *string `field:"optional" json:"certificateChain" yaml:"certificateChain"`
	// The status of the Certificate Authority.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate_authority_activation#status AcmpcaCertificateAuthorityActivation#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

