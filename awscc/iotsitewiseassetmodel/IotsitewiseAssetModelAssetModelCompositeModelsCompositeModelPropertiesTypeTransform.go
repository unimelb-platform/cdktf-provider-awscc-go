package iotsitewiseassetmodel


type IotsitewiseAssetModelAssetModelCompositeModelsCompositeModelPropertiesTypeTransform struct {
	// The mathematical expression that defines the transformation function. You can specify up to 10 functions per expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_asset_model#expression IotsitewiseAssetModel#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// The list of variables used in the expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_asset_model#variables IotsitewiseAssetModel#variables}
	Variables interface{} `field:"required" json:"variables" yaml:"variables"`
}

