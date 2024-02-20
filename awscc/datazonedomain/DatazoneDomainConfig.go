package datazonedomain

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatazoneDomainConfig struct {
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
	// The domain execution role that is created when an Amazon DataZone domain is created.
	//
	// The domain execution role is created in the AWS account that houses the Amazon DataZone domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datazone_domain#domain_execution_role DatazoneDomain#domain_execution_role}
	DomainExecutionRole *string `field:"required" json:"domainExecutionRole" yaml:"domainExecutionRole"`
	// The name of the Amazon DataZone domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datazone_domain#name DatazoneDomain#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the Amazon DataZone domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datazone_domain#description DatazoneDomain#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The identifier of the AWS Key Management Service (KMS) key that is used to encrypt the Amazon DataZone domain, metadata, and reporting data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datazone_domain#kms_key_identifier DatazoneDomain#kms_key_identifier}
	KmsKeyIdentifier *string `field:"optional" json:"kmsKeyIdentifier" yaml:"kmsKeyIdentifier"`
	// The single-sign on configuration of the Amazon DataZone domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datazone_domain#single_sign_on DatazoneDomain#single_sign_on}
	SingleSignOn *DatazoneDomainSingleSignOn `field:"optional" json:"singleSignOn" yaml:"singleSignOn"`
	// The tags specified for the Amazon DataZone domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datazone_domain#tags DatazoneDomain#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

