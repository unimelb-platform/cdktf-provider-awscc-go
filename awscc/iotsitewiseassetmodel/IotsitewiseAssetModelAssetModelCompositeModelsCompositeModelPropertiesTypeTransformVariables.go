package iotsitewiseassetmodel


type IotsitewiseAssetModelAssetModelCompositeModelsCompositeModelPropertiesTypeTransformVariables struct {
	// The friendly name of the variable to be used in the expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_asset_model#name IotsitewiseAssetModel#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The variable that identifies an asset property from which to use values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_asset_model#value IotsitewiseAssetModel#value}
	Value *IotsitewiseAssetModelAssetModelCompositeModelsCompositeModelPropertiesTypeTransformVariablesValue `field:"required" json:"value" yaml:"value"`
}

