package iotsitewiseassetmodel


type IotsitewiseAssetModelAssetModelProperties struct {
	// The data type of the asset model property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_asset_model#data_type IotsitewiseAssetModel#data_type}
	DataType *string `field:"required" json:"dataType" yaml:"dataType"`
	// Customer provided ID for property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_asset_model#logical_id IotsitewiseAssetModel#logical_id}
	LogicalId *string `field:"required" json:"logicalId" yaml:"logicalId"`
	// The name of the asset model property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_asset_model#name IotsitewiseAssetModel#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The property type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_asset_model#type IotsitewiseAssetModel#type}
	Type *IotsitewiseAssetModelAssetModelPropertiesType `field:"required" json:"type" yaml:"type"`
	// The data type of the structure for this property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_asset_model#data_type_spec IotsitewiseAssetModel#data_type_spec}
	DataTypeSpec *string `field:"optional" json:"dataTypeSpec" yaml:"dataTypeSpec"`
	// The unit of the asset model property, such as Newtons or RPM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_asset_model#unit IotsitewiseAssetModel#unit}
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

