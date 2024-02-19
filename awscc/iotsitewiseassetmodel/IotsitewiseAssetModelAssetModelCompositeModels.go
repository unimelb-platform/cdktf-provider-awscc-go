package iotsitewiseassetmodel


type IotsitewiseAssetModelAssetModelCompositeModels struct {
	// A unique, friendly name for the asset composite model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_asset_model#name IotsitewiseAssetModel#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of the composite model. For alarm composite models, this type is AWS/ALARM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_asset_model#type IotsitewiseAssetModel#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The property definitions of the asset model. You can specify up to 200 properties per asset model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_asset_model#composite_model_properties IotsitewiseAssetModel#composite_model_properties}
	CompositeModelProperties interface{} `field:"optional" json:"compositeModelProperties" yaml:"compositeModelProperties"`
	// A description for the asset composite model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_asset_model#description IotsitewiseAssetModel#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

