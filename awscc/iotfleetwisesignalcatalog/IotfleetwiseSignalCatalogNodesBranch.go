package iotfleetwisesignalcatalog


type IotfleetwiseSignalCatalogNodesBranch struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotfleetwise_signal_catalog#fully_qualified_name IotfleetwiseSignalCatalog#fully_qualified_name}.
	FullyQualifiedName *string `field:"required" json:"fullyQualifiedName" yaml:"fullyQualifiedName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotfleetwise_signal_catalog#description IotfleetwiseSignalCatalog#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

