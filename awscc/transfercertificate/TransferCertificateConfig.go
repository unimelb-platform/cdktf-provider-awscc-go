package transfercertificate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TransferCertificateConfig struct {
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
	// Specifies the certificate body to be imported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/transfer_certificate#certificate TransferCertificate#certificate}
	Certificate *string `field:"required" json:"certificate" yaml:"certificate"`
	// Specifies the usage type for the certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/transfer_certificate#usage TransferCertificate#usage}
	Usage *string `field:"required" json:"usage" yaml:"usage"`
	// Specifies the active date for the certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/transfer_certificate#active_date TransferCertificate#active_date}
	ActiveDate *string `field:"optional" json:"activeDate" yaml:"activeDate"`
	// Specifies the certificate chain to be imported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/transfer_certificate#certificate_chain TransferCertificate#certificate_chain}
	CertificateChain *string `field:"optional" json:"certificateChain" yaml:"certificateChain"`
	// A textual description for the certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/transfer_certificate#description TransferCertificate#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies the inactive date for the certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/transfer_certificate#inactive_date TransferCertificate#inactive_date}
	InactiveDate *string `field:"optional" json:"inactiveDate" yaml:"inactiveDate"`
	// Specifies the private key for the certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/transfer_certificate#private_key TransferCertificate#private_key}
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// Key-value pairs that can be used to group and search for certificates.
	//
	// Tags are metadata attached to certificates for any purpose.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/transfer_certificate#tags TransferCertificate#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

