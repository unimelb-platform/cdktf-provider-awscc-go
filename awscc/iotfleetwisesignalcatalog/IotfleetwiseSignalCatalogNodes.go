package iotfleetwisesignalcatalog


type IotfleetwiseSignalCatalogNodes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotfleetwise_signal_catalog#actuator IotfleetwiseSignalCatalog#actuator}.
	Actuator *IotfleetwiseSignalCatalogNodesActuator `field:"optional" json:"actuator" yaml:"actuator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotfleetwise_signal_catalog#attribute IotfleetwiseSignalCatalog#attribute}.
	Attribute *IotfleetwiseSignalCatalogNodesAttribute `field:"optional" json:"attribute" yaml:"attribute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotfleetwise_signal_catalog#branch IotfleetwiseSignalCatalog#branch}.
	Branch *IotfleetwiseSignalCatalogNodesBranch `field:"optional" json:"branch" yaml:"branch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotfleetwise_signal_catalog#sensor IotfleetwiseSignalCatalog#sensor}.
	Sensor *IotfleetwiseSignalCatalogNodesSensor `field:"optional" json:"sensor" yaml:"sensor"`
}

