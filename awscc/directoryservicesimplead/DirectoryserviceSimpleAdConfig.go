package directoryservicesimplead

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DirectoryserviceSimpleAdConfig struct {
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
	// The fully qualified domain name for the AWS Managed Simple AD directory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/directoryservice_simple_ad#name DirectoryserviceSimpleAd#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The size of the directory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/directoryservice_simple_ad#size DirectoryserviceSimpleAd#size}
	Size *string `field:"required" json:"size" yaml:"size"`
	// VPC settings of the Simple AD directory server in AWS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/directoryservice_simple_ad#vpc_settings DirectoryserviceSimpleAd#vpc_settings}
	VpcSettings *DirectoryserviceSimpleAdVpcSettings `field:"required" json:"vpcSettings" yaml:"vpcSettings"`
	// The name of the configuration set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/directoryservice_simple_ad#create_alias DirectoryserviceSimpleAd#create_alias}
	CreateAlias interface{} `field:"optional" json:"createAlias" yaml:"createAlias"`
	// Description for the directory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/directoryservice_simple_ad#description DirectoryserviceSimpleAd#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether to enable single sign-on for a Simple Active Directory in AWS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/directoryservice_simple_ad#enable_sso DirectoryserviceSimpleAd#enable_sso}
	EnableSso interface{} `field:"optional" json:"enableSso" yaml:"enableSso"`
	// The password for the default administrative user named Admin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/directoryservice_simple_ad#password DirectoryserviceSimpleAd#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The NetBIOS name for your domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/directoryservice_simple_ad#short_name DirectoryserviceSimpleAd#short_name}
	ShortName *string `field:"optional" json:"shortName" yaml:"shortName"`
}

