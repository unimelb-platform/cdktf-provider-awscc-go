package codeartifactdomain

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CodeartifactDomainConfig struct {
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
	// The name of the domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codeartifact_domain#domain_name CodeartifactDomain#domain_name}
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The access control resource policy on the provided domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codeartifact_domain#permissions_policy_document CodeartifactDomain#permissions_policy_document}
	PermissionsPolicyDocument *string `field:"optional" json:"permissionsPolicyDocument" yaml:"permissionsPolicyDocument"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codeartifact_domain#tags CodeartifactDomain#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

