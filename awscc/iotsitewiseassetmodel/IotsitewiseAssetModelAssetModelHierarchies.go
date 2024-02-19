package iotsitewiseassetmodel


type IotsitewiseAssetModelAssetModelHierarchies struct {
	// The ID of the asset model.
	//
	// All assets in this hierarchy must be instances of the child AssetModelId asset model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_asset_model#child_asset_model_id IotsitewiseAssetModel#child_asset_model_id}
	ChildAssetModelId *string `field:"required" json:"childAssetModelId" yaml:"childAssetModelId"`
	// Customer provided ID for hierarchy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_asset_model#logical_id IotsitewiseAssetModel#logical_id}
	LogicalId *string `field:"required" json:"logicalId" yaml:"logicalId"`
	// The name of the asset model hierarchy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_asset_model#name IotsitewiseAssetModel#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

