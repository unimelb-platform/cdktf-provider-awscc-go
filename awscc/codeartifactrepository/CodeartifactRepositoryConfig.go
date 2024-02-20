package codeartifactrepository

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CodeartifactRepositoryConfig struct {
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
	// The name of the domain that contains the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codeartifact_repository#domain_name CodeartifactRepository#domain_name}
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The name of the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codeartifact_repository#repository_name CodeartifactRepository#repository_name}
	RepositoryName *string `field:"required" json:"repositoryName" yaml:"repositoryName"`
	// A text description of the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codeartifact_repository#description CodeartifactRepository#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A list of external connections associated with the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codeartifact_repository#external_connections CodeartifactRepository#external_connections}
	ExternalConnections *[]*string `field:"optional" json:"externalConnections" yaml:"externalConnections"`
	// The access control resource policy on the provided repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codeartifact_repository#permissions_policy_document CodeartifactRepository#permissions_policy_document}
	PermissionsPolicyDocument *string `field:"optional" json:"permissionsPolicyDocument" yaml:"permissionsPolicyDocument"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codeartifact_repository#tags CodeartifactRepository#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// A list of upstream repositories associated with the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codeartifact_repository#upstreams CodeartifactRepository#upstreams}
	Upstreams *[]*string `field:"optional" json:"upstreams" yaml:"upstreams"`
}

