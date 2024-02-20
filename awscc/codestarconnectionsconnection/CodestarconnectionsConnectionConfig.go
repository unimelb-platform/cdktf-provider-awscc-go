package codestarconnectionsconnection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CodestarconnectionsConnectionConfig struct {
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
	// The name of the connection. Connection names must be unique in an AWS user account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codestarconnections_connection#connection_name CodestarconnectionsConnection#connection_name}
	ConnectionName *string `field:"required" json:"connectionName" yaml:"connectionName"`
	// The host arn configured to represent the infrastructure where your third-party provider is installed.
	//
	// You must specify either a ProviderType or a HostArn.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codestarconnections_connection#host_arn CodestarconnectionsConnection#host_arn}
	HostArn *string `field:"optional" json:"hostArn" yaml:"hostArn"`
	// The name of the external provider where your third-party code repository is configured.
	//
	// You must specify either a ProviderType or a HostArn.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codestarconnections_connection#provider_type CodestarconnectionsConnection#provider_type}
	ProviderType *string `field:"optional" json:"providerType" yaml:"providerType"`
	// Specifies the tags applied to a connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codestarconnections_connection#tags CodestarconnectionsConnection#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

