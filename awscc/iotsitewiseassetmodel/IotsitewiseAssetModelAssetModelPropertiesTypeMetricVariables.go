package iotsitewiseassetmodel


type IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariables struct {
	// The friendly name of the variable to be used in the expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iotsitewise_asset_model#name IotsitewiseAssetModel#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The variable that identifies an asset property from which to use values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iotsitewise_asset_model#value IotsitewiseAssetModel#value}
	Value *IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValue `field:"optional" json:"value" yaml:"value"`
}

