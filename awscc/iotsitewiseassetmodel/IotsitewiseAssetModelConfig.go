package iotsitewiseassetmodel

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IotsitewiseAssetModelConfig struct {
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
	// A unique, friendly name for the asset model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_asset_model#asset_model_name IotsitewiseAssetModel#asset_model_name}
	AssetModelName *string `field:"required" json:"assetModelName" yaml:"assetModelName"`
	// The composite asset models that are part of this asset model.
	//
	// Composite asset models are asset models that contain specific properties.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_asset_model#asset_model_composite_models IotsitewiseAssetModel#asset_model_composite_models}
	AssetModelCompositeModels interface{} `field:"optional" json:"assetModelCompositeModels" yaml:"assetModelCompositeModels"`
	// A description for the asset model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_asset_model#asset_model_description IotsitewiseAssetModel#asset_model_description}
	AssetModelDescription *string `field:"optional" json:"assetModelDescription" yaml:"assetModelDescription"`
	// The hierarchy definitions of the asset model.
	//
	// Each hierarchy specifies an asset model whose assets can be children of any other assets created from this asset model. You can specify up to 10 hierarchies per asset model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_asset_model#asset_model_hierarchies IotsitewiseAssetModel#asset_model_hierarchies}
	AssetModelHierarchies interface{} `field:"optional" json:"assetModelHierarchies" yaml:"assetModelHierarchies"`
	// The property definitions of the asset model. You can specify up to 200 properties per asset model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_asset_model#asset_model_properties IotsitewiseAssetModel#asset_model_properties}
	AssetModelProperties interface{} `field:"optional" json:"assetModelProperties" yaml:"assetModelProperties"`
	// A list of key-value pairs that contain metadata for the asset model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_asset_model#tags IotsitewiseAssetModel#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

