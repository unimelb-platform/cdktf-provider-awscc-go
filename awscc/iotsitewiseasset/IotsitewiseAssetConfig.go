package iotsitewiseasset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IotsitewiseAssetConfig struct {
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
	// The ID of the asset model from which to create the asset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_asset#asset_model_id IotsitewiseAsset#asset_model_id}
	AssetModelId *string `field:"required" json:"assetModelId" yaml:"assetModelId"`
	// A unique, friendly name for the asset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_asset#asset_name IotsitewiseAsset#asset_name}
	AssetName *string `field:"required" json:"assetName" yaml:"assetName"`
	// A description for the asset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_asset#asset_description IotsitewiseAsset#asset_description}
	AssetDescription *string `field:"optional" json:"assetDescription" yaml:"assetDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_asset#asset_hierarchies IotsitewiseAsset#asset_hierarchies}.
	AssetHierarchies interface{} `field:"optional" json:"assetHierarchies" yaml:"assetHierarchies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_asset#asset_properties IotsitewiseAsset#asset_properties}.
	AssetProperties interface{} `field:"optional" json:"assetProperties" yaml:"assetProperties"`
	// A list of key-value pairs that contain metadata for the asset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_asset#tags IotsitewiseAsset#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

