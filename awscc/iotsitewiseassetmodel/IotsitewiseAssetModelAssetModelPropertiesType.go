package iotsitewiseassetmodel


type IotsitewiseAssetModelAssetModelPropertiesType struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_asset_model#type_name IotsitewiseAssetModel#type_name}.
	TypeName *string `field:"required" json:"typeName" yaml:"typeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_asset_model#attribute IotsitewiseAssetModel#attribute}.
	Attribute *IotsitewiseAssetModelAssetModelPropertiesTypeAttribute `field:"optional" json:"attribute" yaml:"attribute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_asset_model#metric IotsitewiseAssetModel#metric}.
	Metric *IotsitewiseAssetModelAssetModelPropertiesTypeMetric `field:"optional" json:"metric" yaml:"metric"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_asset_model#transform IotsitewiseAssetModel#transform}.
	Transform *IotsitewiseAssetModelAssetModelPropertiesTypeTransform `field:"optional" json:"transform" yaml:"transform"`
}

