package iotfleetwisesignalcatalog


type IotfleetwiseSignalCatalogNodesSensor struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotfleetwise_signal_catalog#data_type IotfleetwiseSignalCatalog#data_type}.
	DataType *string `field:"required" json:"dataType" yaml:"dataType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotfleetwise_signal_catalog#fully_qualified_name IotfleetwiseSignalCatalog#fully_qualified_name}.
	FullyQualifiedName *string `field:"required" json:"fullyQualifiedName" yaml:"fullyQualifiedName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotfleetwise_signal_catalog#allowed_values IotfleetwiseSignalCatalog#allowed_values}.
	AllowedValues *[]*string `field:"optional" json:"allowedValues" yaml:"allowedValues"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotfleetwise_signal_catalog#description IotfleetwiseSignalCatalog#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotfleetwise_signal_catalog#max IotfleetwiseSignalCatalog#max}.
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotfleetwise_signal_catalog#min IotfleetwiseSignalCatalog#min}.
	Min *float64 `field:"optional" json:"min" yaml:"min"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotfleetwise_signal_catalog#unit IotfleetwiseSignalCatalog#unit}.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

