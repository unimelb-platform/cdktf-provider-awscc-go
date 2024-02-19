package iotsitewiseproject

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IotsitewiseProjectConfig struct {
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
	// The ID of the portal in which to create the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_project#portal_id IotsitewiseProject#portal_id}
	PortalId *string `field:"required" json:"portalId" yaml:"portalId"`
	// A friendly name for the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_project#project_name IotsitewiseProject#project_name}
	ProjectName *string `field:"required" json:"projectName" yaml:"projectName"`
	// The IDs of the assets to be associated to the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_project#asset_ids IotsitewiseProject#asset_ids}
	AssetIds *[]*string `field:"optional" json:"assetIds" yaml:"assetIds"`
	// A description for the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_project#project_description IotsitewiseProject#project_description}
	ProjectDescription *string `field:"optional" json:"projectDescription" yaml:"projectDescription"`
	// A list of key-value pairs that contain metadata for the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_project#tags IotsitewiseProject#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

