package iotsitewiseasset


type IotsitewiseAssetAssetHierarchies struct {
	// The ID of the child asset to be associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_asset#child_asset_id IotsitewiseAsset#child_asset_id}
	ChildAssetId *string `field:"required" json:"childAssetId" yaml:"childAssetId"`
	// The LogicalID of a hierarchy in the parent asset's model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_asset#logical_id IotsitewiseAsset#logical_id}
	LogicalId *string `field:"required" json:"logicalId" yaml:"logicalId"`
}

